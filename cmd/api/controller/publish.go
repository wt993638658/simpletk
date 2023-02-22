package controller

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wt993638658/simpletk/cmd/api/rpc"
	"github.com/wt993638658/simpletk/dal/pack"
	"github.com/wt993638658/simpletk/kitex_gen/publish"
	"github.com/wt993638658/simpletk/pkg/errno"
	"io"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// 传递 发布视频操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func Publish(c *gin.Context) {
	var paramVar PublishActionParam
	token := c.PostForm("token")
	title := c.PostForm("title")

	fileHeader, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ErrDecodingFailed))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ErrDecodingFailed))
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		SendResponse(c, pack.BuildPublishResp(err))
		return
	}

	paramVar.Token = token
	paramVar.Title = title

	ctx := context.Background()
	resp, err := rpc.PublishAction(ctx, &publish.DouyinPublishActionRequest{
		Title: paramVar.Title,
		Token: paramVar.Token,
		Data:  buf.Bytes(),
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}

	ctx := context.Background()
	resp, err := rpc.PublishList(ctx, &publish.DouyinPublishListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
