/*
 * @Author: 尤_Ta
 * @Date:  15:50
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  15:50
 */

package dto

// UserLoginReq 映射登录查库
type UserLoginReq struct {
	Username string `json:"username" form:"username" binding:"required" search:"type:exact;column:username;table:user"`
	//Password string `json:"password" form:"password" binding:"required" search:"type:exact;column:password;table:user"`
}

func (userLoginReq *UserLoginReq) LoginSearch() interface{} {
	return *userLoginReq
}

// UserLoginCaptchaReq 映射登录验证-验证码
type UserLoginCaptchaReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Code     string `json:"code" form:"code"  binding:"required"`
	UUID     string `json:"uuid" form:"uuid" binding:"required"`
}
