<template>
  <div id="app">
    <div class="row">
      <div class="col s12">
        <div class="card horizontal">
          <div
            id="chat-messages"
            class="card-content"
            v-html="chatContent"
          ></div>
        </div>
      </div>
    </div>
    <div v-if="joined" class="row">
      <div class="input-field col s8">
        <input v-model="newMsg" type="text" @keyup.enter="send" />
      </div>
      <div class="input-field col s4">
        <button class="waves-effect waves-light btn" @click="send">
          <i class="material-icons right">chat</i>
          Send
        </button>
      </div>
    </div>
    <div v-if="!joined" class="row">
      <div class="input-field col s8">
        <input v-model.trim="username" type="text" placeholder="Username" />
      </div>
      <div class="input-field col s4">
        <button class="waves-effect waves-light btn" @click="join()">
          <i class="material-icons right">done</i>
          Join
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      ws: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      chatContent: '', // A running list of chat messages displayed on the screen
      email: null, // Email address used for grabbing an avatar
      username: null, // Our username
      joined: false, // True if email and username have been filled in
    }
  },
  created() {
    var self = this
    this.ws = new WebSocket('ws://localhost:10180/ws')
    this.ws.addEventListener('message', function(e) {
      console.log('message comin')
      var msg = JSON.parse(e.data)
      self.chatContent += msg.username
      var element = document.getElementById('chat-messages')
      element.scrollTop = element.scrollHeight // Auto scroll to the bottom
    })
  },
  methods: {
    send: function() {
      if (this.newMsg != '') {
        console.log('send', this.newMsg)
        this.ws.send(
          JSON.stringify({
            username: this.username,
            message: `<p>${this.newMsg}</p>`, // Strip out html
          }),
        )
        this.newMsg = '' // Reset newMsg
      }
    },
    join: function() {
      this.email = `<p>${this.email}</p>`
      this.username = `<p>${this.username}</p>`
      this.joined = true
    },
  },
}
</script>

<style>
body {
  display: flex;
  min-height: 100vh;
  flex-direction: column;
}

main {
  flex: 1 0 auto;
}

#chat-messages {
  min-height: 10vh;
  height: 60vh;
  width: 100%;
  overflow-y: scroll;
}
</style>
