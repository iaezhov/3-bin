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

func NewEnvs() (*Envs, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Envs{
		ApiKey:      os.Getenv("API_KEY"),
		ApiUrl:      "https://api.jsonbin.io/v3/b",
		StoragePath: "bin-list.json",
	}, nil
}
