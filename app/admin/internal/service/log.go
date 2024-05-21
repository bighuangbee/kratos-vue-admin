package service

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type LogService struct {
	rc *biz.SysLogCase
}

func NewLogService(rc *biz.SysLogCase) *LogService {
	return &LogService{
		rc: rc,
	}
}

func (r *LogService) Create(ctx context.Context, req *model.SysLog) error {
	return r.rc.Create(ctx, req)
}
