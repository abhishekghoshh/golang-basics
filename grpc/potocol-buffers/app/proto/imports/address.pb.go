// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: address.proto

package imports

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City     *City     `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Street   *Street   `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	Building *Building `protobuf:"bytes,3,opt,name=building,proto3" json:"building,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetCity() *City {
	if x != nil {
		return x.City
	}
	return nil
}

func (x *Address) GetStreet() *Street {
	if x != nil {
		return x.Street
	}
	return nil
}

func (x *Address) GetBuilding() *Building {
	if x != nil {
		return x.Building
	}
	return nil
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x1f, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x25, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_address_proto_goTypes = []interface{}{
	(*Address)(nil),  // 0: Address
	(*City)(nil),     // 1: City
	(*Street)(nil),   // 2: Street
	(*Building)(nil), // 3: Building
}
var file_address_proto_depIdxs = []int32{
	1, // 0: Address.city:type_name -> City
	2, // 1: Address.street:type_name -> Street
	3, // 2: Address.building:type_name -> Building
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	file_building_proto_init()
	file_city_proto_init()
	file_street_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}
