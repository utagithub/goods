/*
 * @Author: 尤_Ta
 * @Date:  22:49
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  22:49
 */

package services

import (
	"errors"
	"github.com/robfig/cron/v3"
	"goods/app/services/dto/common"
	"goods/common/job"
	"goods/common/service"
	"goods/models"
	"time"
)

type SysJobControlServices struct {
	service.Service
	Cron *cron.Cron
}

// StartJob 启动任务
func (e *SysJobControlServices) StartJob(c *common.GeneralGetDto) error {
	var data models.SysJob
	var err error
	err = e.Orm.Debug().Table(data.TableName()).First(&data, c.Id).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err.Error())
		return err
	}
	if data.Status == 1 {
		err = errors.New("当前Job是关闭状态不能被启动，请先启用。")
		return err
	}
	if data.EntryId != 0 {
		err = errors.New("当前Job已在运行中,请勿重复操作")
		return err
	}

	if data.JobType == 1 {
		var j = &job.HttpJob{}
		j.InvokeTarget = data.InvokeTarget     //调用目标
		j.CronExpression = data.CronExpression //cron表达式
		j.JobId = data.JobId                   //任务id
		j.Name = data.JobName                  //任务名称
		data.EntryId, err = job.AddJob(e.Cron, j)
		if err != nil {
			e.Log.Errorf("jobs AddJob[HttpJob] error: %s", err)
		}
	} else {
		var j = &job.ExecJob{}
		j.InvokeTarget = data.InvokeTarget
		j.CronExpression = data.CronExpression
		j.JobId = data.JobId
		j.Name = data.JobName
		j.Args = data.Args
		data.EntryId, err = job.AddJob(e.Cron, j)
		if err != nil {
			e.Log.Errorf("jobs AddJob[ExecJob] error: %s", err)
		}
	}
	if err != nil {
		return err
	}

	err = e.Orm.Debug().Table(data.TableName()).Where(c.Id).Updates(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
	}
	return err
}

// RemoveJob 删除job
func (e *SysJobControlServices) RemoveJob(c *common.GeneralDelDto) error {
	var err error
	var data models.SysJob
	err = e.Orm.Debug().Table(data.TableName()).First(&data, c.Id).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err.Error())
		return err
	}

	if data.EntryId == 0 {
		err = errors.New("当前Job已停止,请勿重复操作")
		return err
	}

	cn := job.Remove(e.Cron, data.EntryId)

	select {
	case res := <-cn:
		if res {
			err = e.Orm.Debug().Table(data.TableName()).Where("entry_id = ?", data.EntryId).Update("entry_id", 0).Error
			if err != nil {
				e.Log.Errorf("db error: %s", err)
			}
			return err
		}
	case <-time.After(time.Second * 1):
		//e.Msg = "操作超时！"
		return nil
	}
	return nil
}
