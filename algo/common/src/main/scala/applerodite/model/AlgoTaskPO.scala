package applerodite.model

import scalikejdbc._

import java.time.ZonedDateTime

// 注意到scalikejdbc会将数据库的驼峰列名转换为小写
// 而会将代码中的驼峰命名转化为下划线命名，需要注意对应
case class AlgoTaskPO(output: String,
                      status: Int,
                      updatetime: ZonedDateTime)

object AlgoTaskPO extends SQLSyntaxSupport[AlgoTaskPO] {
  override val tableName = "exec"

  def apply(o: SyntaxProvider[AlgoTaskPO])(rs: WrappedResultSet): AlgoTaskPO =
    apply(o.resultName)(rs)

  def apply(o: ResultName[AlgoTaskPO])(rs: WrappedResultSet): AlgoTaskPO =
    new AlgoTaskPO(
      rs.string(o.output),
      rs.int(o.status),
      rs.dateTime(o.updatetime)
    )
}