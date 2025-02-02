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

func newNodeAttr(db *gorm.DB, opts ...gen.DOOption) nodeAttr {
	_nodeAttr := nodeAttr{}

	_nodeAttr.nodeAttrDo.UseDB(db, opts...)
	_nodeAttr.nodeAttrDo.UseModel(&model.NodeAttr{})

	tableName := _nodeAttr.nodeAttrDo.TableName()
	_nodeAttr.ALL = field.NewAsterisk(tableName)
	_nodeAttr.ID = field.NewInt64(tableName, "id")
	_nodeAttr.NodeID = field.NewInt64(tableName, "nodeID")
	_nodeAttr.Name = field.NewString(tableName, "name")
	_nodeAttr.Desc = field.NewString(tableName, "desc")
	_nodeAttr.Type = field.NewInt64(tableName, "type")

	_nodeAttr.fillFieldMap()

	return _nodeAttr
}

type nodeAttr struct {
	nodeAttrDo

	ALL    field.Asterisk
	ID     field.Int64
	NodeID field.Int64
	Name   field.String
	Desc   field.String
	Type   field.Int64

	fieldMap map[string]field.Expr
}

func (n nodeAttr) Table(newTableName string) *nodeAttr {
	n.nodeAttrDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n nodeAttr) As(alias string) *nodeAttr {
	n.nodeAttrDo.DO = *(n.nodeAttrDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *nodeAttr) updateTableName(table string) *nodeAttr {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewInt64(table, "id")
	n.NodeID = field.NewInt64(table, "nodeID")
	n.Name = field.NewString(table, "name")
	n.Desc = field.NewString(table, "desc")
	n.Type = field.NewInt64(table, "type")

	n.fillFieldMap()

	return n
}

func (n *nodeAttr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *nodeAttr) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 5)
	n.fieldMap["id"] = n.ID
	n.fieldMap["nodeID"] = n.NodeID
	n.fieldMap["name"] = n.Name
	n.fieldMap["desc"] = n.Desc
	n.fieldMap["type"] = n.Type
}

func (n nodeAttr) clone(db *gorm.DB) nodeAttr {
	n.nodeAttrDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n nodeAttr) replaceDB(db *gorm.DB) nodeAttr {
	n.nodeAttrDo.ReplaceDB(db)
	return n
}

type nodeAttrDo struct{ gen.DO }

type INodeAttrDo interface {
	gen.SubQuery
	Debug() INodeAttrDo
	WithContext(ctx context.Context) INodeAttrDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INodeAttrDo
	WriteDB() INodeAttrDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INodeAttrDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INodeAttrDo
	Not(conds ...gen.Condition) INodeAttrDo
	Or(conds ...gen.Condition) INodeAttrDo
	Select(conds ...field.Expr) INodeAttrDo
	Where(conds ...gen.Condition) INodeAttrDo
	Order(conds ...field.Expr) INodeAttrDo
	Distinct(cols ...field.Expr) INodeAttrDo
	Omit(cols ...field.Expr) INodeAttrDo
	Join(table schema.Tabler, on ...field.Expr) INodeAttrDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INodeAttrDo
	RightJoin(table schema.Tabler, on ...field.Expr) INodeAttrDo
	Group(cols ...field.Expr) INodeAttrDo
	Having(conds ...gen.Condition) INodeAttrDo
	Limit(limit int) INodeAttrDo
	Offset(offset int) INodeAttrDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INodeAttrDo
	Unscoped() INodeAttrDo
	Create(values ...*model.NodeAttr) error
	CreateInBatches(values []*model.NodeAttr, batchSize int) error
	Save(values ...*model.NodeAttr) error
	First() (*model.NodeAttr, error)
	Take() (*model.NodeAttr, error)
	Last() (*model.NodeAttr, error)
	Find() ([]*model.NodeAttr, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NodeAttr, err error)
	FindInBatches(result *[]*model.NodeAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NodeAttr) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INodeAttrDo
	Assign(attrs ...field.AssignExpr) INodeAttrDo
	Joins(fields ...field.RelationField) INodeAttrDo
	Preload(fields ...field.RelationField) INodeAttrDo
	FirstOrInit() (*model.NodeAttr, error)
	FirstOrCreate() (*model.NodeAttr, error)
	FindByPage(offset int, limit int) (result []*model.NodeAttr, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INodeAttrDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n nodeAttrDo) Debug() INodeAttrDo {
	return n.withDO(n.DO.Debug())
}

func (n nodeAttrDo) WithContext(ctx context.Context) INodeAttrDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n nodeAttrDo) ReadDB() INodeAttrDo {
	return n.Clauses(dbresolver.Read)
}

func (n nodeAttrDo) WriteDB() INodeAttrDo {
	return n.Clauses(dbresolver.Write)
}

func (n nodeAttrDo) Session(config *gorm.Session) INodeAttrDo {
	return n.withDO(n.DO.Session(config))
}

func (n nodeAttrDo) Clauses(conds ...clause.Expression) INodeAttrDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n nodeAttrDo) Returning(value interface{}, columns ...string) INodeAttrDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n nodeAttrDo) Not(conds ...gen.Condition) INodeAttrDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n nodeAttrDo) Or(conds ...gen.Condition) INodeAttrDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n nodeAttrDo) Select(conds ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n nodeAttrDo) Where(conds ...gen.Condition) INodeAttrDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n nodeAttrDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INodeAttrDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n nodeAttrDo) Order(conds ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n nodeAttrDo) Distinct(cols ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n nodeAttrDo) Omit(cols ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n nodeAttrDo) Join(table schema.Tabler, on ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n nodeAttrDo) LeftJoin(table schema.Tabler, on ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n nodeAttrDo) RightJoin(table schema.Tabler, on ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n nodeAttrDo) Group(cols ...field.Expr) INodeAttrDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n nodeAttrDo) Having(conds ...gen.Condition) INodeAttrDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n nodeAttrDo) Limit(limit int) INodeAttrDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n nodeAttrDo) Offset(offset int) INodeAttrDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n nodeAttrDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INodeAttrDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n nodeAttrDo) Unscoped() INodeAttrDo {
	return n.withDO(n.DO.Unscoped())
}

func (n nodeAttrDo) Create(values ...*model.NodeAttr) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n nodeAttrDo) CreateInBatches(values []*model.NodeAttr, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n nodeAttrDo) Save(values ...*model.NodeAttr) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n nodeAttrDo) First() (*model.NodeAttr, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodeAttr), nil
	}
}

func (n nodeAttrDo) Take() (*model.NodeAttr, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodeAttr), nil
	}
}

func (n nodeAttrDo) Last() (*model.NodeAttr, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodeAttr), nil
	}
}

func (n nodeAttrDo) Find() ([]*model.NodeAttr, error) {
	result, err := n.DO.Find()
	return result.([]*model.NodeAttr), err
}

func (n nodeAttrDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NodeAttr, err error) {
	buf := make([]*model.NodeAttr, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n nodeAttrDo) FindInBatches(result *[]*model.NodeAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n nodeAttrDo) Attrs(attrs ...field.AssignExpr) INodeAttrDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n nodeAttrDo) Assign(attrs ...field.AssignExpr) INodeAttrDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n nodeAttrDo) Joins(fields ...field.RelationField) INodeAttrDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n nodeAttrDo) Preload(fields ...field.RelationField) INodeAttrDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n nodeAttrDo) FirstOrInit() (*model.NodeAttr, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodeAttr), nil
	}
}

func (n nodeAttrDo) FirstOrCreate() (*model.NodeAttr, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NodeAttr), nil
	}
}

func (n nodeAttrDo) FindByPage(offset int, limit int) (result []*model.NodeAttr, count int64, err error) {
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

func (n nodeAttrDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n nodeAttrDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n nodeAttrDo) Delete(models ...*model.NodeAttr) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *nodeAttrDo) withDO(do gen.Dao) *nodeAttrDo {
	n.DO = *do.(*gen.DO)
	return n
}
