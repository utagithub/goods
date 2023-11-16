/*
 * @Author: 尤_Ta
 * @Date:  22:01
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  22:01
 */

package job

import "github.com/robfig/cron/v3"

type Job interface {
	Run()
	addJob(*cron.Cron) (int, error)
}

type JobsExec interface {
	Exec(arg interface{}) error
}

func CallExec(e JobsExec, arg interface{}) error {
	return e.Exec(arg)
}
