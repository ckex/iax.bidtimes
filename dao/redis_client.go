package dao

import (
	"time"

	"github.com/beego/redigo/redis"
	"github.com/astaxie/beego"
)

var (
	RedisClient *redis.Pool
	REDIS_HOST string
	REDIS_DB int
)

func init() {
	//	config.NewConfig("yaml")
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := beego.AppConfig.String("redisHost")
	//REDIS_HOST = "127.0.0.1:6379"
	//REDIS_HOST = "192.168.118.126:6379"
	REDIS_DB = 0
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
