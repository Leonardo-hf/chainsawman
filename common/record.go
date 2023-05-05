package common

import (
	"strconv"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

func ParseInt(record *nebula.Record, field string) int64 {
	v, _ := record.GetValueByColName(field)
	if v.IsInt() {
		vInt, _ := v.AsInt()
		return vInt
	}
	vStr, _ := v.AsString()
	vInt, _ := strconv.ParseInt(vStr, 10, 64)
	return vInt
}

func Parse(record *nebula.Record, field string) string {
	v, _ := record.GetValueByColName(field)
	vStr, _ := v.AsString()
	return vStr
}
