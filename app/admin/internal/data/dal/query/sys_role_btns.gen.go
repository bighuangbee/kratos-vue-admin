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

func newSysRoleBtn(db *gorm.DB, opts ...gen.DOOption) sysRoleBtn {
	_sysRoleBtn := sysRoleBtn{}

	_sysRoleBtn.sysRoleBtnDo.UseDB(db, opts...)
	_sysRoleBtn.sysRoleBtnDo.UseModel(&model.SysRoleBtn{})

	tableName := _sysRoleBtn.sysRoleBtnDo.TableName()
	_sysRoleBtn.ALL = field.NewAsterisk(tableName)
	_sysRoleBtn.ID = field.NewInt64(tableName, "id")
	_sysRoleBtn.RoleID = field.NewInt64(tableName, "role_id")
	_sysRoleBtn.MenuID = field.NewInt64(tableName, "menu_id")
	_sysRoleBtn.BtnID = field.NewInt64(tableName, "btn_id")

	_sysRoleBtn.fillFieldMap()

	return _sysRoleBtn
}

type sysRoleBtn struct {
	sysRoleBtnDo

	ALL    field.Asterisk
	ID     field.Int64 // 主键id
	RoleID field.Int64 // 角色ID
	MenuID field.Int64 // 菜单ID
	BtnID  field.Int64 // 菜单按钮ID

	fieldMap map[string]field.Expr
}

func (s sysRoleBtn) Table(newTableName string) *sysRoleBtn {
	s.sysRoleBtnDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysRoleBtn) As(alias string) *sysRoleBtn {
	s.sysRoleBtnDo.DO = *(s.sysRoleBtnDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysRoleBtn) updateTableName(table string) *sysRoleBtn {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.RoleID = field.NewInt64(table, "role_id")
	s.MenuID = field.NewInt64(table, "menu_id")
	s.BtnID = field.NewInt64(table, "btn_id")

	s.fillFieldMap()

	return s
}

func (s *sysRoleBtn) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysRoleBtn) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 4)
	s.fieldMap["id"] = s.ID
	s.fieldMap["role_id"] = s.RoleID
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["btn_id"] = s.BtnID
}

func (s sysRoleBtn) clone(db *gorm.DB) sysRoleBtn {
	s.sysRoleBtnDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysRoleBtn) replaceDB(db *gorm.DB) sysRoleBtn {
	s.sysRoleBtnDo.ReplaceDB(db)
	return s
}

type sysRoleBtnDo struct{ gen.DO }

type ISysRoleBtnDo interface {
	gen.SubQuery
	Debug() ISysRoleBtnDo
	WithContext(ctx context.Context) ISysRoleBtnDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysRoleBtnDo
	WriteDB() ISysRoleBtnDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysRoleBtnDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysRoleBtnDo
	Not(conds ...gen.Condition) ISysRoleBtnDo
	Or(conds ...gen.Condition) ISysRoleBtnDo
	Select(conds ...field.Expr) ISysRoleBtnDo
	Where(conds ...gen.Condition) ISysRoleBtnDo
	Order(conds ...field.Expr) ISysRoleBtnDo
	Distinct(cols ...field.Expr) ISysRoleBtnDo
	Omit(cols ...field.Expr) ISysRoleBtnDo
	Join(table schema.Tabler, on ...field.Expr) ISysRoleBtnDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysRoleBtnDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysRoleBtnDo
	Group(cols ...field.Expr) ISysRoleBtnDo
	Having(conds ...gen.Condition) ISysRoleBtnDo
	Limit(limit int) ISysRoleBtnDo
	Offset(offset int) ISysRoleBtnDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysRoleBtnDo
	Unscoped() ISysRoleBtnDo
	Create(values ...*model.SysRoleBtn) error
	CreateInBatches(values []*model.SysRoleBtn, batchSize int) error
	Save(values ...*model.SysRoleBtn) error
	First() (*model.SysRoleBtn, error)
	Take() (*model.SysRoleBtn, error)
	Last() (*model.SysRoleBtn, error)
	Find() ([]*model.SysRoleBtn, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRoleBtn, err error)
	FindInBatches(result *[]*model.SysRoleBtn, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysRoleBtn) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysRoleBtnDo
	Assign(attrs ...field.AssignExpr) ISysRoleBtnDo
	Joins(fields ...field.RelationField) ISysRoleBtnDo
	Preload(fields ...field.RelationField) ISysRoleBtnDo
	FirstOrInit() (*model.SysRoleBtn, error)
	FirstOrCreate() (*model.SysRoleBtn, error)
	FindByPage(offset int, limit int) (result []*model.SysRoleBtn, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysRoleBtnDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysRoleBtnDo) Debug() ISysRoleBtnDo {
	return s.withDO(s.DO.Debug())
}

func (s sysRoleBtnDo) WithContext(ctx context.Context) ISysRoleBtnDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysRoleBtnDo) ReadDB() ISysRoleBtnDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysRoleBtnDo) WriteDB() ISysRoleBtnDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysRoleBtnDo) Session(config *gorm.Session) ISysRoleBtnDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysRoleBtnDo) Clauses(conds ...clause.Expression) ISysRoleBtnDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysRoleBtnDo) Returning(value interface{}, columns ...string) ISysRoleBtnDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysRoleBtnDo) Not(conds ...gen.Condition) ISysRoleBtnDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysRoleBtnDo) Or(conds ...gen.Condition) ISysRoleBtnDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysRoleBtnDo) Select(conds ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysRoleBtnDo) Where(conds ...gen.Condition) ISysRoleBtnDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysRoleBtnDo) Order(conds ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysRoleBtnDo) Distinct(cols ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysRoleBtnDo) Omit(cols ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysRoleBtnDo) Join(table schema.Tabler, on ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysRoleBtnDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysRoleBtnDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysRoleBtnDo) Group(cols ...field.Expr) ISysRoleBtnDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysRoleBtnDo) Having(conds ...gen.Condition) ISysRoleBtnDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysRoleBtnDo) Limit(limit int) ISysRoleBtnDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysRoleBtnDo) Offset(offset int) ISysRoleBtnDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysRoleBtnDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysRoleBtnDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysRoleBtnDo) Unscoped() ISysRoleBtnDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysRoleBtnDo) Create(values ...*model.SysRoleBtn) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysRoleBtnDo) CreateInBatches(values []*model.SysRoleBtn, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysRoleBtnDo) Save(values ...*model.SysRoleBtn) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysRoleBtnDo) First() (*model.SysRoleBtn, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleBtn), nil
	}
}

func (s sysRoleBtnDo) Take() (*model.SysRoleBtn, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleBtn), nil
	}
}

func (s sysRoleBtnDo) Last() (*model.SysRoleBtn, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleBtn), nil
	}
}

func (s sysRoleBtnDo) Find() ([]*model.SysRoleBtn, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysRoleBtn), err
}

func (s sysRoleBtnDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRoleBtn, err error) {
	buf := make([]*model.SysRoleBtn, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysRoleBtnDo) FindInBatches(result *[]*model.SysRoleBtn, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysRoleBtnDo) Attrs(attrs ...field.AssignExpr) ISysRoleBtnDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysRoleBtnDo) Assign(attrs ...field.AssignExpr) ISysRoleBtnDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysRoleBtnDo) Joins(fields ...field.RelationField) ISysRoleBtnDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysRoleBtnDo) Preload(fields ...field.RelationField) ISysRoleBtnDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysRoleBtnDo) FirstOrInit() (*model.SysRoleBtn, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleBtn), nil
	}
}

func (s sysRoleBtnDo) FirstOrCreate() (*model.SysRoleBtn, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleBtn), nil
	}
}

func (s sysRoleBtnDo) FindByPage(offset int, limit int) (result []*model.SysRoleBtn, count int64, err error) {
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

func (s sysRoleBtnDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysRoleBtnDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysRoleBtnDo) Delete(models ...*model.SysRoleBtn) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysRoleBtnDo) withDO(do gen.Dao) *sysRoleBtnDo {
	s.DO = *do.(*gen.DO)
	return s
}
