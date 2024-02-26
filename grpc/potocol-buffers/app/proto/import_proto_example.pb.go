// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: import_proto_example.proto

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

// this is also an imported mesaage
type Human struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *FullAddress `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Human) Reset() {
	*x = Human{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Human) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Human) ProtoMessage() {}

func (x *Human) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Human.ProtoReflect.Descriptor instead.
func (*Human) Descriptor() ([]byte, []int) {
	return file_import_proto_example_proto_rawDescGZIP(), []int{0}
}

func (x *Human) GetAddress() *FullAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

// this is an imported message which is inside a package
type ImportedPackagedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImportedPackagedMessage *PackagedMessage `protobuf:"bytes,1,opt,name=imported_packaged_message,json=importedPackagedMessage,proto3" json:"imported_packaged_message,omitempty"`
}

func (x *ImportedPackagedMessage) Reset() {
	*x = ImportedPackagedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportedPackagedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportedPackagedMessage) ProtoMessage() {}

func (x *ImportedPackagedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportedPackagedMessage.ProtoReflect.Descriptor instead.
func (*ImportedPackagedMessage) Descriptor() ([]byte, []int) {
	return file_import_proto_example_proto_rawDescGZIP(), []int{1}
}

func (x *ImportedPackagedMessage) GetImportedPackagedMessage() *PackagedMessage {
	if x != nil {
		return x.ImportedPackagedMessage
	}
	return nil
}

var File_import_proto_example_proto protoreflect.FileDescriptor

var file_import_proto_example_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x05, 0x48, 0x75, 0x6d, 0x61, 0x6e,
	0x12, 0x26, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6f, 0x0a, 0x17, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x17, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_import_proto_example_proto_rawDescOnce sync.Once
	file_import_proto_example_proto_rawDescData = file_import_proto_example_proto_rawDesc
)

func file_import_proto_example_proto_rawDescGZIP() []byte {
	file_import_proto_example_proto_rawDescOnce.Do(func() {
		file_import_proto_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_import_proto_example_proto_rawDescData)
	})
	return file_import_proto_example_proto_rawDescData
}

var file_import_proto_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_import_proto_example_proto_goTypes = []interface{}{
	(*Human)(nil),                   // 0: Human
	(*ImportedPackagedMessage)(nil), // 1: ImportedPackagedMessage
	(*FullAddress)(nil),             // 2: FullAddress
	(*PackagedMessage)(nil),         // 3: example.PackagedMessage
}
var file_import_proto_example_proto_depIdxs = []int32{
	2, // 0: Human.address:type_name -> FullAddress
	3, // 1: ImportedPackagedMessage.imported_packaged_message:type_name -> example.PackagedMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_import_proto_example_proto_init() }
func file_import_proto_example_proto_init() {
	if File_import_proto_example_proto != nil {
		return
	}
	file_full_address_proto_init()
	file_package_example_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_import_proto_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Human); i {
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
		file_import_proto_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportedPackagedMessage); i {
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
			RawDescriptor: file_import_proto_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_import_proto_example_proto_goTypes,
		DependencyIndexes: file_import_proto_example_proto_depIdxs,
		MessageInfos:      file_import_proto_example_proto_msgTypes,
	}.Build()
	File_import_proto_example_proto = out.File
	file_import_proto_example_proto_rawDesc = nil
	file_import_proto_example_proto_goTypes = nil
	file_import_proto_example_proto_depIdxs = nil
}
