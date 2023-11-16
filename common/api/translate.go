/*
 * @Author: 尤_Ta
 * @Date:  23:13
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:13
 */

package api

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// transInit local 通常取决于 http 请求头的 'Accept-Language'
func transInit(local string) (trans ut.Translator, err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		var o bool
		//register translate
		// 注册翻译器
		switch local {
		case "zh", "zh-CN":
			trans, o = uni.GetTranslator("zh")
			if !o {
				return nil, fmt.Errorf("uni.GetTranslator(%s) failed", "zh")
			}
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		default:
			trans, o = uni.GetTranslator("en")
			if !o {
				return nil, fmt.Errorf("uni.GetTranslator(%s) failed", "en")
			}
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}
