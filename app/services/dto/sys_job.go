/*
 * @Author: 尤_Ta
 * @Date:  16:28
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  16:28
 */

package dto

import (
	"goods/app/services/dto/common"
	"goods/models"
	modelsCommon "goods/models/common"
)

// SysJobGetPageReq 映射列表接口
type SysJobGetPageReq struct {
	common.Pagination `search:"-"`
	//JobId             int    `form:"jobId" search:"type:exact;column:job_id;table:sys_job"`
	JobName        string `form:"jobName" search:"type:icontains;column:job_name;table:sys_job"`
	JobGroup       string `form:"jobGroup" search:"type:exact;column:job_group;table:sys_job"`
	CronExpression string `form:"cronExpression" search:"type:exact;column:cron_expression;table:sys_job"`
	InvokeTarget   string `form:"invokeTarget" search:"type:exact;column:invoke_target;table:sys_job"`
	Status         int    `form:"status" search:"type:exact;column:status;table:sys_job"`
}

func (sysJobGetPageReq *SysJobGetPageReq) GetNeedSearch() interface{} {
	return *sysJobGetPageReq
}

// SysJobInsertReq 映射添加接口
type SysJobInsertReq struct {
	JobId          int    `json:"-" form:"-" comment:""`
	JobName        string `json:"jobName" binding:"required"` // 名称
	JobGroup       string `json:"jobGroup"`                   // 任务分组
	JobType        int    `json:"jobType"`                    // 任务类型
	CronExpression string `json:"cronExpression"`             // cron表达式
	InvokeTarget   string `json:"invokeTarget"`               // 调用目标
	Args           string `json:"args"`                       // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy"`              // 执行策略
	Concurrent     int    `json:"concurrent"`                 // 是否并发
	Status         int    `json:"status"`                     // 状态
	EntryId        int    `json:"entryId"`                    // job启动时返回的id
	modelsCommon.ControlBy
}

func (sysJobInsertReq *SysJobInsertReq) Generate(sysJobModel *models.SysJob) {
	sysJobModel.JobId = sysJobInsertReq.JobId
	sysJobModel.JobName = sysJobInsertReq.JobName
	sysJobModel.JobGroup = sysJobInsertReq.JobGroup
	sysJobModel.JobType = sysJobInsertReq.JobType
	sysJobModel.CronExpression = sysJobInsertReq.CronExpression
	sysJobModel.InvokeTarget = sysJobInsertReq.InvokeTarget
	sysJobModel.Args = sysJobInsertReq.Args
	sysJobModel.MisfirePolicy = sysJobInsertReq.MisfirePolicy
	sysJobModel.Concurrent = sysJobInsertReq.Concurrent
	sysJobModel.Status = sysJobInsertReq.Status
	sysJobModel.EntryId = sysJobInsertReq.EntryId
	sysJobModel.CreateBy = sysJobInsertReq.CreateBy
}

func (sysJobInsertReq *SysJobInsertReq) GetId() interface{} {
	return sysJobInsertReq.JobId
}

// SysJobUpdateReq 映射更新接口
type SysJobUpdateReq struct {
	JobId          int    `json:"jobId" binding:"required"`
	JobName        string `json:"jobName" binding:"required"` // 名称
	JobGroup       string `json:"jobGroup"`                   // 任务分组
	JobType        int    `json:"jobType"`                    // 任务类型
	CronExpression string `json:"cronExpression"`             // cron表达式
	InvokeTarget   string `json:"invokeTarget"`               // 调用目标
	Args           string `json:"args"`                       // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy"`              // 执行策略
	Concurrent     int    `json:"concurrent"`                 // 是否并发
	Status         int    `json:"status"`                     // 状态
	EntryId        int    `json:"entryId"`                    // job启动时返回的id
	modelsCommon.ControlBy
}

func (sysJobUpdateReq *SysJobUpdateReq) Generate(sysJobModel *models.SysJob) {
	sysJobModel.JobId = sysJobUpdateReq.JobId
	sysJobModel.JobName = sysJobUpdateReq.JobName
	sysJobModel.JobGroup = sysJobUpdateReq.JobGroup
	sysJobModel.JobType = sysJobUpdateReq.JobType
	sysJobModel.CronExpression = sysJobUpdateReq.CronExpression
	sysJobModel.InvokeTarget = sysJobUpdateReq.InvokeTarget
	sysJobModel.Args = sysJobUpdateReq.Args
	sysJobModel.MisfirePolicy = sysJobUpdateReq.MisfirePolicy
	sysJobModel.Concurrent = sysJobUpdateReq.Concurrent
	sysJobModel.Status = sysJobUpdateReq.Status
	sysJobModel.EntryId = sysJobUpdateReq.EntryId
	sysJobModel.UpdateBy = sysJobUpdateReq.UpdateBy
}

func (sysJobUpdateReq *SysJobUpdateReq) GetId() interface{} {
	return sysJobUpdateReq.JobId
}

// SysJobGetInfoReq 映射详情和修改页面展示接口
type SysJobGetInfoReq struct {
	//Id int `uri:"jobId" form:"jobId"`
	Id int `uri:"id" form:"id"`
}

func (sysJobGetInfoReq *SysJobGetInfoReq) GetId() interface{} {
	return sysJobGetInfoReq.Id
}

// SysJobDeleteReq 映射删除接口
type SysJobDeleteReq struct {
	//必须是json格式才行,posman测试form传递数组无法解析
	Ids                    []int `json:"ids" form:"ids"`
	modelsCommon.ControlBy `json:"-"`
}

func (sysJobDeleteReq *SysJobDeleteReq) GetId() interface{} {
	return sysJobDeleteReq.Ids
}
