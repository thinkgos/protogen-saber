// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0
// source: nested.proto

package enums

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

// Status 状态值
// #[enum]
type Nested_Status int32

const (
	// "unspecified"
	// aaaa
	Nested_Status_Unspecified Nested_Status = 0
	// nested1
	Nested_Status_Up Nested_Status = 1
	// nested2
	Nested_Status_Down Nested_Status = 2
	// nested3
	Nested_Status_Left Nested_Status = 3
	// nested4
	Nested_Status_Right Nested_Status = 4
	// end
	Nested_Status_End Nested_Status = 999
)

// Enum value maps for Nested_Status.
var (
	Nested_Status_name = map[int32]string{
		0:   "Status_Unspecified",
		1:   "Status_Up",
		2:   "Status_Down",
		3:   "Status_Left",
		4:   "Status_Right",
		999: "Status_End",
	}
	Nested_Status_value = map[string]int32{
		"Status_Unspecified": 0,
		"Status_Up":          1,
		"Status_Down":        2,
		"Status_Left":        3,
		"Status_Right":       4,
		"Status_End":         999,
	}
)

func (x Nested_Status) Enum() *Nested_Status {
	p := new(Nested_Status)
	*p = x
	return p
}

func (x Nested_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Nested_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_nested_proto_enumTypes[0].Descriptor()
}

func (Nested_Status) Type() protoreflect.EnumType {
	return &file_nested_proto_enumTypes[0]
}

func (x Nested_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Nested_Status.Descriptor instead.
func (Nested_Status) EnumDescriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{0, 0}
}

// Type 类型
// #[enum]
type Nested_Nested1_Type int32

const (
	// 禁用
	Nested_Nested1_Type_Disable Nested_Nested1_Type = 0
	// 启用
	Nested_Nested1_Type_Enable Nested_Nested1_Type = 1
)

// Enum value maps for Nested_Nested1_Type.
var (
	Nested_Nested1_Type_name = map[int32]string{
		0: "Type_Disable",
		1: "Type_Enable",
	}
	Nested_Nested1_Type_value = map[string]int32{
		"Type_Disable": 0,
		"Type_Enable":  1,
	}
)

func (x Nested_Nested1_Type) Enum() *Nested_Nested1_Type {
	p := new(Nested_Nested1_Type)
	*p = x
	return p
}

func (x Nested_Nested1_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Nested_Nested1_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_nested_proto_enumTypes[1].Descriptor()
}

func (Nested_Nested1_Type) Type() protoreflect.EnumType {
	return &file_nested_proto_enumTypes[1]
}

func (x Nested_Nested1_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Nested_Nested1_Type.Descriptor instead.
func (Nested_Nested1_Type) EnumDescriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{0}
}

type Nested_Nested1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nested_Nested1) Reset() {
	*x = Nested_Nested1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested_Nested1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested_Nested1) ProtoMessage() {}

func (x *Nested_Nested1) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested_Nested1.ProtoReflect.Descriptor instead.
func (*Nested_Nested1) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{0, 0}
}

var File_nested_proto protoreflect.FileDescriptor

var file_nested_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x5f, 0x73, 0x61, 0x62, 0x65, 0x72, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xb4, 0x01,
	0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x34, 0x0a, 0x07, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x31, 0x22, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x22, 0x74,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x55, 0x70, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x44, 0x6f, 0x77, 0x6e, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x4c, 0x65, 0x66, 0x74, 0x10,
	0x03, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x45, 0x6e,
	0x64, 0x10, 0xe7, 0x07, 0x42, 0x50, 0x0a, 0x1a, 0x63, 0x6e, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2d, 0x67, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x42, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nested_proto_rawDescOnce sync.Once
	file_nested_proto_rawDescData = file_nested_proto_rawDesc
)

func file_nested_proto_rawDescGZIP() []byte {
	file_nested_proto_rawDescOnce.Do(func() {
		file_nested_proto_rawDescData = protoimpl.X.CompressGZIP(file_nested_proto_rawDescData)
	})
	return file_nested_proto_rawDescData
}

var file_nested_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nested_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nested_proto_goTypes = []interface{}{
	(Nested_Status)(0),       // 0: protogen_saber.examples.enums.Nested.Status
	(Nested_Nested1_Type)(0), // 1: protogen_saber.examples.enums.Nested.Nested1.Type
	(*Nested)(nil),           // 2: protogen_saber.examples.enums.Nested
	(*Nested_Nested1)(nil),   // 3: protogen_saber.examples.enums.Nested.Nested1
}
var file_nested_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nested_proto_init() }
func file_nested_proto_init() {
	if File_nested_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nested_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested); i {
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
		file_nested_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested_Nested1); i {
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
			RawDescriptor: file_nested_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nested_proto_goTypes,
		DependencyIndexes: file_nested_proto_depIdxs,
		EnumInfos:         file_nested_proto_enumTypes,
		MessageInfos:      file_nested_proto_msgTypes,
	}.Build()
	File_nested_proto = out.File
	file_nested_proto_rawDesc = nil
	file_nested_proto_goTypes = nil
	file_nested_proto_depIdxs = nil
}
