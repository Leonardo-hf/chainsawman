package applerodite.depth

import applerodite.config.CommonService
import org.apache.spark.graphx.{EdgeDirection, Pregel}
import org.apache.spark.sql.Row

object Main extends Template {
  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    val graph = svc.getGraphClient.loadInitGraphForSoftware(param.graphID)
    if (graph.numVertices == 0) {
      return Seq.empty
    }
    if (graph.numVertices == 1) {
      val fv = graph.vertices.first()
      return Seq.apply(Row.apply(fv._1, fv._2.Artifact, fv._2.Version, 0))
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
    val rs = res.vertices.map(v => ResultRow.apply(
      id = v._1,
      artifact = v._2._2.get.Artifact,
      version = v._2._2.get.Version,
      score = v._2._1
    )).collect()
    val rsMax = rs.map(r => r.`score`).max
    val rsMin = rs.map(r => r.`score`).min
    rs.sortBy(r => -r.score).map(r => r.toRow(score = f"${(r.score - rsMin) / (rsMax - rsMin)}%.4f"))
  }
}
