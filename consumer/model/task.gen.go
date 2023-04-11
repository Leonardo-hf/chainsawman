// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTask = "task"

// Task mapped from table <task>
type Task struct {
	ID         int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Params     string `gorm:"column:params;type:varchar(1024)" json:"params"`
	UpdateTime int64  `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"`
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Status     int64  `gorm:"column:status;type:int;not null" json:"status"`
	Result     string `gorm:"column:result;type:text" json:"result"`
	CreateTime int64  `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"`
}

// TableName Task's table name
func (*Task) TableName() string {
	return TableNameTask
}
