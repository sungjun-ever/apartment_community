package main

import (
	"apart_community/config"
	"apart_community/database"
	"apart_community/internals/common/errUtils"
	"apart_community/registry"
	"apart_community/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	flags := config.ParseFlags()
	env := config.LoadEnv(&flags)

	db := database.ConnectToPostgres(env)
	rdb := database.ConnectToRedis(env)

	em := errUtils.NewErrorManager("../../logs/error.log")

	container := registry.NewContainer(db, rdb)

	r := gin.Default()

	r = router.SetUpRouter(r, em, container)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("<- 서버 종료 신호")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("서버 오류로 강제 종료: %v\n", err)
	}

	log.Println("서버 종료 완료")
}
