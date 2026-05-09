package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName         string `env:"APP_NAME" envDefault:"apart_community"`
	GinMode         string `env:"GIN_MODE" envDefault:"debug"`
	DBUser          string `env:"DB_USER"`
	DBPass          string `env:"DB_PASS"`
	DBHost          string `env:"DB_HOST"`
	DBPort          string `env:"DB_PORT"`
	DBName          string `env:"DB_NAME"`
	MaxIdleConns    int    `env:"MAX_IDLE_CONNS" envDefault:"10"`
	MaxOpenConns    int    `env:"MAX_OPEN_CONNS" envDefault:"100"`
	ConnMaxLifetime int    `env:"CONN_MAX_LIFETIME" envDefault:"2"`
	RedisHost       string `env:"REDIS_HOST"`
	RedisPort       string `env:"REDIS_PORT"`
	JwtSecret       string `env:"JWT_SECRET"`
}

func LoadConfig(environment *string) *Config {
	_ = godotenv.Load("../.env." + *environment)
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println(err.Error())
		log.Fatalln("Error loading .env file")
	}
	return &cfg
}

func isValidEnvironment(environment *string) bool {
	switch *environment {
	case "dev",
		"production":
		return true
	}
	return false
}

func LoadEnv() *Config {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	if !isValidEnvironment(environment) {
		log.Fatalf("Invalid environment: %s\n", *environment)
	}

	config := LoadConfig(environment)

	fmt.Println("env loaded. APP_NAME: ", config.AppName, "")

	return config
}
