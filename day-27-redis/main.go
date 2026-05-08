package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {

	// always initialize context with background
	var ctx = context.Background()

	//redis client initialization
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	//ping the redis server
	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic(err.Error())
	}

	//set key in redis with expiry duration
	er := rdb.Set(ctx, "user", "tony", 10*time.Minute).Err()

	if er != nil {
		panic("Error while set key in redis")
	}

	//get the value of the key from redis
	val, err2 := rdb.Get(ctx, "user").Result()

	if err2 != nil {
		panic("Error while getting key from redis")
	}

	fmt.Printf("The value from redis is : %s", val)

	//Delete the key from redis
	err3 := rdb.Del(ctx, "user").Err()
	if err3 != nil {
		panic("Error while deleting key from redis")
	}

	//get the value of the key from redis
	val2, err4 := rdb.Get(ctx, "user").Result()

	if err4 != nil {
		fmt.Println("Error while getting key from redis")
	}

	fmt.Printf("The value from redis is : %s", val2)

	//Check if the key is exists in redis
	exists, err5 := rdb.Exists(ctx, "user").Result()
	if err5 != nil {
		panic("Error while checking key existence in redis")
	}

	fmt.Printf("The key exists in redis : %d", exists)

	//expire the key from redis
	err6 := rdb.Expire(ctx, "user", 10*time.Minute).Err()
	if err6 != nil {
		panic("Error while expiring key from redis")
	}

}
