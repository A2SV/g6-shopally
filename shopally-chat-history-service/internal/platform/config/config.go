package config

import (
	"os"
	"time"
)

// Config holds application configuration, sourced from environment variables.
type Config struct {
	Port            string
	MongoURI        string
	MongoDBName     string
	MongoCollection string
	ShutdownTimeout time.Duration
}

func FromEnv() Config {
	port := getEnv("CHAT_HISTORY_PORT", "8080")
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	dbName := getEnv("CHAT_HISTORY_MONGO_DB", "shopally_chat_db")
	collection := getEnv("CHAT_HISTORY_MONGO_COLLECTION", "chat_history")

	shutdownSecs := getEnv("CHAT_HISTORY_SHUTDOWN_TIMEOUT", "10")
	timeout, err := time.ParseDuration(shutdownSecs + "s")
	if err != nil {
		timeout = 10 * time.Second
	}

	return Config{
		Port:            port,
		MongoURI:        mongoURI,
		MongoDBName:     dbName,
		MongoCollection: collection,
		ShutdownTimeout: timeout,
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
