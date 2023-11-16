/*
 * @Author: 尤_Ta
 * @Date:  19:16
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  19:16
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"goods/common/api"
	"goods/common/captcha"
)

type CapchaController struct {
	api.Api
}

// GenerateCaptchaHandler 获取验证码
// @Summary 获取验证码
// @Description 获取验证码

func (e CapchaController) GenerateCaptchaHandler(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Error(500, err, "服务初始化失败！")
		return
	}
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		e.Logger.Errorf("DriverDigitFunc error, %s", err.Error())
		e.Error(500, err, "验证码获取失败")
		return
	}
	e.Custom(gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
