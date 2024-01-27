package routes

import (
	"log"

	h "github.com/elliot40404/icache-echo/internal/handlers"
	"github.com/labstack/echo/v4"
)

type ImgRouter struct{}

func NewImgRouter() *ImgRouter {
	return &ImgRouter{}
}

func (t *ImgRouter) RegisterRoutes(e *echo.Group) {
	log.Println("Registering routes for tasks")

	handler := h.NewImgHandler()

	e.GET("/image", handler.GetImage)

	e.GET("/stats", handler.ServerStats)
}
