package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/pangami/auth-service/repo/mongo"
	// "github.com/pangami/auth-service/transport/rest/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failure loading .env file")
	}

	// init repo
	// mongo.InitCon()

	e := echo.New()

	// Example middleware setup
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "user-service here!")
	})

	// Register routes
	// routes.Routes(e)

	// Start the server
	// rest_port := fmt.Sprintf(":%s", os.Getenv("REST_PORT"))

	// e.Logger.Fatal(e.Start(rest_port))
}
