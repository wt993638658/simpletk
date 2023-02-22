// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentsrv

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	comment "github.com/wt993638658/simpletk/kitex_gen/comment"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentActionResponse, err error)
	CommentList(ctx context.Context, req *comment.DouyinCommentListRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentSrvClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentSrvClient struct {
	*kClient
}

func (p *kCommentSrvClient) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, req)
}

func (p *kCommentSrvClient) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, req)
}
