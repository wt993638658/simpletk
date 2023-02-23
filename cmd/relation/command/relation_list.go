package command

import (
	"context"
	"github.com/wt993638658/simpletk/dal/db"
	"github.com/wt993638658/simpletk/dal/pack"
	"github.com/wt993638658/simpletk/kitex_gen/relation"
	"github.com/wt993638658/simpletk/kitex_gen/user"
)

type FollowingListService struct {
	ctx context.Context
}

// NewFollowingListService creates a new FollowingListService
func NewFollowingListService(ctx context.Context) *FollowingListService {
	return &FollowingListService{
		ctx: ctx,
	}
}

// FollowingList returns the following lists
func (s *FollowingListService) FollowingList(req *relation.DouyinRelationFollowListRequest, fromID int64) ([]*user.User, error) {
	FollowingUser, err := db.FollowingList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.FollowingList(s.ctx, FollowingUser, fromID)
}

type FollowerListService struct {
	ctx context.Context
}

// NewFollowerListService creates a new FollowerListService
func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{
		ctx: ctx,
	}
}

// FollowerList returns the Follower Lists
func (s *FollowerListService) FollowerList(req *relation.DouyinRelationFollowerListRequest, fromID int64) ([]*user.User, error) {
	FollowerUser, err := db.FollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return pack.FollowerList(s.ctx, FollowerUser, fromID)
}
