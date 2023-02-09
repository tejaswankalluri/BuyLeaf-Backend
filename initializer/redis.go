package initializer

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"log"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	if !fiber.IsChild() {
		log.Println(pong)
	}
}
