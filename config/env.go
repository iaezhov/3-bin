package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	ApiKey      string
	ApiUrl      string
	StoragePath string
}

func (c *Envs) GetApiKey() string {
	return c.ApiKey
}

func (c *Envs) GetApiUrl() string {
	return c.ApiUrl
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func NewEnvs() (*Envs, error) {
	godotenv.Load()

	return &Envs{
		ApiKey:      getEnv("API_KEY", ""),
		ApiUrl:      getEnv("API_URL", "https://api.jsonbin.io/v3/b"),
		StoragePath: getEnv("STORAGE_PATH", "bin-list.json"),
	}, nil
}
