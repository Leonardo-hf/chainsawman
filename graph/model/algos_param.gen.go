// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAlgosParam = "algos_param"

// AlgosParam mapped from table <algos_param>
type AlgosParam struct {
	ID        int64   `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	AlgoID    int64   `gorm:"column:algoID;type:int;not null" json:"algoID"`
	FieldName string  `gorm:"column:fieldName;type:varchar(255);not null" json:"fieldName"`
	FieldDesc string  `gorm:"column:fieldDesc;type:varchar(255)" json:"fieldDesc"`
	FieldType int64   `gorm:"column:fieldType;type:int;not null" json:"fieldType"`
	InitValue float64 `gorm:"column:initValue;type:double" json:"initValue"`
	Max       float64 `gorm:"column:max;type:double" json:"max"`
	Min       float64 `gorm:"column:min;type:double" json:"min"`
}

// TableName AlgosParam's table name
func (*AlgosParam) TableName() string {
	return TableNameAlgosParam
}
