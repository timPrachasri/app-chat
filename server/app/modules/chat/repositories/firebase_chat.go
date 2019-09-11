package repository

import (
	"context"
	"cloud.google.com/go/firestore"
	"github.com/timPrachasri/app-chat/server/app/modules/chat"
	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

type firestoreChatRepository struct {
	db *firestore.Client
}

func(f *firestoreChatRepository) CreateChatMsg(ctx context.Context, message entities.Message) error {
	_, _, err := f.db.Collection("chat").Doc(message.RoomName).Collection("messages").Add(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

// func(f *firestoreChatRepository) ReadChatMsgByRoomName(ctx context.Context, roomName string) (entities.message, error) {
// 	dsnap, err := f.db.Collection(f.collectionName).Doc(roomName).Get(ctx)
// 	if err != nil {
//         return nil, err
// 	}
// 	var m entities.Message
// 	dsnap.DataTo(&m)
// 	return m, nil
// }


func NewFirestoreChatRepository(firestoreDb *firestore.Client) chat.ChatRepository {
	return &firestoreChatRepository{
		db: firestoreDb,
	}
}