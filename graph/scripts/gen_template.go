package main

import (
	"chainsawman/common"
	"chainsawman/graph/model"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"os"
	"strings"
)

func GetFramework(pack string, params string, resultRow string, nodesDef string, edgesDef string, schema string, input string) string {
	return fmt.Sprintf(`package %s

import applerodite.config.{CommonService, CommonServiceImpl}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.sql.Row
import org.apache.spark.sql.types.{StringType, StructField, StructType}

// 定义算法输入
case class Param(%s)

// 辅助约束算法输出
%s

// 节点和边名称常量
case object GraphView {
  case object Node {
    %s
  }

  case object Edge {
    %s
  }
}

// 模板内部使用的常量
case object Constants {
  %s
}

// 模板
abstract class Template {
  // 输入处理逻辑
  def input(args: Array[String]): Param = {
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")
	%s
  }

  /*
    编写算法需要重载逻辑
    svc: common 包中提供的服务，包含对 spark，minio，mysql，nebula 的访问
    param: 算法输入
   */
  def exec(svc: CommonService, param: Param): Seq[Row]

  // 输出处理逻辑
  def output(svc: CommonService, param: Param, rows: Seq[Row]): Unit = {
    val spark = svc.getSparkSession
    val df = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(rows), Constants.SCHEMA)
    svc.getOSSClient.upload(name = param.target, content = CSVUtil.df2CSV(df))
    svc.getMysqlClient.updateStatusByOutput(param.target)
  }

  // 默认的算法执行流程，即 Input -> Exec -> Output
  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val svc = CommonServiceImpl
    val param = input(args)
    val res = exec(svc, param)
    output(svc, param, res)
    // TODO: 异常处理，defer 关闭 spark
    svc.getSparkSession.stop()
  }
}
`, pack, params, resultRow, nodesDef, edgesDef, schema, input)
}

type AlgoOutputDef struct {
	Id   string
	Name string
	Type int64
}

type AlgoExtra struct {
	model.Algo
	Output []AlgoOutputDef
}

type GroupDocExtra struct {
	model.Group
	Algos   []AlgoExtra
	Extends string
}

func toScalaType(t int64) string {
	switch t {
	case common.TypeInt:
		return "Integer"
	case common.TypeString:
		return "String"
	case common.TypeDouble:
		return "Double"
	case common.TypeStringList:
		return "Seq[String]"
	case common.TypeDoubleList:
		return "Seq[Double]"
	case common.TypeLong:
		return "Long"
	}
	return ""
}
func getParam(fields ...*model.AlgoParam) string {
	params := make([]string, len(fields))
	for i, f := range fields {
		params[i] = fmt.Sprintf("`%s`: %s", f.Name, toScalaType(f.Type))
	}
	params = append(params, "graphID: String")
	params = append(params, "target: String")
	return strings.Join(params, ", ")
}

func getResultRow(defs ...AlgoOutputDef) string {
	c := make([]string, 0)
	tr := make([]string, 0)
	fields := make([]string, 0)
	for _, v := range defs {
		k := v.Id
		fields = append(fields, fmt.Sprintf("`%s`", k))
		c = append(c, fmt.Sprintf("`%s`: %s", k, toScalaType(v.Type)))
		tr = append(tr, fmt.Sprintf("`%s`: String = this.`%s`.toString", k, k))
	}
	return fmt.Sprintf(`
case class ResultRow(%s) {
  def toRow(%s): Row = {
    Row.apply(%s)
  }
}`, strings.Join(c, ", "), strings.Join(tr, ", "), strings.Join(fields, ", "))
}

func getNodeDef(nodes ...*model.Node) string {
	defs := make([]string, len(nodes))
	for i, n := range nodes {
		attrs := make([]string, len(n.Attrs))
		for p, a := range n.Attrs {
			attrs[p] = fmt.Sprintf(`var ATTR_%s = "%s"`, strings.ToUpper(a.Name), a.Name)
		}
		defs[i] = fmt.Sprintf(`
case object %s{
	var NAME = "%s"
	%s
}
`, strings.ToUpper(n.Name), n.Name, strings.Join(attrs, "\n"))
	}
	return strings.Join(defs, "\n")
}

func getEdgeDef(edges ...*model.Edge) string {
	defs := make([]string, len(edges))
	for i, n := range edges {
		attrs := make([]string, len(n.Attrs))
		for p, a := range n.Attrs {
			attrs[p] = fmt.Sprintf(`var ATTR_%s = "%s"`, strings.ToUpper(a.Name), a.Name)
		}
		defs[i] = fmt.Sprintf(`
case object %s{
	var NAME = "%s"
	%s
}
`, strings.ToUpper(n.Name), n.Name, strings.Join(attrs, "\n"))
	}
	return strings.Join(defs, "\n")
}

func getParseScalaParam(name string, t int64) string {
	switch t {
	case common.TypeStringList, common.TypeDoubleList:
		st := "String"
		if t == common.TypeDoubleList {
			st = "Double"
		}
		stat := fmt.Sprintf("val _%s =json.getJSONArray(\"%s\")\n", name, name)
		stat += fmt.Sprintf("val `%s` = Seq.empty"+`
    for (i <- 0 until _%s.size()) {
      %s%s%s :+ _%s.get%s(i)
    }`, name, name, "`", name, "`", name, st)
		return stat
	default:
		st := toScalaType(t)
		return fmt.Sprintf("val `%s`: %s =json.get%s(\"%s\")", name, st, st, name)
	}
}

func getInput(fields ...*model.AlgoParam) string {
	input := make([]string, len(fields))
	c := make([]string, len(fields))
	for i, f := range fields {
		input[i] = getParseScalaParam(f.Name, f.Type)
		c[i] = fmt.Sprintf("%s = %s", f.Name, f.Name)
	}
	c = append(c, "graphID = graphID")
	c = append(c, "target = target")
	return strings.Join(input, "\n") + "\n" + fmt.Sprintf("Param(%s)", strings.Join(c, ", "))
}

func getSchema(def ...AlgoOutputDef) string {
	fieldDef := make([]string, 0)
	structFields := make([]string, 0)
	for _, v := range def {
		k := strings.ToUpper(v.Id)
		fieldDef = append(fieldDef, fmt.Sprintf(`val COL_%s = "%s"`, k, v.Name))
		structFields = append(structFields, fmt.Sprintf(`StructField(COL_%s, StringType, nullable = false)`, k))
	}
	return strings.Join(fieldDef, "\n") + fmt.Sprintf(`
val SCHEMA: StructType = StructType(
    List(
      %s
    )
)`, strings.Join(structFields, ",\n"))
}

func main() {
	indir := flag.String("g", "", "select the path of the dir which defines the group")
	algo := flag.String("a", "", "select the name of the algorithm")
	output := flag.String("o", ".", "select the dir to save the template")
	pack := flag.String("p", "applerodite", "select the package name of the template")
	flag.Parse()
	if *indir == "" {
		fmt.Println("fail: use -g to set the path of the json which defines the group")
		return
	}
	if *algo == "" {
		fmt.Println("fail: use -a to set the name of the algorithm for generating")
		return
	}
	files, err := os.ReadDir(*indir)
	if err != nil {
		fmt.Printf("fail to read dir: %v, err: %v\n", indir, err)
		return
	}
	groups := make(map[string]*GroupDocExtra)
	for _, name := range files {
		g := &GroupDocExtra{}
		p := fmt.Sprintf("%v/%v", *indir, name.Name())
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
		groups[g.Name] = g
	}

	var targetGroup *GroupDocExtra
	var targetAlgo *AlgoExtra
	for _, g := range groups {
		for _, a := range g.Algos {
			if a.Name == *algo {
				targetAlgo = &a
				targetGroup = g
				extends := targetGroup.Extends
				for len(extends) > 0 {
					if parent, ok := groups[extends]; ok {
						if targetGroup.Nodes == nil {
							targetGroup.Nodes = parent.Nodes
						} else {
							targetGroup.Nodes = append(targetGroup.Nodes, parent.Nodes...)
						}
						if targetGroup.Edges == nil {
							targetGroup.Edges = parent.Edges
						} else {
							targetGroup.Edges = append(targetGroup.Edges, parent.Edges...)
						}
						extends = parent.Extends
					} else {
						fmt.Printf("fail to find valid `extends` in group: %v\n", g.Name)
					}
				}
				break
			}
		}
	}
	if targetAlgo == nil {
		fmt.Printf("fail to find algo %v in %v\n", *algo, *indir)
		return
	}
	params := getParam(targetAlgo.Params...)
	input := getInput(targetAlgo.Params...)
	schema := getSchema(targetAlgo.Output...)
	resultRow := getResultRow(targetAlgo.Output...)
	nodesDef := getNodeDef(targetGroup.Nodes...)
	edgesDef := getEdgeDef(targetGroup.Edges...)
	framework := GetFramework(*pack, params, resultRow, nodesDef, edgesDef, schema, input)
	err = os.WriteFile(fmt.Sprintf("%s/Template.scala", *output), []byte(framework), os.ModePerm)
	if err != nil {
		fmt.Printf("fail to write file, path: %v, err: %v\n", *output, err)
		return
	}
}
