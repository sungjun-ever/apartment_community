package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	wd, _ := os.Getwd()

	curr := wd

	for i := 0; i < 5; i++ {
		envPath := filepath.Join(curr, ".env")
		if _, err := os.Stat(envPath); err == nil {
			err := godotenv.Load(envPath)
			if err != nil {
				log.Fatalf(".env 로딩 실패: %v", err)
			}

			log.Printf(".env loaded path: %s", envPath)
			return
		}
		curr = filepath.Dir(curr)
	}
}
