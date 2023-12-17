package applerodite.pagerank

import applerodite.config.{CommonService, CommonServiceImpl}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.sql.Row
import org.apache.spark.sql.types.{StringType, StructField, StructType}

// 定义算法输入
case class Param(`iter`: Integer, `prob`: Double, graphID: String, target: String)

// 辅助约束算法输出

case class ResultRow(`id`: Long, `score`: Double) {
  def toRow(`id`: String = this.`id`.toString, `score`: String = this.`score`.toString): Row = {
    Row.apply(`id`, `score`)
  }
}

// 节点和边名称常量
case object GraphView {
  case object Node {
    
  }

  case object Edge {
    
  }
}

// 模板内部使用的常量
case object Constants {
  val SCHEMA: StructType = 
StructType(
    List(
      StructField("节点id", StringType, nullable = false),
StructField("得分", StringType, nullable = false)
    )
)
}

// 模板
abstract class Template {
  // 输入处理逻辑
  def input(args: Array[String]): Param = {
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")
	val `iter`: Integer =json.getInteger("iter")
val `prob`: Double =json.getDouble("prob")
Param(iter = iter, prob = prob, graphID = graphID, target = target)
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
