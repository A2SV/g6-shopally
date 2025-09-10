package domain

import (
	"context"
)

// ChatRepository defines the interface for database operations related to chat history.
type ChatRepository interface {
	FindByUserEmail(ctx context.Context, userEmail string) (*ChatHistory, error)
	PushChatSession(ctx context.Context, userEmail string, chatSession ChatSession) (string, error)
	PushMessageToSession(ctx context.Context, userEmail, chatID string, message Message) error
	PullChatSession(ctx context.Context, userEmail, chatID string) error
}

// ChatService defines the interface for chat-related business logic operations.
type ChatService interface {
	CreateChat(ctx context.Context, userEmail, chatTitle string) (*ChatSession, error)
	AddMessageToChat(ctx context.Context, userEmail, chatID string, userPrompt string, products []Product) error
	DeleteChat(ctx context.Context, userEmail, chatID string) error
	GetUserChats(ctx context.Context, userEmail string) ([]ChatSession, error)
	GetChatSession(ctx context.Context, userEmail, chatID string) (*ChatSession, error)
}
