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

func newFbReportMemberReg(db *gorm.DB, opts ...gen.DOOption) fbReportMemberReg {
	_fbReportMemberReg := fbReportMemberReg{}

	_fbReportMemberReg.fbReportMemberRegDo.UseDB(db, opts...)
	_fbReportMemberReg.fbReportMemberRegDo.UseModel(&model.FbReportMemberReg{})

	tableName := _fbReportMemberReg.fbReportMemberRegDo.TableName()
	_fbReportMemberReg.ALL = field.NewAsterisk(tableName)
	_fbReportMemberReg.ID = field.NewInt64(tableName, "id")
	_fbReportMemberReg.Day = field.NewInt32(tableName, "day")
	_fbReportMemberReg.Reg = field.NewInt32(tableName, "reg")
	_fbReportMemberReg.RegDetail = field.NewString(tableName, "reg_detail")
	_fbReportMemberReg.RegFirstDeposit = field.NewInt32(tableName, "reg_first_deposit")
	_fbReportMemberReg.RegFirstDepositDetail = field.NewString(tableName, "reg_first_deposit_detail")
	_fbReportMemberReg.ReportType = field.NewBool(tableName, "report_type")
	_fbReportMemberReg.FirstDeposit = field.NewInt32(tableName, "first_deposit")
	_fbReportMemberReg.FirstDepositDetail = field.NewString(tableName, "first_deposit_detail")
	_fbReportMemberReg.FirstDepositAmount = field.NewFloat64(tableName, "first_deposit_amount")
	_fbReportMemberReg.SecondDeposit = field.NewInt32(tableName, "second_deposit")
	_fbReportMemberReg.SecondDepositAmount = field.NewFloat64(tableName, "second_deposit_amount")
	_fbReportMemberReg.ThirdDeposit = field.NewInt32(tableName, "third_deposit")
	_fbReportMemberReg.ThirdDepositAmount = field.NewFloat64(tableName, "third_deposit_amount")
	_fbReportMemberReg.Bonus = field.NewFloat64(tableName, "bonus")
	_fbReportMemberReg.BonusDetail = field.NewString(tableName, "bonus_detail")
	_fbReportMemberReg.CreatedAt = field.NewInt64(tableName, "created_at")

	_fbReportMemberReg.fillFieldMap()

	return _fbReportMemberReg
}

// fbReportMemberReg 新增会员日报表
type fbReportMemberReg struct {
	fbReportMemberRegDo

	ALL                   field.Asterisk
	ID                    field.Int64   // ID
	Day                   field.Int32   // 年月日
	Reg                   field.Int32   // 新增注册数
	RegDetail             field.String  // 注册来源渠道统计
	RegFirstDeposit       field.Int32   // 注册首存数
	RegFirstDepositDetail field.String  // 注册首存明细
	ReportType            field.Bool    // 报表类型，1日报、2周报、3月报
	FirstDeposit          field.Int32   // 首存
	FirstDepositDetail    field.String  // 首存明细
	FirstDepositAmount    field.Float64 // 首存金额
	SecondDeposit         field.Int32   // 二存
	SecondDepositAmount   field.Float64 // 二存金额
	ThirdDeposit          field.Int32   // 三存
	ThirdDepositAmount    field.Float64 // 三存金额
	Bonus                 field.Float64 // 礼金
	BonusDetail           field.String  // 礼金明细
	CreatedAt             field.Int64   // 创建时间

	fieldMap map[string]field.Expr
}

func (f fbReportMemberReg) Table(newTableName string) *fbReportMemberReg {
	f.fbReportMemberRegDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbReportMemberReg) As(alias string) *fbReportMemberReg {
	f.fbReportMemberRegDo.DO = *(f.fbReportMemberRegDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbReportMemberReg) updateTableName(table string) *fbReportMemberReg {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.Day = field.NewInt32(table, "day")
	f.Reg = field.NewInt32(table, "reg")
	f.RegDetail = field.NewString(table, "reg_detail")
	f.RegFirstDeposit = field.NewInt32(table, "reg_first_deposit")
	f.RegFirstDepositDetail = field.NewString(table, "reg_first_deposit_detail")
	f.ReportType = field.NewBool(table, "report_type")
	f.FirstDeposit = field.NewInt32(table, "first_deposit")
	f.FirstDepositDetail = field.NewString(table, "first_deposit_detail")
	f.FirstDepositAmount = field.NewFloat64(table, "first_deposit_amount")
	f.SecondDeposit = field.NewInt32(table, "second_deposit")
	f.SecondDepositAmount = field.NewFloat64(table, "second_deposit_amount")
	f.ThirdDeposit = field.NewInt32(table, "third_deposit")
	f.ThirdDepositAmount = field.NewFloat64(table, "third_deposit_amount")
	f.Bonus = field.NewFloat64(table, "bonus")
	f.BonusDetail = field.NewString(table, "bonus_detail")
	f.CreatedAt = field.NewInt64(table, "created_at")

	f.fillFieldMap()

	return f
}

func (f *fbReportMemberReg) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbReportMemberReg) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 17)
	f.fieldMap["id"] = f.ID
	f.fieldMap["day"] = f.Day
	f.fieldMap["reg"] = f.Reg
	f.fieldMap["reg_detail"] = f.RegDetail
	f.fieldMap["reg_first_deposit"] = f.RegFirstDeposit
	f.fieldMap["reg_first_deposit_detail"] = f.RegFirstDepositDetail
	f.fieldMap["report_type"] = f.ReportType
	f.fieldMap["first_deposit"] = f.FirstDeposit
	f.fieldMap["first_deposit_detail"] = f.FirstDepositDetail
	f.fieldMap["first_deposit_amount"] = f.FirstDepositAmount
	f.fieldMap["second_deposit"] = f.SecondDeposit
	f.fieldMap["second_deposit_amount"] = f.SecondDepositAmount
	f.fieldMap["third_deposit"] = f.ThirdDeposit
	f.fieldMap["third_deposit_amount"] = f.ThirdDepositAmount
	f.fieldMap["bonus"] = f.Bonus
	f.fieldMap["bonus_detail"] = f.BonusDetail
	f.fieldMap["created_at"] = f.CreatedAt
}

func (f fbReportMemberReg) clone(db *gorm.DB) fbReportMemberReg {
	f.fbReportMemberRegDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbReportMemberReg) replaceDB(db *gorm.DB) fbReportMemberReg {
	f.fbReportMemberRegDo.ReplaceDB(db)
	return f
}

type fbReportMemberRegDo struct{ gen.DO }

type IFbReportMemberRegDo interface {
	gen.SubQuery
	Debug() IFbReportMemberRegDo
	WithContext(ctx context.Context) IFbReportMemberRegDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbReportMemberRegDo
	WriteDB() IFbReportMemberRegDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbReportMemberRegDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbReportMemberRegDo
	Not(conds ...gen.Condition) IFbReportMemberRegDo
	Or(conds ...gen.Condition) IFbReportMemberRegDo
	Select(conds ...field.Expr) IFbReportMemberRegDo
	Where(conds ...gen.Condition) IFbReportMemberRegDo
	Order(conds ...field.Expr) IFbReportMemberRegDo
	Distinct(cols ...field.Expr) IFbReportMemberRegDo
	Omit(cols ...field.Expr) IFbReportMemberRegDo
	Join(table schema.Tabler, on ...field.Expr) IFbReportMemberRegDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberRegDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberRegDo
	Group(cols ...field.Expr) IFbReportMemberRegDo
	Having(conds ...gen.Condition) IFbReportMemberRegDo
	Limit(limit int) IFbReportMemberRegDo
	Offset(offset int) IFbReportMemberRegDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbReportMemberRegDo
	Unscoped() IFbReportMemberRegDo
	Create(values ...*model.FbReportMemberReg) error
	CreateInBatches(values []*model.FbReportMemberReg, batchSize int) error
	Save(values ...*model.FbReportMemberReg) error
	First() (*model.FbReportMemberReg, error)
	Take() (*model.FbReportMemberReg, error)
	Last() (*model.FbReportMemberReg, error)
	Find() ([]*model.FbReportMemberReg, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbReportMemberReg, err error)
	FindInBatches(result *[]*model.FbReportMemberReg, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbReportMemberReg) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbReportMemberRegDo
	Assign(attrs ...field.AssignExpr) IFbReportMemberRegDo
	Joins(fields ...field.RelationField) IFbReportMemberRegDo
	Preload(fields ...field.RelationField) IFbReportMemberRegDo
	FirstOrInit() (*model.FbReportMemberReg, error)
	FirstOrCreate() (*model.FbReportMemberReg, error)
	FindByPage(offset int, limit int) (result []*model.FbReportMemberReg, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbReportMemberRegDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbReportMemberRegDo) Debug() IFbReportMemberRegDo {
	return f.withDO(f.DO.Debug())
}

func (f fbReportMemberRegDo) WithContext(ctx context.Context) IFbReportMemberRegDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbReportMemberRegDo) ReadDB() IFbReportMemberRegDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbReportMemberRegDo) WriteDB() IFbReportMemberRegDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbReportMemberRegDo) Session(config *gorm.Session) IFbReportMemberRegDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbReportMemberRegDo) Clauses(conds ...clause.Expression) IFbReportMemberRegDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbReportMemberRegDo) Returning(value interface{}, columns ...string) IFbReportMemberRegDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbReportMemberRegDo) Not(conds ...gen.Condition) IFbReportMemberRegDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbReportMemberRegDo) Or(conds ...gen.Condition) IFbReportMemberRegDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbReportMemberRegDo) Select(conds ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbReportMemberRegDo) Where(conds ...gen.Condition) IFbReportMemberRegDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbReportMemberRegDo) Order(conds ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbReportMemberRegDo) Distinct(cols ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbReportMemberRegDo) Omit(cols ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbReportMemberRegDo) Join(table schema.Tabler, on ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbReportMemberRegDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbReportMemberRegDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbReportMemberRegDo) Group(cols ...field.Expr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbReportMemberRegDo) Having(conds ...gen.Condition) IFbReportMemberRegDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbReportMemberRegDo) Limit(limit int) IFbReportMemberRegDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbReportMemberRegDo) Offset(offset int) IFbReportMemberRegDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbReportMemberRegDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbReportMemberRegDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbReportMemberRegDo) Unscoped() IFbReportMemberRegDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbReportMemberRegDo) Create(values ...*model.FbReportMemberReg) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbReportMemberRegDo) CreateInBatches(values []*model.FbReportMemberReg, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbReportMemberRegDo) Save(values ...*model.FbReportMemberReg) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbReportMemberRegDo) First() (*model.FbReportMemberReg, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberReg), nil
	}
}

func (f fbReportMemberRegDo) Take() (*model.FbReportMemberReg, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberReg), nil
	}
}

func (f fbReportMemberRegDo) Last() (*model.FbReportMemberReg, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberReg), nil
	}
}

func (f fbReportMemberRegDo) Find() ([]*model.FbReportMemberReg, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbReportMemberReg), err
}

func (f fbReportMemberRegDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbReportMemberReg, err error) {
	buf := make([]*model.FbReportMemberReg, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbReportMemberRegDo) FindInBatches(result *[]*model.FbReportMemberReg, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbReportMemberRegDo) Attrs(attrs ...field.AssignExpr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbReportMemberRegDo) Assign(attrs ...field.AssignExpr) IFbReportMemberRegDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbReportMemberRegDo) Joins(fields ...field.RelationField) IFbReportMemberRegDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbReportMemberRegDo) Preload(fields ...field.RelationField) IFbReportMemberRegDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbReportMemberRegDo) FirstOrInit() (*model.FbReportMemberReg, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberReg), nil
	}
}

func (f fbReportMemberRegDo) FirstOrCreate() (*model.FbReportMemberReg, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbReportMemberReg), nil
	}
}

func (f fbReportMemberRegDo) FindByPage(offset int, limit int) (result []*model.FbReportMemberReg, count int64, err error) {
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

func (f fbReportMemberRegDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbReportMemberRegDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbReportMemberRegDo) Delete(models ...*model.FbReportMemberReg) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbReportMemberRegDo) withDO(do gen.Dao) *fbReportMemberRegDo {
	f.DO = *do.(*gen.DO)
	return f
}
