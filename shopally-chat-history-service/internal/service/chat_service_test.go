package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/shopally/chat-history/internal/domain"
	apperrors "github.com/shopally/chat-history/internal/errors"
	mocks "github.com/shopally/chat-history/internal/mocks"
)

// Test suite
type ChatServiceTestSuite struct {
	suite.Suite
	repo *mocks.MockChatRepository
	svc  domain.ChatService
}

func (s *ChatServiceTestSuite) SetupTest() {
	s.repo = &mocks.MockChatRepository{}
	s.svc = NewChatService(s.repo)
}

func (s *ChatServiceTestSuite) TestCreateChat_Success() {
	user := "user@example.com"
	title := "My Chat"

	// Expect PushChatSession to be called with a ChatSession that has the title
	s.repo.On("PushChatSession", mock.Anything, user, mock.MatchedBy(func(arg interface{}) bool {
		cs, ok := arg.(domain.ChatSession)
		if !ok {
			return false
		}
		// At the time PushChatSession is called the ChatID will be empty; repository returns the ID.
		return cs.ChatTitle == title
	})).Return("generated-id", nil)

	chat, err := s.svc.CreateChat(context.Background(), user, title)
	s.Require().NoError(err)
	s.Require().NotNil(chat)
	s.Equal(title, chat.ChatTitle)

	s.repo.AssertExpectations(s.T())
}

func (s *ChatServiceTestSuite) TestCreateChat_InvalidInput() {
	_, err := s.svc.CreateChat(context.Background(), "", "")
	s.ErrorIs(err, apperrors.ErrInvalidInput)
}

func (s *ChatServiceTestSuite) TestAddMessageToChat_Success() {
	user := "u@e.com"
	chatID := "chat123"
	prompt := "Find shoes"

	s.repo.On("PushMessageToSession", mock.Anything, user, chatID, mock.MatchedBy(func(arg interface{}) bool {
		m, ok := arg.(domain.Message)
		if !ok {
			return false
		}
		return m.UserPrompt == prompt
	})).Return(nil)

	err := s.svc.AddMessageToChat(context.Background(), user, chatID, prompt, nil)
	s.Require().NoError(err)
	s.repo.AssertExpectations(s.T())
}

func (s *ChatServiceTestSuite) TestAddMessageToChat_InvalidInput() {
	err := s.svc.AddMessageToChat(context.Background(), "", "", "", nil)
	s.ErrorIs(err, apperrors.ErrInvalidInput)
}

func (s *ChatServiceTestSuite) TestDeleteChat_Success() {
	user := "u@e.com"
	chatID := "c1"

	s.repo.On("PullChatSession", mock.Anything, user, chatID).Return(nil)

	err := s.svc.DeleteChat(context.Background(), user, chatID)
	s.Require().NoError(err)
	s.repo.AssertExpectations(s.T())
}

func (s *ChatServiceTestSuite) TestDeleteChat_InvalidInput() {
	err := s.svc.DeleteChat(context.Background(), "", "")
	s.ErrorIs(err, apperrors.ErrInvalidInput)
}

func (s *ChatServiceTestSuite) TestGetUserChats_NotFound_ReturnsEmpty() {
	user := "nouser@example.com"
	s.repo.On("FindByUserEmail", mock.Anything, user).Return(nil, apperrors.ErrNotFound)

	chats, err := s.svc.GetUserChats(context.Background(), user)
	s.Require().NoError(err)
	s.Empty(chats)
	s.repo.AssertExpectations(s.T())
}

func (s *ChatServiceTestSuite) TestGetUserChats_Success() {
	user := "u@e.com"
	history := &domain.ChatHistory{
		ChatSessions: []domain.ChatSession{{ChatID: "c1", ChatTitle: "t1"}},
	}
	s.repo.On("FindByUserEmail", mock.Anything, user).Return(history, nil)

	chats, err := s.svc.GetUserChats(context.Background(), user)
	s.Require().NoError(err)
	s.Len(chats, 1)
	s.Equal("c1", chats[0].ChatID)
	s.repo.AssertExpectations(s.T())
}

func (s *ChatServiceTestSuite) TestGetChatSession_Success() {
	user := "u@e.com"
	chatID := "c123"
	history := &domain.ChatHistory{
		ChatSessions: []domain.ChatSession{{ChatID: chatID, ChatTitle: "t"}},
	}
	s.repo.On("FindByUserEmail", mock.Anything, user).Return(history, nil)

	cs, err := s.svc.GetChatSession(context.Background(), user, chatID)
	s.Require().NoError(err)
	s.NotNil(cs)
	s.Equal(chatID, cs.ChatID)
}

func (s *ChatServiceTestSuite) TestGetChatSession_NotFound() {
	user := "u@e.com"
	chatID := "missing"
	history := &domain.ChatHistory{
		ChatSessions: []domain.ChatSession{{ChatID: "other", ChatTitle: "t"}},
	}
	s.repo.On("FindByUserEmail", mock.Anything, user).Return(history, nil)

	cs, err := s.svc.GetChatSession(context.Background(), user, chatID)
	s.Nil(cs)
	s.ErrorIs(err, apperrors.ErrNotFound)
}

func TestChatServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ChatServiceTestSuite))
}
