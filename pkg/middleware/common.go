package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/wt993638658/simpletk/pkg/dlog"
)

var _ endpoint.Middleware = CommonMiddleware

func init() {
	var logger = dlog.InitLog(3)
	defer logger.Sync()

	klog.SetLogger(logger)
}

// CommonMiddleware common middleware print some rpc info、real request and real response
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		klog.Debugf("real request: %+v", req)
		// get remote service information
		klog.Debugf("remote service name: %s, remote method: %s", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v", resp)
		return nil
	}
}
