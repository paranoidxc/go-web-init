package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"

	"go.uber.org/zap"
)

const secret = "paranoid"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func QueryUserByID() {
}

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int

	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}

	if count > 0 {
		return ErrorUserExist
		//return errors.New("用户已存在")
	}

	return nil
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	//执行sql
	sqlStr := `insert into user(user_id, username, password) VALUES (?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`

	if err := db.Get(user, sqlStr, user.Username); err != nil {
		if err == sql.ErrNoRows {
			return ErrorUserNotExist
			//return errors.New("用户不存在")
		}
		return err
	}
	zap.L().Debug(user.Password)
	if encryptPassword(oPassword) != user.Password {
		return ErrorInvalidPassword
		//return errors.New("密码错误")
	}

	return nil
}
