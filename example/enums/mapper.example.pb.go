// Code generated by protoc-gen-saber-enum. DO NOT EDIT.
// versions:
//   - protoc-gen-saber-enum v0.0.6
//   - protoc             v4.24.0
// source: nested.proto,non_nested.proto

package enums

func EnumComment() []map[int32]string {
	return []map[int32]string{
		{
			int32(Nested_Status_Unspecified): "\"unspecified\", aaaa",
			int32(Nested_Status_Up):          "nested1",
			int32(Nested_Status_Down):        "nested2",
			int32(Nested_Status_Left):        "nested3",
			int32(Nested_Status_Right):       "nested4",
			int32(Nested_Status_End):         "end",
		},
		{
			int32(Nested_Nested1_Type_Disable): "禁用",
			int32(Nested_Nested1_Type_Enable):  "启用",
		},
		{
			int32(NonNestedStatus_NonNestedStatus_Unspecified): "未定义",
			int32(NonNestedStatus_NonNestedStatus_Up):          "打开",
			int32(NonNestedStatus_NonNestedStatus_Down):        "关闭",
			int32(NonNestedStatus_NonNestedStatus_Left):        "左",
			int32(NonNestedStatus_NonNestedStatus_Right):       "右",
		},
	}
}
