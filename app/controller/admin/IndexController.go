package admin

import (
	"github.com/gin-gonic/gin"
	"goods/common/api"
	"net/http"
)

type IndexController struct {
	api.Api
}

func (e IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/index.tmpl", nil)
}

func (e IndexController) Info(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/info.tmpl", nil)
}
