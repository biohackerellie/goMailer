package env

import (
	"os"
)

type AppConfig struct {
	User           string `json:"user"`
	Pass           string `json:"pass"`
	APIKey         string `json:"api_key"`
	AllowedDomains string `json:"allowed_domains"`
	OAUTH          string `json:"oauth"`
}

const (
	User           = "EMAIL_USER"
	Pass           = "EMAIL_PASSWORD"
	APIKey         = "API_KEY"
	AllowedDomains = "ALLOWED_DOMAINS"
	OAUTH          = "OAUTH"
)

func GetEnv() *AppConfig {
	var config AppConfig
	if val, ok := os.LookupEnv(User); ok {
		config.User = val
	}
	if val, ok := os.LookupEnv(Pass); ok {
		config.Pass = val
	}
	if val, ok := os.LookupEnv(APIKey); ok {
		config.APIKey = val
	}
	if val, ok := os.LookupEnv(AllowedDomains); ok {
		config.AllowedDomains = val
	}
	if val, ok := os.LookupEnv(OAUTH); ok {
		config.OAUTH = val
	}
	return &config
}

func GetConfig() *AppConfig {
	return GetEnv()
}
