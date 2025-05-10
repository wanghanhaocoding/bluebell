package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func LoginHandler(c *gin.Context) {
	//1.获取请求参数及参数校验
	p := new(models.ParamsLogin)
	if err := c.ShouldBind(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		//TODO 先不加翻译校验
		//c.JSON(http.StatusOK,gin.H{
		//	"msg":removeTopStruct(errs.Translate(trans)),//翻译错误
		//})
		//使用到errs
		fmt.Println(errs)
		if errs != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": errs.Error(),
			})
		}
		return
	}
	//2.业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("login.Login failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}
