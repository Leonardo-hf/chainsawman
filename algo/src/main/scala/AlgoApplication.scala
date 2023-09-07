import akka.actor.typed.ActorSystem
import akka.actor.typed.scaladsl.Behaviors
import com.typesafe.config.ConfigFactory
import config.ClientConfig
import server.GrpcServer

object AlgoApplication {

  def main(args: Array[String]): Unit = {
    // important to enable HTTP/2 in ActorSystem's config
    var conf = ConfigFactory.parseString("akka.http.server.preview.enable-http2 = on")
      .withFallback(ConfigFactory.defaultApplication())
    if (System.getenv("CHS_ENV") == "pre") {
      conf = ConfigFactory.load("application-pre.conf")
    }
    ClientConfig.Init(conf)
//    val (df, _) = ecology(3, Seq("edge_1"))
//    println(RankReply.apply(ranks = df.rdd.collect().reverse.slice(0, 100).map(s => Rank.apply(id = s.getLong(0), score = s.getDouble(1))).toSeq))
//    return
    val system = ActorSystem[Nothing](Behaviors.empty, "AlgoServer", conf)
    new GrpcServer(system).run()
  }
}
