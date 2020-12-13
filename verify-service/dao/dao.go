package dao

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"videoWeb/verify-service/config"
)

type Dao struct {
	redis *redis.Pool
}

func New(conf *config.Config) *Dao {
	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", conf.Redis.Address)
		},
		MaxIdle:     conf.Redis.MaxIdle,
		MaxActive:   conf.Redis.MaxActive,
		IdleTimeout: time.Duration(conf.Redis.IdleTimeout),
	}
	return &Dao{redis: pool}
}
