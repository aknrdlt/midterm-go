package db

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

func DB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

func (db *Layer) Get(key string) (string, error) {
	db.M.Lock()
	defer db.M.Unlock()
	val, err := db.DB.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", errors.New("key does not exist")
	} else if err != nil {
		return "", err
	} else {
		return val, nil
	}
}
func (db *Layer) Put(key, value string) error {
	db.M.Lock()
	defer db.M.Unlock()
	if err := db.DB.Set(context.Background(), key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}
