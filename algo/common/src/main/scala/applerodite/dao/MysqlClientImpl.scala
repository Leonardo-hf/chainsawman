package applerodite.dao

import applerodite.model.AlgoTaskPO
import scalikejdbc.{ConnectionPool, DB, update, withSQL}

import java.time.ZonedDateTime

case class MysqlConfig(driver: String, url: String, user: String, password: String)

object MysqlClientImpl extends MysqlClient {

  def Init(config: MysqlConfig): MysqlClient = {
    Class.forName(config.driver)
    ConnectionPool.singleton(config.url, config.user, config.password)
    this
  }

  override def updateStatusByOutput(output: String): (Int, Option[Exception]) = {
    try {
      DB autoCommit { implicit s =>
        val res = withSQL {
          val ap = AlgoTaskPO.column
          update(AlgoTaskPO).set(ap.status -> 1, ap.updateTime -> ZonedDateTime.now()).where.eq(ap.output, output)
        }.update.apply()
        (res, Option.empty)
      }
    }
    catch {
      case e: Exception => (0, Option.apply(e))
    }
  }
}