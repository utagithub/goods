package api

import (
	"github.com/gin-gonic/gin"
	"goods/common/logger"
	"goods/common/runtime"
	"goods/common/tools"
	"strings"
)

type loggerKey struct{}

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logger.Helper {
	var log *logger.Helper
	l, ok := c.Get(tools.LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*logger.Helper)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := tools.GenerateMsgIDFromContext(c)
	log = logger.NewHelper(runtime.App.GetLogger()).WithFields(map[string]interface{}{
		strings.ToLower(tools.TrafficKey): requestId,
	})
	return log
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := tools.GenerateMsgIDFromContext(c)
	log := logger.NewHelper(runtime.App.GetLogger()).WithFields(map[string]interface{}{
		strings.ToLower(tools.TrafficKey): requestId,
	})
	c.Set(tools.LoggerKey, log)
}
