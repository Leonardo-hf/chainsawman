// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGraph = "graphs"

// Graph mapped from table <graphs>
type Graph struct {
	Name   string `gorm:"column:name;type:char(255);not null" json:"name"`
	Nodes  int64  `gorm:"column:nodes;type:int;not null" json:"nodes"`
	Edges  int64  `gorm:"column:edges;type:int;not null" json:"edges"`
	Desc   string `gorm:"column:desc;type:char(255)" json:"desc"`
	Status int64  `gorm:"column:status;type:int;not null" json:"status"`
	ID     int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
}

// TableName Graph's table name
func (*Graph) TableName() string {
	return TableNameGraph
}
