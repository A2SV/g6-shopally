package service

import (
	"context"

	"github.com/shopally/chat-history/internal/domain"
	"github.com/shopally/chat-history/internal/errors"
	"github.com/shopally/chat-history/internal/repository"
)

// ChatService defines the interface for chat-related business logic operations.
type ChatService interface {
	CreateChat(ctx context.Context, userEmail, chatTitle string) (*domain.ChatSession, error)
	AddMessageToChat(ctx context.Context, userEmail, chatID string, userPrompt string, products []domain.Product) error
	DeleteChat(ctx context.Context, userEmail, chatID string) error
	GetUserChats(ctx context.Context, userEmail string) ([]domain.ChatSession, error)
	GetChatSession(ctx context.Context, userEmail, chatID string) (*domain.ChatSession, error)
}

// chatServiceImpl implements the ChatService interface.
type chatServiceImpl struct {
	repo repository.ChatRepository
}

// NewChatService creates and returns a new instance of ChatService.
func NewChatService(repo repository.ChatRepository) ChatService {
	return &chatServiceImpl{
		repo: repo,
	}
}

func (s *chatServiceImpl) CreateChat(ctx context.Context, userEmail, chatTitle string) (*domain.ChatSession, error) {

	return nil, nil
}

// Use Case: Add a complete user-AI interaction turn to an existing chat session.
// This is called after a user submits a prompt and the AI responds with products.
func (s *chatServiceImpl) AddMessageToChat(ctx context.Context, userEmail, chatID string, userPrompt string, products []domain.Product) error {
	return nil
}

// Use Case: Permanently delete an entire chat session.
// This is called when a user wants to remove a conversation from their history.
func (s *chatServiceImpl) DeleteChat(ctx context.Context, userEmail, chatID string) error {

	return nil
}

// Use Case: Retrieve all chat sessions for a given user.
// This is called, for example, when the frontend displays a list of the user's past conversations.
func (s *chatServiceImpl) GetUserChats(ctx context.Context, userEmail string) ([]domain.ChatSession, error) {
	return nil, nil
}

// Use Case: Retrieve a specific chat session and its full message history.
// This is called when a user clicks on a chat title to view the conversation details.
func (s *chatServiceImpl) GetChatSession(ctx context.Context, userEmail, chatID string) (*domain.ChatSession, error) {
	return nil, errors.ErrNotFound // If loop finishes, chat session was not found
}
