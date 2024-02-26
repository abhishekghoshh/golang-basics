// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: street.proto

package import_pkg

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

type Street struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City *City  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *Street) Reset() {
	*x = Street{}
	if protoimpl.UnsafeEnabled {
		mi := &file_street_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Street) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Street) ProtoMessage() {}

func (x *Street) ProtoReflect() protoreflect.Message {
	mi := &file_street_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Street.ProtoReflect.Descriptor instead.
func (*Street) Descriptor() ([]byte, []int) {
	return file_street_proto_rawDescGZIP(), []int{0}
}

func (x *Street) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Street) GetCity() *City {
	if x != nil {
		return x.City
	}
	return nil
}

var File_street_proto protoreflect.FileDescriptor

var file_street_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x1a, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x42, 0x17, 0x5a, 0x15, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_street_proto_rawDescOnce sync.Once
	file_street_proto_rawDescData = file_street_proto_rawDesc
)

func file_street_proto_rawDescGZIP() []byte {
	file_street_proto_rawDescOnce.Do(func() {
		file_street_proto_rawDescData = protoimpl.X.CompressGZIP(file_street_proto_rawDescData)
	})
	return file_street_proto_rawDescData
}

var file_street_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_street_proto_goTypes = []interface{}{
	(*Street)(nil), // 0: street.Street
	(*City)(nil),   // 1: city.City
}
var file_street_proto_depIdxs = []int32{
	1, // 0: street.Street.city:type_name -> city.City
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_street_proto_init() }
func file_street_proto_init() {
	if File_street_proto != nil {
		return
	}
	file_city_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_street_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Street); i {
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
			RawDescriptor: file_street_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_street_proto_goTypes,
		DependencyIndexes: file_street_proto_depIdxs,
		MessageInfos:      file_street_proto_msgTypes,
	}.Build()
	File_street_proto = out.File
	file_street_proto_rawDesc = nil
	file_street_proto_goTypes = nil
	file_street_proto_depIdxs = nil
}
