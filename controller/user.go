package controller

import (
	"errors"
	"net/http"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	// p := new(models.ParamSignUp) 直接指针
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		//判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg": err.Error(),
			// })
			return
		}

		ResponseErrorWithMsg(
			c,
			CodeInvalidParam,
			removeTopStruct(errs.Translate(trans)))
		// c.JSON(http.StatusOK, gin.H{
		// 	//"msg": "请求参数有误",
		// 	"msg": removeTopStruct(errs.Translate(trans)),
		// })
		return

	}
	//fmt.Println(p)
	// 用 Gin 的binding 处理
	// if len(p.Username) == 0 ||
	// 	len(p.Password) == 0 ||
	// 	len(p.RePassword) == 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }

	// 2.业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		/*
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		*/
		return
	}

	// 3.返回响应
	ResponseSuccess(c, nil)
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "success",
	// })
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))

		//判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.login() failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名密码错误",
		})
		return
	}
	ResponseSuccess(c, token)
}
