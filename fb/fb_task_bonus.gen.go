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

func newFbTaskBonu(db *gorm.DB, opts ...gen.DOOption) fbTaskBonu {
	_fbTaskBonu := fbTaskBonu{}

	_fbTaskBonu.fbTaskBonuDo.UseDB(db, opts...)
	_fbTaskBonu.fbTaskBonuDo.UseModel(&model.FbTaskBonu{})

	tableName := _fbTaskBonu.fbTaskBonuDo.TableName()
	_fbTaskBonu.ALL = field.NewAsterisk(tableName)
	_fbTaskBonu.ID = field.NewInt64(tableName, "id")
	_fbTaskBonu.UID = field.NewInt64(tableName, "uid")
	_fbTaskBonu.Username = field.NewString(tableName, "username")
	_fbTaskBonu.ParentUID = field.NewString(tableName, "parent_uid")
	_fbTaskBonu.ParentName = field.NewString(tableName, "parent_name")
	_fbTaskBonu.Pid = field.NewInt64(tableName, "pid")
	_fbTaskBonu.Name = field.NewString(tableName, "name")
	_fbTaskBonu.Ty = field.NewInt32(tableName, "ty")
	_fbTaskBonu.Bonus = field.NewFloat64(tableName, "bonus")
	_fbTaskBonu.State = field.NewInt32(tableName, "state")
	_fbTaskBonu.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbTaskBonu.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_fbTaskBonu.ReviewName = field.NewString(tableName, "review_name")
	_fbTaskBonu.ReviewAt = field.NewInt64(tableName, "review_at")
	_fbTaskBonu.ReviewUID = field.NewInt64(tableName, "review_uid")
	_fbTaskBonu.Remark = field.NewString(tableName, "remark")
	_fbTaskBonu.Day = field.NewInt32(tableName, "day")
	_fbTaskBonu.ReceiveAt = field.NewInt64(tableName, "receive_at")
	_fbTaskBonu.Phone = field.NewString(tableName, "phone")
	_fbTaskBonu.UniqueID = field.NewString(tableName, "unique_id")

	_fbTaskBonu.fillFieldMap()

	return _fbTaskBonu
}

// fbTaskBonu 任务领取记录表
type fbTaskBonu struct {
	fbTaskBonuDo

	ALL        field.Asterisk
	ID         field.Int64
	UID        field.Int64  // 用户ID
	Username   field.String // 用户名
	ParentUID  field.String // 上级UID
	ParentName field.String // 上级用户名
	Pid        field.Int64  // 任务ID
	Name       field.String
	Ty         field.Int32   // 活动类型
	Bonus      field.Float64 // 活动奖金金额
	State      field.Int32   // 状态 1:已领取 2:待审核，3.审核拒绝
	CreatedAt  field.Int64   // 创建时间
	UpdatedAt  field.Int64   // 更新时间
	ReviewName field.String
	ReviewAt   field.Int64
	ReviewUID  field.Int64  // 审核人uid
	Remark     field.String // 备注
	Day        field.Int32
	ReceiveAt  field.Int64
	Phone      field.String
	UniqueID   field.String // 唯一ID

	fieldMap map[string]field.Expr
}

func (f fbTaskBonu) Table(newTableName string) *fbTaskBonu {
	f.fbTaskBonuDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbTaskBonu) As(alias string) *fbTaskBonu {
	f.fbTaskBonuDo.DO = *(f.fbTaskBonuDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbTaskBonu) updateTableName(table string) *fbTaskBonu {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.UID = field.NewInt64(table, "uid")
	f.Username = field.NewString(table, "username")
	f.ParentUID = field.NewString(table, "parent_uid")
	f.ParentName = field.NewString(table, "parent_name")
	f.Pid = field.NewInt64(table, "pid")
	f.Name = field.NewString(table, "name")
	f.Ty = field.NewInt32(table, "ty")
	f.Bonus = field.NewFloat64(table, "bonus")
	f.State = field.NewInt32(table, "state")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")
	f.ReviewName = field.NewString(table, "review_name")
	f.ReviewAt = field.NewInt64(table, "review_at")
	f.ReviewUID = field.NewInt64(table, "review_uid")
	f.Remark = field.NewString(table, "remark")
	f.Day = field.NewInt32(table, "day")
	f.ReceiveAt = field.NewInt64(table, "receive_at")
	f.Phone = field.NewString(table, "phone")
	f.UniqueID = field.NewString(table, "unique_id")

	f.fillFieldMap()

	return f
}

func (f *fbTaskBonu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbTaskBonu) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 20)
	f.fieldMap["id"] = f.ID
	f.fieldMap["uid"] = f.UID
	f.fieldMap["username"] = f.Username
	f.fieldMap["parent_uid"] = f.ParentUID
	f.fieldMap["parent_name"] = f.ParentName
	f.fieldMap["pid"] = f.Pid
	f.fieldMap["name"] = f.Name
	f.fieldMap["ty"] = f.Ty
	f.fieldMap["bonus"] = f.Bonus
	f.fieldMap["state"] = f.State
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["review_name"] = f.ReviewName
	f.fieldMap["review_at"] = f.ReviewAt
	f.fieldMap["review_uid"] = f.ReviewUID
	f.fieldMap["remark"] = f.Remark
	f.fieldMap["day"] = f.Day
	f.fieldMap["receive_at"] = f.ReceiveAt
	f.fieldMap["phone"] = f.Phone
	f.fieldMap["unique_id"] = f.UniqueID
}

func (f fbTaskBonu) clone(db *gorm.DB) fbTaskBonu {
	f.fbTaskBonuDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbTaskBonu) replaceDB(db *gorm.DB) fbTaskBonu {
	f.fbTaskBonuDo.ReplaceDB(db)
	return f
}

type fbTaskBonuDo struct{ gen.DO }

type IFbTaskBonuDo interface {
	gen.SubQuery
	Debug() IFbTaskBonuDo
	WithContext(ctx context.Context) IFbTaskBonuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbTaskBonuDo
	WriteDB() IFbTaskBonuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbTaskBonuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbTaskBonuDo
	Not(conds ...gen.Condition) IFbTaskBonuDo
	Or(conds ...gen.Condition) IFbTaskBonuDo
	Select(conds ...field.Expr) IFbTaskBonuDo
	Where(conds ...gen.Condition) IFbTaskBonuDo
	Order(conds ...field.Expr) IFbTaskBonuDo
	Distinct(cols ...field.Expr) IFbTaskBonuDo
	Omit(cols ...field.Expr) IFbTaskBonuDo
	Join(table schema.Tabler, on ...field.Expr) IFbTaskBonuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbTaskBonuDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbTaskBonuDo
	Group(cols ...field.Expr) IFbTaskBonuDo
	Having(conds ...gen.Condition) IFbTaskBonuDo
	Limit(limit int) IFbTaskBonuDo
	Offset(offset int) IFbTaskBonuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbTaskBonuDo
	Unscoped() IFbTaskBonuDo
	Create(values ...*model.FbTaskBonu) error
	CreateInBatches(values []*model.FbTaskBonu, batchSize int) error
	Save(values ...*model.FbTaskBonu) error
	First() (*model.FbTaskBonu, error)
	Take() (*model.FbTaskBonu, error)
	Last() (*model.FbTaskBonu, error)
	Find() ([]*model.FbTaskBonu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbTaskBonu, err error)
	FindInBatches(result *[]*model.FbTaskBonu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbTaskBonu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbTaskBonuDo
	Assign(attrs ...field.AssignExpr) IFbTaskBonuDo
	Joins(fields ...field.RelationField) IFbTaskBonuDo
	Preload(fields ...field.RelationField) IFbTaskBonuDo
	FirstOrInit() (*model.FbTaskBonu, error)
	FirstOrCreate() (*model.FbTaskBonu, error)
	FindByPage(offset int, limit int) (result []*model.FbTaskBonu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbTaskBonuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbTaskBonuDo) Debug() IFbTaskBonuDo {
	return f.withDO(f.DO.Debug())
}

func (f fbTaskBonuDo) WithContext(ctx context.Context) IFbTaskBonuDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbTaskBonuDo) ReadDB() IFbTaskBonuDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbTaskBonuDo) WriteDB() IFbTaskBonuDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbTaskBonuDo) Session(config *gorm.Session) IFbTaskBonuDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbTaskBonuDo) Clauses(conds ...clause.Expression) IFbTaskBonuDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbTaskBonuDo) Returning(value interface{}, columns ...string) IFbTaskBonuDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbTaskBonuDo) Not(conds ...gen.Condition) IFbTaskBonuDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbTaskBonuDo) Or(conds ...gen.Condition) IFbTaskBonuDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbTaskBonuDo) Select(conds ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbTaskBonuDo) Where(conds ...gen.Condition) IFbTaskBonuDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbTaskBonuDo) Order(conds ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbTaskBonuDo) Distinct(cols ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbTaskBonuDo) Omit(cols ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbTaskBonuDo) Join(table schema.Tabler, on ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbTaskBonuDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbTaskBonuDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbTaskBonuDo) Group(cols ...field.Expr) IFbTaskBonuDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbTaskBonuDo) Having(conds ...gen.Condition) IFbTaskBonuDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbTaskBonuDo) Limit(limit int) IFbTaskBonuDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbTaskBonuDo) Offset(offset int) IFbTaskBonuDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbTaskBonuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbTaskBonuDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbTaskBonuDo) Unscoped() IFbTaskBonuDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbTaskBonuDo) Create(values ...*model.FbTaskBonu) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbTaskBonuDo) CreateInBatches(values []*model.FbTaskBonu, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbTaskBonuDo) Save(values ...*model.FbTaskBonu) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbTaskBonuDo) First() (*model.FbTaskBonu, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbTaskBonu), nil
	}
}

func (f fbTaskBonuDo) Take() (*model.FbTaskBonu, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbTaskBonu), nil
	}
}

func (f fbTaskBonuDo) Last() (*model.FbTaskBonu, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbTaskBonu), nil
	}
}

func (f fbTaskBonuDo) Find() ([]*model.FbTaskBonu, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbTaskBonu), err
}

func (f fbTaskBonuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbTaskBonu, err error) {
	buf := make([]*model.FbTaskBonu, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbTaskBonuDo) FindInBatches(result *[]*model.FbTaskBonu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbTaskBonuDo) Attrs(attrs ...field.AssignExpr) IFbTaskBonuDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbTaskBonuDo) Assign(attrs ...field.AssignExpr) IFbTaskBonuDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbTaskBonuDo) Joins(fields ...field.RelationField) IFbTaskBonuDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbTaskBonuDo) Preload(fields ...field.RelationField) IFbTaskBonuDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbTaskBonuDo) FirstOrInit() (*model.FbTaskBonu, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbTaskBonu), nil
	}
}

func (f fbTaskBonuDo) FirstOrCreate() (*model.FbTaskBonu, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbTaskBonu), nil
	}
}

func (f fbTaskBonuDo) FindByPage(offset int, limit int) (result []*model.FbTaskBonu, count int64, err error) {
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

func (f fbTaskBonuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbTaskBonuDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbTaskBonuDo) Delete(models ...*model.FbTaskBonu) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbTaskBonuDo) withDO(do gen.Dao) *fbTaskBonuDo {
	f.DO = *do.(*gen.DO)
	return f
}
