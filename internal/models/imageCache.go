package models

import (
	"bytes"
)

type CachedImage struct {
	Img   bytes.Buffer
	Ctype string
	Size  int64
}

// TODO: use sync map
type ImageCache map[string]CachedImage

func NewImageCache() *ImageCache {
	var cache ImageCache = make(map[string]CachedImage)
	return &cache
}

func (c *ImageCache) GET(key string) (CachedImage, bool) {
	value, ok := (*c)[key]
	return value, ok
}

func (c *ImageCache) SET(key string, value CachedImage) {
	(*c)[key] = value
}

func (c *ImageCache) HAS(key string) bool {
	_, ok := (*c)[key]
	return ok
}

func (c *ImageCache) DELETE(key string) {
	delete(*c, key)
}
