package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Browser struct {
		Headless bool `yaml:"headless"`
	} `yaml:"browser"`

	Limits struct {
		DailyConnections int `yaml:"daily_connections"`
	} `yaml:"limits"`

	Timing struct {
		MinDelayMs int `yaml:"min_delay_ms"`
		MaxDelayMs int `yaml:"max_delay_ms"`
	} `yaml:"timing"`
}

func Load(path string) *Config {
	cfg := &Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	if err := yaml.Unmarshal(file, cfg); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	return cfg
}
