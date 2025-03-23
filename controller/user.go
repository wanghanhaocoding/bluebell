package controller

import (
	"bluebell/models"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	var p models.ParamsSignUp
	if err := c.ShouldBind(&p); err != nil {

	}
}
