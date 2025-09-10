package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/shopally/chat-history/internal/domain"
)

// MockChatService is a mock implementation of service.ChatService generated-style.
type MockChatService struct {
	mock.Mock
}

func (_m *MockChatService) CreateChat(ctx context.Context, userEmail, chatTitle string) (*domain.ChatSession, error) {
	ret := _m.Called(ctx, userEmail, chatTitle)

	var r0 *domain.ChatSession
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.ChatSession); ok {
		r0 = rf(ctx, userEmail, chatTitle)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ChatSession)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userEmail, chatTitle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockChatService) AddMessageToChat(ctx context.Context, userEmail, chatID string, userPrompt string, products []domain.Product) error {
	ret := _m.Called(ctx, userEmail, chatID, userPrompt, products)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []domain.Product) error); ok {
		r0 = rf(ctx, userEmail, chatID, userPrompt, products)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockChatService) DeleteChat(ctx context.Context, userEmail, chatID string) error {
	ret := _m.Called(ctx, userEmail, chatID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userEmail, chatID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockChatService) GetUserChats(ctx context.Context, userEmail string) ([]domain.ChatSession, error) {
	ret := _m.Called(ctx, userEmail)

	var r0 []domain.ChatSession
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.ChatSession); ok {
		r0 = rf(ctx, userEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ChatSession)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockChatService) GetChatSession(ctx context.Context, userEmail, chatID string) (*domain.ChatSession, error) {
	ret := _m.Called(ctx, userEmail, chatID)

	var r0 *domain.ChatSession
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.ChatSession); ok {
		r0 = rf(ctx, userEmail, chatID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ChatSession)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userEmail, chatID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
