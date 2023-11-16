/*
 * @Author: 尤_Ta
 * @Date:  22:00
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  22:00
 */

package register

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type GoodValidRrgister struct {
}

// goodValidRrgister 验证器
func (g GoodValidRrgister) goodValidRrgister() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("catIdValid", catIdValid)
	}
}

// cat_id字段catIdValid自定义验证
func catIdValid(fl validator.FieldLevel) bool {
	name := fl.Field().Interface().(int)
	var inSomeArray = []int{1, 2, 3}
	for _, value := range inSomeArray {
		if name == value {
			return true
		}
	}
	return false
}
