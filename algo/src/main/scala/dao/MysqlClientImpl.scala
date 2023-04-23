package dao

import model.AlgoPO
import scalikejdbc.config.DBs
import scalikejdbc.{DB, delete, insert, select, withSQL}

object MysqlClientImpl extends MysqlClient {

  def Init(): MysqlClient = {
    DBs.setup()
    this
  }


  override def createAlgo(algo: AlgoPO): (Int, Option[Exception]) = {
    try {
      DB autoCommit { implicit s =>
        val res = withSQL {
          val a = AlgoPO.column
          insert.into(AlgoPO).namedValues(
            a.name -> algo.name,
            a.desc -> algo.desc,
            a.isCustom -> true,
            a.`type` -> algo.`type`
          )
        }.update.apply()
        (res, Option.empty)
      }
    } catch {
      case e: Exception => (0, Option.apply(e))
    }
  }

  override def queryAlgo(): (List[AlgoPO], Option[Exception]) = {
    try {
      DB autoCommit { implicit s =>
        val a = AlgoPO.syntax("a")
        val res: List[AlgoPO] = withSQL {
          select.
            from(AlgoPO as a)
        }.map(AlgoPO(a)).list.apply()
        (res, Option.empty)
      }
    } catch {
      case e: Exception => (List.empty, Option.apply(e))
    }
  }

  override def dropAlgo(name: String): (Int, Option[Exception]) = {
    try {
      DB autoCommit { implicit s =>
        val a = AlgoPO.column
        val res = withSQL {
          delete.from(AlgoPO).where.eq(a.name, name)
        }.update.apply()
        (res, Option.empty)
      }
    } catch {
      case e: Exception => (0, Option.apply(e))
    }
  }
}
