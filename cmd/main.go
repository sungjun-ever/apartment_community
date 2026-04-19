package main

import (
	"apart_community/config"
	"apart_community/internal/database"
	"apart_community/internal/utils"
	"apart_community/registry"
	"apart_community/router"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.LoadEnv()

	db := database.ConnectMySQLDB()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	container := registry.NewContainer(db)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("password_check", utils.ValidatePassword)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	router.SetupRouter(r, container)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("서버 종료 중")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("서버 종료: ", err)
	}

	log.Println("서버 종료 완료")
}
