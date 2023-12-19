package applerodite.stability

import applerodite.config.CommonService
import org.apache.spark.graphx.{EdgeDirection, Graph, Pregel, VertexId}
import org.apache.spark.sql.Row

import scala.collection.mutable.ListBuffer

object Main extends Template {

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

  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    val graph = svc.getGraphClient.loadInitGraphForSoftware(param.graphID)
    graph.cache()

    val r = param.`radius`
    val res = ListBuffer[ResultRow]()

    var preVGraph: Graph[Int, None.type] = null
    var vGraph = setCI(graph.mapVertices((_, _) => 0), r)
    vGraph.cache()

    while (math.pow(vGraph.vertices.map(v => v._2).sum() / vGraph.degrees.map(d => d._2).sum(), 1.0 / (r + 1)) > 1) {
      preVGraph = vGraph
      val (maxVertex, score) = vGraph.vertices.max()(Ordering.by[(VertexId, Int), Double](_._2))
      val meta = graph.vertices.filter(v => v._1 == maxVertex).first()._2
      res.append(ResultRow.apply(`id` = maxVertex, artifact = meta.Artifact, version = meta.Version, score = math.log(score)))
      vGraph = setCI(vGraph.subgraph(epred = e => e.srcId != maxVertex && e.dstId != maxVertex), r)
      vGraph.cache()
      preVGraph.unpersistVertices(blocking = false)
      preVGraph.edges.unpersist(blocking = false)
    }
    val rsMax = res.map(r => r.`score`).max
    val rsMin = res.map(r => r.`score`).min
    res.sortBy(r => -r.score).map(r => r.toRow(score = f"${(r.score - rsMin) / (rsMax - rsMin)}%.4f"))
  }

}
