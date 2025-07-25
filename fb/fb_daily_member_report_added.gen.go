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

func newFbDailyMemberReportAdded(db *gorm.DB, opts ...gen.DOOption) fbDailyMemberReportAdded {
	_fbDailyMemberReportAdded := fbDailyMemberReportAdded{}

	_fbDailyMemberReportAdded.fbDailyMemberReportAddedDo.UseDB(db, opts...)
	_fbDailyMemberReportAdded.fbDailyMemberReportAddedDo.UseModel(&model.FbDailyMemberReportAdded{})

	tableName := _fbDailyMemberReportAdded.fbDailyMemberReportAddedDo.TableName()
	_fbDailyMemberReportAdded.ALL = field.NewAsterisk(tableName)
	_fbDailyMemberReportAdded.ID = field.NewInt64(tableName, "id")
	_fbDailyMemberReportAdded.Day = field.NewInt32(tableName, "day")
	_fbDailyMemberReportAdded.RegCount = field.NewInt64(tableName, "reg_count")
	_fbDailyMemberReportAdded.RegSourceChannel = field.NewString(tableName, "reg_source_channel")
	_fbDailyMemberReportAdded.FirstDepositConversion = field.NewInt64(tableName, "first_deposit_conversion")
	_fbDailyMemberReportAdded.FirstDepositSourceChannel = field.NewString(tableName, "first_deposit_source_channel")
	_fbDailyMemberReportAdded.FirstDepositNew = field.NewInt64(tableName, "first_deposit_new")
	_fbDailyMemberReportAdded.FirstDepositNewChannelRatio = field.NewString(tableName, "first_deposit_new_channel_ratio")
	_fbDailyMemberReportAdded.FirstDepositAmount = field.NewString(tableName, "first_deposit_amount")
	_fbDailyMemberReportAdded.NumberOfSecondDepositors = field.NewInt64(tableName, "number_of_second_depositors")
	_fbDailyMemberReportAdded.AmountOfSecondDeposit = field.NewString(tableName, "amount_of_second_deposit")
	_fbDailyMemberReportAdded.NumberOfThirdDepositors = field.NewInt64(tableName, "number_of_third_depositors")
	_fbDailyMemberReportAdded.AmountOfThirdDeposit = field.NewString(tableName, "amount_of_third_deposit")
	_fbDailyMemberReportAdded.GiftMoneyDay = field.NewString(tableName, "gift_money_day")
	_fbDailyMemberReportAdded.GiftMoneyChannelRatio = field.NewString(tableName, "gift_money_channel_ratio")
	_fbDailyMemberReportAdded.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbDailyMemberReportAdded.ReportType = field.NewInt32(tableName, "report_type")

	_fbDailyMemberReportAdded.fillFieldMap()

	return _fbDailyMemberReportAdded
}

// fbDailyMemberReportAdded 新增会员日报表
type fbDailyMemberReportAdded struct {
	fbDailyMemberReportAddedDo

	ALL                         field.Asterisk
	ID                          field.Int64  // ID
	Day                         field.Int32  // 年月日
	RegCount                    field.Int64  // 新增注册数
	RegSourceChannel            field.String // 注册来源渠道统计
	FirstDepositConversion      field.Int64  // 转化首存数
	FirstDepositSourceChannel   field.String // 首存来源渠道
	FirstDepositNew             field.Int64  // 新增首存数
	FirstDepositNewChannelRatio field.String // 新增首存渠道占比
	FirstDepositAmount          field.String // 首存金额
	NumberOfSecondDepositors    field.Int64  // 二存数
	AmountOfSecondDeposit       field.String // 二存金额
	NumberOfThirdDepositors     field.Int64  // 三存数
	AmountOfThirdDeposit        field.String // 三存金额
	GiftMoneyDay                field.String // 当日礼金
	GiftMoneyChannelRatio       field.String // 渠道礼金占比
	CreatedAt                   field.Int64  // 创建时间
	ReportType                  field.Int32  // 1日报，2周报，3月报

	fieldMap map[string]field.Expr
}

func (f fbDailyMemberReportAdded) Table(newTableName string) *fbDailyMemberReportAdded {
	f.fbDailyMemberReportAddedDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbDailyMemberReportAdded) As(alias string) *fbDailyMemberReportAdded {
	f.fbDailyMemberReportAddedDo.DO = *(f.fbDailyMemberReportAddedDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbDailyMemberReportAdded) updateTableName(table string) *fbDailyMemberReportAdded {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.Day = field.NewInt32(table, "day")
	f.RegCount = field.NewInt64(table, "reg_count")
	f.RegSourceChannel = field.NewString(table, "reg_source_channel")
	f.FirstDepositConversion = field.NewInt64(table, "first_deposit_conversion")
	f.FirstDepositSourceChannel = field.NewString(table, "first_deposit_source_channel")
	f.FirstDepositNew = field.NewInt64(table, "first_deposit_new")
	f.FirstDepositNewChannelRatio = field.NewString(table, "first_deposit_new_channel_ratio")
	f.FirstDepositAmount = field.NewString(table, "first_deposit_amount")
	f.NumberOfSecondDepositors = field.NewInt64(table, "number_of_second_depositors")
	f.AmountOfSecondDeposit = field.NewString(table, "amount_of_second_deposit")
	f.NumberOfThirdDepositors = field.NewInt64(table, "number_of_third_depositors")
	f.AmountOfThirdDeposit = field.NewString(table, "amount_of_third_deposit")
	f.GiftMoneyDay = field.NewString(table, "gift_money_day")
	f.GiftMoneyChannelRatio = field.NewString(table, "gift_money_channel_ratio")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.ReportType = field.NewInt32(table, "report_type")

	f.fillFieldMap()

	return f
}

func (f *fbDailyMemberReportAdded) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbDailyMemberReportAdded) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 17)
	f.fieldMap["id"] = f.ID
	f.fieldMap["day"] = f.Day
	f.fieldMap["reg_count"] = f.RegCount
	f.fieldMap["reg_source_channel"] = f.RegSourceChannel
	f.fieldMap["first_deposit_conversion"] = f.FirstDepositConversion
	f.fieldMap["first_deposit_source_channel"] = f.FirstDepositSourceChannel
	f.fieldMap["first_deposit_new"] = f.FirstDepositNew
	f.fieldMap["first_deposit_new_channel_ratio"] = f.FirstDepositNewChannelRatio
	f.fieldMap["first_deposit_amount"] = f.FirstDepositAmount
	f.fieldMap["number_of_second_depositors"] = f.NumberOfSecondDepositors
	f.fieldMap["amount_of_second_deposit"] = f.AmountOfSecondDeposit
	f.fieldMap["number_of_third_depositors"] = f.NumberOfThirdDepositors
	f.fieldMap["amount_of_third_deposit"] = f.AmountOfThirdDeposit
	f.fieldMap["gift_money_day"] = f.GiftMoneyDay
	f.fieldMap["gift_money_channel_ratio"] = f.GiftMoneyChannelRatio
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["report_type"] = f.ReportType
}

func (f fbDailyMemberReportAdded) clone(db *gorm.DB) fbDailyMemberReportAdded {
	f.fbDailyMemberReportAddedDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbDailyMemberReportAdded) replaceDB(db *gorm.DB) fbDailyMemberReportAdded {
	f.fbDailyMemberReportAddedDo.ReplaceDB(db)
	return f
}

type fbDailyMemberReportAddedDo struct{ gen.DO }

type IFbDailyMemberReportAddedDo interface {
	gen.SubQuery
	Debug() IFbDailyMemberReportAddedDo
	WithContext(ctx context.Context) IFbDailyMemberReportAddedDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbDailyMemberReportAddedDo
	WriteDB() IFbDailyMemberReportAddedDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbDailyMemberReportAddedDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbDailyMemberReportAddedDo
	Not(conds ...gen.Condition) IFbDailyMemberReportAddedDo
	Or(conds ...gen.Condition) IFbDailyMemberReportAddedDo
	Select(conds ...field.Expr) IFbDailyMemberReportAddedDo
	Where(conds ...gen.Condition) IFbDailyMemberReportAddedDo
	Order(conds ...field.Expr) IFbDailyMemberReportAddedDo
	Distinct(cols ...field.Expr) IFbDailyMemberReportAddedDo
	Omit(cols ...field.Expr) IFbDailyMemberReportAddedDo
	Join(table schema.Tabler, on ...field.Expr) IFbDailyMemberReportAddedDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbDailyMemberReportAddedDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbDailyMemberReportAddedDo
	Group(cols ...field.Expr) IFbDailyMemberReportAddedDo
	Having(conds ...gen.Condition) IFbDailyMemberReportAddedDo
	Limit(limit int) IFbDailyMemberReportAddedDo
	Offset(offset int) IFbDailyMemberReportAddedDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbDailyMemberReportAddedDo
	Unscoped() IFbDailyMemberReportAddedDo
	Create(values ...*model.FbDailyMemberReportAdded) error
	CreateInBatches(values []*model.FbDailyMemberReportAdded, batchSize int) error
	Save(values ...*model.FbDailyMemberReportAdded) error
	First() (*model.FbDailyMemberReportAdded, error)
	Take() (*model.FbDailyMemberReportAdded, error)
	Last() (*model.FbDailyMemberReportAdded, error)
	Find() ([]*model.FbDailyMemberReportAdded, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbDailyMemberReportAdded, err error)
	FindInBatches(result *[]*model.FbDailyMemberReportAdded, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbDailyMemberReportAdded) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbDailyMemberReportAddedDo
	Assign(attrs ...field.AssignExpr) IFbDailyMemberReportAddedDo
	Joins(fields ...field.RelationField) IFbDailyMemberReportAddedDo
	Preload(fields ...field.RelationField) IFbDailyMemberReportAddedDo
	FirstOrInit() (*model.FbDailyMemberReportAdded, error)
	FirstOrCreate() (*model.FbDailyMemberReportAdded, error)
	FindByPage(offset int, limit int) (result []*model.FbDailyMemberReportAdded, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbDailyMemberReportAddedDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbDailyMemberReportAddedDo) Debug() IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Debug())
}

func (f fbDailyMemberReportAddedDo) WithContext(ctx context.Context) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbDailyMemberReportAddedDo) ReadDB() IFbDailyMemberReportAddedDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbDailyMemberReportAddedDo) WriteDB() IFbDailyMemberReportAddedDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbDailyMemberReportAddedDo) Session(config *gorm.Session) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbDailyMemberReportAddedDo) Clauses(conds ...clause.Expression) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbDailyMemberReportAddedDo) Returning(value interface{}, columns ...string) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbDailyMemberReportAddedDo) Not(conds ...gen.Condition) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbDailyMemberReportAddedDo) Or(conds ...gen.Condition) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbDailyMemberReportAddedDo) Select(conds ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbDailyMemberReportAddedDo) Where(conds ...gen.Condition) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbDailyMemberReportAddedDo) Order(conds ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbDailyMemberReportAddedDo) Distinct(cols ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbDailyMemberReportAddedDo) Omit(cols ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbDailyMemberReportAddedDo) Join(table schema.Tabler, on ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbDailyMemberReportAddedDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbDailyMemberReportAddedDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbDailyMemberReportAddedDo) Group(cols ...field.Expr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbDailyMemberReportAddedDo) Having(conds ...gen.Condition) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbDailyMemberReportAddedDo) Limit(limit int) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbDailyMemberReportAddedDo) Offset(offset int) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbDailyMemberReportAddedDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbDailyMemberReportAddedDo) Unscoped() IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbDailyMemberReportAddedDo) Create(values ...*model.FbDailyMemberReportAdded) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbDailyMemberReportAddedDo) CreateInBatches(values []*model.FbDailyMemberReportAdded, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbDailyMemberReportAddedDo) Save(values ...*model.FbDailyMemberReportAdded) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbDailyMemberReportAddedDo) First() (*model.FbDailyMemberReportAdded, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDailyMemberReportAdded), nil
	}
}

func (f fbDailyMemberReportAddedDo) Take() (*model.FbDailyMemberReportAdded, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDailyMemberReportAdded), nil
	}
}

func (f fbDailyMemberReportAddedDo) Last() (*model.FbDailyMemberReportAdded, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDailyMemberReportAdded), nil
	}
}

func (f fbDailyMemberReportAddedDo) Find() ([]*model.FbDailyMemberReportAdded, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbDailyMemberReportAdded), err
}

func (f fbDailyMemberReportAddedDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbDailyMemberReportAdded, err error) {
	buf := make([]*model.FbDailyMemberReportAdded, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbDailyMemberReportAddedDo) FindInBatches(result *[]*model.FbDailyMemberReportAdded, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbDailyMemberReportAddedDo) Attrs(attrs ...field.AssignExpr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbDailyMemberReportAddedDo) Assign(attrs ...field.AssignExpr) IFbDailyMemberReportAddedDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbDailyMemberReportAddedDo) Joins(fields ...field.RelationField) IFbDailyMemberReportAddedDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbDailyMemberReportAddedDo) Preload(fields ...field.RelationField) IFbDailyMemberReportAddedDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbDailyMemberReportAddedDo) FirstOrInit() (*model.FbDailyMemberReportAdded, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDailyMemberReportAdded), nil
	}
}

func (f fbDailyMemberReportAddedDo) FirstOrCreate() (*model.FbDailyMemberReportAdded, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbDailyMemberReportAdded), nil
	}
}

func (f fbDailyMemberReportAddedDo) FindByPage(offset int, limit int) (result []*model.FbDailyMemberReportAdded, count int64, err error) {
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

func (f fbDailyMemberReportAddedDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbDailyMemberReportAddedDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbDailyMemberReportAddedDo) Delete(models ...*model.FbDailyMemberReportAdded) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbDailyMemberReportAddedDo) withDO(do gen.Dao) *fbDailyMemberReportAddedDo {
	f.DO = *do.(*gen.DO)
	return f
}
