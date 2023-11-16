package routers

import (
	"github.com/gin-gonic/gin"
	"goods/common/validate/register"
)

func InitEngine() *gin.Engine {
	//生成http engine
	//engine := gin.Default()
	engine := gin.New()

	// 注册一个全局中间件
	//r.Use(m3(true))
	// 注册一个全局中间件 -允许跨域请求
	//r.Use(cors.Default())

	//加载静态文件,访问/static开始的都去./static下找
	engine.Static("/static", "./static")
	//加本地图片文件,访问/storage开始的都去./storage下找
	engine.Static("/storage", "./storage")
	//加载多个模版,需要在模版里面定义模版名称,在绑定模版的时候使用
	engine.LoadHTMLGlob("./template/*/*/*")

	//注册路由
	//RegisterRoutes{}.registerRoutes(engine)

	//注册自定义验证器
	register.ValidRegister()

	return engine
}
