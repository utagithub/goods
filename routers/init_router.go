package routers

import (
	"github.com/gin-gonic/gin"
	log "goods/common/logger"
	common "goods/common/middleware"
	"goods/common/runtime"
	"os"
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := runtime.App.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	// the jwt middleware
	authMiddleware, err := common.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error, %s", err.Error())
	}

	// 注册系统路由
	//InitSysRouter(r, authMiddleware)

	// 注册业务路由
	initRouter(r, authMiddleware)
}
