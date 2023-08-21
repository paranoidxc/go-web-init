package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),        // 数据库
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})

	//_, err = rdb.Ping().Result()
	return nil
}

func Close() {
	_ = rdb.Close()
}
