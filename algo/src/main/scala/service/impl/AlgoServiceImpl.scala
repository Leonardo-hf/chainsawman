package service.impl

//#import

import akka.actor.typed.ActorSystem
import config.ClientConfig
import model.{AlgoPO, AlgoParamPO}
import org.apache.spark.sql.DataFrame
import service._
import util.CSVUtil

import scala.concurrent.Future
import scala.language.postfixOps

//#import

//#service-request-reply
//#service-stream
class AlgoServiceImpl(system: ActorSystem[_]) extends algo {
  private implicit val sys: ActorSystem[_] = system

  private val preview = 10

  //#service-request-reply

  override def createAlgo(in: CreateAlgoReq): Future[AlgoReply] = {
    val (algoID, err) = ClientConfig.mysqlClient.createAlgo(AlgoPO(name = in.name, note = Option.apply(in.desc), `type` = in.`type`.value))
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (_, pErr) = ClientConfig.mysqlClient.multiCreateAlgoParams(in.params.map(e => AlgoParamPO(algoid = algoID, fieldname = e.key, fieldnote = e.keyDesc, fieldtype = e.`type`.value)).seq)
    if (pErr.nonEmpty) {
      return Future.failed(pErr.get)
    }
    Future.successful(AlgoReply(algos = List.apply(Algo(name = in.name, desc = in.desc, `type` = in.`type`, params = in.params))))
  }

  override def queryAlgo(in: Empty): Future[AlgoReply] = {
    val (res, err) = ClientConfig.mysqlClient.queryAlgo()
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    Future.successful(AlgoReply(algos = res.map(a => Algo(name = a.name, desc = a.note.get, `type` = Algo.Type.fromValue(a.`type`),
      params = a.params.map(p => Element.apply(key = p.fieldname, keyDesc = p.fieldnote, `type` = Element.Type.fromValue(p.fieldtype)))))))
  }

  override def dropAlgo(in: DropAlgoReq): Future[AlgoReply] = {
    val (_, err) = ClientConfig.mysqlClient.dropAlgo(in.name)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    Future.successful(AlgoReply())
  }

  def getRank(df: DataFrame, file: String): RankReply = {
    RankReply.apply(ranks = df.rdd.collect().reverse.slice(0, preview).map(s => Rank.apply(id = s.getLong(0), score = s.getDouble(1))).toSeq, file = file)
  }

  override def degree(in: BaseReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.degree(in.graphID, in.edgeTags)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "degree", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }

  override def pagerank(in: PageRankReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.pagerank(in.base.get.graphID, in.base.get.edgeTags, in.cfg.getOrElse(PRConfig.apply(iter = 3, prob = 0.85)))
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "pagerank", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }

  override def voterank(in: VoteRankReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.voterank(in.base.get.graphID, in.base.get.edgeTags, in.cfg.getOrElse(VoteConfig.apply(iter = 2000)))
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "voterank", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }

  override def depth(in: BaseReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.depth(in.graphID, in.edgeTags)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "depth", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }

  override def ecology(in: BaseReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.ecology(in.graphID, in.edgeTags)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "ecology", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }

  override def betweenness(in: BaseReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.betweenness(in.graphID, in.edgeTags)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "betweenness", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }

  override def closeness(in: BaseReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.closeness(in.graphID, in.edgeTags)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "closeness", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }

  override def avgClustering(in: BaseReq): Future[MetricsReply] = {
    val (res, err) = ClientConfig.sparkClient.clusteringCoefficient(in.graphID, in.edgeTags)
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    Future.successful(MetricsReply.apply(score = res))
  }

  override def custom(in: CustomAlgoReq): Future[CustomAlgoReply] = {
    // TODO: 暂时搁置
    //    val graph = in.base.get.graphID
    //    val (algo, err) = ClientConfig.mysqlClient.queryCustomAlgo(in.algoID)
    Future.successful(CustomAlgoReply())
  }

  override def louvain(in: LouvainReq): Future[RankReply] = {
    val (res, err) = ClientConfig.sparkClient.louvain(in.base.get.graphID, in.base.get.edgeTags, in.cfg.getOrElse(LouvainConfig.apply(maxIter = 10, internalIter = 5, tol = 0.5)))
    if (err.nonEmpty) {
      return Future.failed(err.get)
    }
    val (id, _) = ClientConfig.ossClient.upload(name = "louvain", content = CSVUtil.df2CSV(res))
    Future.successful(getRank(res, id))
  }
}
//#service-stream
//#service-request-reply
