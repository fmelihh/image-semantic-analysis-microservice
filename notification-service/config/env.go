package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SMTPHost  string
	SMTPPort  int
	SMTPLogin string
	SMTPToken string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		SMTPHost:  getEnv("SMTP_Host", ""),
		SMTPPort:  int(getEnvAsInt("SMTP_Port", 0)),
		SMTPLogin: getEnv("SMTP_Login", ""),
		SMTPToken: getEnv("SMTP_Token", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
