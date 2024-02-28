// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: account.proto

package enum

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

type AccountType int32

const (
	AccountType_ACCOUNT_TYPE_UNSPECIFIED AccountType = 0
	AccountType_ACCOUNT_TYPE_SAVINGS     AccountType = 1
	AccountType_ACCOUNT_TYPE_CREDIT      AccountType = 2
	AccountType_ACCOUNT_TYPE_BOTH        AccountType = 3
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		0: "ACCOUNT_TYPE_UNSPECIFIED",
		1: "ACCOUNT_TYPE_SAVINGS",
		2: "ACCOUNT_TYPE_CREDIT",
		3: "ACCOUNT_TYPE_BOTH",
	}
	AccountType_value = map[string]int32{
		"ACCOUNT_TYPE_UNSPECIFIED": 0,
		"ACCOUNT_TYPE_SAVINGS":     1,
		"ACCOUNT_TYPE_CREDIT":      2,
		"ACCOUNT_TYPE_BOTH":        3,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_account_proto_enumTypes[0].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_account_proto_enumTypes[0]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Thumbnail  []byte  `protobuf:"bytes,3,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	IsVerified bool    `protobuf:"varint,4,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	Height     float32 `protobuf:"fixed32,5,opt,name=height,proto3" json:"height,omitempty"`
	// repeated element for the array
	Phones     []string    `protobuf:"bytes,6,rep,name=phones,proto3" json:"phones,omitempty"`
	AcountType AccountType `protobuf:"varint,7,opt,name=acountType,proto3,enum=AccountType" json:"acountType,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *Account) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *Account) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Account) GetPhones() []string {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *Account) GetAcountType() AccountType {
	if x != nil {
		return x.AcountType
	}
	return AccountType_ACCOUNT_TYPE_UNSPECIFIED
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xca, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x2c,
	0x0a, 0x0a, 0x61, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x61, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x75, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x41, 0x56, 0x49, 0x4e, 0x47,
	0x53, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x54,
	0x48, 0x10, 0x03, 0x42, 0x1e, 0x5a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_proto_goTypes = []interface{}{
	(AccountType)(0), // 0: AccountType
	(*Account)(nil),  // 1: Account
}
var file_account_proto_depIdxs = []int32{
	0, // 0: Account.acountType:type_name -> AccountType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		EnumInfos:         file_account_proto_enumTypes,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
