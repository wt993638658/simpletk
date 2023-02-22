// 定义 Hertz HTTP API 的 handler
package controller

import (
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
)

// SendResponse pack response
func SendResponse(c *gin.Context, response interface{}) {
	c.JSON(consts.StatusOK, response)
}

// 用户注册 handler 输入参数
type UserRegisterParam struct {
	UserName string `json:"username"` // 用户名
	PassWord string `json:"password"` // 用户密码
}

// 用户信息 输出参数
type UserParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

// 视频流 handler 输入参数
type FeedParam struct {
	LatestTime *int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `json:"token,omitempty"`       // 可选参数，登录用户设置
}

// 发布视频操作 handler 输入参数
type PublishActionParam struct {
	Token string `json:"token,omitempty"` // 用户鉴权token
	Data  []byte `json:"data,omitempty"`  // 视频数据
	Title string `json:"title,omitempty"` // 视频标题
}

// 点赞操作 handler 输入参数
type FavoriteActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	VideoId    int64  `json:"video_id,omitempty"`    // 视频id
	ActionType int32  `json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

// 评论操作  handler 输入参数
type CommentActionParam struct {
	UserId      int64   `json:"user_id,omitempty"`      // 用户id
	Token       string  `json:"token,omitempty"`        // 用户鉴权token
	VideoId     int64   `json:"video_id,omitempty"`     // 视频id
	ActionType  int32   `json:"action_type,omitempty"`  // 1-发布评论，2-删除评论
	CommentText *string `json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
}

// 获取评论列表 handler 输入参数
type CommentListParam struct {
	Token   string `json:"token,omitempty"`    // 用户鉴权token
	VideoId int64  `json:"video_id,omitempty"` // 视频id
}
type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
type Message struct {
	Id         int64  `json:"id,omitempty"`
	FromUserId int64  `json:"from_user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

// 关注操作 handler 输入参数
type RelationActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	ToUserId   int64  `json:"to_user_id,omitempty"`  // 对方用户id
	ActionType int32  `json:"action_type,omitempty"` // 1-关注，2-取消关注
}

// 评论操作  handler 输入参数
type MessageActionParam struct {
	UserId     int64   `json:"user_id,omitempty"`     // 用户id
	Token      string  `json:"token,omitempty"`       // 用户鉴权token
	ToUserId   int64   `json:"to_user_id,omitempty"`  // 闲聊对方id
	ActionType int32   `json:"action_type,omitempty"` // 1-发送消息
	Content    *string `json:"content,omitempty"`     // 消息内容
}

// 获取评论列表 handler 输入参数
type MessageChatParam struct {
	Token    string `json:"token,omitempty"`    // 用户鉴权token
	ToUserId int64  `json:"video_id,omitempty"` // 视频id
}
