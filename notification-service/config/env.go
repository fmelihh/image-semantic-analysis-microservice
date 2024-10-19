package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	splittedPath := strings.Split(pwd, "/notification-service")
	envPath := fmt.Sprintf("%s/%s", splittedPath[0], "notification-service/.env")
	err = godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}
	return Config{
		SMTPHost:  getEnv("SMTP_HOST", ""),
		SMTPPort:  int(getEnvAsInt("SMTP_PORT", 0)),
		SMTPLogin: getEnv("SMTP_LOGIN", ""),
		SMTPToken: getEnv("SMTP_TOKEN", ""),
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
