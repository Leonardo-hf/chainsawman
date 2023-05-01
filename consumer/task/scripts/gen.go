package main

import (
	"chainsawman/consumer/task/config"

	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	flag.Parse()
	var configFile = flag.String("f", "consumer/etc/consumer.yaml", "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	db, err := gorm.Open(mysql.Open(c.Mysql.Addr))
	if err != nil {
		panic(fmt.Errorf("[db] cannot establish db connection, err: %v", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./consumer/db/query",
		ModelPkgPath:      "./consumer/model",
		Mode:              gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable:     false,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	dataMap := map[string]func(detailType string) (dataType string){
		"tinyint":   func(detailType string) (dataType string) { return "int64" },
		"smallint":  func(detailType string) (dataType string) { return "int64" },
		"mediumint": func(detailType string) (dataType string) { return "int64" },
		"bigint":    func(detailType string) (dataType string) { return "int64" },
		"int":       func(detailType string) (dataType string) { return "int64" },
	}

	g.WithDataTypeMap(dataMap)

	//jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `id, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	//})
	autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	//softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")

	fieldOpts := []gen.ModelOpt{autoUpdateTimeField, autoCreateTimeField}

	allModel := g.GenerateAllTable(fieldOpts...)

	g.ApplyBasic(allModel...)

	g.Execute()
}