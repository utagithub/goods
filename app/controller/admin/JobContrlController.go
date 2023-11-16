/*
 * @Author: 尤_Ta
 * @Date:  23:06
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:06
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"goods/app/services"
	"goods/app/services/dto/common"
	"goods/common/api"
	"goods/common/runtime"
)

type JobControlController struct {
	api.Api
}

// StartJobById 启动任务
func (e JobControlController) StartJobById(c *gin.Context) {
	req := common.GeneralGetDto{}
	s := services.SysJobControlServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	s.Cron = runtime.App.GetCrontabKey(c.Request.Host)
	err = s.StartJob(&req)
	if err != nil {
		e.Logger.Errorf("StartJob error, %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.Id, "StartJob Success")

}

// StopJobById 停止任务
func (e JobControlController) StopJobById(c *gin.Context) {
	req := common.GeneralDelDto{}
	s := services.SysJobControlServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	s.Cron = runtime.App.GetCrontabKey(c.Request.Host)
	err = s.RemoveJob(&req)
	if err != nil {
		e.Logger.Errorf("RemoveJob error, %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.Id, "RemoveJob Success")
}
