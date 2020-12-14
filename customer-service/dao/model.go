package dao

import "time"

type Customer struct {
	Id           int64     `json:"id" gorm:"primary_key"`
	AccessKey    string    `json:"accessKey" gorm:"column:access_key"`
	CreateTime   time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime   time.Time `json:"updateTime" gorm:"column:update_time"`
	DeleteStatus int       `json:"deleteStatus" gorm:"column:delete_status"`
	Phone        string    `json:"phone" gorm:"column:phone"`
	Username     string    `json:"username" gorm:"column:username"`
	Password     string    `json:"password" gorm:"column:password"`
	Email        string    `json:"email" gorm:"column:email"`
	Nickname     string    `json:"nickname" gorm:"column:nickname"`
	Status       int       `json:"status" gorm:"column:status"`
}

func (Customer) TableName() string {
	return "customer"
}
