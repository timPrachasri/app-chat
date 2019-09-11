package chat

import (
	entities "github.com/timPrachasri/app-chat/server/app/entities"
	"context"
)

type ChatRepository interface {
	CreateChatMsg(ctx context.Context, message entities.Message) error
}
