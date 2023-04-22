// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package services

@SerialVersionUID(0L)
final case class Element(
    key: _root_.scala.Predef.String = "",
    `type`: services.Element.Type = services.Element.Type.INT64,
    unknownFields: _root_.scalapb.UnknownFieldSet = _root_.scalapb.UnknownFieldSet.empty
    ) extends scalapb.GeneratedMessage with scalapb.lenses.Updatable[Element] {
    @transient
    private[this] var __serializedSizeMemoized: _root_.scala.Int = 0
    private[this] def __computeSerializedSize(): _root_.scala.Int = {
      var __size = 0
      
      {
        val __value = key
        if (!__value.isEmpty) {
          __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(1, __value)
        }
      };
      
      {
        val __value = `type`.value
        if (__value != 0) {
          __size += _root_.com.google.protobuf.CodedOutputStream.computeEnumSize(2, __value)
        }
      };
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
        val __v = key
        if (!__v.isEmpty) {
          _output__.writeString(1, __v)
        }
      };
      {
        val __v = `type`.value
        if (__v != 0) {
          _output__.writeEnum(2, __v)
        }
      };
      unknownFields.writeTo(_output__)
    }
    def withKey(__v: _root_.scala.Predef.String): Element = copy(key = __v)
    def withType(__v: services.Element.Type): Element = copy(`type` = __v)
    def withUnknownFields(__v: _root_.scalapb.UnknownFieldSet) = copy(unknownFields = __v)
    def discardUnknownFields = copy(unknownFields = _root_.scalapb.UnknownFieldSet.empty)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): _root_.scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => {
          val __t = key
          if (__t != "") __t else null
        }
        case 2 => {
          val __t = `type`.javaValueDescriptor
          if (__t.getNumber() != 0) __t else null
        }
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      _root_.scala.Predef.require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => _root_.scalapb.descriptors.PString(key)
        case 2 => _root_.scalapb.descriptors.PEnum(`type`.scalaValueDescriptor)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion: services.Element.type = services.Element
    // @@protoc_insertion_point(GeneratedMessage[services.Element])
}

object Element extends scalapb.GeneratedMessageCompanion[services.Element] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[services.Element] = this
  def parseFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): services.Element = {
    var __key: _root_.scala.Predef.String = ""
    var __type: services.Element.Type = services.Element.Type.INT64
    var `_unknownFields__`: _root_.scalapb.UnknownFieldSet.Builder = null
    var _done__ = false
    while (!_done__) {
      val _tag__ = _input__.readTag()
      _tag__ match {
        case 0 => _done__ = true
        case 10 =>
          __key = _input__.readStringRequireUtf8()
        case 16 =>
          __type = services.Element.Type.fromValue(_input__.readEnum())
        case tag =>
          if (_unknownFields__ == null) {
            _unknownFields__ = new _root_.scalapb.UnknownFieldSet.Builder()
          }
          _unknownFields__.parseField(tag, _input__)
      }
    }
    services.Element(
        key = __key,
        `type` = __type,
        unknownFields = if (_unknownFields__ == null) _root_.scalapb.UnknownFieldSet.empty else _unknownFields__.result()
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[services.Element] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      _root_.scala.Predef.require(__fieldsMap.keys.forall(_.containingMessage eq scalaDescriptor), "FieldDescriptor does not match message type.")
      services.Element(
        key = __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        `type` = services.Element.Type.fromValue(__fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scalapb.descriptors.EnumValueDescriptor]).getOrElse(services.Element.Type.INT64.scalaValueDescriptor).number)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AlgoProto.javaDescriptor.getMessageTypes().get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AlgoProto.scalaDescriptor.messages(1)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_ <: _root_.scalapb.GeneratedMessage]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = {
    (__fieldNumber: @_root_.scala.unchecked) match {
      case 2 => services.Element.Type
    }
  }
  lazy val defaultInstance = services.Element(
    key = "",
    `type` = services.Element.Type.INT64
  )
  sealed abstract class Type(val value: _root_.scala.Int) extends _root_.scalapb.GeneratedEnum {
    type EnumType = Type
    def isInt64: _root_.scala.Boolean = false
    def isDouble: _root_.scala.Boolean = false
    def isString: _root_.scala.Boolean = false
    def isListDouble: _root_.scala.Boolean = false
    def isListString: _root_.scala.Boolean = false
    def companion: _root_.scalapb.GeneratedEnumCompanion[Type] = services.Element.Type
    final def asRecognized: _root_.scala.Option[services.Element.Type.Recognized] = if (isUnrecognized) _root_.scala.None else _root_.scala.Some(this.asInstanceOf[services.Element.Type.Recognized])
  }
  
  object Type extends _root_.scalapb.GeneratedEnumCompanion[Type] {
    sealed trait Recognized extends Type
    implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Type] = this
    
    @SerialVersionUID(0L)
    case object INT64 extends Type(0) with Type.Recognized {
      val index = 0
      val name = "INT64"
      override def isInt64: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object DOUBLE extends Type(1) with Type.Recognized {
      val index = 1
      val name = "DOUBLE"
      override def isDouble: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object STRING extends Type(2) with Type.Recognized {
      val index = 2
      val name = "STRING"
      override def isString: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object LIST_DOUBLE extends Type(3) with Type.Recognized {
      val index = 3
      val name = "LIST_DOUBLE"
      override def isListDouble: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object LIST_STRING extends Type(4) with Type.Recognized {
      val index = 4
      val name = "LIST_STRING"
      override def isListString: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    final case class Unrecognized(unrecognizedValue: _root_.scala.Int) extends Type(unrecognizedValue) with _root_.scalapb.UnrecognizedEnum
    lazy val values = scala.collection.immutable.Seq(INT64, DOUBLE, STRING, LIST_DOUBLE, LIST_STRING)
    def fromValue(__value: _root_.scala.Int): Type = __value match {
      case 0 => INT64
      case 1 => DOUBLE
      case 2 => STRING
      case 3 => LIST_DOUBLE
      case 4 => LIST_STRING
      case __other => Unrecognized(__other)
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = services.Element.javaDescriptor.getEnumTypes().get(0)
    def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = services.Element.scalaDescriptor.enums(0)
  }
  implicit class ElementLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, services.Element]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, services.Element](_l) {
    def key: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.key)((c_, f_) => c_.copy(key = f_))
    def `type`: _root_.scalapb.lenses.Lens[UpperPB, services.Element.Type] = field(_.`type`)((c_, f_) => c_.copy(`type` = f_))
  }
  final val KEY_FIELD_NUMBER = 1
  final val TYPE_FIELD_NUMBER = 2
  def of(
    key: _root_.scala.Predef.String,
    `type`: services.Element.Type
  ): _root_.services.Element = _root_.services.Element(
    key,
    `type`
  )
  // @@protoc_insertion_point(GeneratedMessageCompanion[services.Element])
}
