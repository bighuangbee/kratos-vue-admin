package server

import (
	"context"
	v1 "github.com/bighuangbee/kratos-vue-admin/api/admin/v1"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/conf"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/service"
	pkgMiddleware "github.com/bighuangbee/kratos-vue-admin/pkg/middleware"
	KratosErrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, sysuser *service.SysuserService, logger log.Logger) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			pkgMiddleware.ClientLogFile(logger),
			validate(),
			CheckTokenMiddleWareGrpc(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterSysuserServer(srv, sysuser)

	return srv
}

type Validate interface {
	Validate() error
}

func validate(opts ...recovery.Option) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if _, ok := req.(emptypb.Empty); ok {
				return handler(ctx, req)
			}
			if _, ok := req.(*emptypb.Empty); ok {
				return handler(ctx, req)
			}
			r := req.(Validate)
			if err := r.Validate(); err != nil {
				return nil, KratosErrors.New(http.StatusInternalServerError, "ParamsValidateError", "参数校验错误")
			}
			return handler(ctx, req)
		}
	}
}

// grpc,注册空context值
func CheckTokenMiddleWareGrpc() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			ctx = context.WithValue(ctx, RequestType{}, RequestType{Type: "grpc"})
			return handler(ctx, req)
		}
	}
}

type RequestType struct{ Type string }

type PbErrorReason interface {
	String() string
}
type PbError struct {
	Reason string
}

func (receiver PbError) String() string {
	return receiver.Reason
}

// Deprecated: ReturnErr 返回Error
// 只能返回 500 错误
func ReturnErr(ctx context.Context, reason PbErrorReason, args ...interface{}) *KratosErrors.Error {
	msgId := reason.String()
	return KratosErrors.New(http.StatusInternalServerError, msgId, reason.String())
}

type CustomeErr struct {
	Err error
	// 例 "repo.AddOrgTreeNode CustomeErr"
	ErrTitle string
	//ErrInfo string
}

func (t CustomeErr) Error() string {
	if t.Err != nil {
		return t.Err.Error()
	}
	return ""
}
