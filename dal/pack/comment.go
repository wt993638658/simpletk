package pack

import (
	"context"
	"errors"

	"github.com/wt993638658/simpletk/dal/db"
	"github.com/wt993638658/simpletk/kitex_gen/comment"
	"gorm.io/gorm"
)

// Comment pack Comments info.
func Comments(ctx context.Context, vs []*db.Comment, fromID int64) ([]*comment.Comment, error) {
	comments := make([]*comment.Comment, 0)
	for _, v := range vs {
		user, err := db.GetUserByID(ctx, int64(v.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		packUser, err := User(ctx, user, fromID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment.Comment{
			Id:         int64(v.ID),
			User:       packUser,
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	return comments, nil
}
