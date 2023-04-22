// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package service

@SerialVersionUID(0L)
final case class Cluster(
    name: _root_.scala.Predef.String = "",
    nodes: _root_.scala.Seq[_root_.scala.Predef.String] = _root_.scala.Seq.empty,
    unknownFields: _root_.scalapb.UnknownFieldSet = _root_.scalapb.UnknownFieldSet.empty
    ) extends scalapb.GeneratedMessage with scalapb.lenses.Updatable[Cluster] {
    @transient
    private[this] var __serializedSizeMemoized: _root_.scala.Int = 0
    private[this] def __computeSerializedSize(): _root_.scala.Int = {
      var __size = 0
      
      {
        val __value = name
        if (!__value.isEmpty) {
          __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(1, __value)
        }
      };
      nodes.foreach { __item =>
        val __value = __item
        __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, __value)
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
      {
        val __v = name
        if (!__v.isEmpty) {
          _output__.writeString(1, __v)
        }
      };
      nodes.foreach { __v =>
        val __m = __v
        _output__.writeString(2, __m)
      };
      unknownFields.writeTo(_output__)
    }
    def withName(__v: _root_.scala.Predef.String): Cluster = copy(name = __v)
    def clearNodes = copy(nodes = _root_.scala.Seq.empty)
    def addNodes(__vs: _root_.scala.Predef.String *): Cluster = addAllNodes(__vs)
    def addAllNodes(__vs: Iterable[_root_.scala.Predef.String]): Cluster = copy(nodes = nodes ++ __vs)
    def withNodes(__v: _root_.scala.Seq[_root_.scala.Predef.String]): Cluster = copy(nodes = __v)
    def withUnknownFields(__v: _root_.scalapb.UnknownFieldSet) = copy(unknownFields = __v)
    def discardUnknownFields = copy(unknownFields = _root_.scalapb.UnknownFieldSet.empty)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): _root_.scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => {
          val __t = name
          if (__t != "") __t else null
        }
        case 2 => nodes
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      _root_.scala.Predef.require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => _root_.scalapb.descriptors.PString(name)
        case 2 => _root_.scalapb.descriptors.PRepeated(nodes.iterator.map(_root_.scalapb.descriptors.PString(_)).toVector)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion: service.Cluster.type = service.Cluster
    // @@protoc_insertion_point(GeneratedMessage[services.Cluster])
}

object Cluster extends scalapb.GeneratedMessageCompanion[service.Cluster] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[service.Cluster] = this
  def parseFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): service.Cluster = {
    var __name: _root_.scala.Predef.String = ""
    val __nodes: _root_.scala.collection.immutable.VectorBuilder[_root_.scala.Predef.String] = new _root_.scala.collection.immutable.VectorBuilder[_root_.scala.Predef.String]
    var `_unknownFields__`: _root_.scalapb.UnknownFieldSet.Builder = null
    var _done__ = false
    while (!_done__) {
      val _tag__ = _input__.readTag()
      _tag__ match {
        case 0 => _done__ = true
        case 10 =>
          __name = _input__.readStringRequireUtf8()
        case 18 =>
          __nodes += _input__.readStringRequireUtf8()
        case tag =>
          if (_unknownFields__ == null) {
            _unknownFields__ = new _root_.scalapb.UnknownFieldSet.Builder()
          }
          _unknownFields__.parseField(tag, _input__)
      }
    }
    service.Cluster(
        name = __name,
        nodes = __nodes.result(),
        unknownFields = if (_unknownFields__ == null) _root_.scalapb.UnknownFieldSet.empty else _unknownFields__.result()
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[service.Cluster] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      _root_.scala.Predef.require(__fieldsMap.keys.forall(_.containingMessage eq scalaDescriptor), "FieldDescriptor does not match message type.")
      service.Cluster(
        name = __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        nodes = __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AlgoProto.javaDescriptor.getMessageTypes().get(8)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AlgoProto.scalaDescriptor.messages(8)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_ <: _root_.scalapb.GeneratedMessage]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = service.Cluster(
    name = "",
    nodes = _root_.scala.Seq.empty
  )
  implicit class ClusterLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, service.Cluster]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, service.Cluster](_l) {
    def name: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.name)((c_, f_) => c_.copy(name = f_))
    def nodes: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Seq[_root_.scala.Predef.String]] = field(_.nodes)((c_, f_) => c_.copy(nodes = f_))
  }
  final val NAME_FIELD_NUMBER = 1
  final val NODES_FIELD_NUMBER = 2
  def of(
    name: _root_.scala.Predef.String,
    nodes: _root_.scala.Seq[_root_.scala.Predef.String]
  ): _root_.service.Cluster = _root_.service.Cluster(
    name,
    nodes
  )
  // @@protoc_insertion_point(GeneratedMessageCompanion[services.Cluster])
}
