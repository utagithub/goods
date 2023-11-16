/*
 * @Author: 尤_Ta
 * @Date:  23:37
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:37
 */

package config

import "goods/common/file"

//func NewExtend() (extend *Extend) {
//	extend = &Extend{
//		UpLoad: file.OXS{
//			"123",
//			"123",
//			"123",
//			"123",
//		},
//	}
//	return
//}

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	// 这里配置对应配置文件的结构即可
	//AMap   AMap // 这里配置对应配置文件的结构即可
	//UpLoad file.OXS `yaml:"upload"`
	UpLoad UpLoadConfig `yaml:"upload"`
}

type AMap struct {
	Key string
}

type UpLoadConfig struct {
	Secret file.OXS `yaml:"secret"`
	Driver string   `yaml:"driver"`
}
