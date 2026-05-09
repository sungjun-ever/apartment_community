package config

import (
	"flag"
	"fmt"
	"log"

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

type Flag struct {
	Mode string
	Seed bool
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

func ParseFlags() *Flag {
	environment := flag.String("e", "dev", "")
	seedFlag := flag.Bool("seed", false, "seed table name")
	flag.Parse()

	return &Flag{
		Mode: *environment,
		Seed: *seedFlag,
	}
}

func LoadEnv(parsedFlag *Flag) *Config {
	environment := &parsedFlag.Mode

	if !isValidEnvironment(environment) {
		log.Fatalf("Invalid environment: %s\n", *environment)
	}

	config := LoadConfig(environment)

	fmt.Println("env loaded. APP_NAME: ", config.AppName, "")

	return config
}
