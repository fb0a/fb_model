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

func newFbPlatform(db *gorm.DB, opts ...gen.DOOption) fbPlatform {
	_fbPlatform := fbPlatform{}

	_fbPlatform.fbPlatformDo.UseDB(db, opts...)
	_fbPlatform.fbPlatformDo.UseModel(&model.FbPlatform{})

	tableName := _fbPlatform.fbPlatformDo.TableName()
	_fbPlatform.ALL = field.NewAsterisk(tableName)
	_fbPlatform.ID = field.NewInt64(tableName, "id")
	_fbPlatform.VenueID = field.NewInt64(tableName, "venue_id")
	_fbPlatform.Name = field.NewString(tableName, "name")
	_fbPlatform.EnName = field.NewString(tableName, "en_name")
	_fbPlatform.ZhName = field.NewString(tableName, "zh_name")
	_fbPlatform.TwName = field.NewString(tableName, "tw_name")
	_fbPlatform.GameType = field.NewInt32(tableName, "game_type")
	_fbPlatform.State = field.NewInt32(tableName, "state")
	_fbPlatform.Maintained = field.NewInt32(tableName, "maintained")
	_fbPlatform.Seq = field.NewInt32(tableName, "seq")
	_fbPlatform.Logo = field.NewString(tableName, "logo")
	_fbPlatform.CreatedAt = field.NewInt32(tableName, "created_at")
	_fbPlatform.UpdatedAt = field.NewInt32(tableName, "updated_at")
	_fbPlatform.UpdatedUID = field.NewInt64(tableName, "updated_uid")
	_fbPlatform.UpdatedName = field.NewString(tableName, "updated_name")
	_fbPlatform.MaintainedSt = field.NewInt64(tableName, "maintained_st")
	_fbPlatform.MaintainedEt = field.NewInt64(tableName, "maintained_et")

	_fbPlatform.fillFieldMap()

	return _fbPlatform
}

// fbPlatform 游戏场馆表
type fbPlatform struct {
	fbPlatformDo

	ALL          field.Asterisk
	ID           field.Int64  // id
	VenueID      field.Int64  // 平台id(PP/EVO等)
	Name         field.String // 厂商名
	EnName       field.String // 场馆名称
	ZhName       field.String // 中文名
	TwName       field.String // 繁体中文名
	GameType     field.Int32  // 游戏类型
	State        field.Int32  // 1上线 2下线
	Maintained   field.Int32  // 维护状态 1正常 2维护中
	Seq          field.Int32  // 场馆排序
	Logo         field.String // 场馆logo
	CreatedAt    field.Int32  // 创建时间
	UpdatedAt    field.Int32  // 更新时间
	UpdatedUID   field.Int64  // 更新人uid
	UpdatedName  field.String
	MaintainedSt field.Int64 // 维护开始时间
	MaintainedEt field.Int64 // 维护结束时间

	fieldMap map[string]field.Expr
}

func (f fbPlatform) Table(newTableName string) *fbPlatform {
	f.fbPlatformDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbPlatform) As(alias string) *fbPlatform {
	f.fbPlatformDo.DO = *(f.fbPlatformDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbPlatform) updateTableName(table string) *fbPlatform {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.VenueID = field.NewInt64(table, "venue_id")
	f.Name = field.NewString(table, "name")
	f.EnName = field.NewString(table, "en_name")
	f.ZhName = field.NewString(table, "zh_name")
	f.TwName = field.NewString(table, "tw_name")
	f.GameType = field.NewInt32(table, "game_type")
	f.State = field.NewInt32(table, "state")
	f.Maintained = field.NewInt32(table, "maintained")
	f.Seq = field.NewInt32(table, "seq")
	f.Logo = field.NewString(table, "logo")
	f.CreatedAt = field.NewInt32(table, "created_at")
	f.UpdatedAt = field.NewInt32(table, "updated_at")
	f.UpdatedUID = field.NewInt64(table, "updated_uid")
	f.UpdatedName = field.NewString(table, "updated_name")
	f.MaintainedSt = field.NewInt64(table, "maintained_st")
	f.MaintainedEt = field.NewInt64(table, "maintained_et")

	f.fillFieldMap()

	return f
}

func (f *fbPlatform) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbPlatform) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 17)
	f.fieldMap["id"] = f.ID
	f.fieldMap["venue_id"] = f.VenueID
	f.fieldMap["name"] = f.Name
	f.fieldMap["en_name"] = f.EnName
	f.fieldMap["zh_name"] = f.ZhName
	f.fieldMap["tw_name"] = f.TwName
	f.fieldMap["game_type"] = f.GameType
	f.fieldMap["state"] = f.State
	f.fieldMap["maintained"] = f.Maintained
	f.fieldMap["seq"] = f.Seq
	f.fieldMap["logo"] = f.Logo
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["updated_uid"] = f.UpdatedUID
	f.fieldMap["updated_name"] = f.UpdatedName
	f.fieldMap["maintained_st"] = f.MaintainedSt
	f.fieldMap["maintained_et"] = f.MaintainedEt
}

func (f fbPlatform) clone(db *gorm.DB) fbPlatform {
	f.fbPlatformDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbPlatform) replaceDB(db *gorm.DB) fbPlatform {
	f.fbPlatformDo.ReplaceDB(db)
	return f
}

type fbPlatformDo struct{ gen.DO }

type IFbPlatformDo interface {
	gen.SubQuery
	Debug() IFbPlatformDo
	WithContext(ctx context.Context) IFbPlatformDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbPlatformDo
	WriteDB() IFbPlatformDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbPlatformDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbPlatformDo
	Not(conds ...gen.Condition) IFbPlatformDo
	Or(conds ...gen.Condition) IFbPlatformDo
	Select(conds ...field.Expr) IFbPlatformDo
	Where(conds ...gen.Condition) IFbPlatformDo
	Order(conds ...field.Expr) IFbPlatformDo
	Distinct(cols ...field.Expr) IFbPlatformDo
	Omit(cols ...field.Expr) IFbPlatformDo
	Join(table schema.Tabler, on ...field.Expr) IFbPlatformDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbPlatformDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbPlatformDo
	Group(cols ...field.Expr) IFbPlatformDo
	Having(conds ...gen.Condition) IFbPlatformDo
	Limit(limit int) IFbPlatformDo
	Offset(offset int) IFbPlatformDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbPlatformDo
	Unscoped() IFbPlatformDo
	Create(values ...*model.FbPlatform) error
	CreateInBatches(values []*model.FbPlatform, batchSize int) error
	Save(values ...*model.FbPlatform) error
	First() (*model.FbPlatform, error)
	Take() (*model.FbPlatform, error)
	Last() (*model.FbPlatform, error)
	Find() ([]*model.FbPlatform, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbPlatform, err error)
	FindInBatches(result *[]*model.FbPlatform, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbPlatform) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbPlatformDo
	Assign(attrs ...field.AssignExpr) IFbPlatformDo
	Joins(fields ...field.RelationField) IFbPlatformDo
	Preload(fields ...field.RelationField) IFbPlatformDo
	FirstOrInit() (*model.FbPlatform, error)
	FirstOrCreate() (*model.FbPlatform, error)
	FindByPage(offset int, limit int) (result []*model.FbPlatform, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbPlatformDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbPlatformDo) Debug() IFbPlatformDo {
	return f.withDO(f.DO.Debug())
}

func (f fbPlatformDo) WithContext(ctx context.Context) IFbPlatformDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbPlatformDo) ReadDB() IFbPlatformDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbPlatformDo) WriteDB() IFbPlatformDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbPlatformDo) Session(config *gorm.Session) IFbPlatformDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbPlatformDo) Clauses(conds ...clause.Expression) IFbPlatformDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbPlatformDo) Returning(value interface{}, columns ...string) IFbPlatformDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbPlatformDo) Not(conds ...gen.Condition) IFbPlatformDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbPlatformDo) Or(conds ...gen.Condition) IFbPlatformDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbPlatformDo) Select(conds ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbPlatformDo) Where(conds ...gen.Condition) IFbPlatformDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbPlatformDo) Order(conds ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbPlatformDo) Distinct(cols ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbPlatformDo) Omit(cols ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbPlatformDo) Join(table schema.Tabler, on ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbPlatformDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbPlatformDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbPlatformDo) Group(cols ...field.Expr) IFbPlatformDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbPlatformDo) Having(conds ...gen.Condition) IFbPlatformDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbPlatformDo) Limit(limit int) IFbPlatformDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbPlatformDo) Offset(offset int) IFbPlatformDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbPlatformDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbPlatformDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbPlatformDo) Unscoped() IFbPlatformDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbPlatformDo) Create(values ...*model.FbPlatform) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbPlatformDo) CreateInBatches(values []*model.FbPlatform, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbPlatformDo) Save(values ...*model.FbPlatform) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbPlatformDo) First() (*model.FbPlatform, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPlatform), nil
	}
}

func (f fbPlatformDo) Take() (*model.FbPlatform, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPlatform), nil
	}
}

func (f fbPlatformDo) Last() (*model.FbPlatform, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPlatform), nil
	}
}

func (f fbPlatformDo) Find() ([]*model.FbPlatform, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbPlatform), err
}

func (f fbPlatformDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbPlatform, err error) {
	buf := make([]*model.FbPlatform, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbPlatformDo) FindInBatches(result *[]*model.FbPlatform, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbPlatformDo) Attrs(attrs ...field.AssignExpr) IFbPlatformDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbPlatformDo) Assign(attrs ...field.AssignExpr) IFbPlatformDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbPlatformDo) Joins(fields ...field.RelationField) IFbPlatformDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbPlatformDo) Preload(fields ...field.RelationField) IFbPlatformDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbPlatformDo) FirstOrInit() (*model.FbPlatform, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPlatform), nil
	}
}

func (f fbPlatformDo) FirstOrCreate() (*model.FbPlatform, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbPlatform), nil
	}
}

func (f fbPlatformDo) FindByPage(offset int, limit int) (result []*model.FbPlatform, count int64, err error) {
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

func (f fbPlatformDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbPlatformDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbPlatformDo) Delete(models ...*model.FbPlatform) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbPlatformDo) withDO(do gen.Dao) *fbPlatformDo {
	f.DO = *do.(*gen.DO)
	return f
}
