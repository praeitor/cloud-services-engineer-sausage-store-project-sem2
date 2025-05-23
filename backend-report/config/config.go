package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Configuration contains static info required to run the apps
// It contains DB info
type Configuration struct {
	Port                  string `env:"PORT" envDefault:"8080"`
	DataBaseConnectionURL string `env:"DB,required"`
}

// NewConfig will read the config data from given .env file
func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...) // Loading config from env file

	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}

	cfg := Configuration{}

	// Parse env to configuration
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
