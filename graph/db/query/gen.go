// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	Algo       *algo
	AlgosParam *algosParam
	Edge       *edge
	EdgesAttr  *edgesAttr
	Graph      *graph
	Group      *group
	Node       *node
	NodesAttr  *nodesAttr
	Task       *task
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Algo = &Q.Algo
	AlgosParam = &Q.AlgosParam
	Edge = &Q.Edge
	EdgesAttr = &Q.EdgesAttr
	Graph = &Q.Graph
	Group = &Q.Group
	Node = &Q.Node
	NodesAttr = &Q.NodesAttr
	Task = &Q.Task
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		Algo:       newAlgo(db, opts...),
		AlgosParam: newAlgosParam(db, opts...),
		Edge:       newEdge(db, opts...),
		EdgesAttr:  newEdgesAttr(db, opts...),
		Graph:      newGraph(db, opts...),
		Group:      newGroup(db, opts...),
		Node:       newNode(db, opts...),
		NodesAttr:  newNodesAttr(db, opts...),
		Task:       newTask(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Algo       algo
	AlgosParam algosParam
	Edge       edge
	EdgesAttr  edgesAttr
	Graph      graph
	Group      group
	Node       node
	NodesAttr  nodesAttr
	Task       task
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Algo:       q.Algo.clone(db),
		AlgosParam: q.AlgosParam.clone(db),
		Edge:       q.Edge.clone(db),
		EdgesAttr:  q.EdgesAttr.clone(db),
		Graph:      q.Graph.clone(db),
		Group:      q.Group.clone(db),
		Node:       q.Node.clone(db),
		NodesAttr:  q.NodesAttr.clone(db),
		Task:       q.Task.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Algo:       q.Algo.replaceDB(db),
		AlgosParam: q.AlgosParam.replaceDB(db),
		Edge:       q.Edge.replaceDB(db),
		EdgesAttr:  q.EdgesAttr.replaceDB(db),
		Graph:      q.Graph.replaceDB(db),
		Group:      q.Group.replaceDB(db),
		Node:       q.Node.replaceDB(db),
		NodesAttr:  q.NodesAttr.replaceDB(db),
		Task:       q.Task.replaceDB(db),
	}
}

type queryCtx struct {
	Algo       IAlgoDo
	AlgosParam IAlgosParamDo
	Edge       IEdgeDo
	EdgesAttr  IEdgesAttrDo
	Graph      IGraphDo
	Group      IGroupDo
	Node       INodeDo
	NodesAttr  INodesAttrDo
	Task       ITaskDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Algo:       q.Algo.WithContext(ctx),
		AlgosParam: q.AlgosParam.WithContext(ctx),
		Edge:       q.Edge.WithContext(ctx),
		EdgesAttr:  q.EdgesAttr.WithContext(ctx),
		Graph:      q.Graph.WithContext(ctx),
		Group:      q.Group.WithContext(ctx),
		Node:       q.Node.WithContext(ctx),
		NodesAttr:  q.NodesAttr.WithContext(ctx),
		Task:       q.Task.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
