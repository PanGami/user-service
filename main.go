package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/pangami/user-service/repo/mysql"
	"github.com/pangami/user-service/repo/redis"

	// "github.com/pangami/auth-service/repo/mongo"
	"github.com/pangami/user-service/transport/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failure loading .env file")
	}

	// init repo
	mysql.InitCon()
	redis.InitCon()
	// mongo.InitCon()

	e := echo.New()

	// Example middleware setup
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "user-service here! You shouldn't be here. Port not open :)")
	})

	// Run gRPC server in a separate goroutine
	go func() {
		grpc.Run()
	}()

	// Keep the main function running
	select {}
}
