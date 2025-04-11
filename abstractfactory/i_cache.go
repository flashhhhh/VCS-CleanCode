package main

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type ICache interface {
	SetUpConnection()
	SetQuery(key string, value string) (string, error)
	GetQuery(query string) (string, error)
}

type Cache struct {
	host	 string
	port	 int
	client *redis.Client
}

func (c *Cache) SetUpConnection() {
	// Set up Redis connection
	c.client = redis.NewClient(&redis.Options{
		Addr:     c.host + ":" + strconv.Itoa(c.port),
	})
}

func (c *Cache) SetQuery(key string, value string) (string, error) {
	err := c.client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return "", err
	}
	return "OK", nil
}

func (c *Cache) GetQuery(query string) (string, error) {
	val, err := c.client.Get(context.Background(), query).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}