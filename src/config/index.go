package main

import (
	"os"
)

type Config struct {
	APIPath      string
	Version      string
	GeneratedTime string
}

func main() {
	config := Config{
		APIPath:      os.Getenv("API_PATH"),
		Version:      os.Getenv("VERSION"),
		GeneratedTime: os.Getenv("GENERATED_TIME"),
	}
}