// Code generated by protoc-gen-saber-enum. DO NOT EDIT.
// versions:
//   - protoc-gen-saber-enum v0.0.2
//   - protoc             v3.21.2
// source: nested.proto

package enums

// __Nested_StatusMapping Nested_Status mapping
var __Nested_StatusMapping = map[Nested_Status]string{
	0: "未定义",
	1: "nested1",
	2: "nested2",
	3: "nested3",
	4: "nested4",
}

// GetNested_StatusDesc get mapping description
//
//	Status 状态值, [0:未定义,1:nested1,2:nested2,3:nested3,4:nested4]
func GetNested_StatusDesc(t Nested_Status) string {
	return __Nested_StatusMapping[t]
}

// __Nested_Nested1_TypeMapping Nested_Nested1_Type mapping
var __Nested_Nested1_TypeMapping = map[Nested_Nested1_Type]string{
	0: "禁用",
	1: "启用",
}

// GetNested_Nested1_TypeDesc get mapping description
// [0:禁用,1:启用]
func GetNested_Nested1_TypeDesc(t Nested_Nested1_Type) string {
	return __Nested_Nested1_TypeMapping[t]
}
