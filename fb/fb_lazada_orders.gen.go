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

func newFbLazadaOrder(db *gorm.DB, opts ...gen.DOOption) fbLazadaOrder {
	_fbLazadaOrder := fbLazadaOrder{}

	_fbLazadaOrder.fbLazadaOrderDo.UseDB(db, opts...)
	_fbLazadaOrder.fbLazadaOrderDo.UseModel(&model.FbLazadaOrder{})

	tableName := _fbLazadaOrder.fbLazadaOrderDo.TableName()
	_fbLazadaOrder.ALL = field.NewAsterisk(tableName)
	_fbLazadaOrder.OrderID = field.NewInt64(tableName, "order_id")
	_fbLazadaOrder.OrderStatus = field.NewString(tableName, "order_status")
	_fbLazadaOrder.SmsStatus = field.NewString(tableName, "sms_status")
	_fbLazadaOrder.BuyerID = field.NewString(tableName, "buyer_id")
	_fbLazadaOrder.BuyerPhoneNumber = field.NewString(tableName, "buyer_phone_number")
	_fbLazadaOrder.PaidAmount = field.NewFloat64(tableName, "paid_amount")
	_fbLazadaOrder.PaidAt = field.NewInt64(tableName, "paid_at")
	_fbLazadaOrder.ExchangeCode = field.NewString(tableName, "exchange_code")
	_fbLazadaOrder.ExchangeAmount = field.NewFloat64(tableName, "exchange_amount")
	_fbLazadaOrder.ExchangeStatus = field.NewString(tableName, "exchange_status")
	_fbLazadaOrder.ExchangerPhoneNumber = field.NewString(tableName, "exchanger_phone_number")
	_fbLazadaOrder.ExchangerUsername = field.NewString(tableName, "exchanger_username")
	_fbLazadaOrder.ExchangedAt = field.NewInt64(tableName, "exchanged_at")
	_fbLazadaOrder.Log = field.NewString(tableName, "log")
	_fbLazadaOrder.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbLazadaOrder.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_fbLazadaOrder.fillFieldMap()

	return _fbLazadaOrder
}

type fbLazadaOrder struct {
	fbLazadaOrderDo

	ALL                  field.Asterisk
	OrderID              field.Int64
	OrderStatus          field.String
	SmsStatus            field.String
	BuyerID              field.String
	BuyerPhoneNumber     field.String
	PaidAmount           field.Float64
	PaidAt               field.Int64
	ExchangeCode         field.String
	ExchangeAmount       field.Float64
	ExchangeStatus       field.String
	ExchangerPhoneNumber field.String
	ExchangerUsername    field.String
	ExchangedAt          field.Int64
	Log                  field.String
	CreatedAt            field.Int64
	UpdatedAt            field.Int64

	fieldMap map[string]field.Expr
}

func (f fbLazadaOrder) Table(newTableName string) *fbLazadaOrder {
	f.fbLazadaOrderDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbLazadaOrder) As(alias string) *fbLazadaOrder {
	f.fbLazadaOrderDo.DO = *(f.fbLazadaOrderDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbLazadaOrder) updateTableName(table string) *fbLazadaOrder {
	f.ALL = field.NewAsterisk(table)
	f.OrderID = field.NewInt64(table, "order_id")
	f.OrderStatus = field.NewString(table, "order_status")
	f.SmsStatus = field.NewString(table, "sms_status")
	f.BuyerID = field.NewString(table, "buyer_id")
	f.BuyerPhoneNumber = field.NewString(table, "buyer_phone_number")
	f.PaidAmount = field.NewFloat64(table, "paid_amount")
	f.PaidAt = field.NewInt64(table, "paid_at")
	f.ExchangeCode = field.NewString(table, "exchange_code")
	f.ExchangeAmount = field.NewFloat64(table, "exchange_amount")
	f.ExchangeStatus = field.NewString(table, "exchange_status")
	f.ExchangerPhoneNumber = field.NewString(table, "exchanger_phone_number")
	f.ExchangerUsername = field.NewString(table, "exchanger_username")
	f.ExchangedAt = field.NewInt64(table, "exchanged_at")
	f.Log = field.NewString(table, "log")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fbLazadaOrder) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbLazadaOrder) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 16)
	f.fieldMap["order_id"] = f.OrderID
	f.fieldMap["order_status"] = f.OrderStatus
	f.fieldMap["sms_status"] = f.SmsStatus
	f.fieldMap["buyer_id"] = f.BuyerID
	f.fieldMap["buyer_phone_number"] = f.BuyerPhoneNumber
	f.fieldMap["paid_amount"] = f.PaidAmount
	f.fieldMap["paid_at"] = f.PaidAt
	f.fieldMap["exchange_code"] = f.ExchangeCode
	f.fieldMap["exchange_amount"] = f.ExchangeAmount
	f.fieldMap["exchange_status"] = f.ExchangeStatus
	f.fieldMap["exchanger_phone_number"] = f.ExchangerPhoneNumber
	f.fieldMap["exchanger_username"] = f.ExchangerUsername
	f.fieldMap["exchanged_at"] = f.ExchangedAt
	f.fieldMap["log"] = f.Log
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fbLazadaOrder) clone(db *gorm.DB) fbLazadaOrder {
	f.fbLazadaOrderDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbLazadaOrder) replaceDB(db *gorm.DB) fbLazadaOrder {
	f.fbLazadaOrderDo.ReplaceDB(db)
	return f
}

type fbLazadaOrderDo struct{ gen.DO }

type IFbLazadaOrderDo interface {
	gen.SubQuery
	Debug() IFbLazadaOrderDo
	WithContext(ctx context.Context) IFbLazadaOrderDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbLazadaOrderDo
	WriteDB() IFbLazadaOrderDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbLazadaOrderDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbLazadaOrderDo
	Not(conds ...gen.Condition) IFbLazadaOrderDo
	Or(conds ...gen.Condition) IFbLazadaOrderDo
	Select(conds ...field.Expr) IFbLazadaOrderDo
	Where(conds ...gen.Condition) IFbLazadaOrderDo
	Order(conds ...field.Expr) IFbLazadaOrderDo
	Distinct(cols ...field.Expr) IFbLazadaOrderDo
	Omit(cols ...field.Expr) IFbLazadaOrderDo
	Join(table schema.Tabler, on ...field.Expr) IFbLazadaOrderDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbLazadaOrderDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbLazadaOrderDo
	Group(cols ...field.Expr) IFbLazadaOrderDo
	Having(conds ...gen.Condition) IFbLazadaOrderDo
	Limit(limit int) IFbLazadaOrderDo
	Offset(offset int) IFbLazadaOrderDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbLazadaOrderDo
	Unscoped() IFbLazadaOrderDo
	Create(values ...*model.FbLazadaOrder) error
	CreateInBatches(values []*model.FbLazadaOrder, batchSize int) error
	Save(values ...*model.FbLazadaOrder) error
	First() (*model.FbLazadaOrder, error)
	Take() (*model.FbLazadaOrder, error)
	Last() (*model.FbLazadaOrder, error)
	Find() ([]*model.FbLazadaOrder, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbLazadaOrder, err error)
	FindInBatches(result *[]*model.FbLazadaOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbLazadaOrder) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbLazadaOrderDo
	Assign(attrs ...field.AssignExpr) IFbLazadaOrderDo
	Joins(fields ...field.RelationField) IFbLazadaOrderDo
	Preload(fields ...field.RelationField) IFbLazadaOrderDo
	FirstOrInit() (*model.FbLazadaOrder, error)
	FirstOrCreate() (*model.FbLazadaOrder, error)
	FindByPage(offset int, limit int) (result []*model.FbLazadaOrder, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbLazadaOrderDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbLazadaOrderDo) Debug() IFbLazadaOrderDo {
	return f.withDO(f.DO.Debug())
}

func (f fbLazadaOrderDo) WithContext(ctx context.Context) IFbLazadaOrderDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbLazadaOrderDo) ReadDB() IFbLazadaOrderDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbLazadaOrderDo) WriteDB() IFbLazadaOrderDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbLazadaOrderDo) Session(config *gorm.Session) IFbLazadaOrderDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbLazadaOrderDo) Clauses(conds ...clause.Expression) IFbLazadaOrderDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbLazadaOrderDo) Returning(value interface{}, columns ...string) IFbLazadaOrderDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbLazadaOrderDo) Not(conds ...gen.Condition) IFbLazadaOrderDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbLazadaOrderDo) Or(conds ...gen.Condition) IFbLazadaOrderDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbLazadaOrderDo) Select(conds ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbLazadaOrderDo) Where(conds ...gen.Condition) IFbLazadaOrderDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbLazadaOrderDo) Order(conds ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbLazadaOrderDo) Distinct(cols ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbLazadaOrderDo) Omit(cols ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbLazadaOrderDo) Join(table schema.Tabler, on ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbLazadaOrderDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbLazadaOrderDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbLazadaOrderDo) Group(cols ...field.Expr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbLazadaOrderDo) Having(conds ...gen.Condition) IFbLazadaOrderDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbLazadaOrderDo) Limit(limit int) IFbLazadaOrderDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbLazadaOrderDo) Offset(offset int) IFbLazadaOrderDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbLazadaOrderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbLazadaOrderDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbLazadaOrderDo) Unscoped() IFbLazadaOrderDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbLazadaOrderDo) Create(values ...*model.FbLazadaOrder) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbLazadaOrderDo) CreateInBatches(values []*model.FbLazadaOrder, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbLazadaOrderDo) Save(values ...*model.FbLazadaOrder) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbLazadaOrderDo) First() (*model.FbLazadaOrder, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbLazadaOrder), nil
	}
}

func (f fbLazadaOrderDo) Take() (*model.FbLazadaOrder, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbLazadaOrder), nil
	}
}

func (f fbLazadaOrderDo) Last() (*model.FbLazadaOrder, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbLazadaOrder), nil
	}
}

func (f fbLazadaOrderDo) Find() ([]*model.FbLazadaOrder, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbLazadaOrder), err
}

func (f fbLazadaOrderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbLazadaOrder, err error) {
	buf := make([]*model.FbLazadaOrder, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbLazadaOrderDo) FindInBatches(result *[]*model.FbLazadaOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbLazadaOrderDo) Attrs(attrs ...field.AssignExpr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbLazadaOrderDo) Assign(attrs ...field.AssignExpr) IFbLazadaOrderDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbLazadaOrderDo) Joins(fields ...field.RelationField) IFbLazadaOrderDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbLazadaOrderDo) Preload(fields ...field.RelationField) IFbLazadaOrderDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbLazadaOrderDo) FirstOrInit() (*model.FbLazadaOrder, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbLazadaOrder), nil
	}
}

func (f fbLazadaOrderDo) FirstOrCreate() (*model.FbLazadaOrder, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbLazadaOrder), nil
	}
}

func (f fbLazadaOrderDo) FindByPage(offset int, limit int) (result []*model.FbLazadaOrder, count int64, err error) {
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

func (f fbLazadaOrderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbLazadaOrderDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbLazadaOrderDo) Delete(models ...*model.FbLazadaOrder) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbLazadaOrderDo) withDO(do gen.Dao) *fbLazadaOrderDo {
	f.DO = *do.(*gen.DO)
	return f
}
