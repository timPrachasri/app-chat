package usecase

import (
	"github.com/timPrachasri/app-chat/server/app/modules/chat"
)

type usecase struct {
	chatRepo chat.ChatRepository
}

func NewUsecase(chatRepo chat.ChatRepository) chat.Usecase {
	return &usecase{
		chatRepo: chatRepo,
	}
}