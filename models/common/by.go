package common

import (
	"gorm.io/gorm"
	"time"
)

// ModelTime 如果想忽略某个字段，加上`json:”-”`；如果在值为空时忽略，加上omitempty option，如：`json:”,omitempty”`
type ModelTime struct {
	CreatedAt time.Time      `json:"created_at" form:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

type ControlBy struct {
	CreateBy int `json:"create_by" form:"create_by" gorm:"size:4;DEFAULT:NULL;index;comment:创建者"`
	UpdateBy int `json:"update_by" form:"update_by" gorm:"size:4;index;DEFAULT:NULL;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}
