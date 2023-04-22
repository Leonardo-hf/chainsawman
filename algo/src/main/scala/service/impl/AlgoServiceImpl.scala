package service.impl

//#import
import akka.actor.typed.ActorSystem
import service._

import scala.concurrent.Future

//#import

//#service-request-reply
//#service-stream
class AlgoServiceImpl(system: ActorSystem[_]) extends algo {
  private implicit val sys: ActorSystem[_] = system

  //#service-request-reply

  override def createAlgo(in: CreateAlgoReq): Future[AlgoReply] = {

    Future.successful(new AlgoReply)
  }

  override def queryAlgo(in: Empty): Future[AlgoReply] = ???

  override def dropAlgo(in: DropAlgoReq): Future[AlgoReply] = ???

  override def degree(in: BaseReq): Future[RankReply] = ???

  override def pagerank(in: BaseReq): Future[RankReply] = ???

  override def louvain(in: BaseReq): Future[ClusterReply] = ???

  override def shortestPath(in: ShortestPathReq): Future[ClusterReply] = ???

  override def avgShortestPath(in: BaseReq): Future[MetricsReply] = ???

  override def avgClustering(in: BaseReq): Future[MetricsReply] = ???

  override def custom(in: CustomAlgoReq): Future[CustomAlgoReply] = ???
}
//#service-stream
//#service-request-reply
