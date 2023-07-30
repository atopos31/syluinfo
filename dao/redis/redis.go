package redis

import (
	"cld/settings"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func Init(cfg *settings.RedisConfig) (err error) {
	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	//测试连通性
	ctx := context.Background()
	if _, err := db.Ping(ctx).Result(); err != nil {
		fmt.Println("connect redis  failed, err :" + err.Error())
		return err
	}

	rdb = db
	return
}

func Close() {
	rdb.Close()
}
