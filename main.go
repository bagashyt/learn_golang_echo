package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/cors"
)

func main() {
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https:/novalagung.com", "https://www.google.com"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-type", "X-CSRF-Token"},
		Debug:          true,
	})

	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/index", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
