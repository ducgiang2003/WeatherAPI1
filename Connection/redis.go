package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func RedisCon(addr string, ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
	//Ping to rdb -> pong if ping success
	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Ping to redis server failed : %v", err)
	}
	fmt.Printf("Redis connection succesfully")
	return rdb
}
