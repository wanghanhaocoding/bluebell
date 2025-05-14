package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查询数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	//查询数据库 查找到所有的community 并返回

	detail, err := mysql.GetCommunityDetail(id)
	return detail, err
}
