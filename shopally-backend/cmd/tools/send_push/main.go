package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/shopally-ai/internal/adapter/gateway"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	token := os.Getenv("FCM_TEST_TOKEN")
	if token == "" {
		log.Fatal("FCM_TEST_TOKEN is required")
	}

	gw, err := gateway.NewFCMGateway(ctx, gateway.FCMGatewayConfig{})
	if err != nil {
		log.Fatalf("init FCM: %v", err)
	}

	id, err := gw.Send(ctx, token, "ShopAlly Test", "This is a test push from the backend.", map[string]string{
		"type": "test",
	})
	if err != nil {
		log.Fatalf("send: %v", err)
	}
	log.Printf("Sent message ID: %s", id)
}