/*
 * @Author: 尤_Ta
 * @Date:  23:32
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:32
 */

package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goods/common/api"
	"goods/common/config"
	"goods/common/config/source/file"
	"goods/common/database"
	"goods/common/global"
	"goods/common/job"
	"goods/common/middleware"
	"goods/common/middleware/handler"
	"goods/common/runtime"
	"goods/common/storage/storage"
	"goods/common/tools"
	ext "goods/config"
	"goods/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "goods server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, routers.InitRouter)
}

func setup() {
	// 注入配置扩展项
	//fmt.Printf("66666:%v\n", ext.ExtConfig)
	config.ExtendConfig = &ext.ExtConfig
	//1. 读取配置
	config.Setup(
		file.NewSource(file.WithPath(configYml)),
		database.Setup,
		storage.Setup,
	)
	//fmt.Printf("66666:%v\n", ext.ExtConfig.UpLoad.Driver)
	//注册监听函数
	//queue := sdk.Runtime.GetMemoryQueue("")
	//queue.Register(global.LoginLog, models.SaveLoginLog)
	//queue.Register(global.OperateLog, models.SaveOperaLog)
	//queue.Register(global.ApiCheck, models.SaveSysApi)
	//go queue.Run()

	usageStr := `starting api server...`
	log.Println(usageStr)
}

func run() error {
	//设置开发模式
	if config.ApplicationConfig.Mode == global.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}

	//初始化gin engine
	initRouter()

	//初始化路由,关联上下文gin engine
	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: runtime.App.GetEngine(),
	}

	//初始化job
	go func() {
		job.InitJob() //初始化任务列表,定时任务可以放在job_exec目录下
		job.Setup(runtime.App.GetDb())
	}()

	//if apiCheck {
	//	var routers = runtime.App.GetRouter()
	//	q := runtime.App.GetMemoryQueue("")
	//	mp := make(map[string]interface{}, 0)
	//	mp["List"] = routers
	//	message, err := runtime.App.GetStreamMessage("", global.ApiCheck, mp)
	//	if err != nil {
	//		log.Printf("GetStreamMessage error, %s \n", err.Error())
	//		//日志报错错误，不中断请求
	//	} else {
	//		err = q.Append(message)
	//		if err != nil {
	//			log.Printf("Append message error, %s \n", err.Error())
	//		}
	//	}
	//}

	go func() {
		// 服务连接
		if config.SslConfig.Enable {
			if err := srv.ListenAndServeTLS(config.SslConfig.Pem, config.SslConfig.KeyStr); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		}
	}()
	fmt.Println(tools.Red(string(global.LogoContent)))
	tip()
	fmt.Println(tools.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", tools.GetLocaHonst(), config.ApplicationConfig.Port)
	//fmt.Println(common.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/swagger/admin/index.html \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%d/swagger/admin/index.html \r\n", tools.GetLocaHonst(), config.ApplicationConfig.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Printf("%s Shutdown Server ... \r\n", tools.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	return nil
}

//var Router runtime.Router

func tip() {
	usageStr := `欢迎使用 ` + tools.Green(`goods `+global.Version) + ` 可以使用 ` + tools.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}

func initRouter() {
	var r *gin.Engine
	h := runtime.App.GetEngine()
	if h == nil {
		//h = gin.New()
		h = routers.InitEngine()
		runtime.App.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	if config.SslConfig.Enable {
		r.Use(handler.TlsHandler())
	}

	//r.Use(middleware.Metrics())
	r.Use(middleware.Sentinel()).
		Use(middleware.RequestId(tools.TrafficKey)).
		Use(api.SetRequestLogger)

	middleware.InitMiddleware(r)

}
