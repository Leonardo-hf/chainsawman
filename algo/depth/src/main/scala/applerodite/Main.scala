package applerodite

import applerodite.config.AlgoConstants.{SCHEMA_DEFAULT, SCHEMA_SOFTWARE}
import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.{EdgeDirection, Graph, Pregel}
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions._

object Main {
  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")

    val graph = ClientConfig.graphClient.loadInitGraphForSoftware(graphID)

    val spark = ClientConfig.spark

    if (graph.numVertices == 0) {
      val df = spark.sqlContext
        .createDataFrame(spark.sparkContext.parallelize(List.apply(Row.apply(0, "", "", 0))
        ), SCHEMA_SOFTWARE).sort(AlgoConstants.SCORE_COL)
      ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
      return
    }
    if (graph.numVertices == 1) {
      val df = spark.sqlContext
        .createDataFrame(spark.sparkContext.parallelize(graph.vertices.map(v => Row.apply(v._1, v._2.Artifact, v._2.Version, 0)).collect()
        ), SCHEMA_SOFTWARE).sort(AlgoConstants.SCORE_COL)
      ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
      return
    }
    var vGraph = graph.outerJoinVertices(graph.outDegrees) { (_, _, d) => {
      if (d.isEmpty) {
        // 当前累计分，当前累计出节点数目，出度，应用层级
        (0.0, 1, 0, 0.0)
      } else {
        (0.0, 0, d.get, 0.0)
      }
    }
    }

    // 获得所有节点的应用层级
    vGraph = Pregel(vGraph, initialMsg = (0.0, 0), maxIterations = Int.MaxValue, activeDirection = EdgeDirection.In)(
      (_, attr, msg) => {
        var al = 0.0
        if (attr._3 - msg._2 == 0) {
          al = 1 + math.sqrt((attr._1 + msg._1) / (attr._2 + msg._2))
        }
        (attr._1 + msg._1, attr._2 + msg._2, attr._3 - msg._2, al)
      },
      edge => {
        if (edge.dstAttr._3 == 0) {
          Iterator((edge.srcId, (math.pow(edge.dstAttr._4, 2), 1)))
        } else {
          Iterator.empty
        }
      }, (a, b) => (a._1 + b._1, a._2 + b._2))

    // 获得深度之和
    var vsGraph = vGraph.outerJoinVertices(graph.inDegrees) { (_, t, d) => {
      if (d.isEmpty) {
        // 当前累计分，入度，当前累计总节点数目，应用层级，深度指标
        (t._4, 0, 1, t._4, 0.0)
      } else {
        (0.0, d.get, 0, t._4, 0.0)
      }
    }
    }
    val size = vsGraph.numVertices
    vsGraph = Pregel(vsGraph, initialMsg = (0.0, 0, 0), maxIterations = Int.MaxValue, activeDirection = EdgeDirection.Out)(
      (_, attr, msg) => {
        var ds = 0.0
        if (attr._2 - msg._2 == 0) {
          ds = (attr._1 + msg._1 - (attr._3 + msg._3) * attr._4) / size
        }
        (attr._1 + msg._1, attr._2 - msg._2, attr._3 + msg._3, attr._4, ds)
      },
      edge => {
        if (edge.srcAttr._2 == 0) {
          Iterator((edge.dstId, (edge.srcAttr._4, 1, edge.srcAttr._3)))
        } else {
          Iterator.empty
        }
      }, (a, b) => (a._1 + b._1, a._2 + b._2, a._3 + b._3))

    val res = vsGraph.outerJoinVertices(graph.vertices) {
      (_, vd, release) => {
        (vd._5, release)
      }
    }

    val df = spark.sqlContext.createDataFrame(res.vertices.map(v => Row.apply(v._1, v._2._2.get.Artifact, v._2._2.get.Version, v._2._1)), SCHEMA_DEFAULT).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.close()
  }
}
