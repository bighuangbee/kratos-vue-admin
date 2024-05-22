package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/pkg/authz"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/service"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"net"
	nethttp "net/http"
	"strings"
	"time"
)

// CheckTokenMiddleWare Check Token middleware
func LogFileDB(logger log.Logger, logService *service.LogService, apiService *service.ApiService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				var (
					code      int32
					reason    string
					kind      string
					operation string
					ip        string
				)

				startTime := time.Now()
				kind = tr.Kind().String()
				operation = tr.Operation()

				httpTr := tr.(*http.Transport)
				ip = getIP(httpTr.Request())

				reply, err = handler(ctx, req)
				if se := errors.FromError(err); se != nil {
					code = se.Code
					reason = se.Reason
				}

				body, _ := json.Marshal(req)

				logHelper(ctx, err, logger,
					"kind", "server",
					"component", kind,
					"url", httpTr.Request().URL.String(),
					"operation", operation,
					"method", httpTr.Request().Method,
					"body", string(body),
					"IP", ip,
					"code", code,
					"reason", reason,
					"latency", fmt.Sprintf("%dms", time.Since(startTime).Milliseconds()))

				if api := apiService.GetSysApi(operation); api != nil {

					if claims, err := authz.FromContext(ctx); err != nil {
						logger.Log(log.LevelWarn, "authz.FromContext(ctx)", err)
					} else {
						logService.Create(ctx, &model.SysLog{
							Title:        api.Description,
							BusinessType: 0,
							URL:          httpTr.Request().URL.String(),
							Method:       httpTr.Request().Method,
							UserName:     claims.Nickname,
							UserID:       claims.UserID,
							IP:           ip,
							Latency:      int32(time.Since(startTime).Milliseconds()),
							Resp:         "",
							Status:       code,
							ErrorMessage: reason,
							CreatedAt:    time.Now(),
						})
					}

				} else {
					logger.Log(log.LevelWarn, "operation not exists", operation)
				}
			}
			return
		}
	}
}

func getIP(r *nethttp.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return ""
	}
	remoteIP := net.ParseIP(ip)
	if remoteIP == nil {
		return ""
	}
	return remoteIP.String()
}

func logHelper(ctx context.Context, err error, logger log.Logger, keyvals ...interface{}) {
	helper := log.NewHelper(log.WithContext(ctx, logger))
	if err != nil {
		helper.Errorw(append(keyvals, "stack", fmt.Sprintf("%+v", err))...)
	} else {
		helper.Infow(keyvals...)
	}
}
