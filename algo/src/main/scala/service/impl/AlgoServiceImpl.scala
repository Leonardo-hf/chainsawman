package service.impl

//#import

import akka.actor.typed.ActorSystem
import config.ClientConfig
import model.AlgoPO
import service._

import scala.concurrent.Future

//#import

//#service-request-reply
//#service-stream
class AlgoServiceImpl(system: ActorSystem[_]) extends algo {
  private implicit val sys: ActorSystem[_] = system

  //#service-request-reply

  override def createAlgo(in: CreateAlgoReq): Future[AlgoReply] = {
    val (_, err) = ClientConfig.mysqlClient.createAlgo(AlgoPO(name = in.name, desc = Option.apply(in.desc), `type` = in.`type`.value))
    if (err != null) {
      return Future.failed(err)
    }
    Future.successful(AlgoReply(algos = List.apply(Algo(name = in.name, desc = in.desc, `type` = in.`type`, isCustom = true))))
  }

  override def queryAlgo(in: Empty): Future[AlgoReply] = {
    val (res, err) = ClientConfig.mysqlClient.queryAlgo()
    if (err != null) {
      return Future.failed(err)
    }
    Future.successful(AlgoReply(algos = res.map(a => Algo(name = a.name, desc = a.desc.get, `type` = Algo.Type.fromValue(a.`type`), isCustom = a.isCustom))))
  }

  override def dropAlgo(in: DropAlgoReq): Future[AlgoReply] = {
    val (_, err) = ClientConfig.mysqlClient.dropAlgo(in.name)
    if (err != null) {
      return Future.failed(err)
    }
    Future.successful(AlgoReply())
  }

  override def degree(in: BaseReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.degree(in.graph)
    if (err != null) {
      return Future.failed(err)
    }
    Future.successful(RankReply(ranks = res.ranks.map(r => Rank(name = r.node, score = r.score))))
  }

  override def pagerank(in: BaseReq): Future[RankReply] = ???

  override def louvain(in: BaseReq): Future[ClusterReply] = ???

  override def shortestPath(in: ShortestPathReq): Future[ClusterReply] = ???

  override def avgShortestPath(in: BaseReq): Future[MetricsReply] = ???

  override def avgClustering(in: BaseReq): Future[MetricsReply] = ???

  override def custom(in: CustomAlgoReq): Future[CustomAlgoReply] = ???
}
//#service-stream
//#service-request-reply
