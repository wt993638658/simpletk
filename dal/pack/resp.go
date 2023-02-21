package pack

import (
	"errors"
	"github.com/wt993638658/simpletk/kitex_gen/feed"
	"github.com/wt993638658/simpletk/pkg/errno"
)

// BuildVideoResp build VideoResp from error
func BuildVideoResp(err error) *feed.DouyinFeedResponse {
	if err == nil {
		return videoResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return videoResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return videoResp(s)
}
func videoResp(err errno.ErrNo) *feed.DouyinFeedResponse {
	return &feed.DouyinFeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
