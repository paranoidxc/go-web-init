package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//执行sql
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id
		) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

func GetPostDetailByID(id int64) (p *models.Post, err error) {
	p = new(models.Post)
	//执行sql
	sqlStr := `select 
		post_id, title, content, author_id, community_id, status, create_time
		FROM post 
		WHERE post_id = ? `

	if err = db.Get(p, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is not post detail in db")
			err = ErrorInvalidID
		}
	}

	return
}

func GetPostList(page, size int64) (post []*models.Post, err error) {
	sqlStr := `SELECT
	post_id, title, content, author_id, community_id, create_time
	FROM post
	LIMIT ?,?
	`
	post = make([]*models.Post, 0, 10)
	err = db.Select(&post, sqlStr, (page-1)*size, size)

	return
}
