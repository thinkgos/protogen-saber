// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0
// source: non_nested.proto

package enums

import (
	_ "github.com/things-go/protogen-saber/protosaber/enumerate"
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

// NonNestedStatus 状态值
type NonNestedStatus int32

const (
	// 未定义
	NonNestedStatus_NON_NESTED_STATUS_UNSPECIFIED NonNestedStatus = 0
	// 打开
	NonNestedStatus_NON_NESTED_STATUS_UP NonNestedStatus = 1
	// 关闭
	NonNestedStatus_NON_NESTED_STATUS_DOWN NonNestedStatus = 2
	// 左
	NonNestedStatus_NON_NESTED_STATUS_LEFT NonNestedStatus = 3
	// 右
	NonNestedStatus_NON_NESTED_STATUS_RIGHT NonNestedStatus = 4
)

// Enum value maps for NonNestedStatus.
var (
	NonNestedStatus_name = map[int32]string{
		0: "NON_NESTED_STATUS_UNSPECIFIED",
		1: "NON_NESTED_STATUS_UP",
		2: "NON_NESTED_STATUS_DOWN",
		3: "NON_NESTED_STATUS_LEFT",
		4: "NON_NESTED_STATUS_RIGHT",
	}
	NonNestedStatus_value = map[string]int32{
		"NON_NESTED_STATUS_UNSPECIFIED": 0,
		"NON_NESTED_STATUS_UP":          1,
		"NON_NESTED_STATUS_DOWN":        2,
		"NON_NESTED_STATUS_LEFT":        3,
		"NON_NESTED_STATUS_RIGHT":       4,
	}
)

func (x NonNestedStatus) Enum() *NonNestedStatus {
	p := new(NonNestedStatus)
	*p = x
	return p
}

func (x NonNestedStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NonNestedStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_non_nested_proto_enumTypes[0].Descriptor()
}

func (NonNestedStatus) Type() protoreflect.EnumType {
	return &file_non_nested_proto_enumTypes[0]
}

func (x NonNestedStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NonNestedStatus.Descriptor instead.
func (NonNestedStatus) EnumDescriptor() ([]byte, []int) {
	return file_non_nested_proto_rawDescGZIP(), []int{0}
}

var File_non_nested_proto protoreflect.FileDescriptor

var file_non_nested_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x5f, 0x73, 0x61, 0x62,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x19, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xdd, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x6e, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x4e, 0x4f, 0x4e, 0x5f, 0x4e,
	0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x14, 0x4e, 0x4f,
	0x4e, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x50, 0x10, 0x01, 0x1a, 0x0c, 0xea, 0x8b, 0xb7, 0xeb, 0x02, 0x06, 0xe6, 0x89, 0x93, 0xe5,
	0xbc, 0x80, 0x12, 0x28, 0x0a, 0x16, 0x4e, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x1a, 0x0c,
	0xea, 0x8b, 0xb7, 0xeb, 0x02, 0x06, 0xe5, 0x85, 0xb3, 0xe9, 0x97, 0xad, 0x12, 0x25, 0x0a, 0x16,
	0x4e, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x03, 0x1a, 0x09, 0xea, 0x8b, 0xb7, 0xeb, 0x02, 0x03,
	0xe5, 0xb7, 0xa6, 0x12, 0x26, 0x0a, 0x17, 0x4e, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45,
	0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x04,
	0x1a, 0x09, 0xea, 0x8b, 0xb7, 0xeb, 0x02, 0x03, 0xe5, 0x8f, 0xb3, 0x1a, 0x06, 0xc8, 0x85, 0xb7,
	0xeb, 0x02, 0x01, 0x42, 0x50, 0x0a, 0x1a, 0x63, 0x6e, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2d, 0x67, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x42, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2d, 0x67, 0x6f,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_non_nested_proto_rawDescOnce sync.Once
	file_non_nested_proto_rawDescData = file_non_nested_proto_rawDesc
)

func file_non_nested_proto_rawDescGZIP() []byte {
	file_non_nested_proto_rawDescOnce.Do(func() {
		file_non_nested_proto_rawDescData = protoimpl.X.CompressGZIP(file_non_nested_proto_rawDescData)
	})
	return file_non_nested_proto_rawDescData
}

var file_non_nested_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_non_nested_proto_goTypes = []interface{}{
	(NonNestedStatus)(0), // 0: protogen_saber.enums.NonNestedStatus
}
var file_non_nested_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_non_nested_proto_init() }
func file_non_nested_proto_init() {
	if File_non_nested_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_non_nested_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_non_nested_proto_goTypes,
		DependencyIndexes: file_non_nested_proto_depIdxs,
		EnumInfos:         file_non_nested_proto_enumTypes,
	}.Build()
	File_non_nested_proto = out.File
	file_non_nested_proto_rawDesc = nil
	file_non_nested_proto_goTypes = nil
	file_non_nested_proto_depIdxs = nil
}
