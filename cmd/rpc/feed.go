package rpc

import (
	"context"
	"github.com/wt993638658/simpletk/kitex_gen/feed"
	"github.com/wt993638658/simpletk/kitex_gen/feed/feedsrv"
	"github.com/wt993638658/simpletk/pkg/errno"
)

var feedClient feedsrv.Client

// Feed RPC 客户端初始化
//func initFeedRpc(Config *ttviper.Config) {
//	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
//	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
//	if err != nil {
//		panic(err)
//	}
//	ServiceName := Config.Viper.GetString("Server.Name")
//
//	p := provider.NewOpenTelemetryProvider(
//		provider.WithServiceName(ServiceName),
//		provider.WithExportEndpoint("localhost:4317"),
//		provider.WithInsecure(),
//	)
//	defer p.Shutdown(context.Background())
//
//	c, err := feedsrv.NewClient(
//		ServiceName,
//		client.WithMiddleware(middleware.CommonMiddleware),
//		client.WithInstanceMW(middleware.ClientMiddleware),
//		client.WithMuxConnection(1),                       // mux
//		client.WithRPCTimeout(30*time.Second),             // rpc timeout
//		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
//		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
//		client.WithSuite(tracing.NewClientSuite()),        // tracer
//		client.WithResolver(r),                            // resolver
//		// Please keep the same as provider.WithServiceName
//		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
//	)
//	if err != nil {
//		panic(err)
//	}
//	feedClient = c
//}

// 传递 获取视频流操作 的上下文, 并获取 RPC Server 端的响应.
func GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	resp, err = feedClient.GetUserFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
