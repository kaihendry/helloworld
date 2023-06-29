package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/gateway/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"golang.org/x/exp/slog"
	log "golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		logger = log.New(log.NewJSONHandler(os.Stdout, nil))
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// simulate an error
	e.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError, "A simulated error")
	})

	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		log.SetDefault(log.New(log.NewJSONHandler(os.Stdout, nil)))
		gateway.ListenAndServe("", e)
	} else {
		log.SetDefault(log.New(log.NewTextHandler(os.Stdout, nil)))
		log.Info("local development", "port", os.Getenv("PORT"))
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	}

}
