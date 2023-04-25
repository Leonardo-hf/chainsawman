package service.impl

//#import

import akka.actor.typed.ActorSystem
import config.ClientConfig
import model.{AlgoPO, AlgoParamPO}
import service._

import scala.concurrent.Future

//#import

//#service-request-reply
//#service-stream
class AlgoServiceImpl(system: ActorSystem[_]) extends algo {
  private implicit val sys: ActorSystem[_] = system

  //#service-request-reply

  override def createAlgo(in: CreateAlgoReq): Future[AlgoReply] = {
    val (algoID, err) = ClientConfig.mysqlClient.createAlgo(AlgoPO(name = in.name, note = Option.apply(in.desc), `type` = in.`type`.value))
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (_, pErr) = ClientConfig.mysqlClient.multiCreateAlgoParams(in.params.map(e => AlgoParamPO(algoid = algoID, fieldname = e.key, fieldtype = e.`type`.value)).seq)
    if (pErr.nonEmpty) {
      return Future.failed(pErr.get)
    }
    Future.successful(AlgoReply(algos = List.apply(Algo(name = in.name, desc = in.desc, `type` = in.`type`, isCustom = true, params = in.params))))
  }

  override def queryAlgo(in: Empty): Future[AlgoReply] = {
    val (res, err) = ClientConfig.mysqlClient.queryAlgo()
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    Future.successful(AlgoReply(algos = res.map(a => Algo(name = a.name, desc = a.note.get, `type` = Algo.Type.fromValue(a.`type`), isCustom = a.iscustom,
      params = a.params.map(p => Element.apply(key = p.fieldname, `type` = Element.Type.fromValue(p.fieldtype)))))))
  }

  override def dropAlgo(in: DropAlgoReq): Future[AlgoReply] = {
    val (_, err) = ClientConfig.mysqlClient.dropAlgo(in.name)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    Future.successful(AlgoReply())
  }

  override def degree(in: BaseReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.degree(in.graphID)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    Future.successful(RankReply(ranks = res.ranks.map(r => Rank(id = r.nodeID, score = r.score))))
  }

  override def pagerank(in: PageRankReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.pagerank(in.base.get.graphID, in.cfg.get)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    Future.successful(RankReply(ranks = res.ranks.map(r => Rank(id = r.nodeID, score = r.score))))
  }

  override def louvain(in: BaseReq): Future[ClusterReply] = ???

  override def shortestPath(in: ShortestPathReq): Future[ClusterReply] = ???

  override def avgShortestPath(in: BaseReq): Future[MetricsReply] = ???

  override def avgClustering(in: BaseReq): Future[MetricsReply] = ???

  override def custom(in: CustomAlgoReq): Future[CustomAlgoReply] = {
    // TODO: 需求太小众，不打算写了
    //    val graph = in.base.get.graphID
    //    val (algo, err) = ClientConfig.mysqlClient.queryCustomAlgo(in.algoID)
    Future.successful(CustomAlgoReply())
  }


}
//#service-stream
//#service-request-reply
