// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGroup = "groups"

// Group mapped from table <groups>
type Group struct {
	ID       int64   `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Name     string  `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Desc     string  `gorm:"column:desc;type:text" json:"desc"`
	ParentID int64   `gorm:"column:parentID;type:int;default:1;comment:标识父策略组，子策略组继承父策略组的全部节点与边缘" json:"parentID"`
	Nodes    []*Node `gorm:"foreignKey:groupID" json:"nodes"`
	Edges    []*Edge `gorm:"foreignKey:groupID" json:"edges"`
}

// TableName Group's table name
func (*Group) TableName() string {
	return TableNameGroup
}
