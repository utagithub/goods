/*
 * @Author: 尤_Ta
 * @Date:  16:05
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  16:05
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

type LoginServices struct {
	service.Service
}

// Login 登录
func (e *LoginServices) Login(c *dto.UserLoginReq, model *models.User) error {
	var err error
	//var data models.User
	db := e.Orm.Debug().Debug().Model(model).Scopes(
		common.MakeCondition(c.LoginSearch()),
	).
		Find(model)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		//err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Login Find data error: %s", err.Error())
		return err
	}
	return nil
}
