package controller

import (
	"strconv"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post) //指针
	// post.go binding tag
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Post with invalid param", zap.Error(err))
		//判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(
			c,
			CodeInvalidParam,
			removeTopStruct(errs.Translate(trans)))
		return
	}

	userID, _ := getCurrentUser(c)
	p.AuthorID = userID
	//fmt.Printf("TTTT %v", p)

	// 2.业务处理
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.SignUp() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostDetail(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	ResponseSuccess(c, data)
}
