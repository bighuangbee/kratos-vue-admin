package middleware

import (
	"context"
	"fmt"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/service"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

//
//copy from "github.com/go-kratos/kratos/v2/middleware/logging"
//

// Redacter defines how to log an object
type Redacter interface {
	Redact() string
}

// Server is an server logging middleware.
func Server(logger log.Logger, logService *service.LogService, apiService *service.ApiService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)

			t, err := authz.FromContext(ctx)
			fmt.Println("ddddd", t, err)

			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()

				//t, err := gojwt.Parse(info.RequestHeader().Get("Authorization"),
				//	func(token *gojwt.Token) (interface{}, error) {
				//		return jwtTestDefaultKey, nil
				//	}
				//)
				//fmt.Println("=====", t, err)
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			level, stack := extractError(err)
			log.NewHelper(log.WithContext(ctx, logger)).Log(level,
				"kind", "server",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime).Seconds(),
			)

			if api := apiService.GetSysApi(operation); api != nil {

				if claims, err := authz.FromContext(ctx); err != nil {
					logger.Log(log.LevelWarn, "authz.FromContext(ctx)", err)
				} else {
					logService.Create(ctx, &model.LogOper{
						Title:        api.Description,
						BusinessType: 0,
						URL:          api.Path,
						Method:       api.Method,
						UserName:     claims.Nickname,
						UserID:       claims.UserID,
						//IP:           ip,
						Latency:      int32(time.Since(startTime).Milliseconds()),
						Resp:         "",
						Status:       code,
						ErrorMessage: "",
						CreatedAt:    time.Now(),
					})
				}

			} else {
				logger.Log(log.LevelWarn, "operation not exists", operation)
			}

			return
		}
	}
}

// Client is a client logging middleware.
func Client(logger log.Logger) middleware.Middleware {
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
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			level, stack := extractError(err)
			log.NewHelper(log.WithContext(ctx, logger)).Log(level,
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
	if redacter, ok := req.(Redacter); ok {
		return redacter.Redact()
	}
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
