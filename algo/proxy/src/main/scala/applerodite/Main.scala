package applerodite

import applerodite.config.AlgoConstants.{SCHEMA_DEFAULT, SCHEMA_SOFTWARE}
import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import org.apache.spark.graphx._
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions._

import scala.collection.mutable
import scala.collection.mutable.ListBuffer

object Main {

  /**
   * 构建BetweennessCentrality图，图中顶点属性维护了图中所有顶点id的列表和所有边（srcId， dstId， attr）的列表
   *
   * @param initG 原始数据构造的图
   * @param k     最大迭代次数
   * @return Graph
   * */
  def createBetweenGraph(initG: Graph[None.type, Double], k: Int): Graph[VertexProperty, Double] = {
    val betweenG = initG
      .mapTriplets[Double]({ x: EdgeTriplet[None.type, Double] =>
        x.attr
      })
      .mapVertices((_, _) => new VertexProperty)
      .cache
    //准备进入pregel前的初始化消息、vertexProgram方法、 sendMessage方法、mergeMessage方法
    val initMessage = (List[VertexId](), List[(VertexId, VertexId, Double)]())

    //将发送过来的邻居节点信息以及当前点与邻居点的边，更新到当前点的属性中
    def vertexProgram(
                       id: VertexId,
                       attr: VertexProperty,
                       msgSum: (List[VertexId], List[(VertexId, VertexId, Double)])): VertexProperty = {
      val newAttr = new VertexProperty()
      newAttr.CB = attr.CB
      newAttr.vlist = (msgSum._1 ++ attr.vlist).distinct
      newAttr.elist = (msgSum._2 ++ attr.elist).distinct
      newAttr
    }

    //向邻居节点发送自身节点的id和自身与邻居点的边
    def sendMessage(edge: EdgeTriplet[VertexProperty, Double])
    : Iterator[(VertexId, (List[VertexId], List[(VertexId, VertexId, Double)]))] = Iterator(
      (edge.dstId,
        (edge.srcId +: edge.srcAttr.vlist,
          (edge.srcId, edge.dstId, edge.attr) +: edge.srcAttr.elist)),
      (edge.srcId,
        (edge.dstId +: edge.dstAttr.vlist,
          (edge.srcId, edge.dstId, edge.attr) +: edge.dstAttr.elist))
    )

    //合并接受到的多条消息
    def mergeMessage(a: (List[VertexId], List[(VertexId, VertexId, Double)]),
                     b: (List[VertexId], List[(VertexId, VertexId, Double)]))
    : (List[VertexId], List[(VertexId, VertexId, Double)]) = {
      ((a._1 ++ b._1).distinct, (a._2 ++ b._2).distinct)
    }

    Pregel(betweenG, initMessage, k, EdgeDirection.Either)(vertexProgram, sendMessage, mergeMessage)
  }

  /**
   * 定义顶点的属性类
   * CB: 定义初始的betweennessCentrality
   * vlist： 每个顶点需维护图中所有的顶点信息
   * elist: 每个顶点需维护图中所有的边信息
   *
   * 维护所有边信息是为了在计算介数中心性的时候可以从每个顶点依次根据邻居节点走下去（空间复杂度很高，O(n2)）
   * */
  class VertexProperty() extends Serializable {
    var CB = 0.0
    var vlist: List[VertexId] = List[VertexId]()
    var elist: List[(VertexId, VertexId, Double)] = List[(VertexId, VertexId, Double)]()
  }

  /**
   * betweennessCentralityForUnweightedGraph
   *
   * 对无权图计算顶点的介数中心性
   * 对每个顶点vid计算从该节点出发时其他节点的介数中心性,返回对于顶点vid而言图中节点所计算出来的介数中心性
   *
   * 对图中所有节点都会计算一次全局的介数中心性，最后进行汇总
   *
   * @param vid   :顶点id
   * @param vAttr ：顶点对应的属性信息
   * @return List((Vid, betweennessValue))
   *
   * */
  def betweennessCentralityForUnweightedGraph(vid: VertexId,
                                              vAttr: VertexProperty): List[(VertexId, Double)] = {
    //无权图的计算方法
    println("enter betweennessCentrality for vertex: " + vid)

    //对图中每个顶点做如下操作
    val S = mutable.Stack[VertexId]() //每次访问过的节点入栈
    val P = new mutable.HashMap[VertexId, ListBuffer[VertexId]]() //存储源顶点到某个顶点中间经过哪些顶点
    //如[5,[2,3]]，表示源顶点到顶点5的最短路径会经过顶点2,3
    val Q = mutable.Queue[VertexId]() //BFS遍历时将顶点入队列
    val dist = new mutable.HashMap[VertexId, Double]()
    val sigma = new mutable.HashMap[VertexId, Double]()
    val delta = new mutable.HashMap[VertexId, Double]()
    val neighborMap = getNeighborMap(vAttr.vlist, vAttr.elist)
    val medBC = new ListBuffer[(VertexId, Double)]()

    for (vertex <- vAttr.vlist) {
      dist.put(vertex, -1)
      sigma.put(vertex, 0.0)
      delta.put(vertex, 0.0)
      P.put(vertex, ListBuffer[VertexId]())
    }
    //对于当前节点，有特殊对待
    dist(vid) = 0.0
    sigma(vid) = 1.0
    Q.enqueue(vid)

    while (Q.nonEmpty) {
      val v = Q.dequeue()
      S.push(v)
      for (w <- neighborMap(v)) {
        if (dist(w) < 0) { //节点w未被访问过
          Q.enqueue(w)
          dist(w) = dist(v) + 1
        }
        if (dist(w) == dist(v) + 1) {
          sigma(w) += sigma(v)
          P(w).+=(v)
        }
      }
    }

    while (S.nonEmpty) {
      val w = S.pop()
      for (v <- P(w)) {
        delta(v) += sigma(v) / sigma(w) * (1 + delta(w))
      }
      if (w != vid)
        medBC.append((w, delta(w) / 2)) //一条边会被两个节点各自计算一次，所以需要对求出来的值除以2
    }
    medBC.toList
  }

  /**
   * 为每个顶点收集其邻居节点信息
   * 尝试过用收集邻居节点的api，但在计算介数中心性内部又需要每个节点都维护所有节点信息和边信息，所以采用对每个节点根据边来计算邻居节点的方式
   *
   * @param vlist , elist     所有顶点信息和所有边信息
   * @return [vid, [ 邻居id， 邻居id ...] ]
   * */
  def getNeighborMap(vlist: List[VertexId],
                     elist: List[(VertexId, VertexId, Double)]): mutable.HashMap[VertexId, List[VertexId]] = {
    val neighborList = new mutable.HashMap[VertexId, List[VertexId]]()
    vlist.map(v => {
      val nlist = (elist
        .filter(e => (e._1 == v || e._2 == v)))
        .map(e => {
          if (v == e._1) e._2
          else e._1
        })
      neighborList.+=((v, nlist.distinct))
    })
    neighborList
  }

  /**
   * 将每个节点分别计算出来的BC值进行统计
   * */
  def aggregateBetweennessScores(
                                  BCgraph: Graph[(VertexId, List[(VertexId, Double)]), Double]): Graph[Double, Double] = {
    //将图中顶点属性所维护的listBC信息单独提取出来
    val BCaggregate = BCgraph.vertices.flatMap {
      case (_, (_, listBC)) =>
        listBC.map { case (w, bc) => (w, bc) }
    }
    //对BCaggregate的信息 （w, bc）根据顶点id做汇总
    val vertexBC = BCaggregate.reduceByKey(_ + _)
    val resultG = BCgraph.outerJoinVertices(vertexBC)((vid, oldAttr, vBC) => {
      vBC.getOrElse(0.0)
    })
    resultG
  }

  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")

    val graph = ClientConfig.graphClient.loadInitGraphForSoftware(graphID)

    val graphForBC = Graph.fromEdges(graph.edges.mapValues(_ => 1.0), None)

    val spark = ClientConfig.spark

    val initBCgraph = createBetweenGraph(graphForBC, k = 3)
    val vertexBCGraph = initBCgraph.mapVertices((id, attr) => {
      (id, betweennessCentralityForUnweightedGraph(id, attr))
    })
    val BCGraph = aggregateBetweennessScores(vertexBCGraph).outerJoinVertices(graph.vertices) {
      (_, score, release) => {
        (score, release)
      }
    }

    val df = spark.sqlContext.createDataFrame(BCGraph.vertices.map(r => Row.apply(r._1, r._2._2.get.Artifact, r._2._2.get.Artifact, r._2._1)), SCHEMA_SOFTWARE).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.close()
  }
}
