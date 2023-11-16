/*
 * @Author: 尤_Ta
 * @Date:  23:12
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:12
 */

package job

import (
	"fmt"
	"time"
)

// InitJob 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	JobList = map[string]JobsExec{
		"ExamplesOne": ExamplesOne{},
		"Examples1":   Examples1{},
		"Examples2":   Examples2{},
		// ...
	}
}

// ExamplesOne 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	str := time.Now().Format(TimeFormat) + " [INFO] JobCore ExamplesOne exec success"
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
			fmt.Println(str, arg.(string))
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
		break
	}

	return nil
}

type Examples1 struct {
}

func (t1 Examples1) Exec(arg interface{}) error {

	fmt.Printf("%v--------\n", arg)
	fmt.Printf("Examples1--------\n")

	return nil
}

type Examples2 struct {
}

func (t2 Examples2) Exec(arg interface{}) error {
	fmt.Printf("%v--------\n", arg)
	fmt.Printf("Examples2--------\n")

	return nil
}
