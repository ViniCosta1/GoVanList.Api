package config

import (
	"os"
)

type Config struct {
	DataDir  string
	FileName string
}

func LoadConfig() Config {
	return Config{
		DataDir: getEnv("DATA_DIR", "C:\\Projects\\GoVanList.Api\\data\\"),
		FileName: getEnv("FILE_NAME", "response.json"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}