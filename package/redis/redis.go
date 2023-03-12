package redis

import (
	"fmt"

	redisClient "github.com/go-redis/redis"
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

	res, err := client.Ping().Result()

	if err != nil {
		fmt.Println("There is error while initiating redis client")
		return err
	}

	fmt.Printf("redis client is initialized successfully res:= %v", res)

	return nil
}
