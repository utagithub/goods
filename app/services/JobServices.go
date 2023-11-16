/*
 * @Author: 尤_Ta
 * @Date:  16:48
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  16:48
 */

package services

import (
	"errors"
	"goods/app/services/dto"
	"goods/app/services/dto/common"
	"goods/common/service"
	"goods/models"
	"gorm.io/gorm"
)

type JobServices struct {
	service.Service
}

// GetPage 获取列表
func (e *JobServices) GetPage(c *dto.SysJobGetPageReq, list *[]models.SysJob, count *int64) error {
	var err error
	var data models.SysJob
	err = e.Orm.Debug().Model(&data).Scopes(
		common.OrderDest("job_id", true),
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

// Insert 创建对象
func (e *JobServices) Insert(c *dto.SysJobInsertReq) error {
	var err error
	var data = new(models.SysJob)
	c.Generate(data)
	err = e.Orm.Debug().Create(data).Error
	if err != nil {
		e.Log.Errorf("Insert error: %s", err)
		return err
	}
	return nil
}

// Update 修改对象
func (e *JobServices) Update(c *dto.SysJobUpdateReq) error {
	var err error
	var model = models.SysJob{}
	err = e.Orm.Debug().First(&model, c.GetId()).Error
	if err != nil {
		e.Log.Errorf("Update for find data error: %s", err)
		return err
	}
	c.Generate(&model)
	db := e.Orm.Debug().Save(model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Update error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		e.Log.Error("Update error:更新该数据失败")
		return errors.New("更新该数据失败")
	}
	return nil
}

// Get 获取对象
func (e *JobServices) Get(c *dto.SysJobGetInfoReq, model *models.SysJob) error {
	var err error
	var data models.SysJob
	db := e.Orm.Debug().Model(&data).First(model, c.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		//err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Find data error: %s", err.Error())
		return err
	}
	return nil
}

// Remove 删除
func (e *JobServices) Remove(c *dto.SysJobDeleteReq) error {
	var err error
	var data models.SysJob
	err = e.Orm.Debug().Model(&data).Where("job_id IN ?", c.GetId()).Update("update_by", 1).Error
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
			e.Log.Error("Delete error:数据不存在,或已经被删除")
			err = errors.New("数据不存在,或已经被删除")
			return err
		}
		return nil
	}

}
