package models

type CACHE[T any] interface {
	GET(key string) (T, bool)
	SET(key string, value T)
	HAS(key string) bool
	DELETE(key string)
	GET_ALL_KEYS() []string
	FLUSH()
}
