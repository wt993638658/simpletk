package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wt993638658/simpletk/cmd/api/rpc"
	"github.com/wt993638658/simpletk/dal/pack"
	"github.com/wt993638658/simpletk/kitex_gen/favorite"
	"github.com/wt993638658/simpletk/pkg/errno"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var paramVar FavoriteActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.VideoId = int64(vid)
	paramVar.ActionType = int32(act)
	ctx := context.Background()
	resp, err := rpc.FavoriteAction(ctx, &favorite.DouyinFavoriteActionRequest{
		VideoId:    paramVar.VideoId,
		Token:      paramVar.Token,
		ActionType: paramVar.ActionType,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	ctx := context.Background()
	resp, err := rpc.FavoriteList(ctx, &favorite.DouyinFavoriteListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
