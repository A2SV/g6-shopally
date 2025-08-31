package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shopally-ai/cmd/api/middleware"
	"github.com/shopally-ai/cmd/api/router"
	"github.com/shopally-ai/internal/adapter/gateway"
	"github.com/shopally-ai/internal/adapter/handler"
	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/internal/platform"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to MongoDB using custom db package
	client, err := platform.Connect(cfg.Mongo.URI)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := platform.Disconnect(client); err != nil {
			log.Printf("failed to disconnect MongoDB: %v", err)
		}
	}()
	db := client.Database(cfg.Mongo.Database)
	fmt.Printf("Connected to MongoDB database: %s\n", db.Name())

	// Initialize Redis client
	rdb := platform.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx); err != nil {
		log.Printf("⚠️  Redis connection failed: %v (continuing without Redis)", err)
		rdb = nil
	} else {
		log.Println("✅ Redis connected")
	}

	limiter := middleware.NewRateLimiter(
		cfg.Redis.Host+":"+cfg.Redis.Port,
		cfg.RateLimit.Limit,
		time.Duration(cfg.RateLimit.Window)*time.Second,
	)

	// FX client (provider defaults to exchangerate.host if not configured)
	fxInner := gateway.NewFXHTTPGateway("", "", nil)
	var fxClient domain.IFXClient = fxInner
	// Wrap with Redis cache if available
	if rdb != nil {
		redisCache := gateway.NewRedisCache(rdb.Client, "sa:")
		fxClient = gateway.NewCachedFXClient(fxInner, redisCache, 12*time.Hour)
	}

	// Choose LLM implementation
	var lg domain.LLMGateway
	if os.Getenv("GEMINI_API_KEY") != "" {
		lg = gateway.NewGeminiLLMGateway("", fxClient)
		log.Println("LLM: using Gemini gateway")
	} else {
		lg = gateway.NewMockLLMGateway()
		log.Println("LLM: using Mock gateway (no GEMINI_API_KEY)")
	}

	// Alibaba gateway: use HTTP gateway (real) and pass configuration
	// If you want to force the mock gateway for local development, replace
	// the following line with: ag := gateway.NewMockAlibabaGateway()
	ag := gateway.NewAlibabaHTTPGateway(cfg)

	// Construct usecase and handler
	uc := usecase.NewSearchProductsUseCase(ag, lg, nil)
	searchHandler := handler.NewSearchHandler(uc)

	// Initialize router
	router := router.SetupRouter(cfg, limiter, searchHandler)

	// Start the server
	log.Println("Starting server on port", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
