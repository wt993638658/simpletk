package pack

import (
	"github.com/a76yyyy/tiktok/kitex_gen/message"
	"sync/atomic"

	"github.com/a76yyyy/tiktok/dal/db"
)

var messageIdSequence = int64(1)

// Comment pack Comments info.
func Messages(vs []*db.Message) ([]*message.Message, error) {

	var messageIdSequence = int64(1)
	messages := make([]*message.Message, 0)
	for _, v := range vs {
		//fmt.Println(v)
		atomic.AddInt64(&messageIdSequence, 1)
		messages = append(messages, &message.Message{
			Id:         int64(v.ID),
			ToUserId:   int64(v.ToUserID),
			FromUserId: int64(v.UserID),
			Content:    v.Content,
			CreateTime: messageIdSequence,
			//CreateTime: v.CreatedAt.Format("01-02"),
		})
	}
	return messages, nil
	//return nil, nil
}
