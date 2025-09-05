package worker

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/shopally-ai/internal/adapter/gateway"
	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/internal/platform"
	workerpkg "github.com/shopally-ai/internal/worker"
	"github.com/shopally-ai/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func WorkerStart() {
	log.Println("Starting background worker...")
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

	ctx := context.Background()

	fcm, err := gateway.NewFCMGateway(ctx, gateway.FCMGatewayConfig{})
	if err != nil {
		log.Printf("FCM init failed (alerts disabled): %v", err)
	} else if t := os.Getenv("FCM_TEST_TOKEN"); t != "" {
		if _, err := fcm.Send(ctx, t, "ShopAlly Alerts Ready", "Worker can send push notifications.", nil); err != nil {
			log.Printf("FCM test send failed: %v", err)
		}
	}

	// Mongo setup for alerts collection
	mongoClient, err := platform.Connect(cfg.Mongo.URI)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	db := mongoClient.Database(cfg.Mongo.Database)
	alertsColl := db.Collection(cfg.Mongo.AlertCollection)

	// Ensure an index on IsActive to speed up worker scans
	_, err = alertsColl.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "IsActive", Value: 1}},
		Options: options.Index().SetName("idx_isactive"),
	})
	if err != nil {
		log.Printf("create index on alerts.IsActive failed: %v", err)
	}

	// Alibaba config and price service (use AliExpress lightweight client)
	_ = gateway.NewAlibabaHTTPGateway // keep constructed gateway unused; we rely on PriceService fetcher
	priceSvc := util.NewWithAli(cfg.Aliexpress.AppKey, cfg.Aliexpress.AppSecret, cfg.Aliexpress.BaseURL)

	// Start alerts worker if FCM is available
	if fcm != nil {
		aw := workerpkg.NewAlertsWorker(alertsColl, priceSvc, fcm)
		// run periodic loop alongside FX warming
		go aw.Run(ctx)
	}

	// Keep the process alive: retain FX warm ticker
	ticker := time.NewTicker(6 * time.Hour)
	defer ticker.Stop()
	for range ticker.C {
		warm()
	}
}
