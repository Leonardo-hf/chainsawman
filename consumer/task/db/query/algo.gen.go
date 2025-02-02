// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"chainsawman/consumer/task/model"
)

func newAlgo(db *gorm.DB, opts ...gen.DOOption) algo {
	_algo := algo{}

	_algo.algoDo.UseDB(db, opts...)
	_algo.algoDo.UseModel(&model.Algo{})

	tableName := _algo.algoDo.TableName()
	_algo.ALL = field.NewAsterisk(tableName)
	_algo.ID = field.NewInt64(tableName, "id")
	_algo.Name = field.NewString(tableName, "name")
	_algo.Define = field.NewString(tableName, "define")
	_algo.Detail = field.NewString(tableName, "detail")
	_algo.JarPath = field.NewString(tableName, "jarPath")
	_algo.MainClass = field.NewString(tableName, "mainClass")
	_algo.Tag = field.NewString(tableName, "tag")
	_algo.TagID = field.NewInt64(tableName, "tagID")
	_algo.IsTag = field.NewInt64(tableName, "isTag")
	_algo.GroupID = field.NewInt64(tableName, "groupId")

	_algo.fillFieldMap()

	return _algo
}

type algo struct {
	algoDo

	ALL       field.Asterisk
	ID        field.Int64
	Name      field.String
	Define    field.String
	Detail    field.String
	JarPath   field.String
	MainClass field.String
	Tag       field.String
	TagID     field.Int64
	IsTag     field.Int64
	/*
		约束算法应用于某个策略组的图谱，此外：
		1......应用于全部策略组
	*/
	GroupID field.Int64

	fieldMap map[string]field.Expr
}

func (a algo) Table(newTableName string) *algo {
	a.algoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a algo) As(alias string) *algo {
	a.algoDo.DO = *(a.algoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *algo) updateTableName(table string) *algo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Name = field.NewString(table, "name")
	a.Define = field.NewString(table, "define")
	a.Detail = field.NewString(table, "detail")
	a.JarPath = field.NewString(table, "jarPath")
	a.MainClass = field.NewString(table, "mainClass")
	a.Tag = field.NewString(table, "tag")
	a.TagID = field.NewInt64(table, "tagID")
	a.IsTag = field.NewInt64(table, "isTag")
	a.GroupID = field.NewInt64(table, "groupId")

	a.fillFieldMap()

	return a
}

func (a *algo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *algo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["define"] = a.Define
	a.fieldMap["detail"] = a.Detail
	a.fieldMap["jarPath"] = a.JarPath
	a.fieldMap["mainClass"] = a.MainClass
	a.fieldMap["tag"] = a.Tag
	a.fieldMap["tagID"] = a.TagID
	a.fieldMap["isTag"] = a.IsTag
	a.fieldMap["groupId"] = a.GroupID
}

func (a algo) clone(db *gorm.DB) algo {
	a.algoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a algo) replaceDB(db *gorm.DB) algo {
	a.algoDo.ReplaceDB(db)
	return a
}

type algoDo struct{ gen.DO }

type IAlgoDo interface {
	gen.SubQuery
	Debug() IAlgoDo
	WithContext(ctx context.Context) IAlgoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAlgoDo
	WriteDB() IAlgoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAlgoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAlgoDo
	Not(conds ...gen.Condition) IAlgoDo
	Or(conds ...gen.Condition) IAlgoDo
	Select(conds ...field.Expr) IAlgoDo
	Where(conds ...gen.Condition) IAlgoDo
	Order(conds ...field.Expr) IAlgoDo
	Distinct(cols ...field.Expr) IAlgoDo
	Omit(cols ...field.Expr) IAlgoDo
	Join(table schema.Tabler, on ...field.Expr) IAlgoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAlgoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAlgoDo
	Group(cols ...field.Expr) IAlgoDo
	Having(conds ...gen.Condition) IAlgoDo
	Limit(limit int) IAlgoDo
	Offset(offset int) IAlgoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAlgoDo
	Unscoped() IAlgoDo
	Create(values ...*model.Algo) error
	CreateInBatches(values []*model.Algo, batchSize int) error
	Save(values ...*model.Algo) error
	First() (*model.Algo, error)
	Take() (*model.Algo, error)
	Last() (*model.Algo, error)
	Find() ([]*model.Algo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Algo, err error)
	FindInBatches(result *[]*model.Algo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Algo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAlgoDo
	Assign(attrs ...field.AssignExpr) IAlgoDo
	Joins(fields ...field.RelationField) IAlgoDo
	Preload(fields ...field.RelationField) IAlgoDo
	FirstOrInit() (*model.Algo, error)
	FirstOrCreate() (*model.Algo, error)
	FindByPage(offset int, limit int) (result []*model.Algo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAlgoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a algoDo) Debug() IAlgoDo {
	return a.withDO(a.DO.Debug())
}

func (a algoDo) WithContext(ctx context.Context) IAlgoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a algoDo) ReadDB() IAlgoDo {
	return a.Clauses(dbresolver.Read)
}

func (a algoDo) WriteDB() IAlgoDo {
	return a.Clauses(dbresolver.Write)
}

func (a algoDo) Session(config *gorm.Session) IAlgoDo {
	return a.withDO(a.DO.Session(config))
}

func (a algoDo) Clauses(conds ...clause.Expression) IAlgoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a algoDo) Returning(value interface{}, columns ...string) IAlgoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a algoDo) Not(conds ...gen.Condition) IAlgoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a algoDo) Or(conds ...gen.Condition) IAlgoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a algoDo) Select(conds ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a algoDo) Where(conds ...gen.Condition) IAlgoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a algoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAlgoDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a algoDo) Order(conds ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a algoDo) Distinct(cols ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a algoDo) Omit(cols ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a algoDo) Join(table schema.Tabler, on ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a algoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a algoDo) RightJoin(table schema.Tabler, on ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a algoDo) Group(cols ...field.Expr) IAlgoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a algoDo) Having(conds ...gen.Condition) IAlgoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a algoDo) Limit(limit int) IAlgoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a algoDo) Offset(offset int) IAlgoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a algoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAlgoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a algoDo) Unscoped() IAlgoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a algoDo) Create(values ...*model.Algo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a algoDo) CreateInBatches(values []*model.Algo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a algoDo) Save(values ...*model.Algo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a algoDo) First() (*model.Algo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Algo), nil
	}
}

func (a algoDo) Take() (*model.Algo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Algo), nil
	}
}

func (a algoDo) Last() (*model.Algo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Algo), nil
	}
}

func (a algoDo) Find() ([]*model.Algo, error) {
	result, err := a.DO.Find()
	return result.([]*model.Algo), err
}

func (a algoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Algo, err error) {
	buf := make([]*model.Algo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a algoDo) FindInBatches(result *[]*model.Algo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a algoDo) Attrs(attrs ...field.AssignExpr) IAlgoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a algoDo) Assign(attrs ...field.AssignExpr) IAlgoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a algoDo) Joins(fields ...field.RelationField) IAlgoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a algoDo) Preload(fields ...field.RelationField) IAlgoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a algoDo) FirstOrInit() (*model.Algo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Algo), nil
	}
}

func (a algoDo) FirstOrCreate() (*model.Algo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Algo), nil
	}
}

func (a algoDo) FindByPage(offset int, limit int) (result []*model.Algo, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a algoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a algoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a algoDo) Delete(models ...*model.Algo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *algoDo) withDO(do gen.Dao) *algoDo {
	a.DO = *do.(*gen.DO)
	return a
}
