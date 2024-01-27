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

	"github.com/elliot40404/icache-echo/internal/models"
	"github.com/elliot40404/icache-echo/internal/utils"
	"github.com/h2non/bimg"
)

type ImageService struct {
	Height uint
	Width  uint
	Cache  *models.ImageCache
}

type ImageDownload struct {
	Img    *bytes.Buffer
	Ctype  string
	Size   int64
	Cached bool
}

func NewImageService(width, height uint) ImageService {
	return ImageService{
		Height: height,
		Width:  width,
		Cache:  models.NewImageCache(),
	}
}

func downloadImage(imageURL string) (*ImageDownload, error) {
	slog.Info("Downloading image", "URL", imageURL)
	// Parse the URL
	parsedURL, err := url.Parse(imageURL)
	if err != nil {
		fmt.Println("url.Parse error:", err)
		return nil, err
	}

	// Make an HTTP GET request with redirects followed
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		fmt.Println("http.Get error:", err)
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

func (s *ImageService) GetImage(imageURL string) (*ImageDownload, error) {
	// Check if the image is in the cache
	if s.Cache.HAS(imageURL) {
		img, ok := s.Cache.GET(imageURL)
		if ok {
			return &ImageDownload{
				Img:    &img.Img,
				Ctype:  img.Ctype,
				Size:   img.Size,
				Cached: true,
			}, nil
		}
	}

	// Download the image
	downloadedImg, err := downloadImage(imageURL)
	if err != nil {
		return nil, err
	}

	// resize the image in the background
	s.ResizeImage(downloadedImg, int(s.Width), int(s.Height))

	// Save the image in the cache
	s.Cache.SET(imageURL, models.CachedImage{
		Img:   *downloadedImg.Img,
		Ctype: downloadedImg.Ctype,
		Size:  downloadedImg.Size,
	})

	return downloadedImg, nil
}

func (s *ImageService) ResizeImage(img *ImageDownload, width, height int) {
	slog.Info("Original Image Details", "size", utils.BytesTo(utils.KB, (uint(img.Size))), "content-type", img.Ctype)

	// convert the img to an image.Image
	newImg := bimg.NewImage(img.Img.Bytes())

	// resize the image
	resizedImg, err := newImg.Resize(width, height)
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
