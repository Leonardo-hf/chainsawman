package applerodite

import applerodite.config.AlgoConstants.SCHEMA_RANK
import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.Graph
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions._

object Main {
  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")
    val edgeTags: Seq[String] = json.getJSONArray("edgeTags").toArray().map(a => a.toString)

    val graph: Graph[None.type, Double] = ClientConfig.graphClient.loadInitGraph(graphID, edgeTags, hasWeight = false)

    val spark = ClientConfig.spark

    val df = spark.sqlContext.createDataFrame(graph.degrees.map(r => Row.apply(r._1, r._2.toDouble)), SCHEMA_RANK).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.close()
  }
}
