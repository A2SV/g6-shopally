package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shopally/chat-history/cmd/chat-history-service/handler"
)

// RegisterRoutes wires all HTTP routes to their respective handlers.
func RegisterRoutes(r *gin.Engine, h *handler.ChatHandler) {
	router := r.Group("/api/chat-history/v1")
	{
		router.POST("/users/:user_email/chats", h.CreateChat)
		router.GET("/users/:user_email/chats", h.GetUserChats)
		router.GET("/users/:user_email/chats/:chat_id", h.GetChatSession)
		router.POST("/users/:user_email/chats/:chat_id/messages", h.AddMessageToChat)
		router.DELETE("/users/:user_email/chats/:chat_id", h.DeleteChat)
	}

	log.Println("Chat History API routes registered with Gin.")
}
