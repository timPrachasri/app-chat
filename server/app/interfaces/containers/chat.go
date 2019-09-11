package container

import (
	"cloud.google.com/go/firestore"

	connectors "github.com/timPrachasri/app-chat/server/app/interfaces/connectors"
	chat "github.com/timPrachasri/app-chat/server/app/modules/chat"
	chatRepo "github.com/timPrachasri/app-chat/server/app/modules/chat/repositories"
	chatUsecase "github.com/timPrachasri/app-chat/server/app/modules/chat/usecases"
)

// ChatContainer container is a way for you to build everything from the ground up. It's a bootstrap file for module chat
type ChatContainer struct {
	db                     *firestore.Client
	chatRepository 	chat.ChatRepository
	chatUsecase    chat.Usecase
}

// BootUsecase is a function for getting chat usecase
func (c *ChatContainer) BootUsecase() chat.Usecase {
	if c.chatUsecase == nil {
		c.chatUsecase = chatUsecase.NewUsecase(
			c.BootChatRepo(),
		)
	}
	return c.chatUsecase
}

// BootChatRepo is a function for getting chat repo.
func (c *ChatContainer) BootChatRepo() chat.ChatRepository {
	if c.chatRepository == nil {
		c.chatRepository = chatRepo.NewFirestoreChatRepository(
			c.BootDB(),
		)
	}
	return c.chatRepository
}


// BootDB is a function for getting db instance
func (c *ChatContainer) BootDB() *firestore.Client {
	if c.db != nil {
		return c.db
	}
	c.db = connectors.ConnectFirestore()
	return c.db
}


// NewChatContainer is a function for creating client container instance
func NewChatContainer() *ChatContainer {
	return &ChatContainer{}
}
