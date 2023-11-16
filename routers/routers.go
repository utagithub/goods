/*
 * @Author: 尤_Ta
 * @Date:  00:04
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  00:04
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"goods/app/controller/admin"
	jwt "goods/common/jwtauth"
	"goods/common/middleware/handler"
	"goods/common/ws"
	_ "goods/docs"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCheckRoleRoutes)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckRolerRoutes)
}

func registerNoCheckRolerRoutes(r *gin.RouterGroup) {

	r.GET("/test", admin.BaseController{}.Index)

	//路由组-captcha
	captchaGroup := r.Group("/captcha")
	{
		captchaGroup.GET("", admin.CapchaController{}.GenerateCaptchaHandler)
	}

	//路由组-admin-login
	loginGroup := r.Group("/admin/login")
	{
		loginGroup.GET("/index", admin.LoginController{}.Index)
		loginGroup.POST("oldLogin", admin.LoginController{}.DoLogin)
	}

}

func registerCheckRoleRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	//路由组-admin-login
	loginGroup := r.Group("/admin/login")
	{
		loginGroup.POST("", authMiddleware.LoginHandler)
		loginGroup.GET("/refresh_token", authMiddleware.RefreshHandler)
		loginGroup.POST("/logout", handler.LogOut)
	}

	wss := r.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}

	//路由组-captcha
	fileUpload := r.Group("/file")
	{
		fileUpload.POST("upload", admin.FileUploadController{}.UploadFile)
	}

	//路由组-admin-index
	adminIndexGroup := r.Group("/admin/index")
	{
		adminIndexGroup.GET("/index", admin.IndexController{}.Index)
		adminIndexGroup.GET("/info", admin.IndexController{}.Info)
	}

	//路由组-goods
	//adminGoodsGroup := r.Group("/admin/goods").Use(authMiddleware.MiddlewareFunc())
	adminGoodsGroup := r.Group("/admin/goods")
	//auth中间件-自测时写的
	//adminGoodsGroup.Use(middleware.JWTAuthMiddleware())
	{
		adminGoodsGroup.GET("/create", admin.GoodController{}.Create)
		adminGoodsGroup.POST("/create", admin.GoodController{}.DoCreate)

		adminGoodsGroup.GET("/update/:id", admin.GoodController{}.Update)
		adminGoodsGroup.POST("/update", admin.GoodController{}.DoUpdate)

		adminGoodsGroup.GET("/list", admin.GoodController{}.List)

		adminGoodsGroup.GET("/detail/:id", admin.GoodController{}.GetById)

		adminGoodsGroup.DELETE("", admin.GoodController{}.DeleteById)
	}

	//路由组-job 定时任务
	//adminJobGroup := r.Group("/admin/job").Use(authMiddleware.MiddlewareFunc())
	adminJobGroup := r.Group("/admin/job")
	{
		adminJobGroup.GET("/create", admin.JobController{}.Create)
		adminJobGroup.POST("/create", admin.JobController{}.DoCreate)

		adminJobGroup.GET("/update/:id", admin.JobController{}.Update)
		adminJobGroup.PUT("/update", admin.JobController{}.DoUpdate)

		adminJobGroup.GET("/list", admin.JobController{}.List)

		adminJobGroup.GET("/detail/:id", admin.JobController{}.GetById)

		adminJobGroup.DELETE("", admin.JobController{}.DeleteById)

		adminJobGroup.GET("/start/:id", admin.JobControlController{}.StartJobById)
		adminJobGroup.GET("/stop/:id", admin.JobControlController{}.StopJobById)
	}

}
