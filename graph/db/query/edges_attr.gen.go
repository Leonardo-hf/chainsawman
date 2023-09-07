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

	"chainsawman/graph/model"
)

func newEdgesAttr(db *gorm.DB, opts ...gen.DOOption) edgesAttr {
	_edgesAttr := edgesAttr{}

	_edgesAttr.edgesAttrDo.UseDB(db, opts...)
	_edgesAttr.edgesAttrDo.UseModel(&model.EdgesAttr{})

	tableName := _edgesAttr.edgesAttrDo.TableName()
	_edgesAttr.ALL = field.NewAsterisk(tableName)
	_edgesAttr.ID = field.NewInt64(tableName, "id")
	_edgesAttr.EdgeID = field.NewInt64(tableName, "edgeID")
	_edgesAttr.Name = field.NewString(tableName, "name")
	_edgesAttr.Desc = field.NewString(tableName, "desc")
	_edgesAttr.Type = field.NewInt64(tableName, "type")
	_edgesAttr.Primary = field.NewInt64(tableName, "primary")

	_edgesAttr.fillFieldMap()

	return _edgesAttr
}

type edgesAttr struct {
	edgesAttrDo

	ALL     field.Asterisk
	ID      field.Int64
	EdgeID  field.Int64
	Name    field.String
	Desc    field.String
	Type    field.Int64
	Primary field.Int64

	fieldMap map[string]field.Expr
}

func (e edgesAttr) Table(newTableName string) *edgesAttr {
	e.edgesAttrDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e edgesAttr) As(alias string) *edgesAttr {
	e.edgesAttrDo.DO = *(e.edgesAttrDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *edgesAttr) updateTableName(table string) *edgesAttr {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.EdgeID = field.NewInt64(table, "edgeID")
	e.Name = field.NewString(table, "name")
	e.Desc = field.NewString(table, "desc")
	e.Type = field.NewInt64(table, "type")
	e.Primary = field.NewInt64(table, "primary")

	e.fillFieldMap()

	return e
}

func (e *edgesAttr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *edgesAttr) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["id"] = e.ID
	e.fieldMap["edgeID"] = e.EdgeID
	e.fieldMap["name"] = e.Name
	e.fieldMap["desc"] = e.Desc
	e.fieldMap["type"] = e.Type
	e.fieldMap["primary"] = e.Primary
}

func (e edgesAttr) clone(db *gorm.DB) edgesAttr {
	e.edgesAttrDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e edgesAttr) replaceDB(db *gorm.DB) edgesAttr {
	e.edgesAttrDo.ReplaceDB(db)
	return e
}

type edgesAttrDo struct{ gen.DO }

type IEdgesAttrDo interface {
	gen.SubQuery
	Debug() IEdgesAttrDo
	WithContext(ctx context.Context) IEdgesAttrDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEdgesAttrDo
	WriteDB() IEdgesAttrDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEdgesAttrDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEdgesAttrDo
	Not(conds ...gen.Condition) IEdgesAttrDo
	Or(conds ...gen.Condition) IEdgesAttrDo
	Select(conds ...field.Expr) IEdgesAttrDo
	Where(conds ...gen.Condition) IEdgesAttrDo
	Order(conds ...field.Expr) IEdgesAttrDo
	Distinct(cols ...field.Expr) IEdgesAttrDo
	Omit(cols ...field.Expr) IEdgesAttrDo
	Join(table schema.Tabler, on ...field.Expr) IEdgesAttrDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEdgesAttrDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEdgesAttrDo
	Group(cols ...field.Expr) IEdgesAttrDo
	Having(conds ...gen.Condition) IEdgesAttrDo
	Limit(limit int) IEdgesAttrDo
	Offset(offset int) IEdgesAttrDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEdgesAttrDo
	Unscoped() IEdgesAttrDo
	Create(values ...*model.EdgesAttr) error
	CreateInBatches(values []*model.EdgesAttr, batchSize int) error
	Save(values ...*model.EdgesAttr) error
	First() (*model.EdgesAttr, error)
	Take() (*model.EdgesAttr, error)
	Last() (*model.EdgesAttr, error)
	Find() ([]*model.EdgesAttr, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EdgesAttr, err error)
	FindInBatches(result *[]*model.EdgesAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EdgesAttr) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEdgesAttrDo
	Assign(attrs ...field.AssignExpr) IEdgesAttrDo
	Joins(fields ...field.RelationField) IEdgesAttrDo
	Preload(fields ...field.RelationField) IEdgesAttrDo
	FirstOrInit() (*model.EdgesAttr, error)
	FirstOrCreate() (*model.EdgesAttr, error)
	FindByPage(offset int, limit int) (result []*model.EdgesAttr, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEdgesAttrDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e edgesAttrDo) Debug() IEdgesAttrDo {
	return e.withDO(e.DO.Debug())
}

func (e edgesAttrDo) WithContext(ctx context.Context) IEdgesAttrDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e edgesAttrDo) ReadDB() IEdgesAttrDo {
	return e.Clauses(dbresolver.Read)
}

func (e edgesAttrDo) WriteDB() IEdgesAttrDo {
	return e.Clauses(dbresolver.Write)
}

func (e edgesAttrDo) Session(config *gorm.Session) IEdgesAttrDo {
	return e.withDO(e.DO.Session(config))
}

func (e edgesAttrDo) Clauses(conds ...clause.Expression) IEdgesAttrDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e edgesAttrDo) Returning(value interface{}, columns ...string) IEdgesAttrDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e edgesAttrDo) Not(conds ...gen.Condition) IEdgesAttrDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e edgesAttrDo) Or(conds ...gen.Condition) IEdgesAttrDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e edgesAttrDo) Select(conds ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e edgesAttrDo) Where(conds ...gen.Condition) IEdgesAttrDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e edgesAttrDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IEdgesAttrDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e edgesAttrDo) Order(conds ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e edgesAttrDo) Distinct(cols ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e edgesAttrDo) Omit(cols ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e edgesAttrDo) Join(table schema.Tabler, on ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e edgesAttrDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e edgesAttrDo) RightJoin(table schema.Tabler, on ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e edgesAttrDo) Group(cols ...field.Expr) IEdgesAttrDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e edgesAttrDo) Having(conds ...gen.Condition) IEdgesAttrDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e edgesAttrDo) Limit(limit int) IEdgesAttrDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e edgesAttrDo) Offset(offset int) IEdgesAttrDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e edgesAttrDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEdgesAttrDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e edgesAttrDo) Unscoped() IEdgesAttrDo {
	return e.withDO(e.DO.Unscoped())
}

func (e edgesAttrDo) Create(values ...*model.EdgesAttr) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e edgesAttrDo) CreateInBatches(values []*model.EdgesAttr, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e edgesAttrDo) Save(values ...*model.EdgesAttr) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e edgesAttrDo) First() (*model.EdgesAttr, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EdgesAttr), nil
	}
}

func (e edgesAttrDo) Take() (*model.EdgesAttr, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EdgesAttr), nil
	}
}

func (e edgesAttrDo) Last() (*model.EdgesAttr, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EdgesAttr), nil
	}
}

func (e edgesAttrDo) Find() ([]*model.EdgesAttr, error) {
	result, err := e.DO.Find()
	return result.([]*model.EdgesAttr), err
}

func (e edgesAttrDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EdgesAttr, err error) {
	buf := make([]*model.EdgesAttr, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e edgesAttrDo) FindInBatches(result *[]*model.EdgesAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e edgesAttrDo) Attrs(attrs ...field.AssignExpr) IEdgesAttrDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e edgesAttrDo) Assign(attrs ...field.AssignExpr) IEdgesAttrDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e edgesAttrDo) Joins(fields ...field.RelationField) IEdgesAttrDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e edgesAttrDo) Preload(fields ...field.RelationField) IEdgesAttrDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e edgesAttrDo) FirstOrInit() (*model.EdgesAttr, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EdgesAttr), nil
	}
}

func (e edgesAttrDo) FirstOrCreate() (*model.EdgesAttr, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EdgesAttr), nil
	}
}

func (e edgesAttrDo) FindByPage(offset int, limit int) (result []*model.EdgesAttr, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e edgesAttrDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e edgesAttrDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e edgesAttrDo) Delete(models ...*model.EdgesAttr) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *edgesAttrDo) withDO(do gen.Dao) *edgesAttrDo {
	e.DO = *do.(*gen.DO)
	return e
}
