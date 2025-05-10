package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"errors"
)

func SignUp(p *models.ParamsSignUp) (err error) {
	//1.判断用户存不存在
	mysql.CheckUserExist(p.Username)
	//2.生成UID
	userID, err := snowflake.GetID()
	if err != nil {
		return errors.New("用户已存在")
	}
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamsLogin) error {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
