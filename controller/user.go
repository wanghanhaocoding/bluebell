package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamsSignUp)
	if err := c.ShouldBind(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp Shouldn't Bind Parameters Error", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
