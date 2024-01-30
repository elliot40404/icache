package models

import (
	"bytes"
	"sync"
)

// TODO: Cache invalidation.
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

func (c *ImageCache) Get(key string) (CachedImage, bool) {
	value, ok := c.store.Load(key)
	if !ok {
		return CachedImage{}, false
	}

	return value.(CachedImage), true
}

func (c *ImageCache) Set(key string, value CachedImage) {
	c.store.Store(key, value)
}

func (c *ImageCache) Has(key string) bool {
	_, ok := c.store.Load(key)
	return ok
}

func (c *ImageCache) Delete(key string) {
	c.store.Delete(key)
}

func (c *ImageCache) GetAllKeys() []string {
	keys := make([]string, 0)
	c.store.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

func (c *ImageCache) Flush() {
	c.store.Range(func(key, value interface{}) bool {
		c.store.Delete(key)
		return true
	})
}
