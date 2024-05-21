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

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

func newSysLog(db *gorm.DB, opts ...gen.DOOption) sysLog {
	_sysLog := sysLog{}

	_sysLog.sysLogDo.UseDB(db, opts...)
	_sysLog.sysLogDo.UseModel(&model.SysLog{})

	tableName := _sysLog.sysLogDo.TableName()
	_sysLog.ALL = field.NewAsterisk(tableName)
	_sysLog.ID = field.NewInt64(tableName, "id")
	_sysLog.Title = field.NewString(tableName, "title")
	_sysLog.BusinessType = field.NewInt32(tableName, "business_type")
	_sysLog.URL = field.NewString(tableName, "url")
	_sysLog.Operation = field.NewString(tableName, "operation")
	_sysLog.Method = field.NewString(tableName, "method")
	_sysLog.UserName = field.NewString(tableName, "user_name")
	_sysLog.UserID = field.NewInt64(tableName, "user_id")
	_sysLog.IP = field.NewString(tableName, "ip")
	_sysLog.Latency = field.NewInt32(tableName, "latency")
	_sysLog.Resp = field.NewString(tableName, "resp")
	_sysLog.Status = field.NewInt32(tableName, "status")
	_sysLog.ErrorMessage = field.NewString(tableName, "error_message")
	_sysLog.CreatedAt = field.NewTime(tableName, "created_at")

	_sysLog.fillFieldMap()

	return _sysLog
}

type sysLog struct {
	sysLogDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键id
	Title        field.String // 操作的模块
	BusinessType field.Int32  // 0其它 1新增 2修改 3删除
	URL          field.String // 操作url
	Operation    field.String // 操作路径
	Method       field.String // 请求方法
	UserName     field.String // 操作人员
	UserID       field.Int64  // 用户id
	IP           field.String // 操作IP
	Latency      field.Int32  // 延迟
	Resp         field.String // 请求参数
	Status       field.Int32  // 1=正常 2=异常
	ErrorMessage field.String // 错误信息
	CreatedAt    field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (s sysLog) Table(newTableName string) *sysLog {
	s.sysLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysLog) As(alias string) *sysLog {
	s.sysLogDo.DO = *(s.sysLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysLog) updateTableName(table string) *sysLog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Title = field.NewString(table, "title")
	s.BusinessType = field.NewInt32(table, "business_type")
	s.URL = field.NewString(table, "url")
	s.Operation = field.NewString(table, "operation")
	s.Method = field.NewString(table, "method")
	s.UserName = field.NewString(table, "user_name")
	s.UserID = field.NewInt64(table, "user_id")
	s.IP = field.NewString(table, "ip")
	s.Latency = field.NewInt32(table, "latency")
	s.Resp = field.NewString(table, "resp")
	s.Status = field.NewInt32(table, "status")
	s.ErrorMessage = field.NewString(table, "error_message")
	s.CreatedAt = field.NewTime(table, "created_at")

	s.fillFieldMap()

	return s
}

func (s *sysLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["title"] = s.Title
	s.fieldMap["business_type"] = s.BusinessType
	s.fieldMap["url"] = s.URL
	s.fieldMap["operation"] = s.Operation
	s.fieldMap["method"] = s.Method
	s.fieldMap["user_name"] = s.UserName
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["ip"] = s.IP
	s.fieldMap["latency"] = s.Latency
	s.fieldMap["resp"] = s.Resp
	s.fieldMap["status"] = s.Status
	s.fieldMap["error_message"] = s.ErrorMessage
	s.fieldMap["created_at"] = s.CreatedAt
}

func (s sysLog) clone(db *gorm.DB) sysLog {
	s.sysLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysLog) replaceDB(db *gorm.DB) sysLog {
	s.sysLogDo.ReplaceDB(db)
	return s
}

type sysLogDo struct{ gen.DO }

type ISysLogDo interface {
	gen.SubQuery
	Debug() ISysLogDo
	WithContext(ctx context.Context) ISysLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysLogDo
	WriteDB() ISysLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysLogDo
	Not(conds ...gen.Condition) ISysLogDo
	Or(conds ...gen.Condition) ISysLogDo
	Select(conds ...field.Expr) ISysLogDo
	Where(conds ...gen.Condition) ISysLogDo
	Order(conds ...field.Expr) ISysLogDo
	Distinct(cols ...field.Expr) ISysLogDo
	Omit(cols ...field.Expr) ISysLogDo
	Join(table schema.Tabler, on ...field.Expr) ISysLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysLogDo
	Group(cols ...field.Expr) ISysLogDo
	Having(conds ...gen.Condition) ISysLogDo
	Limit(limit int) ISysLogDo
	Offset(offset int) ISysLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysLogDo
	Unscoped() ISysLogDo
	Create(values ...*model.SysLog) error
	CreateInBatches(values []*model.SysLog, batchSize int) error
	Save(values ...*model.SysLog) error
	First() (*model.SysLog, error)
	Take() (*model.SysLog, error)
	Last() (*model.SysLog, error)
	Find() ([]*model.SysLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysLog, err error)
	FindInBatches(result *[]*model.SysLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysLogDo
	Assign(attrs ...field.AssignExpr) ISysLogDo
	Joins(fields ...field.RelationField) ISysLogDo
	Preload(fields ...field.RelationField) ISysLogDo
	FirstOrInit() (*model.SysLog, error)
	FirstOrCreate() (*model.SysLog, error)
	FindByPage(offset int, limit int) (result []*model.SysLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysLogDo) Debug() ISysLogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysLogDo) WithContext(ctx context.Context) ISysLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysLogDo) ReadDB() ISysLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysLogDo) WriteDB() ISysLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysLogDo) Session(config *gorm.Session) ISysLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysLogDo) Clauses(conds ...clause.Expression) ISysLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysLogDo) Returning(value interface{}, columns ...string) ISysLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysLogDo) Not(conds ...gen.Condition) ISysLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysLogDo) Or(conds ...gen.Condition) ISysLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysLogDo) Select(conds ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysLogDo) Where(conds ...gen.Condition) ISysLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysLogDo) Order(conds ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysLogDo) Distinct(cols ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysLogDo) Omit(cols ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysLogDo) Join(table schema.Tabler, on ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysLogDo) Group(cols ...field.Expr) ISysLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysLogDo) Having(conds ...gen.Condition) ISysLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysLogDo) Limit(limit int) ISysLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysLogDo) Offset(offset int) ISysLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysLogDo) Unscoped() ISysLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysLogDo) Create(values ...*model.SysLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysLogDo) CreateInBatches(values []*model.SysLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysLogDo) Save(values ...*model.SysLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysLogDo) First() (*model.SysLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLog), nil
	}
}

func (s sysLogDo) Take() (*model.SysLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLog), nil
	}
}

func (s sysLogDo) Last() (*model.SysLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLog), nil
	}
}

func (s sysLogDo) Find() ([]*model.SysLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysLog), err
}

func (s sysLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysLog, err error) {
	buf := make([]*model.SysLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysLogDo) FindInBatches(result *[]*model.SysLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysLogDo) Attrs(attrs ...field.AssignExpr) ISysLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysLogDo) Assign(attrs ...field.AssignExpr) ISysLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysLogDo) Joins(fields ...field.RelationField) ISysLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysLogDo) Preload(fields ...field.RelationField) ISysLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysLogDo) FirstOrInit() (*model.SysLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLog), nil
	}
}

func (s sysLogDo) FirstOrCreate() (*model.SysLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLog), nil
	}
}

func (s sysLogDo) FindByPage(offset int, limit int) (result []*model.SysLog, count int64, err error) {
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

func (s sysLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysLogDo) Delete(models ...*model.SysLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysLogDo) withDO(do gen.Dao) *sysLogDo {
	s.DO = *do.(*gen.DO)
	return s
}