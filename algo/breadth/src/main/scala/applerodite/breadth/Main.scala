package applerodite.breadth

import applerodite.config.CommonService
import org.apache.spark.graphx.{Graph, TripletFields, VertexId}
import org.apache.spark.sql.Row

import scala.collection.mutable
import scala.collection.mutable.ListBuffer
import scala.util.control.Breaks

object Main extends Template {

  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    val graph = svc.getGraphClient.loadInitGraphForSoftware(param.graphID)
    var iter = param.`iter`
    graph.cache()

    val res = ListBuffer[ResultRow]()
    val sumDegree = graph.inDegrees.map(v => v._2).sum()
    var f = 1.0
    if (sumDegree != 0) {
      f = graph.numVertices / sumDegree
    }
    val cal = mutable.Set[VertexId]()
    var vGraph = graph.mapVertices((_, _) => 1.0)
    var preVGraph: Graph[Double, None.type] = null
    vGraph.cache()
    val loop = new Breaks
    loop.breakable {
      while (iter > 0 && res.length < graph.numVertices) {
        preVGraph = vGraph
        iter = iter - 1
        val addVertices = vGraph.aggregateMessages[Double](t => t.sendToDst(t.srcAttr), _ + _, TripletFields.Src).filter(vid => !cal.contains(vid._1))
        if (addVertices.isEmpty()) {
          loop.break()
        }
        val (maxVertexId, score) = addVertices.max()(Ordering.by[(VertexId, Double), Double](_._2))
        if (score == 0) {
          loop.break()
        }
        val maxVertex = graph.vertices.filter(v => v._1 == maxVertexId).collect().head
        res.append(ResultRow.apply(id = maxVertexId, artifact = maxVertex._2.Artifact, version = maxVertex._2.Version, score = math.log(score + 1)))
        cal.add(maxVertexId)
        vGraph = vGraph.joinVertices(graph.edges.filter(e => e.dstId == maxVertexId).map[(VertexId, Double)](e => (e.srcId, -f))) {
          (_, oldScore, change) => oldScore + change
        }.mapVertices((vid, score) => {
          if (vid == maxVertexId || score < 0) {
            0
          } else {
            score
          }
        })
        vGraph.cache()
        preVGraph.unpersistVertices(blocking = false)
        preVGraph.edges.unpersist(blocking = false)
      }
    }
    val rsMax = res.map(r => r.`score`).max
    val rsMin = res.map(r => r.`score`).min
    res.sortBy(r => -r.score).map(r => r.toRow(score = f"${(r.score - rsMin) / (rsMax - rsMin)}%.4f"))
  }
}
