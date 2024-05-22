// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/data/dal/model"
)

func newSysJob(db *gorm.DB, opts ...gen.DOOption) sysJob {
	_sysJob := sysJob{}

	_sysJob.sysJobDo.UseDB(db, opts...)
	_sysJob.sysJobDo.UseModel(&model.SysJob{})

	tableName := _sysJob.sysJobDo.TableName()
	_sysJob.ALL = field.NewAsterisk(tableName)
	_sysJob.ID = field.NewInt64(tableName, "id")
	_sysJob.JobName = field.NewString(tableName, "job_name")
	_sysJob.JobGroup = field.NewString(tableName, "job_group")
	_sysJob.JobType = field.NewInt32(tableName, "job_type")
	_sysJob.CronExpression = field.NewString(tableName, "cron_expression")
	_sysJob.InvokeTarget = field.NewString(tableName, "invoke_target")
	_sysJob.Args = field.NewString(tableName, "args")
	_sysJob.MisfirePolicy = field.NewInt32(tableName, "misfire_policy")
	_sysJob.Concurrent = field.NewInt32(tableName, "concurrent")
	_sysJob.Status = field.NewInt32(tableName, "status")
	_sysJob.EntryID = field.NewInt32(tableName, "entry_id")
	_sysJob.CreateBy = field.NewString(tableName, "create_by")
	_sysJob.UpdateBy = field.NewString(tableName, "update_by")
	_sysJob.CreatedAt = field.NewTime(tableName, "created_at")
	_sysJob.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysJob.DeletedAt = field.NewField(tableName, "deleted_at")

	_sysJob.fillFieldMap()

	return _sysJob
}

type sysJob struct {
	sysJobDo

	ALL            field.Asterisk
	ID             field.Int64  // 主键id
	JobName        field.String // 任务名称
	JobGroup       field.String // 任务组
	JobType        field.Int32  // 任务类型
	CronExpression field.String // cron表达式
	InvokeTarget   field.String // 调用目标
	Args           field.String // 目标参数
	MisfirePolicy  field.Int32  // 执行策略
	Concurrent     field.Int32  // 是否并发 1=是 2=否
	Status         field.Int32  // 1=正常 2=异常
	EntryID        field.Int32  // job启动时返回的id
	CreateBy       field.String // 创建人
	UpdateBy       field.String // 更新者
	CreatedAt      field.Time   // 创建时间
	UpdatedAt      field.Time   // 更新时间
	DeletedAt      field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (s sysJob) Table(newTableName string) *sysJob {
	s.sysJobDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysJob) As(alias string) *sysJob {
	s.sysJobDo.DO = *(s.sysJobDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysJob) updateTableName(table string) *sysJob {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.JobName = field.NewString(table, "job_name")
	s.JobGroup = field.NewString(table, "job_group")
	s.JobType = field.NewInt32(table, "job_type")
	s.CronExpression = field.NewString(table, "cron_expression")
	s.InvokeTarget = field.NewString(table, "invoke_target")
	s.Args = field.NewString(table, "args")
	s.MisfirePolicy = field.NewInt32(table, "misfire_policy")
	s.Concurrent = field.NewInt32(table, "concurrent")
	s.Status = field.NewInt32(table, "status")
	s.EntryID = field.NewInt32(table, "entry_id")
	s.CreateBy = field.NewString(table, "create_by")
	s.UpdateBy = field.NewString(table, "update_by")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysJob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysJob) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 16)
	s.fieldMap["id"] = s.ID
	s.fieldMap["job_name"] = s.JobName
	s.fieldMap["job_group"] = s.JobGroup
	s.fieldMap["job_type"] = s.JobType
	s.fieldMap["cron_expression"] = s.CronExpression
	s.fieldMap["invoke_target"] = s.InvokeTarget
	s.fieldMap["args"] = s.Args
	s.fieldMap["misfire_policy"] = s.MisfirePolicy
	s.fieldMap["concurrent"] = s.Concurrent
	s.fieldMap["status"] = s.Status
	s.fieldMap["entry_id"] = s.EntryID
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysJob) clone(db *gorm.DB) sysJob {
	s.sysJobDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysJob) replaceDB(db *gorm.DB) sysJob {
	s.sysJobDo.ReplaceDB(db)
	return s
}

type sysJobDo struct{ gen.DO }

type ISysJobDo interface {
	gen.SubQuery
	Debug() ISysJobDo
	WithContext(ctx context.Context) ISysJobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysJobDo
	WriteDB() ISysJobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysJobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysJobDo
	Not(conds ...gen.Condition) ISysJobDo
	Or(conds ...gen.Condition) ISysJobDo
	Select(conds ...field.Expr) ISysJobDo
	Where(conds ...gen.Condition) ISysJobDo
	Order(conds ...field.Expr) ISysJobDo
	Distinct(cols ...field.Expr) ISysJobDo
	Omit(cols ...field.Expr) ISysJobDo
	Join(table schema.Tabler, on ...field.Expr) ISysJobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysJobDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysJobDo
	Group(cols ...field.Expr) ISysJobDo
	Having(conds ...gen.Condition) ISysJobDo
	Limit(limit int) ISysJobDo
	Offset(offset int) ISysJobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysJobDo
	Unscoped() ISysJobDo
	Create(values ...*model.SysJob) error
	CreateInBatches(values []*model.SysJob, batchSize int) error
	Save(values ...*model.SysJob) error
	First() (*model.SysJob, error)
	Take() (*model.SysJob, error)
	Last() (*model.SysJob, error)
	Find() ([]*model.SysJob, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysJob, err error)
	FindInBatches(result *[]*model.SysJob, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysJob) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysJobDo
	Assign(attrs ...field.AssignExpr) ISysJobDo
	Joins(fields ...field.RelationField) ISysJobDo
	Preload(fields ...field.RelationField) ISysJobDo
	FirstOrInit() (*model.SysJob, error)
	FirstOrCreate() (*model.SysJob, error)
	FindByPage(offset int, limit int) (result []*model.SysJob, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysJobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysJobDo) Debug() ISysJobDo {
	return s.withDO(s.DO.Debug())
}

func (s sysJobDo) WithContext(ctx context.Context) ISysJobDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysJobDo) ReadDB() ISysJobDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysJobDo) WriteDB() ISysJobDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysJobDo) Session(config *gorm.Session) ISysJobDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysJobDo) Clauses(conds ...clause.Expression) ISysJobDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysJobDo) Returning(value interface{}, columns ...string) ISysJobDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysJobDo) Not(conds ...gen.Condition) ISysJobDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysJobDo) Or(conds ...gen.Condition) ISysJobDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysJobDo) Select(conds ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysJobDo) Where(conds ...gen.Condition) ISysJobDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysJobDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysJobDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysJobDo) Order(conds ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysJobDo) Distinct(cols ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysJobDo) Omit(cols ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysJobDo) Join(table schema.Tabler, on ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysJobDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysJobDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysJobDo) Group(cols ...field.Expr) ISysJobDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysJobDo) Having(conds ...gen.Condition) ISysJobDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysJobDo) Limit(limit int) ISysJobDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysJobDo) Offset(offset int) ISysJobDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysJobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysJobDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysJobDo) Unscoped() ISysJobDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysJobDo) Create(values ...*model.SysJob) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysJobDo) CreateInBatches(values []*model.SysJob, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysJobDo) Save(values ...*model.SysJob) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysJobDo) First() (*model.SysJob, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) Take() (*model.SysJob, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) Last() (*model.SysJob, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) Find() ([]*model.SysJob, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysJob), err
}

func (s sysJobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysJob, err error) {
	buf := make([]*model.SysJob, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysJobDo) FindInBatches(result *[]*model.SysJob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysJobDo) Attrs(attrs ...field.AssignExpr) ISysJobDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysJobDo) Assign(attrs ...field.AssignExpr) ISysJobDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysJobDo) Joins(fields ...field.RelationField) ISysJobDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysJobDo) Preload(fields ...field.RelationField) ISysJobDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysJobDo) FirstOrInit() (*model.SysJob, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) FirstOrCreate() (*model.SysJob, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) FindByPage(offset int, limit int) (result []*model.SysJob, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysJobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysJobDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysJobDo) Delete(models ...*model.SysJob) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysJobDo) withDO(do gen.Dao) *sysJobDo {
	s.DO = *do.(*gen.DO)
	return s
}
