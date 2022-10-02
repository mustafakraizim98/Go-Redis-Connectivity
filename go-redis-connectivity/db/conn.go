package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Database struct {
	Client *redis.Client
}

func NewClient(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("%s: %s", context.TODO(), "no matching record found in redis database")
	}

	return &Database{
		Client: client,
	}, nil
}

func Set(databse *Database, key string, value interface{}, expiration time.Duration) {
	err := databse.Client.Set(key, value, expiration).Err()
	Must(err)
}

func Get(databse *Database, key string) string {
	value, err := databse.Client.Get(key).Result()
	Must(err)
	return value
}

func Del(databse *Database, key string) {
	err := databse.Client.Del(key).Err()
	Must(err)
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
