package model

import scalikejdbc._

case class AlgoPO(id: Long = 0L,
                  name: String = "",
                  note: Option[String] = None,
                  iscustom: Boolean = false,
                  `type`: Int = 0)

object AlgoPO extends SQLSyntaxSupport[AlgoPO] {
  override val tableName = "algo"

  def apply(o: SyntaxProvider[AlgoPO])(rs: WrappedResultSet): AlgoPO =
    apply(o.resultName)(rs)

  def apply(o: ResultName[AlgoPO])(rs: WrappedResultSet): AlgoPO =
    new AlgoPO(
      rs.long(o.id),
      rs.string(o.name),
      rs.stringOpt(o.note),
      rs.boolean(o.iscustom),
      rs.int(o.`type`)
    )
}

case class AlgoPOWithParams(id: Long = 0L,
                            name: String = "",
                            note: Option[String] = None,
                            iscustom: Boolean = false,
                            `type`: Int = 0,
                            params: Seq[AlgoParamPO] = List.empty)

object AlgoPOWithParams {
  def apply(a: SyntaxProvider[AlgoPO], ap: SyntaxProvider[AlgoParamPO])(rs: WrappedResultSet): AlgoPOWithParams =
    apply(a.resultName, ap.resultName)(rs)

  def apply(a: ResultName[AlgoPO], ap: ResultName[AlgoParamPO])(rs: WrappedResultSet): AlgoPOWithParams =
    new AlgoPOWithParams(
      rs.long(a.id),
      rs.string(a.name),
      rs.stringOpt(a.note),
      rs.boolean(a.iscustom),
      rs.int(a.`type`)
    ).copy(params = List.apply(AlgoParamPO.apply(ap)(rs)))
}