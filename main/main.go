package main

import (
	"distributed_locking/package/redis"
	"fmt"
)

func main() {
	redisConfig := redis.Config{
		Address:  "127.0.0.1:6379",
		Password: "",
		Db:       0,
	}

	fmt.Println("initiating redis client")
	initRedisClient(redisConfig)
}

func initRedisClient(c redis.Config) {
	err := redis.CreateClient(c)
	if err != nil {
		fmt.Println("failed initiating redis client")
	}

}
