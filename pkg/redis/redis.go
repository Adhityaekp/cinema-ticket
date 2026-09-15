package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Adhityaekp/cinema-ticket/config"
	goredis "github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg *config.Config) (*goredis.Client, error) {
	db, err := strconv.Atoi(cfg.RedisDB)
	if err != nil {
		return nil, err
	}

	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       db,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
