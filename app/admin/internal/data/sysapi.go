package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/biz"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type sysApiRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysApiRepo(data *Data, logger log.Logger) biz.SysApiRepo {
	return &sysApiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *sysApiRepo) Create(ctx context.Context, api *model.SysAPI) error {
	q := a.data.Query(ctx).SysAPI
	return q.WithContext(ctx).Create(api)
}

func (a *sysApiRepo) Save(ctx context.Context, api *model.SysAPI) error {
	q := a.data.Query(ctx).SysAPI
	return q.WithContext(ctx).Save(api)
}

func (a *sysApiRepo) Delete(ctx context.Context, id int64) error {
	q := a.data.Query(ctx).SysAPI
	_, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Delete()
	return err
}

func (a *sysApiRepo) ListPage(ctx context.Context, page, size int32, path string, description, group string, method string) ([]*model.SysAPI, error) {
	q := a.data.Query(ctx).SysAPI
	db := q.WithContext(ctx)

	limit, offset := convertPageSize(page, size)

	if path != "" {
		db = db.Where(q.Path.Like("%" + path + "%"))
	}
	if description != "" {
		db = db.Where(q.Description.Like("%" + description + "%"))
	}
	if group != "" {
		db = db.Where(q.APIGroup.Like("%" + group + "%"))
	}
	if method != "" {
		db = db.Where(q.Method.Like("%" + method + "%"))
	}
	return db.Limit(limit).Offset(offset).Find()
}

func (a *sysApiRepo) ListPageCount(ctx context.Context) (int32, error) {
	q := a.data.Query(ctx).SysAPI
	db := q.WithContext(ctx)
	count, err := db.Count()
	return int32(count), err
}

func (a *sysApiRepo) FindAll(ctx context.Context) ([]*model.SysAPI, error) {
	q := a.data.Query(ctx).SysAPI
	return q.WithContext(ctx).Find()
}

func (a *sysApiRepo) FindByID(ctx context.Context, id int64) (*model.SysAPI, error) {
	q := a.data.Query(ctx).SysAPI
	return q.WithContext(ctx).Where(q.ID.Eq(id)).First()
}
