package applerodite

import applerodite.config.AlgoConstants.SCHEMA_SOFTWARE
import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx.{EdgeDirection, Graph, Pregel, VertexId}
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions._

import scala.collection.mutable.ListBuffer

object Main {

  def setCI(g: Graph[Int, None.type], r: Int): Graph[Int, None.type] = {
    // 准备度数
    var cGraph = g.outerJoinVertices(g.degrees) {
      (_, _, deg) => (0, deg.getOrElse(0), ListBuffer[List[(VertexId, Int)]]())
    }

    // 获得所有入路径
    cGraph = Pregel(cGraph, initialMsg = (0, ListBuffer[List[(VertexId, Int)]]()), maxIterations = Int.MaxValue, activeDirection = EdgeDirection.Out)(
      (_, attr, msg) => {
        // 累计完成的节点
        (attr._1 + msg._1, attr._2, attr._3.union(msg._2))
      },
      edge => {
        if (edge.srcAttr._1 == edge.srcAttr._2) {
          val paths = ListBuffer[List[(VertexId, Int)]]()
          // 如果是空，则将自身直接加入
          if (edge.srcAttr._3.isEmpty) {
            paths.append(List.apply((edge.srcId, edge.srcAttr._2)))
          }
          // 否则，将自己加入到路径的最后
          for (p <- edge.srcAttr._3) {
            paths.append(p :+ (edge.srcId, edge.srcAttr._2))
          }
          Iterator((edge.dstId, (1, paths)))
        } else {
          Iterator.empty
        }
      }, (a, b) => (a._1 + b._1, a._2.union(b._2)))

    // 过滤得到所有长度为r的路径，计算CI值
    cGraph.mapVertices((_, vd) =>
      vd._3.filter(p => p.length >= r).map(p => p(p.length - r)).toSet[(VertexId, Int)].map(p => p._2 - 1).sum * (vd._2 - 1)
    )
  }

  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")

    val graph = ClientConfig.graphClient.loadInitGraphForSoftware(graphID)
    graph.cache()

    val spark = ClientConfig.spark
    val r = 3
    val res = ListBuffer[Row]()

    var preVGraph: Graph[Int, None.type] = null
    var vGraph = setCI(graph.mapVertices((_, _) => 0), r)
    vGraph.cache()

    while (math.pow(vGraph.vertices.map(v => v._2).sum() / vGraph.degrees.map(d => d._2).sum(), 1.0 / (r + 1)) > 1) {
      preVGraph = vGraph
      val (maxVertex, score) = vGraph.vertices.max()(Ordering.by[(VertexId, Int), Double](_._2))
      val meta = graph.vertices.filter(v => v._1 == maxVertex).first()._2
      res.append(Row.apply(maxVertex, meta.Artifact, meta.Version, score))
      vGraph = setCI(vGraph.subgraph(epred = e => e.srcId != maxVertex && e.dstId != maxVertex), r)
      vGraph.cache()
      preVGraph.unpersistVertices(blocking = false)
      preVGraph.edges.unpersist(blocking = false)
    }

    val df = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(res), SCHEMA_SOFTWARE).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.stop()
  }
}
