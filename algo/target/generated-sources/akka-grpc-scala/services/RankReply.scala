// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package services

@SerialVersionUID(0L)
final case class RankReply(
    ranks: _root_.scala.Seq[services.Rank] = _root_.scala.Seq.empty,
    unknownFields: _root_.scalapb.UnknownFieldSet = _root_.scalapb.UnknownFieldSet.empty
    ) extends scalapb.GeneratedMessage with scalapb.lenses.Updatable[RankReply] {
    @transient
    private[this] var __serializedSizeMemoized: _root_.scala.Int = 0
    private[this] def __computeSerializedSize(): _root_.scala.Int = {
      var __size = 0
      ranks.foreach { __item =>
        val __value = __item
        __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(__value.serializedSize) + __value.serializedSize
      }
      __size += unknownFields.serializedSize
      __size
    }
    override def serializedSize: _root_.scala.Int = {
      var __size = __serializedSizeMemoized
      if (__size == 0) {
        __size = __computeSerializedSize() + 1
        __serializedSizeMemoized = __size
      }
      __size - 1
      
    }
    def writeTo(`_output__`: _root_.com.google.protobuf.CodedOutputStream): _root_.scala.Unit = {
      ranks.foreach { __v =>
        val __m = __v
        _output__.writeTag(1, 2)
        _output__.writeUInt32NoTag(__m.serializedSize)
        __m.writeTo(_output__)
      };
      unknownFields.writeTo(_output__)
    }
    def clearRanks = copy(ranks = _root_.scala.Seq.empty)
    def addRanks(__vs: services.Rank *): RankReply = addAllRanks(__vs)
    def addAllRanks(__vs: Iterable[services.Rank]): RankReply = copy(ranks = ranks ++ __vs)
    def withRanks(__v: _root_.scala.Seq[services.Rank]): RankReply = copy(ranks = __v)
    def withUnknownFields(__v: _root_.scalapb.UnknownFieldSet) = copy(unknownFields = __v)
    def discardUnknownFields = copy(unknownFields = _root_.scalapb.UnknownFieldSet.empty)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): _root_.scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ranks
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      _root_.scala.Predef.require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => _root_.scalapb.descriptors.PRepeated(ranks.iterator.map(_.toPMessage).toVector)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion: services.RankReply.type = services.RankReply
    // @@protoc_insertion_point(GeneratedMessage[services.RankReply])
}

object RankReply extends scalapb.GeneratedMessageCompanion[services.RankReply] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[services.RankReply] = this
  def parseFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): services.RankReply = {
    val __ranks: _root_.scala.collection.immutable.VectorBuilder[services.Rank] = new _root_.scala.collection.immutable.VectorBuilder[services.Rank]
    var `_unknownFields__`: _root_.scalapb.UnknownFieldSet.Builder = null
    var _done__ = false
    while (!_done__) {
      val _tag__ = _input__.readTag()
      _tag__ match {
        case 0 => _done__ = true
        case 10 =>
          __ranks += _root_.scalapb.LiteParser.readMessage[services.Rank](_input__)
        case tag =>
          if (_unknownFields__ == null) {
            _unknownFields__ = new _root_.scalapb.UnknownFieldSet.Builder()
          }
          _unknownFields__.parseField(tag, _input__)
      }
    }
    services.RankReply(
        ranks = __ranks.result(),
        unknownFields = if (_unknownFields__ == null) _root_.scalapb.UnknownFieldSet.empty else _unknownFields__.result()
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[services.RankReply] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      _root_.scala.Predef.require(__fieldsMap.keys.forall(_.containingMessage eq scalaDescriptor), "FieldDescriptor does not match message type.")
      services.RankReply(
        ranks = __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).map(_.as[_root_.scala.Seq[services.Rank]]).getOrElse(_root_.scala.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AlgoProto.javaDescriptor.getMessageTypes().get(7)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AlgoProto.scalaDescriptor.messages(7)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = services.Rank
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_ <: _root_.scalapb.GeneratedMessage]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = services.RankReply(
    ranks = _root_.scala.Seq.empty
  )
  implicit class RankReplyLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, services.RankReply]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, services.RankReply](_l) {
    def ranks: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Seq[services.Rank]] = field(_.ranks)((c_, f_) => c_.copy(ranks = f_))
  }
  final val RANKS_FIELD_NUMBER = 1
  def of(
    ranks: _root_.scala.Seq[services.Rank]
  ): _root_.services.RankReply = _root_.services.RankReply(
    ranks
  )
  // @@protoc_insertion_point(GeneratedMessageCompanion[services.RankReply])
}
