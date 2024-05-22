package grpc_client

import (
	"context"
	"fmt"
	"strconv"
	"time"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metrics"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/prometheus/client_golang/prometheus"
)

type (
	JwtToken     struct{}
	RefreshToken struct{}
)

type metricsOpt struct {
	// counter: <client/server>_requests_code_total{kind, operation, code, reason, caller}
	requests metrics.Counter
	// histogram: <client/server>_requests_seconds_bucket{kind, operation, caller}
	seconds metrics.Observer
}

type MetricsOpt func(*metricsOpt)

func WithRequests(c *prometheus.CounterVec) MetricsOpt {
	return func(o *metricsOpt) {
		o.requests = prom.NewCounter(c)
	}
}

func WithSeconds(v *prometheus.HistogramVec) MetricsOpt {
	return func(o *metricsOpt) {
		o.seconds = prom.NewHistogram(v)
	}
}

// ServerMetrics server性能监控中间件
func ServerMetrics(opts ...MetricsOpt) middleware.Middleware {
	op := metricsOpt{}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				code      int
				reason    string
				kind      string
				caller    string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				caller = info.RequestHeader().Get("caller")
				operation = info.Operation()
			}
			reply, err := handler(ctx, req)
			if se := kerrors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}
			if op.requests != nil {
				op.requests.With(kind, operation, strconv.Itoa(code), reason, caller).Inc()
			}
			if op.seconds != nil {
				op.seconds.With(kind, operation, caller).Observe(time.Since(startTime).Seconds())
			}
			return reply, err
		}
	}
}

// ClientMetrics 客户端调用性能监控中间件
func ClientMetrics(opts ...MetricsOpt) middleware.Middleware {
	op := metricsOpt{}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				code      int
				reason    string
				kind      string
				operation string
				caller    string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err := handler(ctx, req)
			if se := kerrors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}
			if op.requests != nil {
				op.requests.With(kind, operation, strconv.Itoa(code), reason, caller).Inc()
			}
			if op.seconds != nil {
				op.seconds.With(kind, operation, caller).Observe(time.Since(startTime).Seconds())
			}
			return reply, err
		}
	}
}

// Server is an server logging middleware.
func KServerLog(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := kerrors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			_, stack := extractError(err)
			_ = log.WithContext(ctx, logger).Log(log.LevelInfo,
				"kind", "server",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}

// Client is an client logging middleware.
func KClientLog(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := kerrors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			_, stack := extractError(err)
			_ = log.WithContext(ctx, logger).Log(log.LevelInfo,
				"kind", "client",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

// extractError returns the string of the error
func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}
