// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoritesrv

import (
	server "github.com/cloudwego/kitex/server"
	favorite "github.com/wt993638658/simpletk/kitex_gen/favorite"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler favorite.FavoriteSrv, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
