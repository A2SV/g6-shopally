package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopally/chat-history/internal/domain"
	"github.com/shopally/chat-history/internal/errors"
	"github.com/shopally/chat-history/internal/service"
)

type ChatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(svc service.ChatService) *ChatHandler {

	return &ChatHandler{
		chatService: svc,
	}
}

// respondWithJSON writes a success JSON response to the client.
func respondWithJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, domain.Response{Data: data, Error: nil})
}

// respondWithError writes an error JSON response to the client and aborts the request.
func respondWithError(c *gin.Context, statusCode int, errorCode, message string) {
	c.AbortWithStatusJSON(statusCode, domain.Response{Data: nil, Error: &domain.APIError{Code: errorCode, Message: message}})
}

func (h *ChatHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/users/:user_email/chats", h.CreateChat)
	router.GET("/users/:user_email/chats", h.GetUserChats)
	router.GET("/users/:user_email/chats/:chat_id", h.GetChatSession)
	router.POST("/users/:user_email/chats/:chat_id/messages", h.AddMessageToChat)
	router.DELETE("/users/:user_email/chats/:chat_id", h.DeleteChat)

	log.Println("Chat History API routes registered with Gin.")
}

func (h *ChatHandler) CreateChat(c *gin.Context) {
	userEmail := c.Param("user_email")

	var req domain.CreateChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind create chat request for user %s: %v", userEmail, err)
		respondWithError(c, http.StatusBadRequest, "INVALID_INPUT", "Invalid request payload")
		return
	}

	if strings.TrimSpace(req.ChatTitle) == "" {
		log.Printf("Missing chat_title for user %s", userEmail)
		respondWithError(c, http.StatusBadRequest, "INVALID_INPUT", "Chat title is required")
		return
	}

	chatSession, err := h.chatService.CreateChat(c, userEmail, req.ChatTitle)
	if err != nil {
		log.Printf("Service failed to create chat for user %s: %v", userEmail, err)
		respondWithError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Failed to create chat")
		return
	}

	respondWithJSON(c, http.StatusCreated, chatSession)
}

func (h *ChatHandler) AddMessageToChat(c *gin.Context) {
	userEmail := c.Param("user_email")
	chatID := c.Param("chat_id")

	var req domain.AddMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind add message request for user %s, chat %s: %v", userEmail, chatID, err)
		respondWithError(c, http.StatusBadRequest, "INVALID_INPUT", "Invalid request payload")
		return
	}

	if strings.TrimSpace(req.UserPrompt) == "" {
		log.Printf("Missing user_prompt for user %s, chat %s", userEmail, chatID)
		respondWithError(c, http.StatusBadRequest, "INVALID_INPUT", "User prompt is required")
		return
	}

	err := h.chatService.AddMessageToChat(c, userEmail, chatID, req.UserPrompt, req.Products)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			log.Printf("Chat %s not found for user %s during message add: %v", chatID, userEmail, err)
			respondWithError(c, http.StatusNotFound, "NOT_FOUND", "Chat session not found")
			return
		}
		log.Printf("Service failed to add message for user %s, chat %s: %v", userEmail, chatID, err)
		respondWithError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Failed to add message to chat")
		return
	}

	respondWithJSON(c, http.StatusCreated, nil)
}

// DeleteChat handles DELETE /users/{user_email}/chats/{chat_id} to delete a chat session.
func (h *ChatHandler) DeleteChat(c *gin.Context) {
	userEmail := c.Param("user_email")
	chatID := c.Param("chat_id")

	err := h.chatService.DeleteChat(c, userEmail, chatID)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			log.Printf("Chat %s not found for user %s during deletion: %v", chatID, userEmail, err)
			respondWithError(c, http.StatusNotFound, "NOT_FOUND", "Chat session not found")
			return
		}
		log.Printf("Service failed to delete chat %s for user %s: %v", chatID, userEmail, err)
		respondWithError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Failed to delete chat")
		return
	}

	respondWithJSON(c, http.StatusNoContent, nil)
}

// GetUserChats handles GET /users/{user_email}/chats to retrieve all chat sessions.
func (h *ChatHandler) GetUserChats(c *gin.Context) {
	userEmail := c.Param("user_email")

	chats, err := h.chatService.GetUserChats(c.Request.Context(), userEmail)
	if err != nil {
		log.Printf("Service failed to get user chats for %s: %v", userEmail, err)
		respondWithError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Failed to retrieve user chats")
		return
	}

	respondWithJSON(c, http.StatusOK, chats)
}

// GetChatSession handles GET /users/{user_email}/chats/{chat_id} to retrieve a specific chat session.
func (h *ChatHandler) GetChatSession(c *gin.Context) {
	userEmail := c.Param("user_email")
	chatID := c.Param("chat_id")

	chat, err := h.chatService.GetChatSession(c, userEmail, chatID)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			log.Printf("Chat %s not found for user %s: %v", chatID, userEmail, err)
			respondWithError(c, http.StatusNotFound, "NOT_FOUND", "Chat session not found")
			return
		}
		log.Printf("Service failed to get chat session %s for user %s: %v", chatID, userEmail, err)
		respondWithError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Failed to retrieve chat session")
		return
	}

	respondWithJSON(c, http.StatusOK, chat)
}
