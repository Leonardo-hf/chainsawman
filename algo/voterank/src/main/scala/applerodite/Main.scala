package applerodite

import applerodite.config.AlgoConstants.SCHEMA_RANK
import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.{Graph, TripletFields, VertexId}
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions._

import scala.collection.mutable
import scala.collection.mutable.ListBuffer
import scala.util.control.Breaks

object Main {
  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")
    var iter: Int = json.getIntValue("iter")
    val edgeTags: Seq[String] = json.getJSONArray("edgeTags").toArray().map(a => a.toString)
    val graph: Graph[None.type, Double] = ClientConfig.graphClient.loadInitGraph(graphID, edgeTags, hasWeight = false)

    val spark = ClientConfig.spark

    val res = ListBuffer[Row]()
    val sumDegree = graph.inDegrees.map(v => v._2).sum()
    var f = 1.0
    if (sumDegree != 0) {
      f = graph.numVertices / sumDegree
    }
    val cal = mutable.Set[VertexId]()
    var vGraph = graph.mapVertices((_, _) => 1.0)
    val loop = new Breaks
    loop.breakable {
      while (iter > 0 && res.length < graph.numVertices) {
        iter = iter - 1
        val addVertices = vGraph.aggregateMessages[Double](t => t.sendToDst(t.srcAttr), _ + _, TripletFields.Src).filter(vid => !cal.contains(vid._1))
        if (addVertices.isEmpty()) {
          loop.break()
        }
        val (maxVertex, score) = addVertices.max()(Ordering.by[(VertexId, Double), Double](_._2))
        res.append(Row.apply(maxVertex, score))
        cal.add(maxVertex)
        vGraph = vGraph.joinVertices(graph.edges.filter(e => e.dstId == maxVertex).map[(VertexId, Double)](e => (e.srcId, -f))) {
          (_, oldScore, change) => oldScore + change
        }.mapVertices((vid, score) => {
          if (vid == maxVertex || score < 0) {
            0
          } else {
            score
          }
        })
      }
    }

    val df = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(res), SCHEMA_RANK).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.close()
  }
}
