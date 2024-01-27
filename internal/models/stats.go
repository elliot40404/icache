package models

type ServerStats struct {
	TotalAlloc uint `json:"total_alloc"`
	Alloc      uint `json:"alloc"`
	Sys        uint `json:"sys"`
	NumGC      uint `json:"num_gc"`
}
