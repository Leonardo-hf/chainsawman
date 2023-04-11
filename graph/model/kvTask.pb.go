// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: kvTask.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KVTask_Status int32

const (
	KVTask_New      KVTask_Status = 0
	KVTask_Finished KVTask_Status = 1
)

// Enum value maps for KVTask_Status.
var (
	KVTask_Status_name = map[int32]string{
		0: "New",
		1: "Finished",
	}
	KVTask_Status_value = map[string]int32{
		"New":      0,
		"Finished": 1,
	}
)

func (x KVTask_Status) Enum() *KVTask_Status {
	p := new(KVTask_Status)
	*p = x
	return p
}

func (x KVTask_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KVTask_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_kvTask_proto_enumTypes[0].Descriptor()
}

func (KVTask_Status) Type() protoreflect.EnumType {
	return &file_kvTask_proto_enumTypes[0]
}

func (x KVTask_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KVTask_Status.Descriptor instead.
func (KVTask_Status) EnumDescriptor() ([]byte, []int) {
	return file_kvTask_proto_rawDescGZIP(), []int{0, 0}
}

// TODO: 这种文件最好放到一个单独的github仓库统一管理
type KVTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Params     string        `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	Status     KVTask_Status `protobuf:"varint,4,opt,name=status,proto3,enum=model.KVTask_Status" json:"status,omitempty"`
	CreateTime int64         `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime int64         `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Result     string        `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *KVTask) Reset() {
	*x = KVTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvTask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVTask) ProtoMessage() {}

func (x *KVTask) ProtoReflect() protoreflect.Message {
	mi := &file_kvTask_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVTask.ProtoReflect.Descriptor instead.
func (*KVTask) Descriptor() ([]byte, []int) {
	return file_kvTask_proto_rawDescGZIP(), []int{0}
}

func (x *KVTask) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KVTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KVTask) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *KVTask) GetStatus() KVTask_Status {
	if x != nil {
		return x.Status
	}
	return KVTask_New
}

func (x *KVTask) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *KVTask) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *KVTask) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_kvTask_proto protoreflect.FileDescriptor

var file_kvTask_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6b, 0x76, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xeb, 0x01, 0x0a, 0x06, 0x4b, 0x56, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2c, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x56, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03,
	0x4e, 0x65, 0x77, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x10, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kvTask_proto_rawDescOnce sync.Once
	file_kvTask_proto_rawDescData = file_kvTask_proto_rawDesc
)

func file_kvTask_proto_rawDescGZIP() []byte {
	file_kvTask_proto_rawDescOnce.Do(func() {
		file_kvTask_proto_rawDescData = protoimpl.X.CompressGZIP(file_kvTask_proto_rawDescData)
	})
	return file_kvTask_proto_rawDescData
}

var file_kvTask_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kvTask_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kvTask_proto_goTypes = []interface{}{
	(KVTask_Status)(0), // 0: model.KVTask.Status
	(*KVTask)(nil),     // 1: model.KVTask
}
var file_kvTask_proto_depIdxs = []int32{
	0, // 0: model.KVTask.status:type_name -> model.KVTask.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kvTask_proto_init() }
func file_kvTask_proto_init() {
	if File_kvTask_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kvTask_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kvTask_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kvTask_proto_goTypes,
		DependencyIndexes: file_kvTask_proto_depIdxs,
		EnumInfos:         file_kvTask_proto_enumTypes,
		MessageInfos:      file_kvTask_proto_msgTypes,
	}.Build()
	File_kvTask_proto = out.File
	file_kvTask_proto_rawDesc = nil
	file_kvTask_proto_goTypes = nil
	file_kvTask_proto_depIdxs = nil
}
