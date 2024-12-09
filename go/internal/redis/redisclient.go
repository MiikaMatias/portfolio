package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/MiikaMatias/portfolio/internal/utils"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	DB struct {
		RedisIp   string `yaml:"redisip"`
		RedisPort string `yaml:"redisport"`
	} `yaml:"db"`
}

const (
	CONFIG_FILE_PATH = "internal/redis/config.yaml"
)

func InitializeRedisClient() redis.Client {
	log.Println("Starting Redis Client")

	config := Config{}
	err := utils.CastConfig(CONFIG_FILE_PATH, &config)
	if err != nil {
		log.Fatalf("Error when opening config: %s", err)
	}

	addr := fmt.Sprintf("%s:%s", config.DB.RedisIp, config.DB.RedisPort)
	log.Print(addr)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	},
	)

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to ping redis upon client start")
	}
	log.Println(ping)

	return *client
}
