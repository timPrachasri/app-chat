package usecase

import (
	"context"
	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

func(c *usecase) ReadChat(ctx context.Context, roomName string) ([]*entities.Message, error) {
	return c.chatRepo.ReadChatMsgByRoomName(ctx, roomName)
}
