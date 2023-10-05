package main

import (
	"flag"
	"fmt"
	"gorm.io/gen/field"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const addr = "root:12345678@(localhost:3306)/graph?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	flag.Parse()
	db, err := gorm.Open(mysql.Open(addr))
	if err != nil {
		panic(fmt.Errorf("[db] cannot establish db connection, err: %v", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./consumer/task/db/query",
		ModelPkgPath:      "./consumer/task/model",
		Mode:              gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable:     false,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)

	autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "update_time")
		tag.Set("type", "int unsigned")
		tag.Set("autoUpdateTime", "")
		return tag
	})
	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "create_time")
		tag.Set("type", "int unsigned")
		tag.Set("autoCreateTime", "")
		return tag
	})
	//softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")

	fieldOpts := []gen.ModelOpt{autoUpdateTimeField, autoCreateTimeField}

	graphModel := g.GenerateModel("graphs", fieldOpts...)
	taskModel := g.GenerateModel("tasks", fieldOpts...)
	nodeAttrModel := g.GenerateModel("nodes_attr")
	edgeAttrModel := g.GenerateModel("edges_attr")
	nodeIDGormTag := field.NewGormTag()
	nodeIDGormTag.Set("foreignKey", "nodeID")
	nodeModel := g.GenerateModel("nodes", gen.FieldRelate(field.HasMany, "NodeAttrs", nodeAttrModel, &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            nodeIDGormTag,
	}))
	edgeIDGormTag := field.NewGormTag()
	edgeIDGormTag.Set("foreignKey", "edgeID")
	edgeModel := g.GenerateModel("edges", gen.FieldRelate(field.HasMany, "EdgeAttrs", edgeAttrModel, &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            edgeIDGormTag,
	}))
	groupIDGormTag := field.NewGormTag()
	groupIDGormTag.Set("foreignKey", "groupID")
	groupModel := g.GenerateModel("groups",
		gen.FieldRelate(field.HasMany, "Nodes", nodeModel, &field.RelateConfig{
			RelateSlicePointer: true,
			GORMTag:            groupIDGormTag,
		}),
		gen.FieldRelate(field.HasMany, "Edges", edgeModel, &field.RelateConfig{
			RelateSlicePointer: true,
			GORMTag:            groupIDGormTag,
		}))
	algoIDGormTag := field.NewGormTag()
	algoIDGormTag.Set("foreignKey", "algoID")
	algoModel := g.GenerateModel("algos")
	g.ApplyBasic(graphModel, taskModel, groupModel, nodeModel, edgeModel, nodeAttrModel, edgeAttrModel, algoModel)
	g.Execute()
}
