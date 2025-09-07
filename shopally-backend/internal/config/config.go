package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
	}

	Mongo struct {
		URI             string
		Database        string
		AlertCollection string
	}

	Redis struct {
		Host            string
		Port            string
		Password        string
		DB              int
		ViewTrackingTTL int
		KeyPrefix       string
	}

	FX struct {
		APIURL          string
		APIKEY          string
		CacheTTLSeconds int
	}

	OAuth struct {
		Google struct {
			ClientID     string
			ClientSecret string
			RedirectURI  string
		}

		Aliexpress struct {
			ClientID     string
			ClientSecret string
			RedirectURI  string
		}
	}

	RateLimit struct {
		Limit  int
		Window int
	}

	Aliexpress struct {
		AppCategory string
		AppKey      string
		AppSecret   string
		BaseURL     string
	}

	Gemini struct {
		APIKey string
	}

	FCM struct {
		Type                    string
		ProjectID               string
		PrivateKeyID            string
		PrivateKey              string
		ClientEmail             string
		ClientID                string
		AuthURI                 string
		TokenURI                string
		AuthProviderX509CertURL string
		ClientX509CertURL       string
		UniverseDomain          string
	}
}

func LoadConfig(path string) (Config, error) {
	// Load .env file
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			// Just log the error but don't fail - this is normal on Render
			log.Println("Note: .env file not found (this is expected in production)")
		}
	}

	var cfg Config

	// Server configuration
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// MongoDB configuration
	cfg.Mongo.URI = getEnv("MONGO_URI", "")
	cfg.Mongo.Database = getEnv("MONGO_DATABASE", "")
	cfg.Mongo.AlertCollection = getEnv("MONGO_ALERT_COLLECTION", "")

	// Redis configuration
	cfg.Redis.Host = getEnv("REDIS_HOST", "localhost")
	cfg.Redis.Port = getEnv("REDIS_PORT", "6379")
	cfg.Redis.Password = getEnv("REDIS_PASSWORD", "")
	cfg.Redis.DB = getEnvAsInt("REDIS_DB", 0)
	cfg.Redis.ViewTrackingTTL = getEnvAsInt("REDIS_VIEW_TRACKING_TTL", 86400)
	cfg.Redis.KeyPrefix = getEnv("REDIS_KEY_PREFIX", "sa:")

	// FX configuration
	cfg.FX.APIURL = getEnv("FX_API_URL", "")
	cfg.FX.APIKEY = getEnv("FX_API_KEY", "")
	cfg.FX.CacheTTLSeconds = getEnvAsInt("FX_CACHE_TTL_SECONDS", 3600)

	// OAuth configuration
	cfg.OAuth.Google.ClientID = getEnv("OAUTH_GOOGLE_CLIENT_ID", "")
	cfg.OAuth.Google.ClientSecret = getEnv("OAUTH_GOOGLE_CLIENT_SECRET", "")
	cfg.OAuth.Google.RedirectURI = getEnv("OAUTH_GOOGLE_REDIRECT_URI", "")
	cfg.OAuth.Aliexpress.ClientID = getEnv("OAUTH_ALIEXPRESS_CLIENT_ID", "")
	cfg.OAuth.Aliexpress.ClientSecret = getEnv("OAUTH_ALIEXPRESS_CLIENT_SECRET", "")
	cfg.OAuth.Aliexpress.RedirectURI = getEnv("OAUTH_ALIEXPRESS_REDIRECT_URI", "")

	// Rate limit configuration
	cfg.RateLimit.Limit = getEnvAsInt("RATE_LIMIT_LIMIT", 3)
	cfg.RateLimit.Window = getEnvAsInt("RATE_LIMIT_WINDOW", 60)

	// Aliexpress configuration
	cfg.Aliexpress.AppCategory = getEnv("ALIEXPRESS_APP_CATEGORY", "")
	cfg.Aliexpress.AppKey = getEnv("ALIEXPRESS_APP_KEY", "")
	cfg.Aliexpress.AppSecret = getEnv("ALIEXPRESS_APP_SECRET", "")
	cfg.Aliexpress.BaseURL = getEnv("ALIEXPRESS_BASE_URL", "")

	// Gemini configuration
	cfg.Gemini.APIKey = getEnv("GEMINI_API_KEY", "")

	// Firebase FCM Configuration
	cfg.FCM.Type = getEnv("FIREBASE_TYPE", "service_account")
	cfg.FCM.ProjectID = getEnv("FIREBASE_PROJECT_ID", "")
	cfg.FCM.PrivateKeyID = getEnv("FIREBASE_PRIVATE_KEY_ID", "")
	cfg.FCM.PrivateKey = getEnv("FIREBASE_PRIVATE_KEY", "")
	cfg.FCM.ClientEmail = getEnv("FIREBASE_CLIENT_EMAIL", "")
	cfg.FCM.ClientID = getEnv("FIREBASE_CLIENT_ID", "")
	cfg.FCM.AuthURI = getEnv("FIREBASE_AUTH_URI", "https://accounts.google.com/o/oauth2/auth")
	cfg.FCM.TokenURI = getEnv("FIREBASE_TOKEN_URI", "https://oauth2.googleapis.com/token")
	cfg.FCM.AuthProviderX509CertURL = getEnv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL", "https://www.googleapis.com/oauth2/v1/certs")
	cfg.FCM.ClientX509CertURL = getEnv("FIREBASE_CLIENT_X509_CERT_URL", "")
	cfg.FCM.UniverseDomain = getEnv("FIREBASE_UNIVERSE_DOMAIN", "googleapis.com")

	return cfg, nil
}

// Helper functions to get environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func (c *Config) BuildFirebaseCredentialsJSON() (string, error) {
	// Check if required fields are present
	if c.FCM.ProjectID == "" || c.FCM.PrivateKeyID == "" || c.FCM.PrivateKey == "" || c.FCM.ClientEmail == "" {
		return "", fmt.Errorf("missing required Firebase credentials")
	}

	// Replace escaped newlines with actual newlines in private key
	privateKey := strings.ReplaceAll(c.FCM.PrivateKey, "\\n", "\n")

	creds := map[string]interface{}{
		"type":                        c.FCM.Type,
		"project_id":                  c.FCM.ProjectID,
		"private_key_id":              c.FCM.PrivateKeyID,
		"private_key":                 privateKey,
		"client_email":                c.FCM.ClientEmail,
		"client_id":                   c.FCM.ClientID,
		"auth_uri":                    c.FCM.AuthURI,
		"token_uri":                   c.FCM.TokenURI,
		"auth_provider_x509_cert_url": c.FCM.AuthProviderX509CertURL,
		"client_x509_cert_url":        c.FCM.ClientX509CertURL,
		"universe_domain":             c.FCM.UniverseDomain,
	}

	jsonBytes, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal Firebase credentials: %w", err)
	}

	return string(jsonBytes), nil
}

// ValidateFirebaseCredentials validates Firebase configuration
func (c *Config) ValidateFirebaseCredentials() error {
	required := map[string]string{
		"FIREBASE_PROJECT_ID":           c.FCM.ProjectID,
		"FIREBASE_PRIVATE_KEY_ID":       c.FCM.PrivateKeyID,
		"FIREBASE_PRIVATE_KEY":          c.FCM.PrivateKey,
		"FIREBASE_CLIENT_EMAIL":         c.FCM.ClientEmail,
		"FIREBASE_CLIENT_X509_CERT_URL": c.FCM.ClientX509CertURL,
	}

	for name, value := range required {
		if value == "" {
			return fmt.Errorf("missing required Firebase configuration: %s", name)
		}
	}

	// Validate private key format
	if !strings.Contains(c.FCM.PrivateKey, "BEGIN PRIVATE KEY") {
		return fmt.Errorf("invalid private key format - should contain 'BEGIN PRIVATE KEY'")
	}

	if !strings.Contains(c.FCM.PrivateKey, "END PRIVATE KEY") {
		return fmt.Errorf("invalid private key format - should contain 'END PRIVATE KEY'")
	}

	return nil
}
