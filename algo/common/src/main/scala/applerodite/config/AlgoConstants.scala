package applerodite.config

import org.apache.spark.sql.types.{DoubleType, LongType, StructField, StructType}

object AlgoConstants {
  val NODE_ID_COL: String = "id"
  val SCORE_COL: String = "score"

  val SCHEMA_RANK: StructType = StructType(
    List(
      StructField(AlgoConstants.NODE_ID_COL, LongType, nullable = false),
      StructField(AlgoConstants.SCORE_COL, DoubleType, nullable = true)
    ))
}