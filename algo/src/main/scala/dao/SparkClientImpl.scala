package dao

import com.facebook.thrift.protocol.TCompactProtocol
import com.vesoft.nebula.connector.NebulaConnectionConfig
import config.AlgoConstants
import org.apache.spark.SparkConf
import org.apache.spark.graphx.{Edge, EdgeDirection, EdgeTriplet, Graph, Pregel, VertexId, VertexRDD}
import org.apache.spark.graphx.lib._
import org.apache.spark.rdd.RDD
import org.apache.spark.sql.{DataFrame, Row, SparkSession}
import org.apache.spark.sql.types.{DoubleType, LongType, StructField, StructType}
import service.{LouvainConfig, PRConfig}
import util.GraphUtil

import scala.collection.mutable
import scala.collection.mutable.ListBuffer


object SparkClientImpl extends SparkClient {

  var spark: SparkSession = _

  var nebulaCfg: NebulaConnectionConfig = _

  private val schema = StructType(
    List(
      StructField(AlgoConstants.NODE_ID_COL, LongType, nullable = false),
      StructField(AlgoConstants.SCORE_COL, DoubleType, nullable = true)
    ))

  def Init(): SparkClient = {
    val sparkConf = new SparkConf()
      .set("spark.serializer", "org.apache.spark.serializer.KryoSerializer")
      .registerKryoClasses(Array[Class[_]](classOf[TCompactProtocol]))
    spark = SparkSession
      .builder()
      .master("local")
      .config(sparkConf)
      .getOrCreate()
    spark.sparkContext.setLogLevel("WARN")
    nebulaCfg =
      NebulaConnectionConfig
        .builder()
        .withMetaAddress("127.0.0.1:9559")
        .withConenctionRetry(2)
        .withExecuteRetry(2)
        .withTimeout(6000)
        .build()
    this
  }

  override def degree(graphID: Long): (DataFrame, Option[Exception]) = {
    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    val schema = StructType(
      List(
        StructField(AlgoConstants.NODE_ID_COL, LongType, nullable = false),
        StructField(AlgoConstants.SCORE_COL, DoubleType, nullable = true)
      ))
    (spark.sqlContext
      .createDataFrame(graph.degrees.map(r => Row.apply(r._1, r._2)), schema).sort(AlgoConstants.SCORE_COL), Option.empty)
  }

  override def pagerank(graphID: Long, cfg: PRConfig): (DataFrame, Option[Exception]) = {
    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    val prResultRDD = PageRank.run(graph, cfg.iter.toInt, cfg.prob).vertices.map(r => Row.apply(r._1, r._2))
    (spark.sqlContext
      .createDataFrame(prResultRDD, schema).sort(AlgoConstants.SCORE_COL), Option.empty)
  }

  override def betweenness(graphID: Long): (DataFrame, Option[Exception]) = {
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
    def getNeighborMap(
                        vlist: List[VertexId],
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

    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    val initBCgraph = createBetweenGraph(graph, k = 100)
    val vertexBCGraph = initBCgraph.mapVertices((id, attr) => {
      (id, betweennessCentralityForUnweightedGraph(id, attr))
    })
    val BCGraph = aggregateBetweennessScores(vertexBCGraph)
    (spark.sqlContext
      .createDataFrame(BCGraph.vertices.map(r => Row.apply(r._1, r._2)), schema).sort(AlgoConstants.SCORE_COL), Option.empty)
  }

  override def closeness(graphID: Long): (DataFrame, Option[Exception]) = {
    type SPMap = Map[VertexId, Double]

    def makeMap(x: (VertexId, Double)*) = Map(x: _*)

    def addMap(spmap: SPMap, weight: Double): SPMap = spmap.map {
      case (v, d) => v -> (d + weight)
    }

    def addMaps(spmap1: SPMap, spmap2: SPMap): SPMap = {
      (spmap1.keySet ++ spmap2.keySet).map { k =>
        k -> math.min(spmap1.getOrElse(k, Double.MaxValue), spmap2.getOrElse(k, Double.MaxValue))
      }(collection.breakOut)
    }

    def vertexProgram(id: VertexId, attr: SPMap, msg: SPMap): SPMap = {
      addMaps(attr, msg)
    }

    def sendMessage(edge: EdgeTriplet[SPMap, Double]): Iterator[(VertexId, SPMap)] = {
      val newAttr = addMap(edge.dstAttr, edge.attr)
      if (edge.srcAttr != addMaps(newAttr, edge.srcAttr)) Iterator((edge.srcId, newAttr))
      else Iterator.empty
    }

    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    val spGraph = graph.mapVertices((vid, _) => makeMap(vid -> 0.0))
    val initialMessage = makeMap()

    val spsGraph = Pregel(spGraph, initialMessage)(vertexProgram, sendMessage, addMaps)
    val closenessRDD = spsGraph.vertices.map(vertex => {
      var dstNum = 0
      var dstDistanceSum = 0.0
      for (distance <- vertex._2.values) {
        dstNum += 1
        dstDistanceSum += distance
      }
      Row(vertex._1, (dstNum - 1) / dstDistanceSum)
    })
    (spark.sqlContext
      .createDataFrame(closenessRDD, schema).sort(AlgoConstants.SCORE_COL), Option.empty)
  }

  override def voterank(graphID: VertexId): (DataFrame, Option[Exception]) = ???

  override def clusteringCoefficient(graphID: VertexId): (Double, Option[Exception]) = {
    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    val closedTriangleNum = graph.triangleCount().vertices.map(_._2).reduce(_ + _)
    // compute the number of open triangle and closed triangle (According to C(n,2)=n*(n-1)/2)
    val triangleNum = graph.degrees.map(vertex => (vertex._2 * (vertex._2 - 1)) / 2.0).reduce(_ + _)
    if (triangleNum == 0)
      (0.0, Option.empty)
    else
      ((closedTriangleNum / triangleNum * 1.0).formatted("%.6f").toDouble, Option.empty)
  }

  override def louvain(graphID: VertexId, cfg: LouvainConfig): (DataFrame, Option[Exception]) = {
    /**
     * Louvain step1：Traverse the vertices and get the new community information of each node遍历节点，获取每个节点对应的所属新社区信息
     *
     *   1. Calculate the information of the community that each vertex currently belongs to.
     *      2. Calculate the community modularity deference and get the community info which has max modularity deference.
     *      3. Count the number of vertices that have changed in the community, used to determine whether the internal iteration can stop.
     *      4. Update vertices' community id and update each community's innerVertices.
     *
     * This step update vertexData's cid and commVertex.
     *
     * @param maxIter max interation
     * @param louvainG
     * @param m       The sum of the weights of all edges in the graph
     * @return (Graph[VertexData,Double],Int)
     */
    def step1(
               maxIter: Int,
               louvainG: Graph[VertexData, Double],
               m: Double,
               tol: Double
             ): (Graph[VertexData, Double], Int) = {
      var G = louvainG
      var iterTime = 0
      var canStop = false
      while (iterTime < maxIter && !canStop) {
        val neighborComm = getNeighCommInfo(G)
        val changeInfo = getChangeInfo(G, neighborComm, m, tol)
        // Count the number of vertices that have changed in the community
        val changeCount =
          G.vertices.zip(changeInfo).filter(x => x._1._2.cId != x._2._2).count()
        if (changeCount == 0)
          canStop = true
        // use connectedComponents algorithm to solve the problem of community attribution delay.
        else {
          val newChangeInfo = Graph
            .fromEdgeTuples(changeInfo.map(x => (x._1, x._2)), 0)
            .connectedComponents()
            .vertices
          G = LouvainGraphUtil.updateGraph(G, newChangeInfo)
          iterTime += 1
        }
      }
      (G, iterTime)
    }

    /**
     * Louvain step 2：Combine the new graph node obtained in the first step into a super node according to
     * the community information to which it belongs.
     *
     * @param G graph
     * @return graph
     */
    def step2(G: Graph[VertexData, Double]): Graph[VertexData, Double] = {
      //get edges between different communities
      val edges = G.triplets
        .filter(trip => trip.srcAttr.cId != trip.dstAttr.cId)
        .map(trip => {
          val cid1 = trip.srcAttr.cId
          val cid2 = trip.dstAttr.cId
          val weight = trip.attr
          ((math.min(cid1, cid2), math.max(cid1, cid2)), weight)
        })
        .reduceByKey(_ + _)
        .map(x => Edge(x._1._1, x._1._2, x._2)) //sum the edge weights between communities

      // sum kin of all vertices within the same community
      val vInnerKin = G.vertices
        .map(v => (v._2.cId, (v._2.innerVertices.toSet, v._2.innerDegree)))
        .reduceByKey((x, y) => {
          val vertices = (x._1 ++ y._1).toSet
          val kIn = x._2 + y._2
          (vertices, kIn)
        })

      // get all edge weights within the same community
      val v2vKin = G.triplets
        .filter(trip => trip.srcAttr.cId == trip.dstAttr.cId)
        .map(trip => {
          val cid = trip.srcAttr.cId
          val vertices1 = trip.srcAttr.innerVertices
          val vertices2 = trip.dstAttr.innerVertices
          val weight = trip.attr * 2
          (cid, (vertices1.union(vertices2).toSet, weight))
        })
        .reduceByKey((x, y) => {
          val vertices = new mutable.HashSet[VertexId].toSet
          val kIn = x._2 + y._2
          (vertices, kIn)
        })

      // new super vertex
      val superVertexInfo = vInnerKin
        .union(v2vKin)
        .reduceByKey((x, y) => {
          val vertices = x._1 ++ y._1
          val kIn = x._2 + y._2
          (vertices, kIn)
        })

      // reconstruct graph based on new edge info
      val initG = Graph.fromEdges(edges, None)
      var louvainGraph = LouvainGraphUtil.createLouvainGraph(initG)
      // get new louvain graph
      louvainGraph = louvainGraph.outerJoinVertices(superVertexInfo)((vid, data, opt) => {
        var innerVerteices = new mutable.HashSet[VertexId]()
        val kIn = opt.get._2
        for (vid <- opt.get._1)
          innerVerteices += vid
        data.innerVertices = innerVerteices
        data.innerDegree = kIn
        data
      })
      louvainGraph
    }

    /**
     * get new community's basic info after the vertex joins the community
     *   1. get each vertex's community id and the community's tot.
     *      2. compute each vertex's k_in. (The sum of the edge weights between vertex and vertex i in the community)
     *
     * @param G
     */
    def getNeighCommInfo(
                          G: Graph[VertexData, Double]
                        ): RDD[(VertexId, Iterable[(Long, Double, Double)])] = {

      val commKIn = G.triplets
        .flatMap(trip => {
          Array(
            (
              trip.srcAttr.cId,
              (
                (trip.dstId -> trip.attr),
                (trip.srcId, trip.srcAttr.innerDegree + trip.srcAttr.degree)
              )
            ),
            (
              trip.dstAttr.cId,
              (
                (trip.srcId -> trip.attr),
                (trip.dstId, trip.dstAttr.innerDegree + trip.dstAttr.degree)
              )
            )
          )
        })
        .groupByKey()
        .map(t => {
          val cid = t._1
          // add the weight of the same vid in one community.
          val m = new mutable.HashMap[VertexId, Double]() // store community's vertexId and vertex kin
          val degrees = new mutable.HashSet[VertexId]() // record if all vertices has computed the tot
          var tot = 0.0
          for (x <- t._2) {
            if (m.contains(x._1._1))
              m(x._1._1) += x._1._2
            else
              m(x._1._1) = x._1._2
            // compute vertex's tot
            if (!degrees.contains(x._2._1)) {
              tot += x._2._2
              degrees += x._2._1
            }
          }
          (cid, (tot, m))
        })

      // convert commKIn
      val neighCommInfo = commKIn
        .flatMap(x => {
          val cid = x._1
          val tot = x._2._1
          x._2._2.map(t => {
            val vid = t._1
            val kIn = t._2
            (vid, (cid, kIn, tot))
          })
        })
        .groupByKey()

      neighCommInfo
    }

    /**
     * Calculate the influence of each vertex on the modularity change of neighbor communities, and find the most suitable community for the vertex
     * △Q = [Kin - Σtot * Ki / m]
     *
     * @param G             graph
     * @param neighCommInfo neighbor community info
     * @param m             broadcast value
     * @param tol           threshold for modularity deference
     * @return RDD
     */
    def getChangeInfo(G: Graph[VertexData, Double],
                      neighCommInfo: RDD[(VertexId, Iterable[(Long, Double, Double)])],
                      m: Double,
                      tol: Double): RDD[(VertexId, Long, Double)] = {
      val changeInfo = G.vertices
        .join(neighCommInfo)
        .map(x => {
          val vid = x._1
          val data = x._2._1
          val commIter = x._2._2
          val vCid = data.cId
          val k_v = data.degree + data.innerDegree

          val dertaQs = commIter.map(t => {
            val nCid = t._1 // neighbor community id
            val k_v_in = t._2
            var tot = t._3
            if (vCid == nCid)
              tot -= k_v
            val q = (k_v_in - tot * k_v / m)
            (vid, nCid, q)
          })
          val maxQ =
            dertaQs.max(Ordering.by[(VertexId, Long, Double), Double](_._3))
          if (maxQ._3 > tol)
            maxQ
          else // if entering other communities reduces its modularity, then stays in the current community
            (vid, vCid, 0.0)
        })

      changeInfo //(vid,wCid,△Q)
    }

    object LouvainGraphUtil {

      /**
       * Construct louvain graph
       *
       * @param initG
       * @return Graph
       */
      def createLouvainGraph(
                              initG: Graph[None.type, Double]
                            ): Graph[VertexData, Double] = {
        // sum of the weights of the links incident to node i
        val nodeWeights: VertexRDD[Double] = initG.aggregateMessages(
          trip => {
            trip.sendToSrc(trip.attr)
            trip.sendToDst(trip.attr)
          },
          (a, b) => a + b
        )
        // update graph vertex's property
        val louvainG = initG.outerJoinVertices(nodeWeights)((vid, oldData, opt) => {
          val vData = new VertexData(vid, vid)
          val weights = opt.getOrElse(0.0)
          vData.degree = weights
          vData.innerVertices += vid
          vData
        })
        louvainG
      }

      /**
       * update graph using new community info
       *
       * @param G          Louvain graph
       * @param changeInfo （vid，new_cid）
       * @return Graph[VertexData, Double]
       */
      def updateGraph(
                       G: Graph[VertexData, Double],
                       changeInfo: RDD[(VertexId, Long)]
                     ): Graph[VertexData, Double] = {
        // update community id
        G.joinVertices(changeInfo)((vid, data, newCid) => {
          val vData = new VertexData(vid, newCid)
          vData.innerDegree = data.innerDegree
          vData.innerVertices = data.innerVertices
          vData.degree = data.degree
          vData
        })
      }
    }

    object CommUtil {
      // return the collections of communities
      def getCommunities(G: Graph[VertexData, Double]): RDD[Row] = {
        val communities = G.vertices
          .flatMapValues(_.innerVertices)
          .map(value => {
            Row(value._2, value._1)
          })
        communities
      }
    }

    class VertexData(val vId: Long, var cId: Long) extends Serializable {
      var innerDegree = 0.0
      var innerVertices = new mutable.HashSet[Long]()
      var degree = 0.0
    }

    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    val sc = spark.sparkContext

    // convert origin graph to Louvain Graph, Louvain Graph records vertex's community、innerVertices and innerDegrees
    var louvainG: Graph[VertexData, Double] = LouvainGraphUtil.createLouvainGraph(graph)

    // compute and broadcast the sum of all edges
    val m = sc.broadcast(louvainG.edges.map(e => e.attr).sum())
    var curIter = 0
    var res = step1(cfg.internalIter.toInt, louvainG, m.value, cfg.tol)
    while (res._2 != 0 && curIter < cfg.maxIter) {
      louvainG = res._1
      louvainG = step2(louvainG)
      res = step1(cfg.internalIter.toInt, louvainG, m.value, cfg.tol)
      curIter += 1
    }
    (spark.sqlContext
      .createDataFrame(CommUtil.getCommunities(louvainG), schema).sort(AlgoConstants.SCORE_COL), Option.empty)
  }
}
