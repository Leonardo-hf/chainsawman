package model

import scalikejdbc._

case class AlgoParamPO(id: Long = 0L,
                       algoid: Long = 0L,
                       fieldname: String = "",
                       fieldnote: String = "",
                       fieldtype: Int = 0
                      )

object AlgoParamPO extends SQLSyntaxSupport[AlgoParamPO] {
  override val tableName = "algoParam"
//
//  def apply(algoID: Long, fieldName: String, fieldType: Int): AlgoParamPO = {
//    new AlgoParamPO(0L, algoID, fieldName, fieldType)
//  }

  def apply(o: SyntaxProvider[AlgoParamPO])(rs: WrappedResultSet): AlgoParamPO =
    apply(o.resultName)(rs)

  def apply(o: ResultName[AlgoParamPO])(rs: WrappedResultSet): AlgoParamPO = {
    if (rs.longOpt(o.id).isEmpty){
      return AlgoParamPO()
    }
    new AlgoParamPO(
      rs.long(o.id),
      rs.long(o.algoid),
      rs.string(o.fieldname),
      rs.string(o.fieldnote),
      rs.int(o.fieldtype)
    )
  }
}
