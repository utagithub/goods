/*
 * @Author: 尤_Ta
 * @Date:  16:48
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  16:48
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/app/services"
	"goods/app/services/dto"
	"goods/common/api"
	"goods/common/jwtauth/user"
	"goods/models"
	"net/http"
)

type JobController struct {
	api.Api
}

// List 列表
// @Summary 任务列表
// @Description 获取JSON
// @Tags 任务数据
// @Param cat_id query string false "cat_id"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/admin/job/list [get]
// @Security Bearer
func (e JobController) List(c *gin.Context) {
	//查询参数数据绑定
	req := dto.SysJobGetPageReq{}
	s := services.JobServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.SysJob, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")

}

// GetById 详情
// @Summary 通过id获取任务数据
// @Description 获取JSON
// @Tags 任务数据
// @Param id path int true "任务ID"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/admin/job/{id} [get]
// @Security Bearer
func (e JobController) GetById(c *gin.Context) {
	req := dto.SysJobGetInfoReq{}
	s := services.JobServices{}
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
	var object models.SysJob
	err = s.Get(&req, &object)
	if err != nil {
		e.Logger.Warnf("Get error: %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(object, "查询成功")
}

// Create 添加页面
func (e JobController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create page",
	})
}

// DoCreate 添加操作
// @Summary 添加任务数据
// @Description 获取JSON
// @Tags 任务数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysJobInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/job/create [post]
// @Security Bearer
func (e JobController) DoCreate(c *gin.Context) {
	req := dto.SysJobInsertReq{}
	s := services.JobServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")

}

// Update 更新页面
func (e JobController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update page",
	})
}

// DoUpdate 更新操作
// @Summary 修改任务数据
// @Description 获取JSON
// @Tags 任务数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysJobUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/job/update/{id} [put]
// @Security Bearer
func (e JobController) DoUpdate(c *gin.Context) {
	req := dto.SysJobUpdateReq{}
	s := services.JobServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// DeleteById 删除
// @Summary 删除任务数据
// @Description 删除数据
// @Tags 任务数据
// @Param dictCode body dto.SysJobDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/job [delete]
// @Security Bearer
func (e JobController) DeleteById(c *gin.Context) {
	//数据绑定
	req := dto.SysJobDeleteReq{}
	s := services.JobServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "删除成功")
}
