// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTask = "tasks"

// Task mapped from table <tasks>
type Task struct {
	ID         int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Params     string    `gorm:"column:params;type:varchar(1024)" json:"params"`
	Status     int64     `gorm:"column:status;type:int;not null" json:"status"`
	Result     string    `gorm:"column:result;type:mediumtext" json:"result"`
	GraphID    int64     `gorm:"column:graphID;type:int;not null" json:"graphID"`
	Visible    int64     `gorm:"column:visible;type:tinyint(1)" json:"visible"`
	Tid        string    `gorm:"column:tid;type:varchar(255)" json:"tid"`
	Idf        string    `gorm:"column:idf;type:varchar(255);not null" json:"idf"`
	CreateTime time.Time `gorm:"column:createTime;type:datetime(0);default:CURRENT_TIMESTAMP;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime;type:datetime(0);default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updateTime"`
}

// TableName Task's table name
func (*Task) TableName() string {
	return TableNameTask
}
