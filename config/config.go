package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func isValidEnvironment(environment *string) bool {
	switch *environment {
	case "dev",
		"prod":
		return true
	}
	return false
}

func LoadEnv() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	if !isValidEnvironment(environment) {
		log.Fatalf("Invalid environment: %s\n", *environment)
	}

	err := godotenv.Load("../.env." + *environment)

	if err != nil {
		log.Println(err.Error())
		log.Fatalln("Error loading .env file")
	}

	fmt.Println("env loaded. APP_NAME: ", os.Getenv("APP_NAME"), "")
}
