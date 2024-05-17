package biz

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type SysLogRepo interface {
	Create(ctx context.Context, log *model.LogOper) error
}

type SysLogCase struct {
	repo SysLogRepo
}

func NewSysLogCase(repo SysLogRepo) *SysLogCase {
	return &SysLogCase{
		repo: repo,
	}
}

func (r *SysLogCase) Create(ctx context.Context, log *model.LogOper) error {
	return r.repo.Create(ctx, log)
}
