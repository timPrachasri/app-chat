package repository

import (
	"log"
	"context"
	"cloud.google.com/go/firestore"
	"github.com/timPrachasri/app-chat/server/app/modules/chat"
	"google.golang.org/api/iterator"
	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

type firestoreChatRepository struct {
	db *firestore.Client
}

func(f *firestoreChatRepository) CreateChatMsg(ctx context.Context, message entities.Message) error {
	_, _, err := f.db.Collection("chat").Doc(message.RoomName).Collection("messages").Add(ctx, message)
	log.Printf("%+v", err)
	if err != nil {
		return err
	}
	return nil
}

func(f *firestoreChatRepository) ReadChatMsgByRoomName(ctx context.Context, roomName string) (messages []*entities.Message, err error) {
	iter := f.db.Collection("chat").Doc(roomName).Collection("messages").Documents(ctx)
	for {
		var message *entities.Message
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                return nil, err
        }
		message, err = new(entities.Message).Parse(doc.Data())
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	for _, ech := range messages {
		log.Printf("%+v", *ech)
	}
	return messages, nil
}


func NewFirestoreChatRepository(firestoreDb *firestore.Client) chat.ChatRepository {
	return &firestoreChatRepository{
		db: firestoreDb,
	}
}