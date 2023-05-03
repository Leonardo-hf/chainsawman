package dao

import model.{AlgoPO, AlgoPOWithParams, AlgoParamPO}
import scalikejdbc.config.DBs
import scalikejdbc.{DB, delete, insert, select, withSQL}

object MysqlClientImpl extends MysqlClient {

  def Init(): MysqlClient = {
    DBs.setup()
    this
  }

  override def multiCreateAlgoParams(params: Seq[AlgoParamPO]): (Int, Option[Exception]) = {
    try {
      (params.map(p => {
        DB autoCommit { implicit s =>
          withSQL {
            val ap = AlgoParamPO.column
            insert.into(AlgoParamPO).namedValues(
              ap.algoid -> p.algoid,
              ap.fieldname -> p.fieldname,
              ap.fieldtype -> p.fieldtype,
            )
          }.update.apply()
        }
      }).sum,
        Option.empty)
    }
    catch {
      case e: Exception => (0, Option.apply(e))
    }
  }

  override def createAlgo(algo: AlgoPO): (Long, Option[Exception]) = {
    try {
      DB autoCommit { implicit s =>
        val res = withSQL {
          val a = AlgoPO.column
          insert.into(AlgoPO).namedValues(
            a.name -> algo.name,
            a.note -> algo.note,
            a.`type` -> algo.`type`,
          )
        }.updateAndReturnGeneratedKey.apply()
        (res, Option.empty)
      }
    } catch {
      case e: Exception => (0, Option.apply(e))
    }
  }

  override def queryAlgo(): (List[AlgoPOWithParams], Option[Exception]) = {
    try {
      DB autoCommit { implicit s =>
        val a = AlgoPO.syntax("a")
        val ap = AlgoParamPO.syntax("ap")
        val res: List[AlgoPOWithParams] = withSQL {
          select.
            from(AlgoPO as a).leftJoin(AlgoParamPO as ap).on(a.id, ap.algoid)
        }.map(AlgoPOWithParams(a, ap)).list.apply()
        (res.groupBy(a => a.id).mapValues(ps => ps.reduce((p1, p2) => p1.copy(params = p1.params ++: p2.params))).values.toList, Option.empty)
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
