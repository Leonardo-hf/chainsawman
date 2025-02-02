// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGraph = "graph"

// Graph mapped from table <graph>
type Graph struct {
	ID         int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Status     int64     `gorm:"column:status;type:int;not null" json:"status"`
	NumNode    int64     `gorm:"column:numNode;type:int;not null" json:"numNode"`
	NumEdge    int64     `gorm:"column:numEdge;type:int;not null" json:"numEdge"`
	GroupID    int64     `gorm:"column:groupID;type:int;not null" json:"groupID"`
	CreateTime time.Time `gorm:"column:createTime;type:timestamp;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime;type:timestamp;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName Graph's table name
func (*Graph) TableName() string {
	return TableNameGraph
}
