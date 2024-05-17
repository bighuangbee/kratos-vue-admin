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

func newSysDictType(db *gorm.DB, opts ...gen.DOOption) sysDictType {
	_sysDictType := sysDictType{}

	_sysDictType.sysDictTypeDo.UseDB(db, opts...)
	_sysDictType.sysDictTypeDo.UseModel(&model.SysDictType{})

	tableName := _sysDictType.sysDictTypeDo.TableName()
	_sysDictType.ALL = field.NewAsterisk(tableName)
	_sysDictType.DictID = field.NewInt64(tableName, "dict_id")
	_sysDictType.DictName = field.NewString(tableName, "dict_name")
	_sysDictType.DictType = field.NewString(tableName, "dict_type")
	_sysDictType.Status = field.NewInt32(tableName, "status")
	_sysDictType.CreateBy = field.NewString(tableName, "create_by")
	_sysDictType.UpdateBy = field.NewString(tableName, "update_by")
	_sysDictType.Remark = field.NewString(tableName, "remark")
	_sysDictType.CreatedAt = field.NewTime(tableName, "created_at")
	_sysDictType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysDictType.DeletedAt = field.NewField(tableName, "deleted_at")

	_sysDictType.fillFieldMap()

	return _sysDictType
}

type sysDictType struct {
	sysDictTypeDo

	ALL       field.Asterisk
	DictID    field.Int64
	DictName  field.String // 名称
	DictType  field.String // 类型
	Status    field.Int32  // 状态
	CreateBy  field.String
	UpdateBy  field.String
	Remark    field.String // 备注
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s sysDictType) Table(newTableName string) *sysDictType {
	s.sysDictTypeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDictType) As(alias string) *sysDictType {
	s.sysDictTypeDo.DO = *(s.sysDictTypeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDictType) updateTableName(table string) *sysDictType {
	s.ALL = field.NewAsterisk(table)
	s.DictID = field.NewInt64(table, "dict_id")
	s.DictName = field.NewString(table, "dict_name")
	s.DictType = field.NewString(table, "dict_type")
	s.Status = field.NewInt32(table, "status")
	s.CreateBy = field.NewString(table, "create_by")
	s.UpdateBy = field.NewString(table, "update_by")
	s.Remark = field.NewString(table, "remark")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysDictType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDictType) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["dict_id"] = s.DictID
	s.fieldMap["dict_name"] = s.DictName
	s.fieldMap["dict_type"] = s.DictType
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysDictType) clone(db *gorm.DB) sysDictType {
	s.sysDictTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDictType) replaceDB(db *gorm.DB) sysDictType {
	s.sysDictTypeDo.ReplaceDB(db)
	return s
}

type sysDictTypeDo struct{ gen.DO }

type ISysDictTypeDo interface {
	gen.SubQuery
	Debug() ISysDictTypeDo
	WithContext(ctx context.Context) ISysDictTypeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysDictTypeDo
	WriteDB() ISysDictTypeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysDictTypeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysDictTypeDo
	Not(conds ...gen.Condition) ISysDictTypeDo
	Or(conds ...gen.Condition) ISysDictTypeDo
	Select(conds ...field.Expr) ISysDictTypeDo
	Where(conds ...gen.Condition) ISysDictTypeDo
	Order(conds ...field.Expr) ISysDictTypeDo
	Distinct(cols ...field.Expr) ISysDictTypeDo
	Omit(cols ...field.Expr) ISysDictTypeDo
	Join(table schema.Tabler, on ...field.Expr) ISysDictTypeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictTypeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysDictTypeDo
	Group(cols ...field.Expr) ISysDictTypeDo
	Having(conds ...gen.Condition) ISysDictTypeDo
	Limit(limit int) ISysDictTypeDo
	Offset(offset int) ISysDictTypeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictTypeDo
	Unscoped() ISysDictTypeDo
	Create(values ...*model.SysDictType) error
	CreateInBatches(values []*model.SysDictType, batchSize int) error
	Save(values ...*model.SysDictType) error
	First() (*model.SysDictType, error)
	Take() (*model.SysDictType, error)
	Last() (*model.SysDictType, error)
	Find() ([]*model.SysDictType, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysDictType, err error)
	FindInBatches(result *[]*model.SysDictType, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysDictType) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysDictTypeDo
	Assign(attrs ...field.AssignExpr) ISysDictTypeDo
	Joins(fields ...field.RelationField) ISysDictTypeDo
	Preload(fields ...field.RelationField) ISysDictTypeDo
	FirstOrInit() (*model.SysDictType, error)
	FirstOrCreate() (*model.SysDictType, error)
	FindByPage(offset int, limit int) (result []*model.SysDictType, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysDictTypeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysDictTypeDo) Debug() ISysDictTypeDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDictTypeDo) WithContext(ctx context.Context) ISysDictTypeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDictTypeDo) ReadDB() ISysDictTypeDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDictTypeDo) WriteDB() ISysDictTypeDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDictTypeDo) Session(config *gorm.Session) ISysDictTypeDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDictTypeDo) Clauses(conds ...clause.Expression) ISysDictTypeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDictTypeDo) Returning(value interface{}, columns ...string) ISysDictTypeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDictTypeDo) Not(conds ...gen.Condition) ISysDictTypeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDictTypeDo) Or(conds ...gen.Condition) ISysDictTypeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDictTypeDo) Select(conds ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDictTypeDo) Where(conds ...gen.Condition) ISysDictTypeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDictTypeDo) Order(conds ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDictTypeDo) Distinct(cols ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDictTypeDo) Omit(cols ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDictTypeDo) Join(table schema.Tabler, on ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDictTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDictTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDictTypeDo) Group(cols ...field.Expr) ISysDictTypeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDictTypeDo) Having(conds ...gen.Condition) ISysDictTypeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDictTypeDo) Limit(limit int) ISysDictTypeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDictTypeDo) Offset(offset int) ISysDictTypeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDictTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictTypeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDictTypeDo) Unscoped() ISysDictTypeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDictTypeDo) Create(values ...*model.SysDictType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDictTypeDo) CreateInBatches(values []*model.SysDictType, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDictTypeDo) Save(values ...*model.SysDictType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDictTypeDo) First() (*model.SysDictType, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) Take() (*model.SysDictType, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) Last() (*model.SysDictType, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) Find() ([]*model.SysDictType, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysDictType), err
}

func (s sysDictTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysDictType, err error) {
	buf := make([]*model.SysDictType, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDictTypeDo) FindInBatches(result *[]*model.SysDictType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDictTypeDo) Attrs(attrs ...field.AssignExpr) ISysDictTypeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDictTypeDo) Assign(attrs ...field.AssignExpr) ISysDictTypeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDictTypeDo) Joins(fields ...field.RelationField) ISysDictTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDictTypeDo) Preload(fields ...field.RelationField) ISysDictTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDictTypeDo) FirstOrInit() (*model.SysDictType, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) FirstOrCreate() (*model.SysDictType, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) FindByPage(offset int, limit int) (result []*model.SysDictType, count int64, err error) {
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

func (s sysDictTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDictTypeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDictTypeDo) Delete(models ...*model.SysDictType) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDictTypeDo) withDO(do gen.Dao) *sysDictTypeDo {
	s.DO = *do.(*gen.DO)
	return s
}
