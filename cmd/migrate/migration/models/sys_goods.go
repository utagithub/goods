package models

import "goods/models/common"

type SysGoods struct {
	Id         int    `json:"id" form:"id" gorm:"primaryKey;autoIncrement;column:id;comment:主键id"`
	GoodsName  string `json:"goods_name" form:"goods_name" gorm:"size:32;DEFAULT:NULL;column:goods_name;comment:商品名称"`
	CatId      int    `json:"cat_id" form:"cat_id" gorm:"size:4;DEFAULT:NULL;column:cat_id;comment:商品分类"`
	GoodsPrice string `json:"goods_price" form:"goods_price" gorm:"size:128;DEFAULT:NULL;column:goods_price;comment:商品价格"`
	GoodsThumb string `json:"goods_thumb" form:"goods_thumb" gorm:"size:255;DEFAULT:NULL;column:goods_thumb;comment:商品缩略图"`
	GoodsDesc  string `json:"goods_desc" form:"goods_desc" gorm:"size:255;DEFAULT:NULL;column:goods_desc;comment:商品描述"`
	common.ModelTime
	common.ControlBy
}

func (SysGoods) TableName() string {
	return "good"
}
