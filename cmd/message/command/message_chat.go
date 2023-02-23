package command

import (
	"context"
	"github.com/wt993638658/simpletk/kitex_gen/message"

	"github.com/wt993638658/simpletk/dal/db"
	"github.com/wt993638658/simpletk/dal/pack"
)

type MessageChatService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{
		ctx: ctx,
	}
}

// CommentList return comment list
func (s *MessageChatService) MessageChat(req *message.DouyinMessageChatRequest, uid int64) ([]*message.Message, error) {
	Messages, err := db.GetMessage(s.ctx, uid, req.ToUserId)
	if err != nil {
		return nil, err
	}

	messages, err := pack.Messages(Messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
