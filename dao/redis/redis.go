package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),        // 数据库
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("rdb.Ping() failed", zap.Error(err))
	}
	zap.L().Debug("rdb.Ping() Ok ========")
	return nil
}

func Close() {
	_ = rdb.Close()
}
