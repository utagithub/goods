package services

import (
	"errors"
	"goods/app/services/dto"
	"goods/app/services/dto/common"
	"goods/common/service"
	"goods/models"
	"gorm.io/gorm"
)

type GoodServices struct {
	service.Service
}

// GetPage 获取列表
func (e *GoodServices) GetPage(c *dto.GoodGetPageReq, list *[]models.Good, count *int64) error {
	var err error
	var data models.Good
	err = e.Orm.Debug().Model(&data).
		Scopes(
			common.OrderDest("id", true),
			common.MakeCondition(c.GetNeedSearch()),
			common.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Get list data errorr: %s", err)
		return err
	}
	return nil
}

// Get 获取对象
func (e *GoodServices) Get(c *dto.GoodGetInfoReq, model *models.Good) error {
	var err error
	var data models.Good
	db := e.Orm.Debug().Model(&data).First(model, c.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看数据不存在或无权查看")
		return err
	}
	return nil
}

// Insert 创建对象
func (e *GoodServices) Insert(c *dto.GoodInsertReq) error {
	var err error
	var data = new(models.Good)
	c.Generate(data)
	err = e.Orm.Create(data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改对象
func (e *GoodServices) Update(c *dto.GoodUpdateReq) error {
	var err error
	var model = models.Good{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.Orm.Save(model)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败,该数据不存在或已被删除")

	}
	return nil

}

// Remove 删除
func (e *GoodServices) Remove(c *dto.GoodDeleteReq) error {
	var err error
	var data models.Good
	err = e.Orm.Debug().Model(&data).Where("id IN ?", c.GetId()).Update("update_by", 1).Error
	if err != nil {
		e.Log.Errorf("Delete Before Update: %s", err.Error())
		return err
	} else {
		db := e.Orm.Debug().Delete(&data, c.GetId())
		if err = db.Error; err != nil {
			err = db.Error
			e.Log.Errorf("Delete error: %s", err)
			return err
		}
		if db.RowsAffected == 0 {
			err = errors.New("删除失败,该数据不存在或已被删除")
			return err
		}
		return nil
	}

}
