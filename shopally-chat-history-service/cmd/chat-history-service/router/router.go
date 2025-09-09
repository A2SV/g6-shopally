package handler

import (
	"net/http"

	"github.com/gorilla/mux" // Example router, you can use net/http, chi, gin etc.
	"github.com/shopally/chat-history/internal/service"
	"github.com/spiffe/go-spiffe/v2/logger"
)

// ChatHandler handles HTTP requests related to chat history.
type ChatHandler struct {
	chatService service.ChatService
}

// NewChatHandler creates and returns a new instance of ChatHandler.
func NewChatHandler(svc service.ChatService, logger logger.Logger) *ChatHandler {
	return &ChatHandler{
		chatService: svc,
	}
}

// CreateChatRequest represents the request body for creating a new chat session.
type CreateChatRequest struct {
	ChatTitle string
}

// RegisterRoutes registers the chat history API endpoints with the provided router.
func (h *ChatHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/{user_email}/chats", h.CreateChat).Methods("POST")
	router.HandleFunc("/users/{user_email}/chats", h.GetUserChats).Methods("GET")
	router.HandleFunc("/users/{user_email}/chats/{chat_id}", h.GetChatSession).Methods("GET")
	router.HandleFunc("/users/{user_email}/chats/{chat_id}/messages", h.AddMessageToChat).Methods("POST")
	router.HandleFunc("/users/{user_email}/chats/{chat_id}", h.DeleteChat).Methods("DELETE")
}

// CreateChat handles POST /users/{user_email}/chats to create a new chat session.
func (h *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
}

// AddMessageToChat handles POST /users/{user_email}/chats/{chat_id}/messages to add a new message turn.
func (h *ChatHandler) AddMessageToChat(w http.ResponseWriter, r *http.Request) {
}

// DeleteChat handles DELETE /users/{user_email}/chats/{chat_id} to delete a chat session.
func (h *ChatHandler) DeleteChat(w http.ResponseWriter, r *http.Request) {

}

// GetUserChats handles GET /users/{user_email}/chats to retrieve all chat sessions.
func (h *ChatHandler) GetUserChats(w http.ResponseWriter, r *http.Request) {
}

// GetChatSession handles GET /users/{user_email}/chats/{chat_id} to retrieve a specific chat session.
func (h *ChatHandler) GetChatSession(w http.ResponseWriter, r *http.Request) {

}
