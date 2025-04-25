package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Key string
}

func (c *Config) GetKey() string {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	key := os.Getenv("KEY")

	c.Key = key

	return c.Key
}
