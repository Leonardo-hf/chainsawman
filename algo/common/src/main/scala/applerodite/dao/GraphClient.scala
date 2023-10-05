package applerodite.dao

import org.apache.spark.graphx.Graph

trait GraphClient {
  def loadInitGraph(graphName: String, edgeTags: Seq[String], hasWeight: Boolean): Graph[None.type, Double]
}
