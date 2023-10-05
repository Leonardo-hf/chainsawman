package applerodite

import applerodite.config.ClientConfig
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.Graph

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

    val closedTriangleNum = graph.triangleCount().vertices.map(_._2).reduce(_ + _)
    // compute the number of open triangle and closed triangle (According to C(n,2)=n*(n-1)/2)
    val triangleNum = graph.degrees.map(vertex => (vertex._2 * (vertex._2 - 1)) / 2.0).reduce(_ + _)
    var score = 0.0
    if (triangleNum != 0)
      score = (closedTriangleNum / triangleNum * 1.0).formatted("%.6f").toDouble

    ClientConfig.ossClient.upload(name = target, content = score.toString)
    spark.close()
  }
}
