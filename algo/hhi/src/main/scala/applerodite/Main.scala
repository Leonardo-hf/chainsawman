package applerodite

import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.{Edge, Graph, VertexRDD}
import org.apache.spark.rdd.RDD
import org.apache.spark.sql.Row
import org.apache.spark.sql.types.{DoubleType, StringType, StructField, StructType}

object Main {

  val SCHEMA_HHI: StructType = StructType(
    List(
      StructField("topic", StringType, nullable = false),
      StructField(AlgoConstants.SCORE_COL, StringType, nullable = false)
    ))

  def main(args: Array[String]): Unit = {
    val spark = ClientConfig.spark
    if (args.length < 1) {
      spark.stop()
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")

    var hhi: RDD[Row] = spark.sparkContext.emptyRDD

    val ctrRes = ClientConfig.graphClient.gql(graphID,
      "match (s:library)<-[b1:belong2]-(r1:release)-[d:depend]->(r2:release)-[b2:belong2]->(t:library) " +
        "return distinct id(s) as source, s.library.topic as s_topic, id(t) as target, t.library.topic as t_topic;")
    if (ctrRes.isDefined) {
      val res = ctrRes.get
      val edges = spark.sparkContext.parallelize((0 until res.rowsSize()).map(
        i => {
          Edge.apply(res.rowValues(i).get("source").asLong(), res.rowValues(i).get("target").asLong(), None)
        }
      ))
      val vertices: VertexRDD[String] = VertexRDD.apply(spark.sparkContext.parallelize((0 until res.rowsSize()).map(
        i => {
          List.apply(
            (res.rowValues(i).get("source").asLong(), res.rowValues(i).get("s_topic").asString()),
            (res.rowValues(i).get("target").asLong(), res.rowValues(i).get("t_topic").asString())
          )
        }
      ).toList.flatten).distinct())
      val graph = Graph.apply(vertices, edges)
      hhi = graph.outerJoinVertices(graph.degrees) {
        (_, t, d) => (t, d.getOrElse(0))
      }.vertices.map(v => v._2).groupBy(vd => vd._1).map(vds =>
        Row.apply(vds._1, f"${100 * vds._2.map(vd => math.pow(vd._2, 2)).sum / math.pow(vds._2.map(vd => vd._2).sum, 2)}%.2f" + "%")
      )
    }
    val df = spark.sqlContext.createDataFrame(hhi, SCHEMA_HHI).orderBy(AlgoConstants.SCORE_COL)
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.stop()
  }
}
