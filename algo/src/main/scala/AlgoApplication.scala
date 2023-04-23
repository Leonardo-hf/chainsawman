import akka.actor.typed.ActorSystem
import akka.actor.typed.scaladsl.Behaviors
import com.typesafe.config.ConfigFactory
import config.ClientConfig
import server.GrpcServer

object AlgoApplication {
  def main(args: Array[String]): Unit = {
    // important to enable HTTP/2 in ActorSystem's config
    val conf = ConfigFactory.parseString("akka.http.server.preview.enable-http2 = on")
      .withFallback(ConfigFactory.defaultApplication())
    ClientConfig.Init()
    val system = ActorSystem[Nothing](Behaviors.empty, "AlgoServer", conf)
    new GrpcServer(system).run()
  }
}
