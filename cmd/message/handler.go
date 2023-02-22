package main

import (
	"context"
	message "github.com/wt993638658/simpletk/kitex_gen/message"
)

// MessageSrvImpl implements the last service interface defined in the IDL.
type MessageSrvImpl struct{}

// MessageAction implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageCharResponse, err error) {
	// TODO: Your code here...
	return
}
