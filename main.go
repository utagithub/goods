package main

import (
	"goods/cmd"
)

//go:generate swag init --parseDependency --parseDepth=3

func main() {
	//设置开发环境
	//gin.SetMode(gin.ReleaseMode)

	//初始化数据库连接
	//dao.MysqlInit()
	//_ = dao.DB.AutoMigrate(&models.Good{})
	//_ = dao.DB.AutoMigrate(&models.User{})
	//_ = dao.DB.AutoMigrate(&models.SysJob{})

	//初始化日志-项目中没有使用logurs
	//InitLogurs()

	//设置命令行
	cmd.Execute()
	//生成http engine
	//engine := routers.InitEngine()

	//运行http engine
	//_ = engine.Run("localhost:8080")

}
0



