package applerodite.dao

import com.vesoft.nebula.client.graph.data.{HostAddress, ResultSet}
import com.vesoft.nebula.client.graph.{SessionPool, SessionPoolConfig}
import com.vesoft.nebula.connector.connector.NebulaDataFrameReader
import com.vesoft.nebula.connector.{NebulaConnectionConfig, ReadNebulaConfig, VertexID}
import org.apache.spark.graphx.{Edge, Graph}
import org.apache.spark.sql.{Encoder, SparkSession}

import java.util
import scala.collection.mutable.ListBuffer

case class NebulaConfig(metaHost: String, metaPort: Int, graphdHost: String, graphdPort: Int, user: String, passwd: String)

class NebulaClientImpl extends GraphClient {

  var connectorConfig: NebulaConnectionConfig = _

  var config: NebulaConfig = _

  var spark: SparkSession = _

  def GetGraphClient(nc: NebulaConfig, spark: SparkSession): GraphClient = {
    config = nc
    connectorConfig =
      NebulaConnectionConfig
        .builder()
        .withMetaAddress(s"${config.metaHost}:${config.metaPort}")
        .withConenctionRetry(2)
        .withExecuteRetry(2)
        .withTimeout(6000)
        .build()
    this.spark = spark
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
    sessionPool.close()
    Option.apply(res)
  }

  override def loadInitGraph(graphName: String, hasWeight: Boolean): Graph[None.type, Double] = {
    val res = gql(graphName, "show edges;")
    val edgeTags = ListBuffer[String]()
    if (res.isDefined) {
      for (i <- 0 until res.get.rowsSize()) {
        edgeTags.append(res.get.rowValues(i).get("Name").asString())
      }
    }
    val edges = edgeTags.map(et => {
      val nebulaReadEdgesConfig: ReadNebulaConfig = ReadNebulaConfig
        .builder()
        .withSpace(graphName)
        .withLabel(et)
        .withNoColumn(true)
        //      .withPartitionNum(10)
        .build()
      val dataSet = spark.read.nebula(connectorConfig, nebulaReadEdgesConfig).loadEdgesToDF()
      implicit val encoder: Encoder[Edge[Double]] = org.apache.spark.sql.Encoders.kryo[Edge[Double]]
      dataSet
        .map(row => {
          if (hasWeight) {
            Edge(row.getString(0).toLong, row.getString(1).toLong, row.getString(2).toDouble)
          } else {
            Edge(row.getString(0).toLong, row.getString(1).toLong, 1.0)
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
    val edges = spark.read.nebula(connectorConfig, nebulaReadDependsConfig).loadEdgesToDF()
    implicit val edgeEncoder: Encoder[Edge[None.type]] = org.apache.spark.sql.Encoders.kryo[Edge[None.type]]
    val edgesRDD = edges
      .map(row => {
        Edge(row.getString(0).toLong, row.getString(1).toLong, None)
      })(edgeEncoder).rdd
    val nebulaReadReleaseConfig: ReadNebulaConfig = ReadNebulaConfig
      .builder()
      .withSpace(graphName)
      .withLabel("release")
      .withNoColumn(false)
      .withReturnCols(List("artifact", "version"))
      //      .withPartitionNum(10)
      .build()
    val nodes = spark.read.nebula(connectorConfig, nebulaReadReleaseConfig).loadVerticesToDF()
    implicit val nodeEncoder: Encoder[(VertexID, Release)] = org.apache.spark.sql.Encoders.kryo[(VertexID, Release)]
    val nodesRDD = nodes.map(row => (row.getString(0).toLong, Release(row.getString(1), row.getString(2))))(nodeEncoder).rdd
    Graph(nodesRDD, edgesRDD)
  }
}
