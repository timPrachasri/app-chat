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
      ws: null,
      newMsg: '',
      chatContent: '',
      email: null,
      username: null,
      joined: false,
    }
  },
  created() {
    var self = this
    this.ws = new WebSocket('ws://localhost:10180/ws')
    this.ws.addEventListener('message', function(e) {
      var msg = JSON.parse(e.data)
      msg.forEach((ech) => {
        self.chatContent += ech.message
      })
      var element = document.getElementById('chat-messages')
      element.scrollTop = element.scrollHeight
    })
  },
  methods: {
    send: function() {
      if (this.newMsg != '') {
        this.ws.send(
          JSON.stringify({
            username: this.username,
            message: `<p>${this.newMsg}</p>`, // Strip out html
            roomName: '1st room',
          }),
        )
        this.newMsg = '' // Reset newMsg
      }
    },
    join: function() {
      this.email = `${this.email}`
      this.username = `${this.username}`
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
