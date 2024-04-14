package applerodite.dao

import org.apache.log4j.Logger
import applerodite.model.AlgoTaskPO
import com.typesafe.config.ConfigFactory
import scalikejdbc.{ConnectionPool, DB, GlobalSettings, LoggingSQLAndTimeSettings, scalikejdbcSQLInterpolationImplicitDef, update, withSQL}

import java.time.ZonedDateTime

case class MysqlConfig(driver: String, url: String, user: String, password: String)

object MysqlClientImpl extends MysqlClient {

  def Init(config: MysqlConfig): MysqlClient = {
    Class.forName(config.driver)
    ConnectionPool.singleton(config.url, config.user, config.password)
    GlobalSettings.loggingSQLAndTime = LoggingSQLAndTimeSettings(
      enabled = true,
      singleLineMode = true,
      logLevel = "info"
    )
    this
  }

  override def updateStatusByOutput(output: String): (Int, Option[Exception]) = {
    DB localTx { implicit s =>
      val res = withSQL {
        val ap = AlgoTaskPO.column
        update(AlgoTaskPO).set(ap.status -> 1, ap.updatetime -> ZonedDateTime.now()).where.eq(ap.output, output)
      }.update.apply()
      (res, Option.empty)
    }
  }
}