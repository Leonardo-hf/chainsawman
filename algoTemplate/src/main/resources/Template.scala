package applerodite

import applerodite.config.{CommonService, CommonServiceImpl}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.sql.Row
import org.apache.spark.sql.types.{StringType, StructField, StructType}

case class Param(graphID: String, target: String)

case class ResultRow()

case object Constants {
  var COL_SOFTWARE_ID = "ID"
  var COL_SOFTWARE_ARTIFACT = "工件名"
  var COL_SOFTWARE_VERSION = "版本号"
  var COL_SOFTWARE_SCORE = "得分"
}

abstract class Template {

  val SCHEMA: StructType = StructType(
    List(
      StructField(Constants.COL_SOFTWARE_ID, StringType, nullable = false),
      StructField(Constants.COL_SOFTWARE_ARTIFACT, StringType, nullable = false),
      StructField(Constants.COL_SOFTWARE_VERSION, StringType, nullable = false),
      StructField(Constants.COL_SOFTWARE_SCORE, StringType, nullable = false)
    ))

  def input(args: Array[String]): Param = {
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")
    Param(graphID = graphID, target = target)
  }

  def exec(svc: CommonService, param: Param): Seq[Row]

  def output(svc: CommonService, param: Param, rows: Seq[Row]): Unit = {
    val spark = svc.getSparkSession
    val df = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(rows), SCHEMA)
    svc.getOSSClient.upload(name = param.target, content = CSVUtil.df2CSV(df))
    svc.getMysqlClient.updateStatusByOutput(param.target)
  }

  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val svc = CommonServiceImpl
    val param = input(args)
    val res = exec(svc, param)
    output(svc, param, res)
    svc.getSparkSession.stop()
  }
}
