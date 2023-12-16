package applerodite.model

import scalikejdbc._

case class AlgoTaskPO(output: String,
                      status: Int)

object AlgoTaskPO extends SQLSyntaxSupport[AlgoTaskPO] {
  override val tableName = "exec"

  def apply(o: SyntaxProvider[AlgoTaskPO])(rs: WrappedResultSet): AlgoTaskPO =
    apply(o.resultName)(rs)

  def apply(o: ResultName[AlgoTaskPO])(rs: WrappedResultSet): AlgoTaskPO =
    new AlgoTaskPO(
      rs.string(o.output),
      rs.int(o.status)
    )
}