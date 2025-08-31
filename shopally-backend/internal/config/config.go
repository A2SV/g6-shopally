package config

import (
	"github.com/spf13/viper"
)

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
	}

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
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("configs/config.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
