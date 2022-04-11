package db

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

type Layer struct {
	DB *redis.Client
	M  sync.Mutex
}

func NewDataBaseLayers(db *redis.Client) *Layer {
	return &Layer{
		DB: db,
	}
}
