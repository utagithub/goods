package dto

import (
	"goods/app/services/dto/common"
	"goods/models"
	modelsCommon "goods/models/common"
)

// GoodGetPageReq 映射列表接口
type GoodGetPageReq struct {
	common.Pagination `search:"-"`
	CatId             int `form:"cat_id" search:"type:exact;column:cat_id;table:good" comment:""`
}

func (goodGetPageReq *GoodGetPageReq) GetNeedSearch() interface{} {
	return *goodGetPageReq
}

// GoodInsertReq 映射添加接口
type GoodInsertReq struct {
	Id         int    `json:"-" form:"-" comment:""`
	GoodsName  string `json:"goods_name" form:"goods_name" binding:"required" comment:""`
	CatId      int    `json:"cat_id" form:"cat_id" binding:"required,catIdValid" catIdValidMsg:"商品类别超出范围" comment:""`
	GoodsPrice string `json:"goods_price" form:"goods_price" binding:"required" comment:""`
	GoodsThumb string `json:"goods_thumb" form:"goods_thumb" binding:"required" comment:""`
	GoodsDesc  string `json:"goods_desc" form:"goods_desc" binding:"required" comment:""`
	modelsCommon.ControlBy
}

func (goodInsertReq *GoodInsertReq) Generate(goodModel *models.Good) {
	goodModel.Id = goodInsertReq.Id
	goodModel.GoodsName = goodInsertReq.GoodsName
	goodModel.CatId = goodInsertReq.CatId
	goodModel.GoodsPrice = goodInsertReq.GoodsPrice
	goodModel.GoodsThumb = goodInsertReq.GoodsThumb
	goodModel.GoodsDesc = goodInsertReq.GoodsDesc
	goodModel.GoodsDesc = goodInsertReq.GoodsDesc
	goodModel.CreateBy = goodInsertReq.CreateBy
}

func (goodInsertReq *GoodInsertReq) GetId() interface{} {
	return goodInsertReq.Id
}

// GoodGetInfoReq 映射详情和修改页面展示接口
type GoodGetInfoReq struct {
	Id int `uri:"id" form:"id"`
}

func (goodGetInfoReq *GoodGetInfoReq) GetId() interface{} {
	return goodGetInfoReq.Id
}

// GoodUpdateReq 映射更新接口
type GoodUpdateReq struct {
	Id         int    `json:"id" form:"id" comment:""`
	GoodsName  string `json:"goods_name" form:"goods_name" comment:""`
	CatId      int    `json:"cat_id" form:"cat_id" comment:""`
	GoodsPrice string `json:"goods_price" form:"goods_price" comment:""`
	GoodsThumb string `json:"goods_thumb" form:"goods_thumb" comment:""`
	GoodsDesc  string `json:"goods_desc" form:"goods_desc" comment:""`
	modelsCommon.ControlBy
}

func (goodUpdateReq *GoodUpdateReq) Generate(goodModel *models.Good) {
	goodModel.Id = goodUpdateReq.Id
	goodModel.GoodsName = goodUpdateReq.GoodsName
	goodModel.CatId = goodUpdateReq.CatId
	goodModel.GoodsPrice = goodUpdateReq.GoodsPrice
	goodModel.GoodsThumb = goodUpdateReq.GoodsThumb
	goodModel.GoodsDesc = goodUpdateReq.GoodsDesc
	goodModel.GoodsDesc = goodUpdateReq.GoodsDesc
	goodModel.UpdateBy = goodUpdateReq.UpdateBy
}

func (goodUpdateReq *GoodUpdateReq) GetId() interface{} {
	return goodUpdateReq.Id
}

// GoodDeleteReq 映射删除接口
type GoodDeleteReq struct {
	//必须是json格式才行,posman测试form传递数组无法解析
	Ids                    []int `json:"ids" form:"ids"`
	modelsCommon.ControlBy `json:"-"`
}

func (goodDeleteReq *GoodDeleteReq) GetId() interface{} {
	return goodDeleteReq.Ids
}
