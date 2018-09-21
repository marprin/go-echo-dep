package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve any static file
	e.Static("/static", "assets")

	// Serve robots.txt file
	e.File("/robots.txt", "assets/robots.txt")

	// Customize Logger
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	e.GET("/", func(c echo.Context) error { 
		return c.String(http.StatusOK, "Hello, world")
	})

	// Start Server
	e.Logger.Fatal(e.Start(":7600"))
}