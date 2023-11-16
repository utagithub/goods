package models

import "goods/models/common"

type User struct {
	Id       int    `json:"id" form:"id" gorm:"primaryKey;autoIncrement;column:id;comment:主键id"`
	Username string `json:"username" form:"username" gorm:"size:64;DEFAULT:NULL;column:username;comment:用户名"`
	Password string `json:"password" form:"password" gorm:"size:128;DEFAULT:NULL;column:password;comment:密码"`
	common.ModelTime
}

func (User) TableName() string {
	return "user"
}
