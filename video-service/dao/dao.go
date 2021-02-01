// author: JGZ
// time:   2021-01-25 15:52
package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"videoWeb/video-service/config"
	"videoWeb/video-service/dao/repository"
)

type Dao struct {
	Repo repository.Repository
}

func New(conf *config.Config) *Dao {
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
	return &Dao{Repo: repository.NewRepo(db)}
}
