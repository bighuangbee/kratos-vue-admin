package middleware

import (
	"context"
	"encoding/json"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/service"
	"net"
	nhttp "net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type OpLog struct {
	OpLogGrpcAddr, Dir string
	Logger             log.Logger
	AppId              int32
}

func GetHostIp(transport interface{}) (ip string) {
	tr := transport.(*http.Transport)
	req := tr.Request()
	strs := strings.Split(req.Host, ":")
	if len(strs) > 0 {
		remoteIP := net.ParseIP(strs[0])
		if remoteIP == nil {
			return ""
		}
		return remoteIP.String()
	}

	return ""
}

func RequestUrl(transport interface{}) (uri, method, ip string) {
	tr := transport.(*http.Transport)
	req := tr.Request()
	method = strings.ToUpper(req.Method)
	uri = tr.PathTemplate()
	ip = GetIP(req)
	tr.ReplyHeader()
	return
}
func GetIP(r *nhttp.Request) string {
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

// CheckTokenMiddleWare Check Token middleware
func SaveOpLog(logService *service.LogService, logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				var (
					//corpId   uint64
					userId   uint64
					userName string
					//account  string
					bodyStr string
				)
				uri, method, ip := RequestUrl(tr)

				bs, err := json.Marshal(req)
				if err != nil {
					log.NewHelper(logger).Errorw("req json.Marshal err", err)
				} else {
					bodyStr = string(bs)
				}

				if uri == "/api/v1.0/account/login" {
					obj := make(map[string]interface{})
					json.Unmarshal([]byte(bodyStr), &obj)
					if _, ok := obj["password"]; ok {
						obj["password"] = "***"
					} else if _, ok := obj["token"]; ok {
						obj["token"] = "***"
					}
					if _, ok := obj["account"]; ok {
						//account = obj["account"].(string)
					}
					// 更新body string
					bs, err := json.Marshal(obj)
					if err != nil {
						bodyStr = ""
					} else {
						bodyStr = string(bs)
					}
				}

				//claims, err := authz.FromContext(ctx)

				logService.Create(ctx, &model.LogOper{
					Title:        "",
					BusinessType: 0,
					URL:          uri,
					Method:       method,
					UserName:     userName,
					UserID:       int64(userId),
					IP:           ip,
					Agent:        "",
					Latency:      0,
					Resp:         "",
					Status:       0,
					ErrorMessage: "",
					CreatedAt:    time.Now(),
				})

				return handler(ctx, req)
			}
			return nil, nil
		}
	}
}
