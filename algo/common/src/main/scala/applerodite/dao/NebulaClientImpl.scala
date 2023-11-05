package applerodite.dao

import applerodite.config.AlgoConstants.{ARTIFACT_COL, VERSION_COL}
import applerodite.config.ClientConfig
import com.vesoft.nebula.client.graph.data.{HostAddress, ResultSet}
import com.vesoft.nebula.client.graph.{SessionPool, SessionPoolConfig}
import com.vesoft.nebula.connector.connector.NebulaDataFrameReader
import com.vesoft.nebula.connector.{NebulaConnectionConfig, ReadNebulaConfig, VertexID}
import org.apache.spark.graphx.{Edge, Graph}
import org.apache.spark.sql.Encoder

import java.util

object NebulaClientImpl extends GraphClient {

  case class NebulaConfig(metaHost: String, metaPort: Int, graphdHost: String, graphdPort: Int, user: String, passwd: String)

  var connectorConfig: NebulaConnectionConfig = _

  var config: NebulaConfig = _

  def Init(nc: NebulaConfig): GraphClient = {
    config = nc
    connectorConfig =
      NebulaConnectionConfig
        .builder()
        .withMetaAddress(s"${config.metaHost}:${config.metaPort}")
        .withConenctionRetry(2)
        .withExecuteRetry(2)
        .withTimeout(6000)
        .build()
    this
  }

  override def gql(graphName: String, stmt: String): Option[ResultSet] = {
    val sessionPoolConfig: SessionPoolConfig = new SessionPoolConfig(
      util.Arrays.asList(new HostAddress(config.graphdHost, config.graphdPort)), graphName, config.user, config.passwd).setWaitTime(100).setRetryTimes(3).setIntervalTime(100)
    val sessionPool: SessionPool = new SessionPool(sessionPoolConfig)
    if (!sessionPool.init) {
      println("#### fail to init nebula session")
      return Option.empty
    }
    val res = sessionPool.execute(stmt)
    if (!res.isSucceeded) {
      println("#### " + res.getErrorMessage)
    }
    Option.apply(res)
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
      val dataSet = ClientConfig.spark.read.nebula(connectorConfig, nebulaReadEdgesConfig).loadEdgesToDF()
      implicit val encoder: Encoder[Edge[Double]] = org.apache.spark.sql.Encoders.kryo[Edge[Double]]
      dataSet
        .map(row => {
          if (hasWeight) {
            Edge(row.getLong(0), row.getLong(1), row.getDouble(2))
          } else {
            Edge(row.getLong(0), row.getLong(1), 1.0)
          }
        })(encoder).rdd
    }).reduce((e1, e2) => e1.union(e2))
    Graph.fromEdges(edges, None)
  }


  override def loadInitGraphForSoftware(graphName: String): Graph[Release, None.type] = {
    val nebulaReadDependsConfig: ReadNebulaConfig = ReadNebulaConfig
      .builder()
      .withSpace(graphName)
      .withLabel("depend")
      .withNoColumn(true)
      //      .withPartitionNum(10)
      .build()
    val edges = ClientConfig.spark.read.nebula(connectorConfig, nebulaReadDependsConfig).loadEdgesToDF()
    implicit val edgeEncoder: Encoder[Edge[None.type]] = org.apache.spark.sql.Encoders.kryo[Edge[None.type]]
    val edgesRDD = edges
      .map(row => {
        Edge(row.getLong(0), row.getLong(1), None)
      })(edgeEncoder).rdd
    val nebulaReadReleaseConfig: ReadNebulaConfig = ReadNebulaConfig
      .builder()
      .withSpace(graphName)
      .withLabel("release")
      .withNoColumn(false)
      .withReturnCols(List(ARTIFACT_COL, VERSION_COL))
      //      .withPartitionNum(10)
      .build()
    val nodes = ClientConfig.spark.read.nebula(connectorConfig, nebulaReadReleaseConfig).loadVerticesToDF()
    implicit val nodeEncoder: Encoder[(VertexID, Release)] = org.apache.spark.sql.Encoders.kryo[(VertexID, Release)]
    val nodesRDD = nodes.map(row => (row.getLong(0), Release(row.getString(1), row.getString(2))))(nodeEncoder).rdd
    Graph(nodesRDD, edgesRDD)
  }
}
