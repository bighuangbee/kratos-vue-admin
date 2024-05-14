package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/tx7do/kratos-casbin/authz/casbin"
	"time"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

func AuthWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.admin.v1.Sysuser/Login"] = struct{}{}
	whiteList["/api.admin.v1.Sysuser/GetCaptcha"] = struct{}{}
	whiteList["/api.admin.v1.TencentCallback/TencentCallback"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func Auth(s *conf.Auth, repo biz.CasbinRuleRepo) middleware.Middleware {
	return selector.Server(
		jwt.Server(
			func(token *gojwt.Token) (interface{}, error) { return []byte(s.JwtKey), nil },
			jwt.WithSigningMethod(gojwt.SigningMethodHS256),
			jwt.WithClaims(func() gojwt.Claims { return &authz.TokenClaims{} }),
		),
		casbin.Server(
			casbin.WithCasbinModel(repo.GetModel()),
			casbin.WithCasbinPolicy(repo.GetAdapter()),
			casbin.WithSecurityUserCreator(authz.NewSecurityUser),
			casbin.WithAutoLoadPolicy(true, 30*time.Second),
		),
	).Match(AuthWhiteListMatcher()).Build()
}
