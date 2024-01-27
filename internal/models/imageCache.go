package models

import (
	"bytes"
	"sync"
)

type CachedImage struct {
	Img   bytes.Buffer
	Ctype string
	Size  int64
}

type ImageCache struct {
	store sync.Map
}

func NewImageCache() *ImageCache {
	return &ImageCache{}
}

func (c *ImageCache) GET(key string) (CachedImage, bool) {
	value, ok := c.store.Load(key)
	if !ok {
		return CachedImage{}, false
	}

	return value.(CachedImage), true
}

func (c *ImageCache) SET(key string, value CachedImage) {
	c.store.Store(key, value)
}

func (c *ImageCache) HAS(key string) bool {
	_, ok := c.store.Load(key)
	return ok
}

func (c *ImageCache) DELETE(key string) {
	c.store.Delete(key)
}