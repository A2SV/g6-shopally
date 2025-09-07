package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/shopally-ai/internal/adapter/gateway"
	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/internal/platform"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	rc := platform.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)
	if err := rc.Ping(context.Background()); err != nil {
		log.Fatalf("redis ping: %v", err)
	}
	cache := gateway.NewRedisCache(rc.Client, cfg.Redis.KeyPrefix)

	fxHTTP := gateway.NewFXHTTPGateway(cfg.FX.APIURL, cfg.FX.APIKEY, nil)
	ttl := time.Duration(cfg.FX.CacheTTLSeconds) * time.Second
	fx := gateway.NewCachedFXClient(fxHTTP, cache, ttl)

	// Optional: pre-warm a common FX pair periodically
	warm := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if rate, err := fx.GetRate(ctx, "USD", "ETB"); err != nil {
			log.Printf("worker warm fx error: %v", err)
		} else {
			log.Printf("worker warm fx USD->ETB: %.6f", rate)
		}
	}

	warm()
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		warm()
	}

	ctx := context.Background()

	fcm, err := gateway.NewFCMGateway(ctx, gateway.FCMGatewayConfig{})
	if err != nil {
		log.Printf("FCM init failed (alerts disabled): %v", err)
	} else if t := os.Getenv("FCM_TEST_TOKEN"); t != "" {
		if _, err := fcm.Send(ctx, t, "ShopAlly Alerts Ready", "Worker can send push notifications.", nil); err != nil {
			log.Printf("FCM test send failed: %v", err)
		}
	}

	// TODO: Pass `fcm` into the alerts worker when B2.4 is ready

}
