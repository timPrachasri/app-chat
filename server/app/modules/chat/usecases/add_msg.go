package usecase

import (
	"context"
	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

func(c *usecase) AddChat(ctx context.Context, message entities.Message) error {
	return c.chatRepo.CreateChatMsg(ctx, message)
}
