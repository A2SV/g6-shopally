package gateway

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/shopally-ai/internal/config"
	"google.golang.org/api/option"
)

type FCMGatewayConfig struct {
	Config *config.Config
}

type FCMClient interface {
	Send(ctx context.Context, msg *messaging.Message) (string, error)
}

type FCMGateway struct {
	client FCMClient
}

func NewFCMGateway(ctx context.Context, cfg FCMGatewayConfig) (*FCMGateway, error) {
	var opt option.ClientOption
	var err error

	// First try to use the config object
	if cfg.Config.ValidateFirebaseCredentials() == nil {
		credsJSON, err := cfg.Config.BuildFirebaseCredentialsJSON()
		if err != nil {
			return nil, fmt.Errorf("building Firebase credentials from config: %w", err)
		}
		opt = option.WithCredentialsJSON([]byte(credsJSON))
	} else if os.Getenv("FIREBASE_CREDENTIALS_FILE") != "" {
		opt = option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_FILE"))
	} else if os.Getenv("FIREBASE_CREDENTIALS_JSON") != "" {
		opt = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS_JSON")))
	} else {
		return nil, fmt.Errorf("missing Firebase credentials: set FIREBASE_CREDENTIALS_FILE or FIREBASE_CREDENTIALS_JSON environment variables, or provide Config with Firebase credentials")
	}

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("firebase.NewApp: %w", err)
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("app.Messaging: %w", err)
	}
	return &FCMGateway{client: client}, nil
}

// NewFCMGatewayWithClient is for tests where we inject a fake client.
func NewFCMGatewayWithClient(client FCMClient) *FCMGateway {
	return &FCMGateway{client: client}
}

func (g *FCMGateway) Send(ctx context.Context, token, title, body string, data map[string]string) (string, error) {
	msg := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				ChannelID: "alerts",
				Sound:     "default",
			},
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{"apns-priority": "10"},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{Sound: "default"},
			},
		},
	}

	id, err := g.client.Send(ctx, msg)
	if err != nil {
		return "", err
	}

	return id, nil
}
