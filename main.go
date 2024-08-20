package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	e.Use(middlewareOne)
	e.Use(middlewareTwo)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middlewareLogging)
	e.HTTPErrorHandler = errorHandler

	e.GET("/index", func(ctx echo.Context) (err error) {
		fmt.Println("threee!")
		return ctx.JSON(http.StatusOK, true)
	})

	lock := make(chan error)
	go func(lock chan error) {
		lock <- e.Start(":9000")
	}(lock)

	time.Sleep(1 * time.Millisecond)
	makeLogEntry(nil).Warning("application starting without ssl/tls enabled")

	err := <-lock
	if err != nil {
		makeLogEntry(nil).Panic("failed to start application")
	}

}

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("from middleware one")
		return next(ctx)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("from middleware two")
		return next(ctx)
	}
}

func makeLogEntry(ctx echo.Context) *log.Entry {
	if ctx == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": ctx.Request().Method,
		"uri":    ctx.Request().URL.String(),
		"ip":     ctx.Request().RemoteAddr,
	})
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		makeLogEntry(ctx).Info("incoming request")
		return next(ctx)
	}
}

func errorHandler(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(ctx).Error(report.Message)
	ctx.HTML(report.Code, report.Message.(string))
}
