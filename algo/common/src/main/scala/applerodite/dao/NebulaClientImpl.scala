package applerodite.dao

import com.typesafe.config.Config
import com.vesoft.nebula.connector.connector.NebulaDataFrameReader
import com.vesoft.nebula.connector.{NebulaConnectionConfig, ReadNebulaConfig}
import applerodite.config.ClientConfig
import org.apache.spark.graphx.{Edge, Graph}
import org.apache.spark.sql.Encoder

object NebulaClientImpl extends GraphClient {

  var nebulaCfg: NebulaConnectionConfig = _

  def Init(config: Config): GraphClient = {
    nebulaCfg =
      NebulaConnectionConfig
        .builder()
        .withMetaAddress(config.getString("url"))
        .withConenctionRetry(2)
        .withExecuteRetry(2)
        .withTimeout(6000)
        .build()
    this
  }

  override def loadInitGraph(graphName: String, edgeTags: Seq[String], hasWeight: Boolean): Graph[None.type, Double] = {
    val edges = edgeTags.map(et => {
      val nebulaReadEdgesConfig: ReadNebulaConfig = ReadNebulaConfig
        .builder()
        .withSpace(graphName)
        .withLabel(et)
        .withNoColumn(true)
        //      .withPartitionNum(10)
        .build()
      val dataSet = ClientConfig.spark.read.nebula(nebulaCfg, nebulaReadEdgesConfig).loadEdgesToDF()
      implicit val encoder: Encoder[Edge[Double]] = org.apache.spark.sql.Encoders.kryo[Edge[Double]]
      dataSet
        .map(row => {
          if (hasWeight) {
            Edge(row.get(0).toString.toLong, row.get(1).toString.toLong, row.get(2).toString.toDouble)
          } else {
            Edge(row.get(0).toString.toLong, row.get(1).toString.toLong, 1.0)
          }
        })(encoder).rdd
    }).reduce((e1, e2) => e1.union(e2))
    Graph.fromEdges(edges, None)
  }
}
