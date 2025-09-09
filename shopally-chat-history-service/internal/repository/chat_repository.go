package repository

import (
	"context"

	"github.com/shopally/chat-history/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// ChatRepository defines the interface for database operations related to chat history.
type ChatRepository interface {
	FindByUserEmail(ctx context.Context, userEmail string) (*domain.ChatHistory, error)
	PushChatSession(ctx context.Context, userEmail string, chatSession domain.ChatSession) error
	PushMessageToSession(ctx context.Context, userEmail, chatID string, message domain.Message) error
	PullChatSession(ctx context.Context, userEmail, chatID string) error
}

// mongoChatRepository implements ChatRepository for MongoDB.
type mongoChatRepository struct {
	collection *mongo.Collection
}

// NewMongoChatRepository creates a new MongoDB implementation of ChatRepository.
func NewMongoChatRepository(collection *mongo.Collection) ChatRepository {
	return &mongoChatRepository{
		collection: collection,
	}
}

// FindByUserEmail retrieves the entire ChatHistory document for a given user email.
func (r *mongoChatRepository) FindByUserEmail(ctx context.Context, userEmail string) (*domain.ChatHistory, error) {
	var chatHistory domain.ChatHistory
	return &chatHistory, nil
}

// PushChatSession adds a new chat session to a user's chat history document.
// It creates the top-level document if it doesn't exist.
func (r *mongoChatRepository) PushChatSession(ctx context.Context, userEmail string, chatSession domain.ChatSession) error {
	return nil
}

// PushMessageToSession adds a new message turn to a specific chat session within a user's history.
func (r *mongoChatRepository) PushMessageToSession(ctx context.Context, userEmail, chatID string, message domain.Message) error {
	return nil
}

// PullChatSession removes a specific chat session from a user's history.
func (r *mongoChatRepository) PullChatSession(ctx context.Context, userEmail, chatID string) error {
	return nil
}
