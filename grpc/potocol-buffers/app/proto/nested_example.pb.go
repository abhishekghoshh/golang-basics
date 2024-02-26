// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: nested_example.proto

package proto

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

type OuterMessage_InnerEnum int32

const (
	OuterMessage_INNER_TYPE_UNSPECIFIED OuterMessage_InnerEnum = 0
	OuterMessage_INNER_TYPE_FIRST       OuterMessage_InnerEnum = 1
	OuterMessage_INNER_TYPE_SECOND      OuterMessage_InnerEnum = 2
)

// Enum value maps for OuterMessage_InnerEnum.
var (
	OuterMessage_InnerEnum_name = map[int32]string{
		0: "INNER_TYPE_UNSPECIFIED",
		1: "INNER_TYPE_FIRST",
		2: "INNER_TYPE_SECOND",
	}
	OuterMessage_InnerEnum_value = map[string]int32{
		"INNER_TYPE_UNSPECIFIED": 0,
		"INNER_TYPE_FIRST":       1,
		"INNER_TYPE_SECOND":      2,
	}
)

func (x OuterMessage_InnerEnum) Enum() *OuterMessage_InnerEnum {
	p := new(OuterMessage_InnerEnum)
	*p = x
	return p
}

func (x OuterMessage_InnerEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OuterMessage_InnerEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_nested_example_proto_enumTypes[0].Descriptor()
}

func (OuterMessage_InnerEnum) Type() protoreflect.EnumType {
	return &file_nested_example_proto_enumTypes[0]
}

func (x OuterMessage_InnerEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OuterMessage_InnerEnum.Descriptor instead.
func (OuterMessage_InnerEnum) EnumDescriptor() ([]byte, []int) {
	return file_nested_example_proto_rawDescGZIP(), []int{0, 0}
}

type OuterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Student *OuterMessage_InnerMessage `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	Type    OuterMessage_InnerEnum     `protobuf:"varint,2,opt,name=type,proto3,enum=OuterMessage_InnerEnum" json:"type,omitempty"`
}

func (x *OuterMessage) Reset() {
	*x = OuterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage) ProtoMessage() {}

func (x *OuterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nested_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return file_nested_example_proto_rawDescGZIP(), []int{0}
}

func (x *OuterMessage) GetStudent() *OuterMessage_InnerMessage {
	if x != nil {
		return x.Student
	}
	return nil
}

func (x *OuterMessage) GetType() OuterMessage_InnerEnum {
	if x != nil {
		return x.Type
	}
	return OuterMessage_INNER_TYPE_UNSPECIFIED
}

type OuterMessage_InnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OuterMessage_InnerMessage) Reset() {
	*x = OuterMessage_InnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage_InnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage_InnerMessage) ProtoMessage() {}

func (x *OuterMessage_InnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nested_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage_InnerMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage_InnerMessage) Descriptor() ([]byte, []int) {
	return file_nested_example_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OuterMessage_InnerMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_nested_example_proto protoreflect.FileDescriptor

var file_nested_example_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x4f, 0x75,
	0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x22, 0x0a, 0x0c, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54,
	0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x16, 0x49,
	0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x4e, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nested_example_proto_rawDescOnce sync.Once
	file_nested_example_proto_rawDescData = file_nested_example_proto_rawDesc
)

func file_nested_example_proto_rawDescGZIP() []byte {
	file_nested_example_proto_rawDescOnce.Do(func() {
		file_nested_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_nested_example_proto_rawDescData)
	})
	return file_nested_example_proto_rawDescData
}

var file_nested_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nested_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nested_example_proto_goTypes = []interface{}{
	(OuterMessage_InnerEnum)(0),       // 0: OuterMessage.InnerEnum
	(*OuterMessage)(nil),              // 1: OuterMessage
	(*OuterMessage_InnerMessage)(nil), // 2: OuterMessage.InnerMessage
}
var file_nested_example_proto_depIdxs = []int32{
	2, // 0: OuterMessage.student:type_name -> OuterMessage.InnerMessage
	0, // 1: OuterMessage.type:type_name -> OuterMessage.InnerEnum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nested_example_proto_init() }
func file_nested_example_proto_init() {
	if File_nested_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nested_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage); i {
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
		file_nested_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage_InnerMessage); i {
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
			RawDescriptor: file_nested_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nested_example_proto_goTypes,
		DependencyIndexes: file_nested_example_proto_depIdxs,
		EnumInfos:         file_nested_example_proto_enumTypes,
		MessageInfos:      file_nested_example_proto_msgTypes,
	}.Build()
	File_nested_example_proto = out.File
	file_nested_example_proto_rawDesc = nil
	file_nested_example_proto_goTypes = nil
	file_nested_example_proto_depIdxs = nil
}
