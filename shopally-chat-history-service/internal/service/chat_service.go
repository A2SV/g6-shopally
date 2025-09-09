package service

import (
	"context"
	"strings"
	"time"

	"github.com/shopally/chat-history/internal/domain"
	"github.com/shopally/chat-history/internal/errors"
	"github.com/shopally/chat-history/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// Basic validation (service-level guardrails, even if handlers validate too)
	if strings.TrimSpace(userEmail) == "" || strings.TrimSpace(chatTitle) == "" {
		return nil, errors.ErrInvalidInput
	}

	now := primitive.NewDateTimeFromTime(time.Now())

	chat := domain.ChatSession{
		ChatTitle:   chatTitle,
		StartTime:   now,
		LastUpdated: now,
		Messages:    []domain.Message{},
	}

	chatID, err := s.repo.PushChatSession(ctx, userEmail, chat)
	if err != nil {
		return nil, err
	}
	chat.ChatID = chatID

	return &chat, nil
}

// Use Case: Add a complete user-AI interaction turn to an existing chat session.
// This is called after a user submits a prompt and the AI responds with products.
func (s *chatServiceImpl) AddMessageToChat(ctx context.Context, userEmail, chatID string, userPrompt string, products []domain.Product) error {
	if strings.TrimSpace(userEmail) == "" || strings.TrimSpace(chatID) == "" || strings.TrimSpace(userPrompt) == "" {
		return errors.ErrInvalidInput
	}

	now := primitive.NewDateTimeFromTime(time.Now())

	// Ensure products is non-nil for consistent downstream behavior
	if products == nil {
		products = []domain.Product{}
	}

	msg := domain.Message{
		UserPrompt: userPrompt,
		CreatedAt:  now,
		Products:   products,
	}

	if err := s.repo.PushMessageToSession(ctx, userEmail, chatID, msg); err != nil {
		return err
	}

	return nil
}

// Use Case: Permanently delete an entire chat session.
// This is called when a user wants to remove a conversation from their history.
func (s *chatServiceImpl) DeleteChat(ctx context.Context, userEmail, chatID string) error {
	if strings.TrimSpace(userEmail) == "" || strings.TrimSpace(chatID) == "" {
		return errors.ErrInvalidInput
	}

	if err := s.repo.PullChatSession(ctx, userEmail, chatID); err != nil {
		return err
	}
	return nil
}

// Use Case: Retrieve all chat sessions for a given user.
// This is called, for example, when the frontend displays a list of the user's past conversations.
func (s *chatServiceImpl) GetUserChats(ctx context.Context, userEmail string) ([]domain.ChatSession, error) {
	if strings.TrimSpace(userEmail) == "" {
		return nil, errors.ErrInvalidInput
	}

	history, err := s.repo.FindByUserEmail(ctx, userEmail)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return []domain.ChatSession{}, nil
		}
		return nil, err
	}
	return history.ChatSessions, nil
}

// Use Case: Retrieve a specific chat session and its full message history.
// This is called when a user clicks on a chat title to view the conversation details.
func (s *chatServiceImpl) GetChatSession(ctx context.Context, userEmail, chatID string) (*domain.ChatSession, error) {
	if strings.TrimSpace(userEmail) == "" || strings.TrimSpace(chatID) == "" {
		return nil, errors.ErrInvalidInput
	}

	history, err := s.repo.FindByUserEmail(ctx, userEmail)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	for i := range history.ChatSessions {
		if history.ChatSessions[i].ChatID == chatID {
			return &history.ChatSessions[i], nil
		}
	}
	return nil, errors.ErrNotFound //If loop finishes, chat session was not found!
}
