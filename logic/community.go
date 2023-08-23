// Package l provides ...
package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommuntityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommuntityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
