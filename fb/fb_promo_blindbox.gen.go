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

func newFbPromoBlindbox(db *gorm.DB, opts ...gen.DOOption) fbPromoBlindbox {
	_fbPromoBlindbox := fbPromoBlindbox{}

	_fbPromoBlindbox.fbPromoBlindboxDo.UseDB(db, opts...)
	_fbPromoBlindbox.fbPromoBlindboxDo.UseModel(&model.FbPromoBlindbox{})

	tableName := _fbPromoBlindbox.fbPromoBlindboxDo.TableName()
	_fbPromoBlindbox.ALL = field.NewAsterisk(tableName)
	_fbPromoBlindbox.ID = field.NewInt64(tableName, "id")
	_fbPromoBlindbox.UID = field.NewInt64(tableName, "uid")
	_fbPromoBlindbox.Username = field.NewString(tableName, "username")
	_fbPromoBlindbox.Phone = field.NewString(tableName, "phone")
	_fbPromoBlindbox.ParentUID = field.NewString(tableName, "parent_uid")
	_fbPromoBlindbox.ParentName = field.NewString(tableName, "parent_name")
	_fbPromoBlindbox.Pid = field.NewInt64(tableName, "pid")
	_fbPromoBlindbox.Bonus = field.NewFloat64(tableName, "bonus")
	_fbPromoBlindbox.BetAmount = field.NewFloat64(tableName, "bet_amount")
	_fbPromoBlindbox.NetAmount = field.NewFloat64(tableName, "net_amount")
	_fbPromoBlindbox.State = field.NewInt32(tableName, "state")
	_fbPromoBlindbox.BonusRate = field.NewFloat64(tableName, "bonus_rate")
	_fbPromoBlindbox.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbPromoBlindbox.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_fbPromoBlindbox.Ymd = field.NewInt32(tableName, "ymd")

	_fbPromoBlindbox.fillFieldMap()

	return _fbPromoBlindbox
}

// fbPromoBlindbox 盲盒记录表
type fbPromoBlindbox struct {
	fbPromoBlindboxDo

	ALL        field.Asterisk
	ID         field.Int64
	UID        field.Int64   // 用户ID
	Username   field.String  // 用户名
	Phone      field.String  // 手机号
	ParentUID  field.String  // 上级UID
	ParentName field.String  // 上级用户名
	Pid        field.Int64   // 活动ID
	Bonus      field.Float64 // 奖金金额
	BetAmount  field.Float64 // 投注金额
	NetAmount  field.Float64
	State      field.Int32   // 状态 1:未领取 2:已领取 3:已过期 4:待审核
	BonusRate  field.Float64 // 奖金比例(0.001-0.002之间)
	CreatedAt  field.Int64   // 创建时间
	UpdatedAt  field.Int64   // 更新时间
	Ymd        field.Int32   // 20060102

	fieldMap map[string]field.Expr
}

func (f fbPromoBlindbox) Table(newTableName string) *fbPromoBlindbox {
	f.fbPromoBlindboxDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbPromoBlindbox) As(alias string) *fbPromoBlindbox {
	f.fbPromoBlindboxDo.DO = *(f.fbPromoBlindboxDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbPromoBlindbox) updateTableName(table string) *fbPromoBlindbox {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.UID = field.NewInt64(table, "uid")
	f.Username = field.NewString(table, "username")
	f.Phone = field.NewString(table, "phone")
	f.ParentUID = field.NewString(table, "parent_uid")
	f.ParentName = field.NewString(table, "parent_name")
	f.Pid = field.NewInt64(table, "pid")
	f.Bonus = field.NewFloat64(table, "bonus")
	f.BetAmount = field.NewFloat64(table, "bet_amount")
	f.NetAmount = field.NewFloat64(table, "net_amount")
	f.State = field.NewInt32(table, "state")
	f.BonusRate = field.NewFloat64(table, "bonus_rate")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")
	f.Ymd = field.NewInt32(table, "ymd")

	f.fillFieldMap()

	return f
}

func (f *fbPromoBlindbox) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbPromoBlindbox) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 15)
	f.fieldMap["id"] = f.ID
	f.fieldMap["uid"] = f.UID
	f.fieldMap["username"] = f.Username
	f.fieldMap["phone"] = f.Phone
	f.fieldMap["parent_uid"] = f.ParentUID
	f.fieldMap["parent_name"] = f.ParentName
	f.fieldMap["pid"] = f.Pid
	f.fieldMap["bonus"] = f.Bonus
	f.fieldMap["bet_amount"] = f.BetAmount
	f.fieldMap["net_amount"] = f.NetAmount
	f.fieldMap["state"] = f.State
	f.fieldMap["bonus_rate"] = f.BonusRate
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["ymd"] = f.Ymd
}

func (f fbPromoBlindbox) clone(db *gorm.DB) fbPromoBlindbox {
	f.fbPromoBlindboxDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbPromoBlindbox) replaceDB(db *gorm.DB) fbPromoBlindbox {
	f.fbPromoBlindboxDo.ReplaceDB(db)
	return f
}

type fbPromoBlindboxDo struct{ gen.DO }

type IFbPromoBlindboxDo interface {
	gen.SubQuery
	Debug() IFbPromoBlindboxDo
	WithContext(ctx context.Context) IFbPromoBlindboxDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbPromoBlindboxDo
	WriteDB() IFbPromoBlindboxDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbPromoBlindboxDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbPromoBlindboxDo
	Not(conds ...gen.Condition) IFbPromoBlindboxDo
	Or(conds ...gen.Condition) IFbPromoBlindboxDo
	Select(conds ...field.Expr) IFbPromoBlindboxDo
	Where(conds ...gen.Condition) IFbPromoBlindboxDo
	Order(conds ...field.Expr) IFbPromoBlindboxDo
	Distinct(cols ...field.Expr) IFbPromoBlindboxDo
	Omit(cols ...field.Expr) IFbPromoBlindboxDo
	Join(table schema.Tabler, on ...field.Expr) IFbPromoBlindboxDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbPromoBlindboxDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbPromoBlindboxDo
	Group(cols ...field.Expr) IFbPromoBlindboxDo
	Having(conds ...gen.Condition) IFbPromoBlindboxDo
	Limit(limit int) IFbPromoBlindboxDo
	Offset(offset int) IFbPromoBlindboxDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbPromoBlindboxDo
	Unscoped() IFbPromoBlindboxDo
	Create(values ...*model.FbPromoBlindbox) error
	CreateInBatches(values []*model.FbPromoBlindbox, batchSize int) error
	Save(values ...*model.FbPromoBlindbox) error
	First() (*model.FbPromoBlindbox, error)
	Take() (*model.FbPromoBlindbox, error)
	Last() (*model.FbPromoBlindbox, error)
	Find() ([]*model.FbPromoBlindbox, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbPromoBlindbox, err error)
	FindInBatches(result *[]*model.FbPromoBlindbox, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbPromoBlindbox) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbPromoBlindboxDo
	Assign(attrs ...field.AssignExpr) IFbPromoBlindboxDo
	Joins(fields ...field.RelationField) IFbPromoBlindboxDo
	Preload(fields ...field.RelationField) IFbPromoBlindboxDo
	FirstOrInit() (*model.FbPromoBlindbox, error)
	FirstOrCreate() (*model.FbPromoBlindbox, error)
	FindByPage(offset int, limit int) (result []*model.FbPromoBlindbox, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbPromoBlindboxDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbPromoBlindboxDo) Debug() IFbPromoBlindboxDo {
	return f.withDO(f.DO.Debug())
}

func (f fbPromoBlindboxDo) WithContext(ctx context.Context) IFbPromoBlindboxDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbPromoBlindboxDo) ReadDB() IFbPromoBlindboxDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbPromoBlindboxDo) WriteDB() IFbPromoBlindboxDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbPromoBlindboxDo) Session(config *gorm.Session) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbPromoBlindboxDo) Clauses(conds ...clause.Expression) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbPromoBlindboxDo) Returning(value interface{}, columns ...string) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbPromoBlindboxDo) Not(conds ...gen.Condition) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbPromoBlindboxDo) Or(conds ...gen.Condition) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbPromoBlindboxDo) Select(conds ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbPromoBlindboxDo) Where(conds ...gen.Condition) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbPromoBlindboxDo) Order(conds ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbPromoBlindboxDo) Distinct(cols ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbPromoBlindboxDo) Omit(cols ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbPromoBlindboxDo) Join(table schema.Tabler, on ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbPromoBlindboxDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbPromoBlindboxDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbPromoBlindboxDo) Group(cols ...field.Expr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbPromoBlindboxDo) Having(conds ...gen.Condition) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbPromoBlindboxDo) Limit(limit int) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbPromoBlindboxDo) Offset(offset int) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbPromoBlindboxDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbPromoBlindboxDo) Unscoped() IFbPromoBlindboxDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbPromoBlindboxDo) Create(values ...*model.FbPromoBlindbox) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbPromoBlindboxDo) CreateInBatches(values []*model.FbPromoBlindbox, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbPromoBlindboxDo) Save(values ...*model.FbPromoBlindbox) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbPromoBlindboxDo) First() (*model.FbPromoBlindbox, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPromoBlindbox), nil
	}
}

func (f fbPromoBlindboxDo) Take() (*model.FbPromoBlindbox, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPromoBlindbox), nil
	}
}

func (f fbPromoBlindboxDo) Last() (*model.FbPromoBlindbox, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPromoBlindbox), nil
	}
}

func (f fbPromoBlindboxDo) Find() ([]*model.FbPromoBlindbox, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbPromoBlindbox), err
}

func (f fbPromoBlindboxDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbPromoBlindbox, err error) {
	buf := make([]*model.FbPromoBlindbox, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbPromoBlindboxDo) FindInBatches(result *[]*model.FbPromoBlindbox, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbPromoBlindboxDo) Attrs(attrs ...field.AssignExpr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbPromoBlindboxDo) Assign(attrs ...field.AssignExpr) IFbPromoBlindboxDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbPromoBlindboxDo) Joins(fields ...field.RelationField) IFbPromoBlindboxDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbPromoBlindboxDo) Preload(fields ...field.RelationField) IFbPromoBlindboxDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbPromoBlindboxDo) FirstOrInit() (*model.FbPromoBlindbox, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPromoBlindbox), nil
	}
}

func (f fbPromoBlindboxDo) FirstOrCreate() (*model.FbPromoBlindbox, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPromoBlindbox), nil
	}
}

func (f fbPromoBlindboxDo) FindByPage(offset int, limit int) (result []*model.FbPromoBlindbox, count int64, err error) {
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

func (f fbPromoBlindboxDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbPromoBlindboxDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbPromoBlindboxDo) Delete(models ...*model.FbPromoBlindbox) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbPromoBlindboxDo) withDO(do gen.Dao) *fbPromoBlindboxDo {
	f.DO = *do.(*gen.DO)
	return f
}
