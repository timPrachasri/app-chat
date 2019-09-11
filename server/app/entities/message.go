package entities

import (
	"time"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

// Message Define our message object
type Message struct {
	RoomName  string    `json:"roomName" structs:"roomName" firestore:"-"`
	UserName  string    `json:"username" structs:"userID" firestore:"userName"`
	Content   string    `json:"message" structs:"data" firestore:"content"`
	CreatedAt time.Time `json:"createdAt" structs:"createdAt" firestore:"createdAt"`
}

func (m *Message) Parse(data interface{}) (msg *Message, err error) {
	if err := mapstructure.Decode(data, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Message) ToMap() map[string]interface{} {
	return structs.Map(m)
}
