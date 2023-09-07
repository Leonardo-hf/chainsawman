// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEdge = "edges"

// Edge mapped from table <edges>
type Edge struct {
	ID        int64        `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	GroupID   int64        `gorm:"column:groupID;type:int;not null" json:"groupID"`
	Name      string       `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Desc      string       `gorm:"column:desc;type:text" json:"desc"`
	Direct    int64        `gorm:"column:direct;type:tinyint(1);not null" json:"direct"`
	Display   string       `gorm:"column:display;type:varchar(255)" json:"display"`
	EdgeAttrs []*EdgesAttr `gorm:"foreignKey:edgeID" json:"edge_attrs"`
}

// TableName Edge's table name
func (*Edge) TableName() string {
	return TableNameEdge
}