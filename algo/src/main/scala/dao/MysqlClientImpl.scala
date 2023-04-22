package dao

import model.AlgoPO
import scalikejdbc.config.DBs
import scalikejdbc.{DB, delete, insert, select, withSQL}

object MysqlClientImpl extends MysqlClient {

  def Init(): MysqlClient = {
    DBs.setup()
    this
  }


  override def createAlgo(algo: AlgoPO): (Int, Exception) = {
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
        (res, null)
      }
    } catch {
      case e: Exception => (0, e)
    }
  }

  override def queryAlgo(): (List[AlgoPO], Exception) = {
    try {
      DB autoCommit { implicit s =>
        val a = AlgoPO.syntax("a")
        val res: List[AlgoPO] = withSQL {
          select.
            from(AlgoPO as a)
        }.map(AlgoPO(a)).list.apply()
        (res, null)
      }
    } catch {
      case e: Exception => (null, e)
    }
  }

  override def dropAlgo(name: String): (Int, Exception) = {
    try {
      DB autoCommit { implicit s =>
        val a = AlgoPO.column
        val res = withSQL {
          delete.from(AlgoPO).where.eq(a.name, name)
        }.update.apply()
        (res, null)
      }
    } catch {
      case e: Exception => (null, e)
    }
  }
}
