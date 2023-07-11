package config

import (
	"log"
	"strings"

	_ "github.com/joho/godotenv/autoload" // secret variable replacement for local test
	"github.com/spf13/viper"
)

type Config struct {
	Gin struct {
		Mode string // enum: debug, release, test
	}
	Server struct {
		Port               string
		ShutdownTimeoutSec int // 0 would shut down immediately
	}
}

func New() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file:", err)
	}

	cfg := new(Config)
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Failed to unmarshal config:", err)
	}

	return cfg
}
