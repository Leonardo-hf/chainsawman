package model

import scalikejdbc._

case class AlgoPO(id: Long = 0L,
                  name: String = "",
                  desc: Option[String] = None,
                  isCustom: Boolean = false,
                  `type`: Int = 0)

object AlgoPO extends SQLSyntaxSupport[AlgoPO] {
  override val tableName = "algo"

  def apply(o: SyntaxProvider[AlgoPO])(rs: WrappedResultSet): AlgoPO =
    apply(o.resultName)(rs)

  def apply(o: ResultName[AlgoPO])(rs: WrappedResultSet): AlgoPO =
    new AlgoPO(
      rs.long(o.id),
      rs.string(o.name),
      rs.stringOpt(o.desc),
      rs.boolean(o.isCustom),
      rs.int(o.`type`)
    )
}