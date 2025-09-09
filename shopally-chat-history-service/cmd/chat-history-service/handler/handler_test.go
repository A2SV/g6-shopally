package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/shopally/chat-history/internal/domain"
	apperrors "github.com/shopally/chat-history/internal/errors"
	mocks "github.com/shopally/chat-history/internal/mocks"
)

type HandlerTestSuite struct {
	suite.Suite
	svc *mocks.MockChatService
	h   *ChatHandler
	r   *gin.Engine
}

func (s *HandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.svc = &mocks.MockChatService{}
	s.h = NewChatHandler(s.svc)
	s.r = gin.New()
	s.r.POST("/users/:user_email/chats", s.h.CreateChat)
	s.r.POST("/users/:user_email/chats/:chat_id/messages", s.h.AddMessageToChat)
	s.r.DELETE("/users/:user_email/chats/:chat_id", s.h.DeleteChat)
	s.r.GET("/users/:user_email/chats", s.h.GetUserChats)
	s.r.GET("/users/:user_email/chats/:chat_id", s.h.GetChatSession)
}

func (s *HandlerTestSuite) TestCreateChat_Success() {
	user := "u@e.com"
	body := domain.CreateChatRequest{ChatTitle: "title"}
	b, _ := json.Marshal(body)

	s.svc.On("CreateChat", mock.Anything, user, body.ChatTitle).Return(&domain.ChatSession{ChatID: "cid", ChatTitle: body.ChatTitle}, nil)

	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	s.Equal(http.StatusCreated, w.Code)
}

func (s *HandlerTestSuite) TestCreateChat_InvalidJSON() {
	user := "u@e.com"
	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats", bytes.NewReader([]byte("not-json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *HandlerTestSuite) TestCreateChat_MissingTitle() {
	user := "u@e.com"
	body := domain.CreateChatRequest{ChatTitle: ""}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *HandlerTestSuite) TestCreateChat_ServiceError() {
	user := "u@e.com"
	body := domain.CreateChatRequest{ChatTitle: "title"}
	b, _ := json.Marshal(body)

	s.svc.On("CreateChat", mock.Anything, user, body.ChatTitle).Return(nil, errors.New("boom"))

	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *HandlerTestSuite) TestAddMessageToChat_Success() {
	user := "u@e.com"
	chatID := "c1"
	body := domain.AddMessageRequest{UserPrompt: "hi", Products: []domain.Product{}}
	b, _ := json.Marshal(body)

	s.svc.On("AddMessageToChat", mock.Anything, user, chatID, body.UserPrompt, body.Products).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats/"+chatID+"/messages", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
}

func (s *HandlerTestSuite) TestAddMessageToChat_InvalidJSON() {
	user := "u@e.com"
	chatID := "c1"
	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats/"+chatID+"/messages", bytes.NewReader([]byte("nope")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *HandlerTestSuite) TestAddMessageToChat_MissingPrompt() {
	user := "u@e.com"
	chatID := "c1"
	body := domain.AddMessageRequest{UserPrompt: ""}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats/"+chatID+"/messages", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *HandlerTestSuite) TestAddMessageToChat_NotFound() {
	user := "u@e.com"
	chatID := "c1"
	body := domain.AddMessageRequest{UserPrompt: "hi"}
	b, _ := json.Marshal(body)

	s.svc.On("AddMessageToChat", mock.Anything, user, chatID, body.UserPrompt, mock.Anything).Return(apperrors.ErrNotFound)

	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats/"+chatID+"/messages", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusNotFound, w.Code)
}

func (s *HandlerTestSuite) TestAddMessageToChat_ServiceError() {
	user := "u@e.com"
	chatID := "c1"
	body := domain.AddMessageRequest{UserPrompt: "hi"}
	b, _ := json.Marshal(body)

	s.svc.On("AddMessageToChat", mock.Anything, user, chatID, body.UserPrompt, mock.Anything).Return(errors.New("boom"))

	req := httptest.NewRequest(http.MethodPost, "/users/"+user+"/chats/"+chatID+"/messages", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *HandlerTestSuite) TestDeleteChat_Success() {
	user := "u@e.com"
	chatID := "c1"

	s.svc.On("DeleteChat", mock.Anything, user, chatID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/users/"+user+"/chats/"+chatID, nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusNoContent, w.Code)
}

func (s *HandlerTestSuite) TestDeleteChat_NotFound() {
	user := "u@e.com"
	chatID := "c1"

	s.svc.On("DeleteChat", mock.Anything, user, chatID).Return(apperrors.ErrNotFound)

	req := httptest.NewRequest(http.MethodDelete, "/users/"+user+"/chats/"+chatID, nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusNotFound, w.Code)
}

func (s *HandlerTestSuite) TestDeleteChat_ServiceError() {
	user := "u@e.com"
	chatID := "c1"

	s.svc.On("DeleteChat", mock.Anything, user, chatID).Return(errors.New("boom"))

	req := httptest.NewRequest(http.MethodDelete, "/users/"+user+"/chats/"+chatID, nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *HandlerTestSuite) TestGetUserChats_Success() {
	user := "u@e.com"
	s.svc.On("GetUserChats", mock.Anything, user).Return([]domain.ChatSession{{ChatID: "c1"}}, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/"+user+"/chats", nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *HandlerTestSuite) TestGetUserChats_ServiceError() {
	user := "u@e.com"
	s.svc.On("GetUserChats", mock.Anything, user).Return(nil, errors.New("boom"))

	req := httptest.NewRequest(http.MethodGet, "/users/"+user+"/chats", nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *HandlerTestSuite) TestGetChatSession_Success() {
	user := "u@e.com"
	chatID := "c1"
	s.svc.On("GetChatSession", mock.Anything, user, chatID).Return(&domain.ChatSession{ChatID: chatID}, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/"+user+"/chats/"+chatID, nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *HandlerTestSuite) TestGetChatSession_NotFound() {
	user := "u@e.com"
	chatID := "c1"
	s.svc.On("GetChatSession", mock.Anything, user, chatID).Return(nil, apperrors.ErrNotFound)

	req := httptest.NewRequest(http.MethodGet, "/users/"+user+"/chats/"+chatID, nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusNotFound, w.Code)
}

func (s *HandlerTestSuite) TestGetChatSession_ServiceError() {
	user := "u@e.com"
	chatID := "c1"
	s.svc.On("GetChatSession", mock.Anything, user, chatID).Return(nil, errors.New("boom"))

	req := httptest.NewRequest(http.MethodGet, "/users/"+user+"/chats/"+chatID, nil)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
