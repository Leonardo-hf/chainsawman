package applerodite.stability

import applerodite.config.CommonService
import org.apache.spark.graphx.{EdgeDirection, Graph, Pregel, VertexId}
import org.apache.spark.sql.Row

import scala.collection.mutable.ListBuffer

object Main extends Template  {

  private case class NSkip(var vertex: VertexId, var skip: Int)

  private case class VertexState(state: ListBuffer[NSkip], toSend: Seq[NSkip])

  private def setCI(g: Graph[Int, None.type], r: Int): Graph[Int, None.type] = {
    val degMap = g.degrees.collect().toMap
    // 计算图g中距离每个节点r跳的节点的个数
    Pregel(
      g.mapVertices((vid, _) => VertexState(state = ListBuffer.empty, toSend = Seq.apply(NSkip(vid, r)))),
      initialMsg = Seq.empty[NSkip],
      maxIterations = Int.MaxValue,
      activeDirection = EdgeDirection.Out)(
      vprog = (_, attr, msg) => {
        // state 存放skip为0（达到终点站）的节点，toSend存放发往下一跳的节点
        if (msg.isEmpty) {
          attr
        } else {
          VertexState(state = attr.state ++= msg.filter(_.skip == 0), toSend = msg.filter(_.skip > 0))
        }
      },
      sendMsg = edge => {
        // 如果当前节点有待发送的节点，将待发送的节点发送给下一跳
        if (edge.srcAttr.toSend.nonEmpty) {
          Iterator((edge.dstId, edge.srcAttr.toSend.map(ns => NSkip(ns.vertex, ns.skip - 1))))
        } else {
          Iterator.empty
        }
      },
      mergeMsg = (a, b) => a.union(b)
    )
      // 计算每个节点的CI值
      .mapVertices((vid, attr) => attr.state.map(v => degMap.getOrElse(v.vertex, 0) - 1).sum * (degMap.getOrElse(vid, 0) - 1))
  }

  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    svc.getSparkSession.sparkContext.setCheckpointDir("s3a://tmp")
    val graph = svc.getGraphClient.loadInitGraphForSoftware(param.graphID)
    val metaMap = graph.vertices.collect().toMap
    val r = param.`radius`
    val res = ListBuffer.empty[ResultRow]

    var preVGraph: Graph[Int, None.type] = null
    var vGraph = setCI(graph.mapVertices((_, _) => 0), r)
    vGraph.cache()
    // 多次迭代，每次迭代选出最优的节点的CI值
    while (math.pow(vGraph.vertices.map(v => v._2).sum() / vGraph.degrees.map(d => d._2).sum(), 1.0 / (r + 1)) > 1) {
      preVGraph = vGraph
      val (maxVertex, score) = vGraph.vertices.max()(Ordering.by[(VertexId, Int), Double](_._2))
      val meta = metaMap(maxVertex)
      res.append(ResultRow.apply(`id` = maxVertex, artifact = meta.Artifact, version = meta.Version, score = math.log(score)))
      vGraph = setCI(vGraph.subgraph(epred = e => e.srcId != maxVertex && e.dstId != maxVertex), r)
      vGraph.cache()
      vGraph.checkpoint()
      preVGraph.unpersistVertices()
      preVGraph.edges.unpersist()
    }
    if (res.isEmpty) {
      return Seq.empty
    }
    val rsMax = res.map(r => r.`score`).max
    val rsMin = res.map(r => r.`score`).min
    res.sortBy(r => -r.score).map(r => r.toRow(score = f"${(r.score - rsMin) / (rsMax - rsMin)}%.4f"))
  }

}

