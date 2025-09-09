package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shopally/chat-history/cmd/chat-history-service/handler"
)

// RegisterRoutes wires all HTTP routes to their respective handlers.
func RegisterRoutes(r *gin.Engine, h *handler.ChatHandler) {
	r.POST("/users/:user_email/chats", h.CreateChat)
	r.GET("/users/:user_email/chats", h.GetUserChats)
	r.GET("/users/:user_email/chats/:chat_id", h.GetChatSession)
	r.POST("/users/:user_email/chats/:chat_id/messages", h.AddMessageToChat)
	r.DELETE("/users/:user_email/chats/:chat_id", h.DeleteChat)

	log.Println("Chat History API routes registered with Gin.")
}
