package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const CtxUserIDKey = "userID"

// getCurrentUser 获取当前登录的用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID1, ok := uid.(uint64)
	userID = int64(userID1)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
