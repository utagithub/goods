/*
 * @Author: 尤_Ta
 * @Date:  16:02
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  16:02
 */

package models

import (
	"goods/models/common"
	"gorm.io/gorm"
)

type SysJob struct {
	JobId          int    `json:"jobId" gorm:"primaryKey;autoIncrement;comment:任务id"` // 任务编码
	JobName        string `json:"jobName" gorm:"size:255;comment:任务名称"`               // 任务名称
	JobGroup       string `json:"jobGroup" gorm:"size:255;comment:任务分组"`              // 任务分组 defalut/system
	JobType        int    `json:"jobType" gorm:"size:1;comment:任务类型"`                 // 任务类型  接口/函数
	CronExpression string `json:"cronExpression" gorm:"size:255;comment:cron表达式"`     // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"size:255;comment:调用目标"`          // 调用目标
	Args           string `json:"args" gorm:"size:255;comment:目标参数"`                  // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;comment:执行策略"`         // 执行策略  1立即执行/2执行一次/3废弃
	Concurrent     int    `json:"concurrent" gorm:"size:1;comment:是否并发"`              // 是否并发 1否/2是
	Status         int    `json:"status" gorm:"size:1;comment:状态"`                    // 状态  1停用/2正常  正常状态下才能运行
	EntryId        int    `json:"entry_id" gorm:"size:11;comment:job启动时返回的id"`        // job启动时返回的id
	common.ControlBy
	common.ModelTime

	//DataScope string `json:"dataScope" gorm:"-"`
}

// TableName ...
func (e *SysJob) TableName() string {
	return "sys_job"
}

// GetList 服务启动时,获取状态可用的列表
func (e *SysJob) GetList(db *gorm.DB, list interface{}) (err error) {
	return db.Debug().Table(e.TableName()).Where("status = ?", 2).Find(list).Error
}

// RemoveAllEntryID 服务启动时,设置entry_id=0
func (e *SysJob) RemoveAllEntryID(db *gorm.DB) (update SysJob, err error) {
	if err = db.Debug().Table(e.TableName()).Where("entry_id > ?", 0).Update("entry_id", 0).Error; err != nil {
		return
	}
	return
}

// Update  更新Job
func (e *SysJob) Update(db *gorm.DB, id interface{}) (err error) {
	return db.Debug().Table(e.TableName()).Where(id).Updates(&e).Error
}

//func (e *SysJob) Generate() common.ActiveRecord {
//	o := *e
//	return &o
//}
//
//func (e *SysJob) GetId() interface{} {
//	return e.JobId
//}
//
//func (e *SysJob) SetCreateBy(createBy int) {
//	e.CreateBy = createBy
//}
//
//func (e *SysJob) SetUpdateBy(updateBy int) {
//	e.UpdateBy = updateBy
//}
//
//
