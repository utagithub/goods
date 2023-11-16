package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"goods/app/services"
	"goods/app/services/dto"
	"goods/common/api"
	"goods/common/captcha"
	"goods/common/jwt"
	"goods/models"
	"net/http"
)

type LoginController struct {
	api.Api
}

func (e LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/index.tmpl", nil)
}

// DoLogin 原user表登录
func (e LoginController) DoLogin(c *gin.Context) {
	//验证码参数绑定及验证
	reqCaptcha := dto.UserLoginCaptchaReq{}
	s := services.LoginServices{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&reqCaptcha, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if ok := captcha.Verify(reqCaptcha.UUID, reqCaptcha.Code, false); ok == true {
		// 数据库验证
		reqDb := dto.UserLoginReq{
			reqCaptcha.Username,
		}
		var user models.User
		err = s.Login(&reqDb, &user)
		if err != nil {
			e.Logger.Error(c, "Login Find data error: ", err.Error())
			return
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqCaptcha.Username))
			if err != nil {
				e.Logger.Error(c, "Login Password error", err.Error())
				return
			}
			tokenInfo := jwt.GenToken(user.Username)
			//e.OK(tokenInfo, "success")
			c.HTML(http.StatusOK, "admin/index/index.tmpl", tokenInfo)
		}
	} else {
		e.Logger.Error(c, "Cpatcha verification error", false)
	}

}
