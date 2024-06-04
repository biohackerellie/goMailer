package env

import (
	"log"

	"github.com/joho/godotenv"
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

func readEnvFile(filePath string) map[string]string {
	envMap, err := godotenv.Read(filePath)
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	return envMap
}

func GetEnv(envMap map[string]string) *AppConfig {
	var config AppConfig
	if val, ok := envMap[User]; ok {
		config.User = val
	}
	if val, ok := envMap[Pass]; ok {
		config.Pass = val
	}
	if val, ok := envMap[APIKey]; ok {
		config.APIKey = val
	}
	if val, ok := envMap[AllowedDomains]; ok {
		config.AllowedDomains = val
	}
	if val, ok := envMap[OAUTH]; ok {
		config.OAUTH = val
	}
	return &config
}

func GetConfig() *AppConfig {
	return GetEnv(readEnvFile(".env"))
}
