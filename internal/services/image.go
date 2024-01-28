package services

import (
	"bytes"
	"log/slog"

	"github.com/elliot40404/icache-echo/internal/models"
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

func (s *ImageService) GetImage(imageURL string, width, height uint, webp, optimized bool) (*ImageDownload, error) {
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
	downloadedImg, err := s.DownloadImage(imageURL)
	if err != nil {
		return nil, err
	}

	// add default values for width and height
	if optimized {
		slog.Info("Applying default optimizations")
		if width == 0 {
			width = s.Width
		}
		if height == 0 {
			height = s.Height
		}
	}

	// resize the image in the background
	if width != 0 && height != 0 {
		s.ResizeImage(downloadedImg, int(width), int(height))
	}

	// convert the image to webp
	if webp {
		s.ConvertWebP(downloadedImg)
	}

	// Save the image in the cache
	s.Cache.SET(imageURL, models.CachedImage{
		Img:   *downloadedImg.Img,
		Ctype: downloadedImg.Ctype,
		Size:  downloadedImg.Size,
	})

	return downloadedImg, nil
}

func (s *ImageService) GetImages() []string {
	return s.Cache.GET_ALL_KEYS()
}

func (s *ImageService) DeleteImage(imageURL string) error {
	s.Cache.DELETE(imageURL)
	return nil
}

func (s *ImageService) DeleteImages() error {
	s.Cache.FLUSH()
	return nil
}
