package db

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Ssl      string `yaml:"sslmode"`
}

func LoadConfig(filePath string) (DBConfig, error) {
	var config DBConfig
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ямлю не нашли: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Ямля косячная: %v", err)
	}

	return config, nil
}
