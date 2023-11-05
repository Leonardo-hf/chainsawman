package applerodite.config

import org.apache.spark.sql.types.{DoubleType, LongType, StringType, StructField, StructType}

object AlgoConstants {
  val NODE_ID_COL: String = "id"
  val SCORE_COL: String = "score"
  val ARTIFACT_COL: String = "artifact"
  val VERSION_COL: String = "version"

  val SCHEMA_DEFAULT: StructType = StructType(
    List(
      StructField(AlgoConstants.NODE_ID_COL, LongType, nullable = false),
      StructField(AlgoConstants.SCORE_COL, DoubleType, nullable = false)
    ))

  val SCHEMA_SCORE: StructType = StructType(
    List(
      StructField(AlgoConstants.SCORE_COL, DoubleType, nullable = false)
    ))

  val SCHEMA_SOFTWARE: StructType = StructType(
    List(
      StructField(AlgoConstants.NODE_ID_COL, LongType, nullable = false),
      StructField(AlgoConstants.ARTIFACT_COL, StringType, nullable = false),
      StructField(AlgoConstants.VERSION_COL, StringType, nullable = false),
      StructField(AlgoConstants.SCORE_COL, DoubleType, nullable = false)
    ))

}