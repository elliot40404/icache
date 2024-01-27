package server

import (
	"os"
	"strconv"

	"github.com/elliot40404/icache-echo/internal/routes"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port  int
}

func Run() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:  port,
	}
	e := echo.New()

	// Global middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	v1Grp := e.Group("/api/v1")
	routes.NewImgRouter().RegisterRoutes(v1Grp)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(NewServer.port)))
}