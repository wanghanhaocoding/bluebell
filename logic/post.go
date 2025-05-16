package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1.生成post id
	id, err := snowflake.GetID()
	if err != nil {
		return
	}
	p.ID = int64(id)
	// 2.保存到数据库
	return mysql.CreatePost(p)
}
