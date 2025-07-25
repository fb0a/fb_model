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

func newFbDeposit(db *gorm.DB, opts ...gen.DOOption) fbDeposit {
	_fbDeposit := fbDeposit{}

	_fbDeposit.fbDepositDo.UseDB(db, opts...)
	_fbDeposit.fbDepositDo.UseModel(&model.FbDeposit{})

	tableName := _fbDeposit.fbDepositDo.TableName()
	_fbDeposit.ALL = field.NewAsterisk(tableName)
	_fbDeposit.ID = field.NewInt64(tableName, "id")
	_fbDeposit.UID = field.NewInt64(tableName, "uid")
	_fbDeposit.Username = field.NewString(tableName, "username")
	_fbDeposit.Amount = field.NewFloat64(tableName, "amount")
	_fbDeposit.ChannelCategory = field.NewString(tableName, "channel_category")
	_fbDeposit.ChannelName = field.NewString(tableName, "channel_name")
	_fbDeposit.Currency = field.NewString(tableName, "currency")
	_fbDeposit.ExternalOrderID = field.NewString(tableName, "external_order_id")
	_fbDeposit.PaymentMethod = field.NewString(tableName, "payment_method")
	_fbDeposit.RolloverMultiplier = field.NewFloat64(tableName, "rollover_multiplier")
	_fbDeposit.Status = field.NewString(tableName, "status")
	_fbDeposit.PaidAt = field.NewInt64(tableName, "paid_at")
	_fbDeposit.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbDeposit.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_fbDeposit.CancelReason = field.NewString(tableName, "cancel_reason")
	_fbDeposit.Gt = field.NewInt32(tableName, "gt")
	_fbDeposit.Remark = field.NewString(tableName, "remark")
	_fbDeposit.PaymentChannelID = field.NewInt64(tableName, "payment_channel_id")
	_fbDeposit.ExternalTransID = field.NewString(tableName, "external_trans_id")
	_fbDeposit.Phone = field.NewString(tableName, "phone")
	_fbDeposit.Sid = field.NewInt32(tableName, "sid")
	_fbDeposit.Domain = field.NewString(tableName, "domain")
	_fbDeposit.Sq = field.NewInt64(tableName, "sq")
	_fbDeposit.PaidAccountID = field.NewString(tableName, "paid_account_id")
	_fbDeposit.PaidAmount = field.NewFloat64(tableName, "paid_amount")
	_fbDeposit.Fee = field.NewFloat64(tableName, "fee")

	_fbDeposit.fillFieldMap()

	return _fbDeposit
}

// fbDeposit 存款表
type fbDeposit struct {
	fbDepositDo

	ALL                field.Asterisk
	ID                 field.Int64
	UID                field.Int64
	Username           field.String  // 用户名
	Amount             field.Float64 // 金额
	ChannelCategory    field.String  // 通道分类
	ChannelName        field.String  // 通道名
	Currency           field.String  // 币种
	ExternalOrderID    field.String  // 三方订单号
	PaymentMethod      field.String  // 支付方式
	RolloverMultiplier field.Float64 // 打码倍数
	Status             field.String  // 状态
	PaidAt             field.Int64   // 支付时间
	CreatedAt          field.Int64   // 创建时间
	UpdatedAt          field.Int64   // 创建时间
	CancelReason       field.String  // 取消原因
	Gt                 field.Int32   // 充值次数类型 0 其他 1 首冲 2 二充 3 三充
	Remark             field.String
	PaymentChannelID   field.Int64  // 支付通道id
	ExternalTransID    field.String // 三方订单号
	Phone              field.String
	Sid                field.Int32  // 店铺id
	Domain             field.String // 域名
	Sq                 field.Int64
	PaidAccountID      field.String  // 入账账户ID
	PaidAmount         field.Float64 // 到账金额
	Fee                field.Float64 // 手续费

	fieldMap map[string]field.Expr
}

func (f fbDeposit) Table(newTableName string) *fbDeposit {
	f.fbDepositDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbDeposit) As(alias string) *fbDeposit {
	f.fbDepositDo.DO = *(f.fbDepositDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbDeposit) updateTableName(table string) *fbDeposit {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.UID = field.NewInt64(table, "uid")
	f.Username = field.NewString(table, "username")
	f.Amount = field.NewFloat64(table, "amount")
	f.ChannelCategory = field.NewString(table, "channel_category")
	f.ChannelName = field.NewString(table, "channel_name")
	f.Currency = field.NewString(table, "currency")
	f.ExternalOrderID = field.NewString(table, "external_order_id")
	f.PaymentMethod = field.NewString(table, "payment_method")
	f.RolloverMultiplier = field.NewFloat64(table, "rollover_multiplier")
	f.Status = field.NewString(table, "status")
	f.PaidAt = field.NewInt64(table, "paid_at")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")
	f.CancelReason = field.NewString(table, "cancel_reason")
	f.Gt = field.NewInt32(table, "gt")
	f.Remark = field.NewString(table, "remark")
	f.PaymentChannelID = field.NewInt64(table, "payment_channel_id")
	f.ExternalTransID = field.NewString(table, "external_trans_id")
	f.Phone = field.NewString(table, "phone")
	f.Sid = field.NewInt32(table, "sid")
	f.Domain = field.NewString(table, "domain")
	f.Sq = field.NewInt64(table, "sq")
	f.PaidAccountID = field.NewString(table, "paid_account_id")
	f.PaidAmount = field.NewFloat64(table, "paid_amount")
	f.Fee = field.NewFloat64(table, "fee")

	f.fillFieldMap()

	return f
}

func (f *fbDeposit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbDeposit) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 26)
	f.fieldMap["id"] = f.ID
	f.fieldMap["uid"] = f.UID
	f.fieldMap["username"] = f.Username
	f.fieldMap["amount"] = f.Amount
	f.fieldMap["channel_category"] = f.ChannelCategory
	f.fieldMap["channel_name"] = f.ChannelName
	f.fieldMap["currency"] = f.Currency
	f.fieldMap["external_order_id"] = f.ExternalOrderID
	f.fieldMap["payment_method"] = f.PaymentMethod
	f.fieldMap["rollover_multiplier"] = f.RolloverMultiplier
	f.fieldMap["status"] = f.Status
	f.fieldMap["paid_at"] = f.PaidAt
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["cancel_reason"] = f.CancelReason
	f.fieldMap["gt"] = f.Gt
	f.fieldMap["remark"] = f.Remark
	f.fieldMap["payment_channel_id"] = f.PaymentChannelID
	f.fieldMap["external_trans_id"] = f.ExternalTransID
	f.fieldMap["phone"] = f.Phone
	f.fieldMap["sid"] = f.Sid
	f.fieldMap["domain"] = f.Domain
	f.fieldMap["sq"] = f.Sq
	f.fieldMap["paid_account_id"] = f.PaidAccountID
	f.fieldMap["paid_amount"] = f.PaidAmount
	f.fieldMap["fee"] = f.Fee
}

func (f fbDeposit) clone(db *gorm.DB) fbDeposit {
	f.fbDepositDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbDeposit) replaceDB(db *gorm.DB) fbDeposit {
	f.fbDepositDo.ReplaceDB(db)
	return f
}

type fbDepositDo struct{ gen.DO }

type IFbDepositDo interface {
	gen.SubQuery
	Debug() IFbDepositDo
	WithContext(ctx context.Context) IFbDepositDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbDepositDo
	WriteDB() IFbDepositDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbDepositDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbDepositDo
	Not(conds ...gen.Condition) IFbDepositDo
	Or(conds ...gen.Condition) IFbDepositDo
	Select(conds ...field.Expr) IFbDepositDo
	Where(conds ...gen.Condition) IFbDepositDo
	Order(conds ...field.Expr) IFbDepositDo
	Distinct(cols ...field.Expr) IFbDepositDo
	Omit(cols ...field.Expr) IFbDepositDo
	Join(table schema.Tabler, on ...field.Expr) IFbDepositDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbDepositDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbDepositDo
	Group(cols ...field.Expr) IFbDepositDo
	Having(conds ...gen.Condition) IFbDepositDo
	Limit(limit int) IFbDepositDo
	Offset(offset int) IFbDepositDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbDepositDo
	Unscoped() IFbDepositDo
	Create(values ...*model.FbDeposit) error
	CreateInBatches(values []*model.FbDeposit, batchSize int) error
	Save(values ...*model.FbDeposit) error
	First() (*model.FbDeposit, error)
	Take() (*model.FbDeposit, error)
	Last() (*model.FbDeposit, error)
	Find() ([]*model.FbDeposit, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbDeposit, err error)
	FindInBatches(result *[]*model.FbDeposit, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbDeposit) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbDepositDo
	Assign(attrs ...field.AssignExpr) IFbDepositDo
	Joins(fields ...field.RelationField) IFbDepositDo
	Preload(fields ...field.RelationField) IFbDepositDo
	FirstOrInit() (*model.FbDeposit, error)
	FirstOrCreate() (*model.FbDeposit, error)
	FindByPage(offset int, limit int) (result []*model.FbDeposit, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbDepositDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbDepositDo) Debug() IFbDepositDo {
	return f.withDO(f.DO.Debug())
}

func (f fbDepositDo) WithContext(ctx context.Context) IFbDepositDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbDepositDo) ReadDB() IFbDepositDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbDepositDo) WriteDB() IFbDepositDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbDepositDo) Session(config *gorm.Session) IFbDepositDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbDepositDo) Clauses(conds ...clause.Expression) IFbDepositDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbDepositDo) Returning(value interface{}, columns ...string) IFbDepositDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbDepositDo) Not(conds ...gen.Condition) IFbDepositDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbDepositDo) Or(conds ...gen.Condition) IFbDepositDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbDepositDo) Select(conds ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbDepositDo) Where(conds ...gen.Condition) IFbDepositDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbDepositDo) Order(conds ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbDepositDo) Distinct(cols ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbDepositDo) Omit(cols ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbDepositDo) Join(table schema.Tabler, on ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbDepositDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbDepositDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbDepositDo) Group(cols ...field.Expr) IFbDepositDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbDepositDo) Having(conds ...gen.Condition) IFbDepositDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbDepositDo) Limit(limit int) IFbDepositDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbDepositDo) Offset(offset int) IFbDepositDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbDepositDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbDepositDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbDepositDo) Unscoped() IFbDepositDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbDepositDo) Create(values ...*model.FbDeposit) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbDepositDo) CreateInBatches(values []*model.FbDeposit, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbDepositDo) Save(values ...*model.FbDeposit) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbDepositDo) First() (*model.FbDeposit, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDeposit), nil
	}
}

func (f fbDepositDo) Take() (*model.FbDeposit, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDeposit), nil
	}
}

func (f fbDepositDo) Last() (*model.FbDeposit, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDeposit), nil
	}
}

func (f fbDepositDo) Find() ([]*model.FbDeposit, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbDeposit), err
}

func (f fbDepositDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbDeposit, err error) {
	buf := make([]*model.FbDeposit, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbDepositDo) FindInBatches(result *[]*model.FbDeposit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbDepositDo) Attrs(attrs ...field.AssignExpr) IFbDepositDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbDepositDo) Assign(attrs ...field.AssignExpr) IFbDepositDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbDepositDo) Joins(fields ...field.RelationField) IFbDepositDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbDepositDo) Preload(fields ...field.RelationField) IFbDepositDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbDepositDo) FirstOrInit() (*model.FbDeposit, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDeposit), nil
	}
}

func (f fbDepositDo) FirstOrCreate() (*model.FbDeposit, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDeposit), nil
	}
}

func (f fbDepositDo) FindByPage(offset int, limit int) (result []*model.FbDeposit, count int64, err error) {
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

func (f fbDepositDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbDepositDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbDepositDo) Delete(models ...*model.FbDeposit) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbDepositDo) withDO(do gen.Dao) *fbDepositDo {
	f.DO = *do.(*gen.DO)
	return f
}
