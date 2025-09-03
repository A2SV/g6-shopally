package config

import (
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
}

func LoadConfig(path string) (*Config, error) {
	// Load .env file
	if err := godotenv.Load(path + "/.env"); err != nil {
		return nil, err
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

	return &cfg, nil
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

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	return strings.Split(valueStr, sep)
}