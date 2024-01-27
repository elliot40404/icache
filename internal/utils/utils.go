package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

type Data int

const (
	KB Data = iota + 1
	MB
	GB
)

func BytesTo(data Data, b uint) uint {
	switch data {
	case KB:
		return b / 1024
	case MB:
		return b / 1024 / 1024
	case GB:
		return b / 1024 / 1024 / 1024
	default:
		return b
	}
}	

func WriteImageToDisk(img *bytes.Buffer) {
	// write the image to file
	f, err := os.Create("test/wi.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// copy the image bytes to the file
	_, err = io.Copy(f, img)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func AddCacheHeaders(c *echo.Context, cached bool) {
	if cached {
		(*c).Response().Header().Set("X-Cache", "HIT")
	} else {
		(*c).Response().Header().Set("X-Cache", "MISS")
	}
}