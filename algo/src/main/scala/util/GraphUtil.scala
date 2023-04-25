package util

import com.vesoft.nebula.connector.ReadNebulaConfig
import com.vesoft.nebula.connector.connector.NebulaDataFrameReader
import dao.SparkClientImpl.{nebulaCfg, spark}
import org.apache.spark.graphx.{Edge, Graph}
import org.apache.spark.rdd.RDD
import org.apache.spark.sql.Encoder

object GraphUtil {

  def loadInitGraph(graphID: Long, hasWeight: Boolean): Graph[None.type, Double] = {
    val nebulaReadVertexConfig: ReadNebulaConfig = ReadNebulaConfig
      .builder()
      .withSpace("G%s".format(graphID))
      .withLabel("sedge")
      .withNoColumn(true)
      //      .withPartitionNum(10)
      .build()
    val dataSet = spark.read.nebula(nebulaCfg, nebulaReadVertexConfig).loadEdgesToDF()
    implicit val encoder: Encoder[Edge[Double]] = org.apache.spark.sql.Encoders.kryo[Edge[Double]]
    val edges: RDD[Edge[Double]] = dataSet
      .map(row => {
        if (hasWeight) {
          Edge(row.get(0).toString.toLong, row.get(1).toString.toLong, row.get(2).toString.toDouble)
        } else {
          Edge(row.get(0).toString.toLong, row.get(1).toString.toLong, 1.0)
        }
      })(encoder)
      .rdd
    Graph.fromEdges(edges, None)
  }
}
