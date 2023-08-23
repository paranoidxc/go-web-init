// Package controller provides ...
package controller

import (
	"strconv"
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommuntityList()

	if err != nil {
		zap.L().Error("logic.GetCommuntityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetCommuntityDetail(id)

	if err != nil {
		zap.L().Error("logic.GetCommuntityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	ResponseSuccess(c, data)
}
