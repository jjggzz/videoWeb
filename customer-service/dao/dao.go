package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"videoWeb/customer-service/config"
)

type Dao struct {
	db *gorm.DB
}

func New(conf *config.Config) *Dao {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB.Mysql.Username,
		conf.DB.Mysql.Password,
		conf.DB.Mysql.Host,
		conf.DB.Mysql.Port,
		conf.DB.Mysql.Schema)
	c, err := gorm.Open("mysql", addr)
	if err != nil {
		log.Println("数据库连接失败")
		panic(err)
	}
	return &Dao{db: c}
}
