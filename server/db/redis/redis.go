package redis

import (
	"github.com/go-redis/redis" // 实现了redis连接池
	"github.com/name5566/leaf/log"
	"server/conf"
	"time"
)

// 定义redis链接池
var RedisDb *redis.Client

// 初始化redis链接池
func init() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:        conf.RedisHost,     // Redis地址
		Password:    conf.RedisPwd,      // Redis账号
		DB:          conf.RedisDB,       // Redis库
		PoolSize:    conf.RedisPoolSize, // Redis连接池大小
		MaxRetries:  3,                  // 最大重试次数
		IdleTimeout: 10 * time.Second,   // 空闲链接超时时间
	})
	pong, err := RedisDb.Ping().Result()
	if err == redis.Nil {
		panic("Redis异常")
	} else if err != nil {
		panic("失败")
	} else {
		log.Debug(pong)
	}
}
