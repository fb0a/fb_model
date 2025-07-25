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

func newFbReportMemberGt(db *gorm.DB, opts ...gen.DOOption) fbReportMemberGt {
	_fbReportMemberGt := fbReportMemberGt{}

	_fbReportMemberGt.fbReportMemberGtDo.UseDB(db, opts...)
	_fbReportMemberGt.fbReportMemberGtDo.UseModel(&model.FbReportMemberGt{})

	tableName := _fbReportMemberGt.fbReportMemberGtDo.TableName()
	_fbReportMemberGt.ALL = field.NewAsterisk(tableName)
	_fbReportMemberGt.ID = field.NewInt64(tableName, "id")
	_fbReportMemberGt.Day = field.NewInt32(tableName, "day")
	_fbReportMemberGt.UID = field.NewInt64(tableName, "uid")
	_fbReportMemberGt.Gt = field.NewInt32(tableName, "gt")
	_fbReportMemberGt.Phone = field.NewString(tableName, "phone")
	_fbReportMemberGt.Username = field.NewString(tableName, "username")
	_fbReportMemberGt.ValidBetAmount = field.NewFloat64(tableName, "valid_bet_amount")
	_fbReportMemberGt.SettleAmount = field.NewFloat64(tableName, "settle_amount")
	_fbReportMemberGt.Ggr = field.NewFloat64(tableName, "ggr")
	_fbReportMemberGt.CreatedAt = field.NewInt64(tableName, "created_at")

	_fbReportMemberGt.fillFieldMap()

	return _fbReportMemberGt
}

// fbReportMemberGt 会员游戏分类日报表
type fbReportMemberGt struct {
	fbReportMemberGtDo

	ALL            field.Asterisk
	ID             field.Int64
	Day            field.Int32   // 日期
	UID            field.Int64   // 会员ID
	Gt             field.Int32   // 游戏类型
	Phone          field.String  // 手机号
	Username       field.String  // 用户名
	ValidBetAmount field.Float64 // 投注额
	SettleAmount   field.Float64 // 派彩额
	Ggr            field.Float64 // GGR
	CreatedAt      field.Int64

	fieldMap map[string]field.Expr
}

func (f fbReportMemberGt) Table(newTableName string) *fbReportMemberGt {
	f.fbReportMemberGtDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbReportMemberGt) As(alias string) *fbReportMemberGt {
	f.fbReportMemberGtDo.DO = *(f.fbReportMemberGtDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbReportMemberGt) updateTableName(table string) *fbReportMemberGt {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.Day = field.NewInt32(table, "day")
	f.UID = field.NewInt64(table, "uid")
	f.Gt = field.NewInt32(table, "gt")
	f.Phone = field.NewString(table, "phone")
	f.Username = field.NewString(table, "username")
	f.ValidBetAmount = field.NewFloat64(table, "valid_bet_amount")
	f.SettleAmount = field.NewFloat64(table, "settle_amount")
	f.Ggr = field.NewFloat64(table, "ggr")
	f.CreatedAt = field.NewInt64(table, "created_at")

	f.fillFieldMap()

	return f
}

func (f *fbReportMemberGt) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbReportMemberGt) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 10)
	f.fieldMap["id"] = f.ID
	f.fieldMap["day"] = f.Day
	f.fieldMap["uid"] = f.UID
	f.fieldMap["gt"] = f.Gt
	f.fieldMap["phone"] = f.Phone
	f.fieldMap["username"] = f.Username
	f.fieldMap["valid_bet_amount"] = f.ValidBetAmount
	f.fieldMap["settle_amount"] = f.SettleAmount
	f.fieldMap["ggr"] = f.Ggr
	f.fieldMap["created_at"] = f.CreatedAt
}

func (f fbReportMemberGt) clone(db *gorm.DB) fbReportMemberGt {
	f.fbReportMemberGtDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbReportMemberGt) replaceDB(db *gorm.DB) fbReportMemberGt {
	f.fbReportMemberGtDo.ReplaceDB(db)
	return f
}

type fbReportMemberGtDo struct{ gen.DO }

type IFbReportMemberGtDo interface {
	gen.SubQuery
	Debug() IFbReportMemberGtDo
	WithContext(ctx context.Context) IFbReportMemberGtDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbReportMemberGtDo
	WriteDB() IFbReportMemberGtDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbReportMemberGtDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbReportMemberGtDo
	Not(conds ...gen.Condition) IFbReportMemberGtDo
	Or(conds ...gen.Condition) IFbReportMemberGtDo
	Select(conds ...field.Expr) IFbReportMemberGtDo
	Where(conds ...gen.Condition) IFbReportMemberGtDo
	Order(conds ...field.Expr) IFbReportMemberGtDo
	Distinct(cols ...field.Expr) IFbReportMemberGtDo
	Omit(cols ...field.Expr) IFbReportMemberGtDo
	Join(table schema.Tabler, on ...field.Expr) IFbReportMemberGtDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberGtDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberGtDo
	Group(cols ...field.Expr) IFbReportMemberGtDo
	Having(conds ...gen.Condition) IFbReportMemberGtDo
	Limit(limit int) IFbReportMemberGtDo
	Offset(offset int) IFbReportMemberGtDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbReportMemberGtDo
	Unscoped() IFbReportMemberGtDo
	Create(values ...*model.FbReportMemberGt) error
	CreateInBatches(values []*model.FbReportMemberGt, batchSize int) error
	Save(values ...*model.FbReportMemberGt) error
	First() (*model.FbReportMemberGt, error)
	Take() (*model.FbReportMemberGt, error)
	Last() (*model.FbReportMemberGt, error)
	Find() ([]*model.FbReportMemberGt, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbReportMemberGt, err error)
	FindInBatches(result *[]*model.FbReportMemberGt, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbReportMemberGt) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbReportMemberGtDo
	Assign(attrs ...field.AssignExpr) IFbReportMemberGtDo
	Joins(fields ...field.RelationField) IFbReportMemberGtDo
	Preload(fields ...field.RelationField) IFbReportMemberGtDo
	FirstOrInit() (*model.FbReportMemberGt, error)
	FirstOrCreate() (*model.FbReportMemberGt, error)
	FindByPage(offset int, limit int) (result []*model.FbReportMemberGt, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbReportMemberGtDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbReportMemberGtDo) Debug() IFbReportMemberGtDo {
	return f.withDO(f.DO.Debug())
}

func (f fbReportMemberGtDo) WithContext(ctx context.Context) IFbReportMemberGtDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbReportMemberGtDo) ReadDB() IFbReportMemberGtDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbReportMemberGtDo) WriteDB() IFbReportMemberGtDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbReportMemberGtDo) Session(config *gorm.Session) IFbReportMemberGtDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbReportMemberGtDo) Clauses(conds ...clause.Expression) IFbReportMemberGtDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbReportMemberGtDo) Returning(value interface{}, columns ...string) IFbReportMemberGtDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbReportMemberGtDo) Not(conds ...gen.Condition) IFbReportMemberGtDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbReportMemberGtDo) Or(conds ...gen.Condition) IFbReportMemberGtDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbReportMemberGtDo) Select(conds ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbReportMemberGtDo) Where(conds ...gen.Condition) IFbReportMemberGtDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbReportMemberGtDo) Order(conds ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbReportMemberGtDo) Distinct(cols ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbReportMemberGtDo) Omit(cols ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbReportMemberGtDo) Join(table schema.Tabler, on ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbReportMemberGtDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbReportMemberGtDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbReportMemberGtDo) Group(cols ...field.Expr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbReportMemberGtDo) Having(conds ...gen.Condition) IFbReportMemberGtDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbReportMemberGtDo) Limit(limit int) IFbReportMemberGtDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbReportMemberGtDo) Offset(offset int) IFbReportMemberGtDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbReportMemberGtDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbReportMemberGtDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbReportMemberGtDo) Unscoped() IFbReportMemberGtDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbReportMemberGtDo) Create(values ...*model.FbReportMemberGt) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbReportMemberGtDo) CreateInBatches(values []*model.FbReportMemberGt, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbReportMemberGtDo) Save(values ...*model.FbReportMemberGt) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbReportMemberGtDo) First() (*model.FbReportMemberGt, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberGt), nil
	}
}

func (f fbReportMemberGtDo) Take() (*model.FbReportMemberGt, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberGt), nil
	}
}

func (f fbReportMemberGtDo) Last() (*model.FbReportMemberGt, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberGt), nil
	}
}

func (f fbReportMemberGtDo) Find() ([]*model.FbReportMemberGt, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbReportMemberGt), err
}

func (f fbReportMemberGtDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbReportMemberGt, err error) {
	buf := make([]*model.FbReportMemberGt, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbReportMemberGtDo) FindInBatches(result *[]*model.FbReportMemberGt, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbReportMemberGtDo) Attrs(attrs ...field.AssignExpr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbReportMemberGtDo) Assign(attrs ...field.AssignExpr) IFbReportMemberGtDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbReportMemberGtDo) Joins(fields ...field.RelationField) IFbReportMemberGtDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbReportMemberGtDo) Preload(fields ...field.RelationField) IFbReportMemberGtDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbReportMemberGtDo) FirstOrInit() (*model.FbReportMemberGt, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberGt), nil
	}
}

func (f fbReportMemberGtDo) FirstOrCreate() (*model.FbReportMemberGt, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberGt), nil
	}
}

func (f fbReportMemberGtDo) FindByPage(offset int, limit int) (result []*model.FbReportMemberGt, count int64, err error) {
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

func (f fbReportMemberGtDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbReportMemberGtDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbReportMemberGtDo) Delete(models ...*model.FbReportMemberGt) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbReportMemberGtDo) withDO(do gen.Dao) *fbReportMemberGtDo {
	f.DO = *do.(*gen.DO)
	return f
}
