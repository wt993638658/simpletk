package command

import (
	"context"

	"github.com/wt993638658/simpletk/dal/db"
	"github.com/wt993638658/simpletk/kitex_gen/favorite"
	"github.com/wt993638658/simpletk/pkg/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction action favorite.
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	// 1-点赞
	if req.ActionType == 1 {
		return db.Favorite(s.ctx, req.UserId, req.VideoId)
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		return db.DisFavorite(s.ctx, req.UserId, req.VideoId)
	}
	return errno.ErrBind
}
