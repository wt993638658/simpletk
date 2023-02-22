package pack

import (
	"context"
	"errors"
	"github.com/wt993638658/simpletk/kitex_gen/comment"

	"github.com/wt993638658/simpletk/dal/db"
	"gorm.io/gorm"
)

// User pack user info
// db.User结构体是用于和数据库交互的
// user.User结构体是用于RPC传输信息的
func User(ctx context.Context, u *db.User, fromID int64) (*comment.User, error) {
	if u == nil {
		return &comment.User{
			Name: "已注销用户",
		}, nil
	}

	follow_count := int64(u.FollowingCount)
	follower_count := int64(u.FollowerCount)

	// true->fromID已关注u.ID，false-fromID未关注u.ID
	isFollow := false
	relation, err := db.GetRelation(ctx, fromID, int64(u.ID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if relation != nil {
		isFollow = true
	}
	return &comment.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   &follow_count,
		FollowerCount: &follower_count,
		IsFollow:      isFollow,
	}, nil
}

// Users pack list of user info
func Users(ctx context.Context, us []*db.User, fromID int64) ([]*comment.User, error) {
	users := make([]*comment.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u, fromID)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
