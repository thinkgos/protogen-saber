// Code generated by protoc-gen-saber-enum. DO NOT EDIT.
// versions:
//   - protoc-gen-saber-enum v0.0.3
//   - protoc             v3.21.2
// source: nested.proto,non_nested.proto

package enums

func EnumComment() []map[int32]string {
	return []map[int32]string{
		{
			int32(Nested_STATUS_UNSPECIFIED): "\"unspecified\", aaaa",
			int32(Nested_STATUS_UP):          "nested1",
			int32(Nested_STATUS_DOWN):        "nested2",
			int32(Nested_STATUS_LEFT):        "nested3",
			int32(Nested_STATUS_RIGHT):       "nested4",
		},
		{
			int32(Nested_Nested1_TYPE_DISABLE): "禁用",
			int32(Nested_Nested1_TYPE_Enable):  "启用",
		},
		{
			int32(NonNestedStatus_NON_NESTED_STATUS_UNSPECIFIED): "未定义",
			int32(NonNestedStatus_NON_NESTED_STATUS_UP):          "打开",
			int32(NonNestedStatus_NON_NESTED_STATUS_DOWN):        "关闭",
			int32(NonNestedStatus_NON_NESTED_STATUS_LEFT):        "左",
			int32(NonNestedStatus_NON_NESTED_STATUS_RIGHT):       "右",
		},
	}
}
