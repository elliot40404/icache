package services

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/elliot40404/icache-echo/internal/utils"
	"github.com/h2non/bimg"
)

func (s ImageService) DownloadImage(imageURL string) (*ImageDownload, error) {
	slog.Info("Downloading image", "URL", imageURL)
	// Parse the URL
	parsedURL, err := url.Parse(imageURL)
	if err != nil {
		fmt.Println("url.Parse error:", err)
		return nil, err
	}

	// Make an HTTP GET request with redirects followed
	// resp, err := http.Get(parsedURL.String())
	// if err != nil {
	// 	fmt.Println("http.Get error:", err)
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	// Create a new HTTP client with a default timeout and context
	resp, err := http.DefaultClient.Do(&http.Request{
		Method: http.MethodGet,
		URL:    parsedURL,
	})
	if err != nil {
		fmt.Println("http.DefaultClient.Do error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check for a successful response status
	if resp.StatusCode != http.StatusOK {
		slog.Info("response status code is", "code", resp.StatusCode)
		return nil, errors.New(resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "image") {
		slog.Info("Content-Type is not image", "content-type", contentType)
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &ImageDownload{
		Img:    buf,
		Ctype:  contentType,
		Size:   resp.ContentLength,
		Cached: false,
	}, nil
}

func (s ImageService) ResizeImage(img *ImageDownload, width, height int) {
	slog.Info("Original Image Details", "size", utils.BytesTo(utils.KB, (uint(img.Size))), "content-type", img.Ctype)

	// convert the img to an image.Image
	newImg := bimg.NewImage(img.Img.Bytes())

	// calculate the new image size based on aspect ratio
	size, err := newImg.Size()
	if err != nil {
		slog.Info("Cannot get image size", "error", err.Error())
	}
	owidth := size.Width
	oheight := size.Height

	slog.Info("Original image dimensions", "width", owidth, "height", oheight)

	newWidth, newHeight := utils.CalcNewImageSize(owidth, oheight, width, height)

	slog.Info("Resized image dimensions", "width", newWidth, "height", newHeight)

	// resize the image
	resizedImg, err := newImg.Resize(newWidth, newHeight)
	if err != nil {
		slog.Info("Cannot resize image", "error", err.Error())
		return
	}

	// convert the resized image to a buffer
	buf := bytes.NewBuffer(resizedImg)

	// update the image buffer
	img.Img = buf
	img.Size = int64(buf.Len())

	slog.Info("Resized image details", "size", utils.BytesTo(utils.KB, (uint(img.Size))), "content-type", img.Ctype)
}

func (s ImageService) ConvertWebP(img *ImageDownload) {
	slog.Info("Converting image to webp", "size", utils.BytesTo(utils.KB, (uint(img.Size))), "content-type", img.Ctype)

	// convert the img to an image.Image
	newImg := bimg.NewImage(img.Img.Bytes())

	// convert the image to webp
	convertedImg, err := newImg.Convert(bimg.WEBP)
	if err != nil {
		slog.Info("Cannot convert image to webp", "error", err.Error())
		return
	}

	// convert the resized image to a buffer
	buf := bytes.NewBuffer(convertedImg)

	// update the image buffer
	img.Img = buf
	img.Size = int64(buf.Len())
	img.Ctype = "image/webp"

	slog.Info("Converted image details", "size", utils.BytesTo(utils.KB, (uint(img.Size))), "content-type", img.Ctype)
}
