package main

import (
	"fmt"

	"github.com/Medveddo/team-booster/backend/internal/api"
	"github.com/Medveddo/team-booster/backend/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "net/http"
)

func main()  {
	fmt.Println("Hello, world!")
	api.ApiFunc()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
  
	// Routes
	e.GET("/", hello)
  
	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(200, []models.Skill{models.RedisSkill})
	// return c.String(http.StatusOK, "Hello, World!")
  }