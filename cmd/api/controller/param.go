package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 视频流 handler 输入参数
type FeedParam struct {
	LatestTime *int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `json:"token,omitempty"`       // 可选参数，登录用户设置
}

// SendResponse pack response
func SendResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}
