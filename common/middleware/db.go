package middleware

import (
	"github.com/gin-gonic/gin"
	"goods/common/runtime"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", runtime.App.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}
