package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1.生成 post id
	p.ID = snowflake.GenID()

	// 2. save to db
	return mysql.CreatePost(p)
}

func GetPostDetail(id int64) (data *models.ApiPostDetail, err error) {

	//查询并拼接接口想用的数据
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(id) failed",
			zap.Int64("pid", id),
			zap.Error(err))
		return
	}

	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById((id) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetUserById((id) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return
}
