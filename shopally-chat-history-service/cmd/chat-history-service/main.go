package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/shopally/chat-history/cmd/chat-history-service/handler"
	"github.com/shopally/chat-history/cmd/chat-history-service/router"
	"github.com/shopally/chat-history/internal/platform/config"
	"github.com/shopally/chat-history/internal/platform/mongodb"
	"github.com/shopally/chat-history/internal/repository"
	"github.com/shopally/chat-history/internal/service"
)

func main() {
	// Load configuration
	cfg := config.FromEnv()

	// Configure Gin mode from env if set (GIN_MODE). Defaults to debug otherwise.
	if mode := os.Getenv(gin.EnvGinMode); mode != "" {
		gin.SetMode(mode)
	}

	// Connect to MongoDB
	client, err := mongodb.Connect(cfg.MongoURI, cfg.MongoDBName, cfg.MongoCollection)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer func() {
		if err := mongodb.Disconnect(client); err != nil {
			log.Printf("error disconnecting mongo client: %v", err)
		}
	}()

	// Initialize repository -> service -> handler
	collection := client.Database(cfg.MongoDBName).Collection(cfg.MongoCollection)
	chatRepo := repository.NewMongoChatRepository(collection)
	chatSvc := service.NewChatService(chatRepo)
	chatHandler := handler.NewChatHandler(chatSvc)

	// Setup router
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.RegisterRoutes(r, chatHandler)

	// HTTP server with graceful shutdown
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Chat History Service listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdown: %v", err)
	}

	log.Println("Server exited cleanly.")
}
