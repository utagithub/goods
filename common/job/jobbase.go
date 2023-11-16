/*
 * @Author: 尤_Ta
 * @Date:  21:59
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  21:59
 */

package job

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"goods/common/logger"
	"goods/common/runtime"
	"goods/common/tools"
	"goods/models"
	"gorm.io/gorm"
	"time"
)

var TimeFormat = "2006-01-02 15:04:05"
var retryCount = 3

var JobList map[string]JobsExec

//var lock sync.Mutex

type JobCore struct {
	InvokeTarget   string
	Name           string
	JobId          int
	EntryId        int
	CronExpression string
	Args           string
}

// HttpJob 任务类型 http
type HttpJob struct {
	JobCore
}

type ExecJob struct {
	JobCore
}

func (e *ExecJob) Run() {
	startTime := time.Now()
	var obj = JobList[e.InvokeTarget]
	if obj == nil {
		logger.Warn("[Job] ExecJob Run job nil")
		return
	}
	err := CallExec(obj.(JobsExec), e.Args)
	if err != nil {
		// 如果失败暂停一段时间重试
		fmt.Println(time.Now().Format(TimeFormat), " [ERROR] mission failed! ", err)
	}
	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//TODO: 待完善部分
	//str := time.Now().Format(TimeFormat) + " [INFO] JobCore " + string(e.EntryId) + "exec success , spend :" + latencyTime.String()
	//ws.SendAll(str)
	logger.Info("[Job] JobCore %s exec success , spend :%v", e.Name, latencyTime)
	return
}

// Run http任务接口
func (h *HttpJob) Run() {

	startTime := time.Now()
	var count = 0
	var err error
	var str string
	/* 循环 */
LOOP:
	if count < retryCount {
		/* 跳过迭代 */
		str, err = tools.Get(h.InvokeTarget)
		if err != nil {
			// 如果失败暂停一段时间重试
			fmt.Println(time.Now().Format(TimeFormat), " [ERROR] mission failed! ", err)
			fmt.Printf(time.Now().Format(TimeFormat)+" [INFO] Retry after the task fails %d seconds! %s \n", (count+1)*5, str)
			time.Sleep(time.Duration(count+1) * 5 * time.Second)
			count = count + 1
			goto LOOP
		}
	}
	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//TODO: 待完善部分

	logger.Infof("[Job] JobCore %s exec success , spend :%v", h.Name, latencyTime)
	return
}

// Setup 初始化
func Setup(dbs map[string]*gorm.DB) {

	fmt.Println(time.Now().Format(TimeFormat), " [INFO] JobCore Starting...")
	for k, db := range dbs {
		runtime.App.SetCrontab(k, NewWithSeconds())
		setup(k, db)
	}
}

func setup(key string, db *gorm.DB) {
	crontab := runtime.App.GetCrontabKey(key)
	sysJob := models.SysJob{}
	jobList := make([]models.SysJob, 0)
	err := sysJob.GetList(db, &jobList)
	if err != nil {
		fmt.Println(time.Now().Format(TimeFormat), " [ERROR] JobCore init error", err)
	}
	if len(jobList) == 0 {
		fmt.Println(time.Now().Format(TimeFormat), " [INFO] JobCore total:0")
	}

	_, err = sysJob.RemoveAllEntryID(db)
	if err != nil {
		fmt.Println(time.Now().Format(TimeFormat), " [ERROR] JobCore remove entry_id error", err)
	}

	for i := 0; i < len(jobList); i++ {
		if jobList[i].JobType == 1 {
			j := &HttpJob{}
			j.InvokeTarget = jobList[i].InvokeTarget
			j.CronExpression = jobList[i].CronExpression
			j.JobId = jobList[i].JobId
			j.Name = jobList[i].JobName

			sysJob.EntryId, err = AddJob(crontab, j)
		} else if jobList[i].JobType == 2 {
			j := &ExecJob{}
			j.InvokeTarget = jobList[i].InvokeTarget
			j.CronExpression = jobList[i].CronExpression
			j.JobId = jobList[i].JobId
			j.Name = jobList[i].JobName
			j.Args = jobList[i].Args
			sysJob.EntryId, err = AddJob(crontab, j)
		}
		err = sysJob.Update(db, jobList[i].JobId)
	}

	// 启动任务
	crontab.Start()
	fmt.Println(time.Now().Format(TimeFormat), " [INFO] JobCore start success.")
	// 关闭任务
	defer crontab.Stop()
	select {}
}

// AddJob (invokeTarget string, jobId int, jobName string, cronExpression string) 添加任务
func AddJob(c *cron.Cron, job Job) (int, error) {
	if job == nil {
		fmt.Println("unknown")
		return 0, nil
	}
	return job.addJob(c)
}

func (h *HttpJob) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(h.CronExpression, h)
	if err != nil {
		fmt.Println(time.Now().Format(TimeFormat), " [ERROR] JobCore AddJob error", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}

func (e *ExecJob) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(e.CronExpression, e)
	if err != nil {
		fmt.Println(time.Now().Format(TimeFormat), " [ERROR] JobCore AddJob error", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}

// Remove 移除任务
func Remove(c *cron.Cron, entryID int) chan bool {
	ch := make(chan bool)
	go func() {
		c.Remove(cron.EntryID(entryID))
		fmt.Println(time.Now().Format(TimeFormat), " [INFO] JobCore Remove success ,info entryID :", entryID)
		ch <- true
	}()
	return ch
}

// 任务停止
//func Stop() chan bool {
//	ch := make(chan bool)
//	go func() {
//		global.GADMCron.Stop()
//		ch <- true
//	}()
//	return ch
//}

func NewWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
