package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Medveddo/team-booster/backend/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	// "net/http"
)

func main() {
	api := api.NewAPI()

	// Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/health", api.Health)

	v1 := e.Group("/api/v1")
	{
		skills := v1.Group("/skills")
		{
			skills.GET("", api.GetSkills)
			skills.GET("/:id", api.GetSkill)
			skills.PUT("/:id", api.UpdateSkill)
		}
	}

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Graceful Shutdown
	// https://echo.labstack.com/cookbook/graceful-shutdown/
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := api.Close()
	if err != nil {
		e.Logger.Error("Error while closing API", err)
	}
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
