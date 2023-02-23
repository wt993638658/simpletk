package command

import (
	"context"
	"time"

	"github.com/wt993638658/simpletk/dal/db"
	"github.com/wt993638658/simpletk/dal/pack"
	"github.com/wt993638658/simpletk/kitex_gen/feed"
)

const (
	LIMIT = 30 // 单次返回最大视频数
)

type GetUserFeedService struct {
	ctx context.Context
}

// NewGetUserFeedService new GetUserFeedService
func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{ctx: ctx}
}

// GetUserFeed get feed info.
func (s *GetUserFeedService) GetUserFeed(req *feed.DouyinFeedRequest, fromID int64) (vis []*feed.Video, nextTime int64, err error) {
	videos, err := db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	if err != nil {
		return vis, nextTime, err
	}

	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
		return vis, nextTime, nil
	} else {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}

	if vis, err = pack.FeedVideos(s.ctx, videos, &fromID); err != nil {
		nextTime = time.Now().UnixMilli()
		return vis, nextTime, err
	}

	return vis, nextTime, nil
}
