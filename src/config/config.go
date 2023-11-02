package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DB DB `yaml:"db"`
}

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

func Load() (*Config, error) {
	config := Config{}

	b, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	fmt.Println(os.Getenv("CONFIG_PATH"))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}