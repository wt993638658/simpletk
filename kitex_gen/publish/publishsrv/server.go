// Code generated by Kitex v0.4.4. DO NOT EDIT.
package publishsrv

import (
	server "github.com/cloudwego/kitex/server"
	publish "github.com/wt993638658/simpletk/kitex_gen/publish"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler publish.PublishSrv, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}