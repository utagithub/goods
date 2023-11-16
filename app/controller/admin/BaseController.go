package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goods/common/utils"
	"net/http"
)

type BaseController struct {
}

func (base BaseController) Index(c *gin.Context) {
	tt, _ := utils.LocalTime{}.MarshalJSON()
	fmt.Printf("%v\n", string(tt))
}

func (base BaseController) SuccessList(c *gin.Context, msg string, data interface{}, pageInfo map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": gin.H{
			"listInfo": data,
			"pageInfo": pageInfo,
		},
	})
}

func (base BaseController) Success(c *gin.Context, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}

func (base BaseController) Error(c *gin.Context, msg interface{}, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  msg,
		"data": data,
	})
}
