package applerodite

import applerodite.config.AlgoConstants.SCHEMA_DEFAULT
import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.{EdgeTriplet, Graph, Pregel, VertexId}
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions._

object Main {

  type SPMap = Map[VertexId, Double]

  def makeMap(x: (VertexId, Double)*) = Map(x: _*)

  def addMap(spmap: SPMap, weight: Double): SPMap = spmap.map {
    case (v, d) => v -> (d + weight)
  }

  def addMaps(spmap1: SPMap, spmap2: SPMap): SPMap = {
    (spmap1.keySet ++ spmap2.keySet).map { k =>
      k -> math.min(spmap1.getOrElse(k, Double.MaxValue), spmap2.getOrElse(k, Double.MaxValue))
    }(collection.breakOut)
  }

  def vertexProgram(id: VertexId, attr: SPMap, msg: SPMap): SPMap = {
    addMaps(attr, msg)
  }

  def sendMessage(edge: EdgeTriplet[SPMap, Double]): Iterator[(VertexId, SPMap)] = {
    val newAttr = addMap(edge.dstAttr, edge.attr)
    if (edge.srcAttr != addMaps(newAttr, edge.srcAttr)) Iterator((edge.srcId, newAttr))
    else Iterator.empty
  }

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

    val spGraph = graph.mapVertices((vid, _) => makeMap(vid -> 0.0))
    val initialMessage = makeMap()

    val spsGraph = Pregel(spGraph, initialMessage)(vertexProgram, sendMessage, addMaps)
    val closenessRDD = spsGraph.vertices.map(vertex => {
      var dstNum = 0
      var dstDistanceSum = 0.0
      for (distance <- vertex._2.values) {
        dstNum += 1
        dstDistanceSum += distance
      }
      Row(vertex._1, (dstNum - 1) / dstDistanceSum)
    })

    val df = spark.sqlContext.createDataFrame(closenessRDD, SCHEMA_DEFAULT).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.close()
  }
}
