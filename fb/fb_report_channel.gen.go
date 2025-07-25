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

func newFbReportChannel(db *gorm.DB, opts ...gen.DOOption) fbReportChannel {
	_fbReportChannel := fbReportChannel{}

	_fbReportChannel.fbReportChannelDo.UseDB(db, opts...)
	_fbReportChannel.fbReportChannelDo.UseModel(&model.FbReportChannel{})

	tableName := _fbReportChannel.fbReportChannelDo.TableName()
	_fbReportChannel.ALL = field.NewAsterisk(tableName)
	_fbReportChannel.ID = field.NewInt64(tableName, "id")
	_fbReportChannel.Channel = field.NewString(tableName, "channel")
	_fbReportChannel.Day = field.NewInt32(tableName, "day")
	_fbReportChannel.Regs = field.NewInt32(tableName, "regs")
	_fbReportChannel.FirstDeposit = field.NewInt32(tableName, "first_deposit")
	_fbReportChannel.Bets = field.NewInt32(tableName, "bets")
	_fbReportChannel.BetOrders = field.NewInt32(tableName, "bet_orders")
	_fbReportChannel.RegsFirstDeposit = field.NewInt32(tableName, "regs_first_deposit")
	_fbReportChannel.ValidBetAmount = field.NewFloat64(tableName, "valid_bet_amount")
	_fbReportChannel.AvgBetAmount = field.NewFloat64(tableName, "avg_bet_amount")
	_fbReportChannel.SettleAmount = field.NewFloat64(tableName, "settle_amount")
	_fbReportChannel.Ggr = field.NewFloat64(tableName, "ggr")
	_fbReportChannel.GgrRatio = field.NewFloat64(tableName, "ggr_ratio")
	_fbReportChannel.CreatedAt = field.NewInt64(tableName, "created_at")

	_fbReportChannel.fillFieldMap()

	return _fbReportChannel
}

// fbReportChannel 渠道业绩日报表
type fbReportChannel struct {
	fbReportChannelDo

	ALL              field.Asterisk
	ID               field.Int64
	Channel          field.String  // 渠道名称
	Day              field.Int32   // 年月日
	Regs             field.Int32   // 新增注册
	FirstDeposit     field.Int32   // 新增首存
	Bets             field.Int32   // 活跃人数
	BetOrders        field.Int32   // 总注单数
	RegsFirstDeposit field.Int32   // 注册人数首冲数
	ValidBetAmount   field.Float64 // 总有效投注额
	AvgBetAmount     field.Float64 // 笔均注单额
	SettleAmount     field.Float64 // 总派彩
	Ggr              field.Float64 // ggr
	GgrRatio         field.Float64 // 杀率
	CreatedAt        field.Int64   // 创建时间

	fieldMap map[string]field.Expr
}

func (f fbReportChannel) Table(newTableName string) *fbReportChannel {
	f.fbReportChannelDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbReportChannel) As(alias string) *fbReportChannel {
	f.fbReportChannelDo.DO = *(f.fbReportChannelDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbReportChannel) updateTableName(table string) *fbReportChannel {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.Channel = field.NewString(table, "channel")
	f.Day = field.NewInt32(table, "day")
	f.Regs = field.NewInt32(table, "regs")
	f.FirstDeposit = field.NewInt32(table, "first_deposit")
	f.Bets = field.NewInt32(table, "bets")
	f.BetOrders = field.NewInt32(table, "bet_orders")
	f.RegsFirstDeposit = field.NewInt32(table, "regs_first_deposit")
	f.ValidBetAmount = field.NewFloat64(table, "valid_bet_amount")
	f.AvgBetAmount = field.NewFloat64(table, "avg_bet_amount")
	f.SettleAmount = field.NewFloat64(table, "settle_amount")
	f.Ggr = field.NewFloat64(table, "ggr")
	f.GgrRatio = field.NewFloat64(table, "ggr_ratio")
	f.CreatedAt = field.NewInt64(table, "created_at")

	f.fillFieldMap()

	return f
}

func (f *fbReportChannel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbReportChannel) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 14)
	f.fieldMap["id"] = f.ID
	f.fieldMap["channel"] = f.Channel
	f.fieldMap["day"] = f.Day
	f.fieldMap["regs"] = f.Regs
	f.fieldMap["first_deposit"] = f.FirstDeposit
	f.fieldMap["bets"] = f.Bets
	f.fieldMap["bet_orders"] = f.BetOrders
	f.fieldMap["regs_first_deposit"] = f.RegsFirstDeposit
	f.fieldMap["valid_bet_amount"] = f.ValidBetAmount
	f.fieldMap["avg_bet_amount"] = f.AvgBetAmount
	f.fieldMap["settle_amount"] = f.SettleAmount
	f.fieldMap["ggr"] = f.Ggr
	f.fieldMap["ggr_ratio"] = f.GgrRatio
	f.fieldMap["created_at"] = f.CreatedAt
}

func (f fbReportChannel) clone(db *gorm.DB) fbReportChannel {
	f.fbReportChannelDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbReportChannel) replaceDB(db *gorm.DB) fbReportChannel {
	f.fbReportChannelDo.ReplaceDB(db)
	return f
}

type fbReportChannelDo struct{ gen.DO }

type IFbReportChannelDo interface {
	gen.SubQuery
	Debug() IFbReportChannelDo
	WithContext(ctx context.Context) IFbReportChannelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbReportChannelDo
	WriteDB() IFbReportChannelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbReportChannelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbReportChannelDo
	Not(conds ...gen.Condition) IFbReportChannelDo
	Or(conds ...gen.Condition) IFbReportChannelDo
	Select(conds ...field.Expr) IFbReportChannelDo
	Where(conds ...gen.Condition) IFbReportChannelDo
	Order(conds ...field.Expr) IFbReportChannelDo
	Distinct(cols ...field.Expr) IFbReportChannelDo
	Omit(cols ...field.Expr) IFbReportChannelDo
	Join(table schema.Tabler, on ...field.Expr) IFbReportChannelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbReportChannelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbReportChannelDo
	Group(cols ...field.Expr) IFbReportChannelDo
	Having(conds ...gen.Condition) IFbReportChannelDo
	Limit(limit int) IFbReportChannelDo
	Offset(offset int) IFbReportChannelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbReportChannelDo
	Unscoped() IFbReportChannelDo
	Create(values ...*model.FbReportChannel) error
	CreateInBatches(values []*model.FbReportChannel, batchSize int) error
	Save(values ...*model.FbReportChannel) error
	First() (*model.FbReportChannel, error)
	Take() (*model.FbReportChannel, error)
	Last() (*model.FbReportChannel, error)
	Find() ([]*model.FbReportChannel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbReportChannel, err error)
	FindInBatches(result *[]*model.FbReportChannel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbReportChannel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbReportChannelDo
	Assign(attrs ...field.AssignExpr) IFbReportChannelDo
	Joins(fields ...field.RelationField) IFbReportChannelDo
	Preload(fields ...field.RelationField) IFbReportChannelDo
	FirstOrInit() (*model.FbReportChannel, error)
	FirstOrCreate() (*model.FbReportChannel, error)
	FindByPage(offset int, limit int) (result []*model.FbReportChannel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbReportChannelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbReportChannelDo) Debug() IFbReportChannelDo {
	return f.withDO(f.DO.Debug())
}

func (f fbReportChannelDo) WithContext(ctx context.Context) IFbReportChannelDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbReportChannelDo) ReadDB() IFbReportChannelDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbReportChannelDo) WriteDB() IFbReportChannelDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbReportChannelDo) Session(config *gorm.Session) IFbReportChannelDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbReportChannelDo) Clauses(conds ...clause.Expression) IFbReportChannelDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbReportChannelDo) Returning(value interface{}, columns ...string) IFbReportChannelDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbReportChannelDo) Not(conds ...gen.Condition) IFbReportChannelDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbReportChannelDo) Or(conds ...gen.Condition) IFbReportChannelDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbReportChannelDo) Select(conds ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbReportChannelDo) Where(conds ...gen.Condition) IFbReportChannelDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbReportChannelDo) Order(conds ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbReportChannelDo) Distinct(cols ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbReportChannelDo) Omit(cols ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbReportChannelDo) Join(table schema.Tabler, on ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbReportChannelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbReportChannelDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbReportChannelDo) Group(cols ...field.Expr) IFbReportChannelDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbReportChannelDo) Having(conds ...gen.Condition) IFbReportChannelDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbReportChannelDo) Limit(limit int) IFbReportChannelDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbReportChannelDo) Offset(offset int) IFbReportChannelDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbReportChannelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbReportChannelDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbReportChannelDo) Unscoped() IFbReportChannelDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbReportChannelDo) Create(values ...*model.FbReportChannel) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbReportChannelDo) CreateInBatches(values []*model.FbReportChannel, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbReportChannelDo) Save(values ...*model.FbReportChannel) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbReportChannelDo) First() (*model.FbReportChannel, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportChannel), nil
	}
}

func (f fbReportChannelDo) Take() (*model.FbReportChannel, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportChannel), nil
	}
}

func (f fbReportChannelDo) Last() (*model.FbReportChannel, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportChannel), nil
	}
}

func (f fbReportChannelDo) Find() ([]*model.FbReportChannel, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbReportChannel), err
}

func (f fbReportChannelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbReportChannel, err error) {
	buf := make([]*model.FbReportChannel, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbReportChannelDo) FindInBatches(result *[]*model.FbReportChannel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbReportChannelDo) Attrs(attrs ...field.AssignExpr) IFbReportChannelDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbReportChannelDo) Assign(attrs ...field.AssignExpr) IFbReportChannelDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbReportChannelDo) Joins(fields ...field.RelationField) IFbReportChannelDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbReportChannelDo) Preload(fields ...field.RelationField) IFbReportChannelDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbReportChannelDo) FirstOrInit() (*model.FbReportChannel, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportChannel), nil
	}
}

func (f fbReportChannelDo) FirstOrCreate() (*model.FbReportChannel, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportChannel), nil
	}
}

func (f fbReportChannelDo) FindByPage(offset int, limit int) (result []*model.FbReportChannel, count int64, err error) {
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

func (f fbReportChannelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbReportChannelDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbReportChannelDo) Delete(models ...*model.FbReportChannel) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbReportChannelDo) withDO(do gen.Dao) *fbReportChannelDo {
	f.DO = *do.(*gen.DO)
	return f
}
