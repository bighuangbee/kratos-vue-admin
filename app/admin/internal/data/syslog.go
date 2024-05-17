package data

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type syslogRepo struct {
	data *Data
}

func NewSysLogRepo(data *Data) biz.SysLogRepo {
	return &syslogRepo{
		data: data,
	}
}

func (r *syslogRepo) Create(ctx context.Context, g *model.LogOper) error {
	q := r.data.Query(ctx).LogOper
	err := q.WithContext(ctx).Clauses().Create(g)
	return err
}
