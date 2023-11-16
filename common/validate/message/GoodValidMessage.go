/*
 * @Author: 尤_Ta
 * @Date:  22:51
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  22:51
 */

package message

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GoodValidMsg(err error, obj interface{}) error {
	// obj为结构体指针
	getObj := reflect.TypeOf(obj)
	// 断言为具体的类型，err是一个接口
	if valideErrors, ok := err.(validator.ValidationErrors); ok {
		// valideErrors 是一个存储错误验证信息的切片,错误信息可能会有多个
		for _, err := range valideErrors {
			if structFiledName, exist := getObj.Elem().FieldByName(err.Field()); exist {
				// 错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
				// err.Field()验证不同的字段,err.Tag()字段哪个tag的验证
				switch err.Field() {
				case "CatId":
					if err.Tag() == "catIdValid" {
						tagMsg := structFiledName.Tag.Get("catIdValidMsg")
						return errors.New(tagMsg)
					}
				}
			}

		}
	}
	//return err.Error()
	return err
}
