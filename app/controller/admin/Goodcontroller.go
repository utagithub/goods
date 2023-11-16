package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/app/services"
	"goods/app/services/dto"
	"goods/common/api"
	"goods/common/jwtauth/user"
	"goods/common/validate/message"
	"goods/models"
	"net/http"
)

type GoodController struct {
	api.Api
}

// List 列表
// @Summary 商品列表
// @Description 获取JSON
// @Tags 商品数据
// @Param cat_id query string false "cat_id"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/admin/goods/list [get]
// @Security Bearer
func (e GoodController) List(c *gin.Context) {
	//查询参数数据绑定
	req := dto.GoodGetPageReq{}
	s := services.GoodServices{}

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

	list := make([]models.Good, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	//e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
	c.HTML(http.StatusOK, "admin/goods/list.tmpl", 111)

}

// GetById 详情
// @Summary 通过id获取商品数据
// @Description 获取JSON
// @Tags 商品数据
// @Param id path int true "商品ID"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/admin/goods/{id} [get]
// @Security Bearer
func (e GoodController) GetById(c *gin.Context) {
	req := dto.GoodGetInfoReq{}
	s := services.GoodServices{}
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
	var object models.Good
	err = s.Get(&req, &object)
	if err != nil {
		e.Logger.Warnf("Get error: %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(object, "查询成功")
}

// Create 添加页面
func (e GoodController) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/create.tmpl", nil)
}

// DoCreate 添加操作
// @Summary 添加商品数据
// @Description 获取JSON
// @Tags 商品数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.GoodInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/goods/create [post]
// @Security Bearer
func (e GoodController) DoCreate(c *gin.Context) {
	//数据绑定
	req := dto.GoodInsertReq{}
	s := services.GoodServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		// 显示自定义的错误验证信息
		err = message.GoodValidMsg(err, &req)
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
func (e GoodController) Update(c *gin.Context) {
	req := dto.GoodGetInfoReq{}
	s := services.GoodServices{}
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
	var object models.Good
	err = s.Get(&req, &object)
	if err != nil {
		e.Logger.Warnf("Get error: %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	c.HTML(http.StatusOK, "admin/goods/update.tmpl", object)
}

// DoUpdate 更新操作
// @Summary 修改商品数据
// @Description 获取JSON
// @Tags 商品数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.GoodUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/goods/update/{id} [put]
// @Security Bearer
func (e GoodController) DoUpdate(c *gin.Context) {
	//绑定数据
	req := dto.GoodUpdateReq{}
	s := services.GoodServices{}
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
// @Summary 删除商品数据
// @Description 删除数据
// @Tags 商品数据
// @Param dictCode body dto.GoodDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/goods [delete]
// @Security Bearer
func (e GoodController) DeleteById(c *gin.Context) {
	//数据绑定
	req := dto.GoodDeleteReq{}
	s := services.GoodServices{}
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
