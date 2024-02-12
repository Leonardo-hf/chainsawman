package main

import (
	"chainsawman/graph/model"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"os"
	"strings"
)

type GroupDoc struct {
	model.Group
	Tags    []model.Algo
	Algos   []model.Algo
	Extends string
}

func blank2Null(s string) string {
	if len(s) == 0 {
		return "null"
	}
	return strings.Replace(fmt.Sprintf("'%v'", s), "\\", "\\\\", -1)
}

func main() {
	indir := "./graph/scripts/gjson"
	outDir := "./graph/scripts/gsql"

	files, err := os.ReadDir(indir)
	if err != nil {
		fmt.Printf("fail to read dir: %v, err: %v\n", indir, err)
		return
	}
	graphs := make(map[string]*GroupDoc)
	tags := make(map[string]*model.Algo)
	// 收集所有文件中的GROUP
	for _, name := range files {
		g := &GroupDoc{}
		p := fmt.Sprintf("%v/%v", indir, name.Name())
		if !strings.HasSuffix(p, ".json") {
			continue
		}
		c, _ := os.ReadFile(p)
		err = jsonx.UnmarshalFromString(string(c), g)
		if err != nil {
			fmt.Printf("fail to parse %v, err: %v\n", p, err)
			return
		}
		if len(g.Name) == 0 {
			fmt.Printf("`name` is required in %v", p)
		}
		graphs[g.Name] = g
		for _, t := range g.Tags {
			tags[t.Name] = &t
		}
	}

	// 创建输出目录文件夹
	_, err = os.ReadDir(outDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			fmt.Printf("fail to create output dir, path: %v, err: %v", outDir, err)
			return
		}
	} else if err != nil {
		fmt.Printf("fail to read dir: %v, err: %v\n", outDir, err)
		return
	}

	gid, nid, eid, aid := 1, 1, 1, 1

	// 收集所有JSON中的TAG，统一写入文件
	tagSql := ""
	for _, tag := range tags {
		if len(tag.Name) == 0 {
			fmt.Printf("`tag.name` is required in group %v\n", tag.Name)
		}
		tag.ID = int64(aid)
		tagSql += fmt.Sprintf("INSERT INTO graph.algo(id, name, define, detail, isTag) VALUES (%v, %v, %v, %v, 1);\n\n",
			aid, blank2Null(tag.Name), blank2Null(tag.Define), blank2Null(tag.Detail))
		aid += 1
	}
	p := fmt.Sprintf("%v/tag.sql", outDir)
	err = os.WriteFile(p, []byte(tagSql), os.ModePerm)
	if err != nil {
		fmt.Printf("fail to write file, path: %v, err: %v\n", p, err)
		return
	}

	// 将每个GROUP的JSON转化为SQL写入文件
	var handle func(v *GroupDoc)
	handle = func(v *GroupDoc) {
		// 已经处理过的GROUP，直接返回
		if v.ID != 0 {
			return
		}
		sql := ""
		// 生成写入GROUP SCHEMA的SQL
		if len(v.Extends) > 0 {
			if t, ok := graphs[v.Extends]; !ok {
				fmt.Printf("fail to find valid `extends` in group: %v\n", v.Name)
			} else {
				// 优先处理继承的GROUP的JSON
				handle(t)
			}
			sql += fmt.Sprintf("INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (%v, %v, %v, %v);\n\n",
				gid, blank2Null(v.Name), blank2Null(v.Desc), int(graphs[v.Extends].ID))
		} else {
			sql += fmt.Sprintf("INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (%v, %v, %v, null);\n\n",
				gid, blank2Null(v.Name), blank2Null(v.Desc))
		}
		v.ID = int64(gid)
		// 生成写入NODE, NODE_ATTR SCHEMA的SQL
		handleNode := func(nodes []*model.Node) {
			for _, node := range nodes {
				if len(node.Name) == 0 {
					fmt.Printf("`node.name` is required in group %v\n", v.Name)
					return
				}
				sql += fmt.Sprintf("INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (%v, %v, %v, %v, %v);\n",
					nid, gid, blank2Null(node.Name), blank2Null(node.Desc), blank2Null(node.Primary))
				for _, attr := range node.Attrs {
					if len(attr.Name) == 0 {
						fmt.Printf("`node.attr.name` is required in group %v\n", v.Name)
						return
					}
					sql += fmt.Sprintf("INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (%v, %v, %v, %v);\n",
						nid, blank2Null(attr.Name), blank2Null(attr.Desc), attr.Type)
				}
				sql += "\n"
				nid++
			}

		}
		// 生成写入EDGE, EDGE_ATTR SCHEMA的SQL
		handleEdge := func(edges []*model.Edge) {
			for _, edge := range edges {
				if len(edge.Name) == 0 {
					fmt.Printf("`edge.name` is required in group %v\n", v.Name)
					return
				}
				sql += fmt.Sprintf("INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (%v, %v, %v, %v, %v, %v);\n\n",
					eid, gid, blank2Null(edge.Name), blank2Null(edge.Desc), blank2Null(edge.Primary), edge.Direct)
				for _, attr := range edge.Attrs {
					if len(attr.Name) == 0 {
						fmt.Printf("`edge.attr.name` is required in group %v\n", v.Name)
						return
					}
					sql += fmt.Sprintf("INSERT INTO graph.edgeAttr(edgeID, name, `desc`, type) VALUES (%v, %v, %v, %v);\n",
						eid, blank2Null(attr.Name), blank2Null(attr.Desc), attr.Type)
				}
				sql += "\n"
				eid++
			}
		}
		handleNode(v.Nodes)
		// 处理继承来的NODE
		if len(v.Extends) > 0 {
			handleNode(graphs[v.Extends].Nodes)
		}
		handleEdge(v.Edges)
		// 处理继承来的EDGE
		if len(v.Extends) > 0 {
			handleEdge(graphs[v.Extends].Edges)
		}
		// 处理GROUP持有的ALGO及ALGO_PARAM
		for _, algo := range v.Algos {
			if len(algo.Define) == 0 {
				fmt.Printf("`algo.define` is required in group %v\n", v.Name)
				return
			}
			if _, ok := tags[algo.Tag]; !ok {
				fmt.Printf("`algo.tag` should be defined before used in group %v: %v\n", v.Name, algo.Tag)
				return
			}
			sql += fmt.Sprintf("INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v);\n\n",
				aid, blank2Null(algo.Name), blank2Null(algo.Define), blank2Null(algo.Detail), gid, blank2Null(algo.Tag), tags[algo.Tag].ID, blank2Null(algo.JarPath), blank2Null(algo.MainClass))
			for _, param := range algo.Params {
				if len(param.Name) == 0 {
					fmt.Printf("`param.name` is required in group %v\n", v.Name)
					return
				}
				sql += fmt.Sprintf("INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (%v, %v, %v, %v, %v, %v, %v);\n",
					aid, blank2Null(param.Name), blank2Null(param.Desc), param.Type, blank2Null(param.Default), blank2Null(param.Min), blank2Null(param.Max))
			}
			sql += "\n"
			aid++
		}
		gid++
		// SQL写入文件
		p := fmt.Sprintf("%v/%v.sql", outDir, v.Name)
		err = os.WriteFile(p, []byte(sql), os.ModePerm)
		if err != nil {
			fmt.Printf("fail to write file, path: %v, err: %v\n", p, err)
			return
		}
	}
	for _, v := range graphs {
		handle(v)
	}
}
