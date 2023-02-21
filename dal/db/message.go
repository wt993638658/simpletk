package db

import (
	"context"

	"gorm.io/gorm"
)

// Comment 属于 Video, VideoID是外键(belongs to)
// Comment 也属于 User,UserID是外键(belongs to)

type Message struct {
	gorm.Model
	UserID   int    `gorm:"index:idx_userid;not null"`
	ToUserID int    `gorm:"index:idx_userid;not null"`
	Content  string `gorm:"type:varchar(255);not null"`
}

func (Message) TableName() string {
	return "message"
}

// NewComment creates a new Comment
func NewMessage(ctx context.Context, message *Message) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 1. 新增评论数据
		err := tx.Create(message).Error
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// DelComment deletes a comment from the database.
func DelMessage(ctx context.Context, messageID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		message := new(Message)
		if err := tx.First(&message, messageID).Error; err != nil {
			return err
		}

		// 1. 删除评论数据
		// 因为 Comment中包含了gorm.Model所以拥有软删除能力
		// 而tx.Unscoped().Delete()将永久删除记录
		err := tx.Unscoped().Delete(&message).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetVideoComments returns a list of video comments.
func GetMessage(ctx context.Context, uid, tuid int64) ([]*Message, error) {
	var messages []*Message
	err := DB.WithContext(ctx).Model(&Message{}).Where("user_id = ? AND to_user_id = ?", uid, tuid).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
