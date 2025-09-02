package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shopally-ai/cmd/api/middleware"
	"github.com/shopally-ai/cmd/api/router"
	"github.com/shopally-ai/internal/adapter/gateway"
	"github.com/shopally-ai/internal/adapter/handler"
	repo "github.com/shopally-ai/internal/adapter/repository"
	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/internal/platform"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"
	"github.com/shopally-ai/pkg/util"
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
		cfg.Redis.Password, // Add password
		cfg.RateLimit.Limit,
		time.Duration(cfg.RateLimit.Window)*time.Second,
		false, // Add TLS flag
	)

	// FX client (provider defaults to exchangerate.host if not configured)
	fxInner := gateway.NewFXHTTPGateway("", "", nil)
	var fxClient domain.IFXClient = fxInner
	// Wrap with Redis cache if available
	if rdb != nil {
		redisCache := gateway.NewRedisCache(rdb.Client, "sa:")
		// Register cache with fxutil so USD->ETB conversions can read the FX rate
		util.SetFXCache(redisCache)
		fxClient = gateway.NewCachedFXClient(fxInner, redisCache, 12*time.Hour)
	}

	// Choose LLM implementation
	lg := gateway.NewGeminiLLMGateway(cfg.Gemini.APIKey, fxClient)

	// Alibaba gateway: use HTTP gateway (real) and pass configuration
	// If you want to force the mock gateway for local development, replace
	// the following line with: ag := gateway.NewMockAlibabaGateway()
	ag := gateway.NewAlibabaHTTPGateway(cfg)

	// Construct usecase and handler for search
	uc := usecase.NewSearchProductsUseCase(ag, lg, nil)
	searchHandler := handler.NewSearchHandler(uc)

	// Alerts: set up Mongo repository and handler
	collName := cfg.Mongo.AlertCollection
	if collName == "" {
		collName = "alerts"
	}
	alertsColl := db.Collection(collName)
	alertRepo := repo.NewMongoAlertRepository(alertsColl)
	alertMgr := usecase.NewAlertManager(alertRepo)

	alertHandler := handler.NewAlertHandler(alertMgr)
	compareHandler := handler.NewCompareHandler(usecase.NewCompareProductsUseCase(lg))

	// Price service & handler
	priceSvc := util.New(ag)
	priceHandler := handler.NewPriceHandler(priceSvc)

	// Initialize router
	router := router.SetupRouter(cfg, limiter, searchHandler, compareHandler, alertHandler, priceHandler)

	// Start the server
	log.Println("Starting server on port", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
