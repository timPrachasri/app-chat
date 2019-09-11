package chat

import (
	"context"

	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

type ChatRepository interface {
	CreateChatMsg(ctx context.Context, message entities.Message) error
	ReadChatMsgByRoomName(ctx context.Context, roomName string) ([]*entities.Message, error)
}
