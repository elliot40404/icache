package server

import (
	"log"
	"os"
	"strconv"

	"github.com/elliot40404/icache-echo/internal/routes"
	// Load env variables from .env file.
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port int
	addr string
}

func (s Server) listenAddr() string {
	return s.addr + ":" + strconv.Itoa(s.port)
}

func Run() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("Error parsing port from env, using default port 8080")
		port = 8080
	}
	addr := os.Getenv("ADDR")
	NewServer := &Server{
		port: port,
		addr: addr,
	}
	e := echo.New()
	e.HideBanner = true
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

	e.Logger.Fatal(e.Start(NewServer.listenAddr()))
}
