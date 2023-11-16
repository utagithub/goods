package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"goods/common/api"
	"goods/common/runtime"
)

func LoadPolicy(c *gin.Context) (*casbin.SyncedEnforcer, error) {
	log := api.GetRequestLogger(c)
	if err := runtime.App.GetCasbinKey(c.Request.Host).LoadPolicy(); err == nil {
		return runtime.App.GetCasbinKey(c.Request.Host), err
	} else {
		log.Errorf("casbin rbac_model or policy init error, %s ", err.Error())
		return nil, err
	}
}
