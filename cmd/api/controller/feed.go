package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wt993638658/simpletk/cmd/api/rpc"
	"github.com/wt993638658/simpletk/dal/pack"
	"github.com/wt993638658/simpletk/kitex_gen/feed"
	"github.com/wt993638658/simpletk/pkg/errno"
	"strconv"
)

//type FeedResponse struct {
//	Response
//	VideoList []Video `json:"video_list,omitempty"`
//	NextTime  int64   `json:"next_time,omitempty"`
//}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var feedVar FeedParam
	var laststTime int64
	var token string
	lastst_time := c.Query("latest_time")
	if len(lastst_time) != 0 {
		if latesttime, err := strconv.Atoi(lastst_time); err != nil {
			SendResponse(c, pack.BuildVideoResp(errno.ErrDecodingFailed))
			return
		} else {
			laststTime = int64(latesttime)
		}
	}

	feedVar.LatestTime = &laststTime

	token = c.Query("token")
	feedVar.Token = &token

	ctx := context.Background()
	resp, err := rpc.GetUserFeed(ctx, &feed.DouyinFeedRequest{
		LatestTime: feedVar.LatestTime,
		Token:      feedVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildVideoResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
