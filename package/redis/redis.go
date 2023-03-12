package redis

import (
	"context"
	"fmt"

	redisClient "github.com/go-redis/redis/v8"
)

var client *redisClient.Client

func init() {
	client = &redisClient.Client{}
}

type Config struct {
	Address  string
	Password string
	Db       int
}

func CreateClient(c Config) error {
	client = redisClient.NewClient(&redisClient.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.Db,
	})

	res, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println("There is error while initiating redis client")
		return err
	}

	fmt.Println(fmt.Sprintf("redis client is initialized successfully res:= %v", res))

	return nil
}
