// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package fb

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/fb0a/fb_model/model"
)

func newFbSm(db *gorm.DB, opts ...gen.DOOption) fbSm {
	_fbSm := fbSm{}

	_fbSm.fbSmDo.UseDB(db, opts...)
	_fbSm.fbSmDo.UseModel(&model.FbSm{})

	tableName := _fbSm.fbSmDo.TableName()
	_fbSm.ALL = field.NewAsterisk(tableName)
	_fbSm.ID = field.NewInt32(tableName, "id")
	_fbSm.Name = field.NewString(tableName, "name")
	_fbSm.Config = field.NewString(tableName, "config")
	_fbSm.State = field.NewInt32(tableName, "state")

	_fbSm.fillFieldMap()

	return _fbSm
}

type fbSm struct {
	fbSmDo

	ALL    field.Asterisk
	ID     field.Int32
	Name   field.String
	Config field.String
	State  field.Int32 // 1:开启 2:关闭

	fieldMap map[string]field.Expr
}

func (f fbSm) Table(newTableName string) *fbSm {
	f.fbSmDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSm) As(alias string) *fbSm {
	f.fbSmDo.DO = *(f.fbSmDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSm) updateTableName(table string) *fbSm {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.Name = field.NewString(table, "name")
	f.Config = field.NewString(table, "config")
	f.State = field.NewInt32(table, "state")

	f.fillFieldMap()

	return f
}

func (f *fbSm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSm) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 4)
	f.fieldMap["id"] = f.ID
	f.fieldMap["name"] = f.Name
	f.fieldMap["config"] = f.Config
	f.fieldMap["state"] = f.State
}

func (f fbSm) clone(db *gorm.DB) fbSm {
	f.fbSmDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSm) replaceDB(db *gorm.DB) fbSm {
	f.fbSmDo.ReplaceDB(db)
	return f
}

type fbSmDo struct{ gen.DO }

type IFbSmDo interface {
	gen.SubQuery
	Debug() IFbSmDo
	WithContext(ctx context.Context) IFbSmDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSmDo
	WriteDB() IFbSmDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSmDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSmDo
	Not(conds ...gen.Condition) IFbSmDo
	Or(conds ...gen.Condition) IFbSmDo
	Select(conds ...field.Expr) IFbSmDo
	Where(conds ...gen.Condition) IFbSmDo
	Order(conds ...field.Expr) IFbSmDo
	Distinct(cols ...field.Expr) IFbSmDo
	Omit(cols ...field.Expr) IFbSmDo
	Join(table schema.Tabler, on ...field.Expr) IFbSmDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSmDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSmDo
	Group(cols ...field.Expr) IFbSmDo
	Having(conds ...gen.Condition) IFbSmDo
	Limit(limit int) IFbSmDo
	Offset(offset int) IFbSmDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSmDo
	Unscoped() IFbSmDo
	Create(values ...*model.FbSm) error
	CreateInBatches(values []*model.FbSm, batchSize int) error
	Save(values ...*model.FbSm) error
	First() (*model.FbSm, error)
	Take() (*model.FbSm, error)
	Last() (*model.FbSm, error)
	Find() ([]*model.FbSm, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSm, err error)
	FindInBatches(result *[]*model.FbSm, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSm) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSmDo
	Assign(attrs ...field.AssignExpr) IFbSmDo
	Joins(fields ...field.RelationField) IFbSmDo
	Preload(fields ...field.RelationField) IFbSmDo
	FirstOrInit() (*model.FbSm, error)
	FirstOrCreate() (*model.FbSm, error)
	FindByPage(offset int, limit int) (result []*model.FbSm, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSmDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSmDo) Debug() IFbSmDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSmDo) WithContext(ctx context.Context) IFbSmDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSmDo) ReadDB() IFbSmDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSmDo) WriteDB() IFbSmDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSmDo) Session(config *gorm.Session) IFbSmDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSmDo) Clauses(conds ...clause.Expression) IFbSmDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSmDo) Returning(value interface{}, columns ...string) IFbSmDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSmDo) Not(conds ...gen.Condition) IFbSmDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSmDo) Or(conds ...gen.Condition) IFbSmDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSmDo) Select(conds ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSmDo) Where(conds ...gen.Condition) IFbSmDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSmDo) Order(conds ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSmDo) Distinct(cols ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSmDo) Omit(cols ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSmDo) Join(table schema.Tabler, on ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSmDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSmDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSmDo) Group(cols ...field.Expr) IFbSmDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSmDo) Having(conds ...gen.Condition) IFbSmDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSmDo) Limit(limit int) IFbSmDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSmDo) Offset(offset int) IFbSmDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSmDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSmDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSmDo) Unscoped() IFbSmDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSmDo) Create(values ...*model.FbSm) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSmDo) CreateInBatches(values []*model.FbSm, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSmDo) Save(values ...*model.FbSm) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSmDo) First() (*model.FbSm, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSm), nil
	}
}

func (f fbSmDo) Take() (*model.FbSm, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSm), nil
	}
}

func (f fbSmDo) Last() (*model.FbSm, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSm), nil
	}
}

func (f fbSmDo) Find() ([]*model.FbSm, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSm), err
}

func (f fbSmDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSm, err error) {
	buf := make([]*model.FbSm, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSmDo) FindInBatches(result *[]*model.FbSm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSmDo) Attrs(attrs ...field.AssignExpr) IFbSmDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSmDo) Assign(attrs ...field.AssignExpr) IFbSmDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSmDo) Joins(fields ...field.RelationField) IFbSmDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSmDo) Preload(fields ...field.RelationField) IFbSmDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSmDo) FirstOrInit() (*model.FbSm, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSm), nil
	}
}

func (f fbSmDo) FirstOrCreate() (*model.FbSm, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSm), nil
	}
}

func (f fbSmDo) FindByPage(offset int, limit int) (result []*model.FbSm, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fbSmDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSmDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSmDo) Delete(models ...*model.FbSm) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSmDo) withDO(do gen.Dao) *fbSmDo {
	f.DO = *do.(*gen.DO)
	return f
}
