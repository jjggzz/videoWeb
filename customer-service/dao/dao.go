package dao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
	"videoWeb/customer-service/config"
	"videoWeb/customer-service/dao/repository"
)

type Dao struct {
	Repo  repository.Repository
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

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB.Mysql.Username,
		conf.DB.Mysql.Password,
		conf.DB.Mysql.Host,
		conf.DB.Mysql.Port,
		conf.DB.Mysql.Schema)
	db, err := sqlx.Open("mysql", addr)
	if err != nil {
		log.Println("数据库连接失败")
		panic(err)
	}
	return &Dao{Repo: repository.NewRepo(db), redis: pool}
}
