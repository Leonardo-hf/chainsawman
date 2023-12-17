package applerodite

import applerodite.config.ClientConfig
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.Graph
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions.desc

object Main {
  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")

    val graph: Graph[None.type, Double] = ClientConfig.graphClient.loadInitGraph(graphID, hasWeight = false)

    val spark = ClientConfig.spark

    val closedTriangleNum = graph.triangleCount().vertices.map(_._2).reduce(_ + _)
    // compute the number of open triangle and closed triangle (According to C(n,2)=n*(n-1)/2)
    val triangleNum = graph.degrees.map(vertex => (vertex._2 * (vertex._2 - 1)) / 2.0).reduce(_ + _)
    var score = 0.0
    if (triangleNum != 0)
      score = (closedTriangleNum / triangleNum * 1.0).formatted("%.6f").toDouble

    val res = List(Row.apply(score))
    val df = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(res), SCHEMA_SCORE).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.stop()
  }
}
