package applerodite

import .SCHEMA_SOFTWARE
import applerodite.config.ClientConfig
import com.alibaba.fastjson.JSON
import org.apache.hadoop.shaded.org.eclipse.jetty.websocket.common.frames.DataFrame
import org.apache.spark.graphx.{Graph, TripletFields, VertexId}
import org.apache.spark.sql.functions._
import org.apache.spark.sql.{DataFrame, Dataset, Row}

import scala.collection.mutable
import scala.collection.mutable.ListBuffer
import scala.util.control.Breaks

trait Impact {
  def main(args: Array[String]): Dataset[Row]
}

object Voterank extends Impact {
  def main(args: Array[String]): Dataset[Row] = {
    if (args.length < 1) {
      return null
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    var iter: Int = json.getIntValue("iter")
    val graph = ClientConfig.graphClient.loadInitGraphForSoftware(graphID)
    graph.cache()

    val spark = ClientConfig.spark

    val res = ListBuffer[Row]()
    val sumDegree = graph.inDegrees.map(v => v._2).sum()
    var f = 1.0
    if (sumDegree != 0) {
      f = graph.numVertices / sumDegree
    }
    val cal = mutable.Set[VertexId]()
    var vGraph = graph.mapVertices((_, _) => 1.0)
    var preVGraph: Graph[Double, None.type] = null
    vGraph.cache()
    val loop = new Breaks
    loop.breakable {
      while (iter > 0 && res.length < graph.numVertices) {
        preVGraph = vGraph
        iter = iter - 1
        val addVertices = vGraph.aggregateMessages[Double](t => t.sendToDst(t.srcAttr), _ + _, TripletFields.Src).filter(vid => !cal.contains(vid._1))
        if (addVertices.isEmpty()) {
          loop.break()
        }
        val (maxVertexId, score) = addVertices.max()(Ordering.by[(VertexId, Double), Double](_._2))
        val maxVertex = graph.vertices.filter(v => v._1 == maxVertexId).collect().apply(0)
        res.append(Row.apply(maxVertexId, maxVertex._2.Artifact, maxVertex._2.Version, score))
        cal.add(maxVertexId)
        vGraph = vGraph.joinVertices(graph.edges.filter(e => e.dstId == maxVertexId).map[(VertexId, Double)](e => (e.srcId, -f))) {
          (_, oldScore, change) => oldScore + change
        }.mapVertices((vid, score) => {
          if (vid == maxVertexId || score < 0) {
            0
          } else {
            score
          }
        })
        vGraph.cache()
        preVGraph.unpersistVertices(blocking = false)
        preVGraph.edges.unpersist(blocking = false)
      }
    }

    spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(res), SCHEMA_SOFTWARE).orderBy(desc(AlgoConstants.SCORE_COL))
  }
}
