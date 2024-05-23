package kitGrpc

import (
	"context"
	middleware2 "github.com/bighuangbee/kratos-vue-admin/pkg/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	grpcSelector "github.com/go-kratos/kratos/v2/selector"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
	"time"
)

type GetConnOption struct {
	Endpoint     string
	Timeout      time.Duration
	Logger       log.Logger
	Discovery    registry.Discovery
	SelectFilter grpcSelector.NodeFilter
	Caller       string
}

// Deprecated: Use GetGrpcClient
func GetConn(ctx context.Context, endpoint string) (*grpc.ClientConn, error) {
	return kgrpc.DialInsecure(
		ctx,
		kgrpc.WithEndpoint(endpoint),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(
				tracing.WithTracerProvider(GetTracerProvider()),
			),
		),
		kgrpc.WithTimeout(time.Second*3),
	)
}

// GetGrpcClient 获取grpc客户端
// Timeout 默认3秒
func GetGrpcClient(ctx context.Context, opt GetConnOption) (*grpc.ClientConn, error) {
	if opt.Timeout < 1 {
		opt.Timeout = time.Second * 3
	}

	middlewares := []middleware.Middleware{
		recovery.Recovery(),
		tracing.Client(
			tracing.WithTracerProvider(GetTracerProvider()),
		),
		mmd.Client(mmd.WithConstants(metadata.New(map[string][]string{
			"caller": []string{opt.Caller},
		}))),
	}
	if opt.Logger != nil {
		middlewares = append(middlewares, middleware2.ClientLogFile(opt.Logger))
	}

	clientOpts := []kgrpc.ClientOption{
		kgrpc.WithEndpoint(opt.Endpoint),
		kgrpc.WithMiddleware(middlewares...),
		kgrpc.WithTimeout(opt.Timeout),
	}
	if opt.Discovery != nil {
		clientOpts = append(clientOpts, kgrpc.WithDiscovery(opt.Discovery))
		if opt.SelectFilter != nil {
			clientOpts = append(clientOpts, kgrpc.WithNodeFilter(opt.SelectFilter))
		}
	}
	return kgrpc.DialInsecure(ctx, clientOpts...)
}
