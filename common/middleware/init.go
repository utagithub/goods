package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	JwtTokenCheck string = "JwtToken"
	//RoleCheck       string = "AuthCheckRole"
	//PermissionCheck string = "PermissionAction"
)

func InitMiddleware(r *gin.Engine) {
	// 数据库链接
	r.Use(WithContextDb)
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	// NoCache is a middleware function that appends headers
	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	//r.Use(DemoEvn())
	// 链路追踪
	//r.Use(middleware.Trace())

	//runtime.App.SetMiddleware(JwtTokenCheck, (*jwt.GinJWTMiddleware).MiddlewareFunc)
	//runtime.App.SetMiddleware(RoleCheck, AuthCheckRole())
	//runtime.App.SetMiddleware(PermissionCheck, actions.PermissionAction())
}
