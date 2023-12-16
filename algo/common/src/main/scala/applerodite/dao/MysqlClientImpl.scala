package applerodite.dao

import applerodite.model.AlgoTaskPO
import scalikejdbc.config.DBs
import scalikejdbc.{DB, update, withSQL}

object MysqlClientImpl extends MysqlClient {

  def Init(): MysqlClient = {
    DBs.setup()
    this
  }

  override def updateStatusByOutput(output: String): (Int, Option[Exception]) = {
    try {
      DB autoCommit { implicit s =>
        val res = withSQL {
          val ap = AlgoTaskPO.column
          update(AlgoTaskPO).set(ap.status -> 1).where.eq(ap.output, output)
        }.update.apply()
        (res, Option.empty)
      }
    }
    catch {
      case e: Exception => (0, Option.apply(e))
    }
  }
}