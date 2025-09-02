package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config struct remains exactly as you provided.
// I've added the missing `mapstructure:"fx"` tag for consistency,
// but otherwise, it's untouched to meet your requirements.
type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`

	Mongo struct {
		URI             string `mapstructure:"uri"`
		Database        string `mapstructure:"database"`
		AlertCollection string `mapstructure:"alert_collection"`
	} `mapstructure:"mongo"`

	Redis struct {
		Host            string `mapstructure:"host"`
		Port            string `mapstructure:"port"`
		Password        string `mapstructure:"password"`
		DB              int    `mapstructure:"db"`
		ViewTrackingTTL int    `mapstructure:"view_tracking_ttl"`
		KeyPrefix       string `mapstructure:"key_prefix"`
	} `mapstructure:"redis"`

	FX struct {
		APIURL          string `mapstructure:"api_url"`
		APIKEY          string `mapstructure:"api_key"`
		CacheTTLSeconds int    `mapstructure:"cache_ttl_seconds"`
	} `mapstructure:"fx"` // Added this missing mapstructure tag for consistency

	OAuth struct {
		Google struct {
			ClientID     string `mapstructure:"client_id"`
			ClientSecret string `mapstructure:"client_secret"`
			RedirectURI  string `mapstructure:"redirect_uri"`
		} `mapstructure:"google"`

		Aliexpress struct {
			ClientID     string `mapstructure:"client_id"`
			ClientSecret string `mapstructure:"client_secret"`
			RedirectURI  string `mapstructure:"redirect_uri"`
		} `mapstructure:"aliexpress"`
	} `mapstructure:"oauth"`

	RateLimit struct {
		Limit  int `mapstructure:"limit"`
		Window int `mapstructure:"window"`
	} `mapstructure:"rate-limit"`

	Aliexpress struct {
		AppCategory string `mapstructure:"app_category"`
		AppKey      string `mapstructure:"app_key"`
		AppSecret   string `mapstructure:"app_secret"`
		BaseURL     string `mapstructure:"base_url"` // Assuming this field exists or will be added
	} `mapstructure:"aliexpress"`

	Gemini struct {
		APIKey string `mapstructure:"api_key"`
	} `mapstructure:"gemini"`
}

func LoadConfig(path string) (*Config, error) {
	// Configure Viper to look for environment variables.
	// This replacer converts internal Viper keys (e.g., "rate-limit.window")
	// into the format expected for environment variables (e.g., "RATE_LIMIT_WINDOW").
	// It replaces '.' with '_' and '-' with '_' to match common UPPER_SNAKE_CASE environment variable naming.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv() // Automatically bind environment variables to config keys

	// --- Optional: Keep YAML for default values, but environment variables will override ---
	// If you want to completely remove the YAML file and only use environment variables,
	// you can comment out or remove the following block.
	// The `viper.ReadInConfig()` function will attempt to read the specified YAML file.
	// If it's not found, we handle it as a non-fatal warning, allowing the application
	// to proceed using only environment variables. This is useful for providing defaults
	// in YAML that can be overridden by secrets in environment variables.
	viper.SetConfigName("config.dev") // Name of your config file (e.g., "config.dev", "config.prod")
	viper.SetConfigType("yaml")       // Type of the config file
	viper.AddConfigPath(path)         // Path where Viper should look for the config file

	// Attempt to read the config file. If it's not found, it's not an error
	// because environment variables are now the primary source of truth for secrets.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, log a warning and proceed, relying on environment variables.
			log.Printf("Warning: Config file 'config.dev.yaml' not found in path '%s'. Relying solely on environment variables.\n", path)
		} else {
			// Some other error occurred while reading the config file (e.g., parsing error).
			// This is a fatal error.
			return nil, err
		}
	}
	// --- End Optional YAML loading ---

	var cfg Config
	// Unmarshal the configuration into your Config struct.
	// Environment variables will automatically override any values read from the YAML file.
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}