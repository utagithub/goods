package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 项目中没有使用该db,而是yaml文件配置的db
var (
	DB         *gorm.DB
	connectErr error
)

// MysqlInit 数据库初始化
func MysqlInit() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goods?charset=utf8mb4&parseTime=True&loc=Local"
	DB, connectErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if connectErr != nil {
		panic("failed to connect database")
	}

	//设置存储引擎和字符集
	DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

}
