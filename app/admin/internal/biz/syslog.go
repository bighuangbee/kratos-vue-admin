package biz

import (
	"context"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type SysLogRepo interface {
	Create(ctx context.Context, log *model.SysLog) error
}

type SysLogCase struct {
	repo SysLogRepo
}

func NewSysLogCase(repo SysLogRepo) *SysLogCase {
	return &SysLogCase{
		repo: repo,
	}
}

func (r *SysLogCase) Create(ctx context.Context, log *model.SysLog) error {
	return r.repo.Create(ctx, log)
}
