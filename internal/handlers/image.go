package handlers

import (
	"net/http"
	"runtime"

	"github.com/elliot40404/icache-echo/internal/models"
	"github.com/elliot40404/icache-echo/internal/services"
	"github.com/elliot40404/icache-echo/internal/utils"
	"github.com/labstack/echo/v4"
)

type ImgHandler struct {
	ImgService services.ImageService
}

func NewImgHandler() *ImgHandler {
	return &ImgHandler{
		ImgService: services.NewImageService(300, 300),
	}
}

func (h *ImgHandler) ServerStats(c echo.Context) error {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	stats := models.ServerStats{
		TotalAlloc: utils.BytesTo(utils.MB, uint(mem.TotalAlloc)),
		Alloc:      utils.BytesTo(utils.MB, uint(mem.Alloc)),
		Sys:        utils.BytesTo(utils.MB, uint(mem.Sys)),
		NumGC:      uint(mem.NumGC),
	}
	return c.JSON(http.StatusOK, stats)
}

func (h *ImgHandler) GetImage(c echo.Context) error {
	imgURL := c.QueryParam("url")
	if imgURL == "" {
		return c.JSON(http.StatusBadRequest, "url query param is required")
	}
	width, height, err := utils.GetWidthHeight(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	isWebP := c.QueryParam("format") == "webp"
	bypassCache := c.QueryParam("bypass") == "true"
	resize := c.QueryParam("resize") != "false"

	img, err := h.ImgService.GetImage(imgURL, width, height, isWebP, resize, bypassCache)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	utils.AddCacheHeaders(&c, img.Cached)
	return c.Blob(http.StatusOK, img.Ctype, img.Img.Bytes())
}

func (h *ImgHandler) GetImages(c echo.Context) error {
	images := h.ImgService.GetImages()
	return c.JSON(http.StatusOK, images)
}

func (h *ImgHandler) DeleteImage(c echo.Context) error {
	imgURL := c.QueryParam("url")
	if imgURL == "" {
		return c.JSON(http.StatusBadRequest, "url query param is required")
	}
	err := h.ImgService.DeleteImage(imgURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Image deleted")
}

func (h *ImgHandler) DeleteImages(c echo.Context) error {
	err := h.ImgService.DeleteImages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Images deleted")
}
