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

func newNodesAttr(db *gorm.DB, opts ...gen.DOOption) nodesAttr {
	_nodesAttr := nodesAttr{}

	_nodesAttr.nodesAttrDo.UseDB(db, opts...)
	_nodesAttr.nodesAttrDo.UseModel(&model.NodesAttr{})

	tableName := _nodesAttr.nodesAttrDo.TableName()
	_nodesAttr.ALL = field.NewAsterisk(tableName)
	_nodesAttr.ID = field.NewInt64(tableName, "id")
	_nodesAttr.NodeID = field.NewInt64(tableName, "nodeID")
	_nodesAttr.Name = field.NewString(tableName, "name")
	_nodesAttr.Desc = field.NewString(tableName, "desc")
	_nodesAttr.Type = field.NewInt64(tableName, "type")
	_nodesAttr.Primary = field.NewInt64(tableName, "primary")

	_nodesAttr.fillFieldMap()

	return _nodesAttr
}

type nodesAttr struct {
	nodesAttrDo

	ALL     field.Asterisk
	ID      field.Int64
	NodeID  field.Int64
	Name    field.String
	Desc    field.String
	Type    field.Int64
	Primary field.Int64

	fieldMap map[string]field.Expr
}

func (n nodesAttr) Table(newTableName string) *nodesAttr {
	n.nodesAttrDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n nodesAttr) As(alias string) *nodesAttr {
	n.nodesAttrDo.DO = *(n.nodesAttrDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *nodesAttr) updateTableName(table string) *nodesAttr {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewInt64(table, "id")
	n.NodeID = field.NewInt64(table, "nodeID")
	n.Name = field.NewString(table, "name")
	n.Desc = field.NewString(table, "desc")
	n.Type = field.NewInt64(table, "type")
	n.Primary = field.NewInt64(table, "primary")

	n.fillFieldMap()

	return n
}

func (n *nodesAttr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *nodesAttr) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 6)
	n.fieldMap["id"] = n.ID
	n.fieldMap["nodeID"] = n.NodeID
	n.fieldMap["name"] = n.Name
	n.fieldMap["desc"] = n.Desc
	n.fieldMap["type"] = n.Type
	n.fieldMap["primary"] = n.Primary
}

func (n nodesAttr) clone(db *gorm.DB) nodesAttr {
	n.nodesAttrDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n nodesAttr) replaceDB(db *gorm.DB) nodesAttr {
	n.nodesAttrDo.ReplaceDB(db)
	return n
}

type nodesAttrDo struct{ gen.DO }

type INodesAttrDo interface {
	gen.SubQuery
	Debug() INodesAttrDo
	WithContext(ctx context.Context) INodesAttrDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INodesAttrDo
	WriteDB() INodesAttrDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INodesAttrDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INodesAttrDo
	Not(conds ...gen.Condition) INodesAttrDo
	Or(conds ...gen.Condition) INodesAttrDo
	Select(conds ...field.Expr) INodesAttrDo
	Where(conds ...gen.Condition) INodesAttrDo
	Order(conds ...field.Expr) INodesAttrDo
	Distinct(cols ...field.Expr) INodesAttrDo
	Omit(cols ...field.Expr) INodesAttrDo
	Join(table schema.Tabler, on ...field.Expr) INodesAttrDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INodesAttrDo
	RightJoin(table schema.Tabler, on ...field.Expr) INodesAttrDo
	Group(cols ...field.Expr) INodesAttrDo
	Having(conds ...gen.Condition) INodesAttrDo
	Limit(limit int) INodesAttrDo
	Offset(offset int) INodesAttrDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INodesAttrDo
	Unscoped() INodesAttrDo
	Create(values ...*model.NodesAttr) error
	CreateInBatches(values []*model.NodesAttr, batchSize int) error
	Save(values ...*model.NodesAttr) error
	First() (*model.NodesAttr, error)
	Take() (*model.NodesAttr, error)
	Last() (*model.NodesAttr, error)
	Find() ([]*model.NodesAttr, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NodesAttr, err error)
	FindInBatches(result *[]*model.NodesAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NodesAttr) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INodesAttrDo
	Assign(attrs ...field.AssignExpr) INodesAttrDo
	Joins(fields ...field.RelationField) INodesAttrDo
	Preload(fields ...field.RelationField) INodesAttrDo
	FirstOrInit() (*model.NodesAttr, error)
	FirstOrCreate() (*model.NodesAttr, error)
	FindByPage(offset int, limit int) (result []*model.NodesAttr, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INodesAttrDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n nodesAttrDo) Debug() INodesAttrDo {
	return n.withDO(n.DO.Debug())
}

func (n nodesAttrDo) WithContext(ctx context.Context) INodesAttrDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n nodesAttrDo) ReadDB() INodesAttrDo {
	return n.Clauses(dbresolver.Read)
}

func (n nodesAttrDo) WriteDB() INodesAttrDo {
	return n.Clauses(dbresolver.Write)
}

func (n nodesAttrDo) Session(config *gorm.Session) INodesAttrDo {
	return n.withDO(n.DO.Session(config))
}

func (n nodesAttrDo) Clauses(conds ...clause.Expression) INodesAttrDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n nodesAttrDo) Returning(value interface{}, columns ...string) INodesAttrDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n nodesAttrDo) Not(conds ...gen.Condition) INodesAttrDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n nodesAttrDo) Or(conds ...gen.Condition) INodesAttrDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n nodesAttrDo) Select(conds ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n nodesAttrDo) Where(conds ...gen.Condition) INodesAttrDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n nodesAttrDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INodesAttrDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n nodesAttrDo) Order(conds ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n nodesAttrDo) Distinct(cols ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n nodesAttrDo) Omit(cols ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n nodesAttrDo) Join(table schema.Tabler, on ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n nodesAttrDo) LeftJoin(table schema.Tabler, on ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n nodesAttrDo) RightJoin(table schema.Tabler, on ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n nodesAttrDo) Group(cols ...field.Expr) INodesAttrDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n nodesAttrDo) Having(conds ...gen.Condition) INodesAttrDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n nodesAttrDo) Limit(limit int) INodesAttrDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n nodesAttrDo) Offset(offset int) INodesAttrDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n nodesAttrDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INodesAttrDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n nodesAttrDo) Unscoped() INodesAttrDo {
	return n.withDO(n.DO.Unscoped())
}

func (n nodesAttrDo) Create(values ...*model.NodesAttr) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n nodesAttrDo) CreateInBatches(values []*model.NodesAttr, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n nodesAttrDo) Save(values ...*model.NodesAttr) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n nodesAttrDo) First() (*model.NodesAttr, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodesAttr), nil
	}
}

func (n nodesAttrDo) Take() (*model.NodesAttr, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodesAttr), nil
	}
}

func (n nodesAttrDo) Last() (*model.NodesAttr, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodesAttr), nil
	}
}

func (n nodesAttrDo) Find() ([]*model.NodesAttr, error) {
	result, err := n.DO.Find()
	return result.([]*model.NodesAttr), err
}

func (n nodesAttrDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NodesAttr, err error) {
	buf := make([]*model.NodesAttr, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n nodesAttrDo) FindInBatches(result *[]*model.NodesAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n nodesAttrDo) Attrs(attrs ...field.AssignExpr) INodesAttrDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n nodesAttrDo) Assign(attrs ...field.AssignExpr) INodesAttrDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n nodesAttrDo) Joins(fields ...field.RelationField) INodesAttrDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n nodesAttrDo) Preload(fields ...field.RelationField) INodesAttrDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n nodesAttrDo) FirstOrInit() (*model.NodesAttr, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodesAttr), nil
	}
}

func (n nodesAttrDo) FirstOrCreate() (*model.NodesAttr, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodesAttr), nil
	}
}

func (n nodesAttrDo) FindByPage(offset int, limit int) (result []*model.NodesAttr, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n nodesAttrDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n nodesAttrDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n nodesAttrDo) Delete(models ...*model.NodesAttr) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *nodesAttrDo) withDO(do gen.Dao) *nodesAttrDo {
	n.DO = *do.(*gen.DO)
	return n
}
