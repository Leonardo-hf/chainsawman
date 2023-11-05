package applerodite.dao


import com.vesoft.nebula.client.graph.data.ResultSet
import org.apache.spark.graphx.Graph

trait GraphClient {
  def loadInitGraph(graphName: String, edgeTags: Seq[String], hasWeight: Boolean): Graph[None.type, Double]
  def loadInitGraphForSoftware(graphName: String): Graph[Release, None.type]
  def gql(graphName: String, stmt: String): Option[ResultSet]
}

case class Release(Artifact: String, Version: String)
