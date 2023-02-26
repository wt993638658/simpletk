package main

import (
	"context"
	"github.com/wt993638658/simpletk/cmd/feed/command"
	"github.com/wt993638658/simpletk/dal/pack"
	feed "github.com/wt993638658/simpletk/kitex_gen/feed"
	"github.com/wt993638658/simpletk/pkg/errno"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetUserFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	var uid int64 = 0
	if *req.Token != "" {
		claim, err := Jwt.ParseToken(*req.Token)
		if err != nil {
			resp = pack.BuildVideoResp(errno.ErrTokenInvalid)
			return resp, nil
		} else {
			uid = claim.Id
		}
	}

	vis, nextTime, err := command.NewGetUserFeedService(ctx).GetUserFeed(req, uid, ServiceIPAddr)
	if err != nil {
		resp = pack.BuildVideoResp(err)
		return resp, nil
	}

	resp = pack.BuildVideoResp(errno.Success)

	resp.VideoList = vis
	resp.NextTime = &nextTime
	return resp, nil
}

// GetVideoById implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	// TODO: Your code here...
	return
}
