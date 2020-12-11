package main

import (
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"videoWeb/verify-service/config"
)

func main() {
	flag.Parse()
	config.Init()

	pool := &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")

			return err
		},
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Conf.Redis.Address)
		},
	}
	conn := pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	//写入数据
	_, e := conn.Do("set", "k1", "zs")
	if e != nil {
		fmt.Println(e)
		return
	}
	//读取数据
	value, e := redis.String(conn.Do("get", "k1"))
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(value)

}
