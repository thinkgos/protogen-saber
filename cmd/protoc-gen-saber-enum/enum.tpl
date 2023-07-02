// Code generated by protoc-gen-saber-enum. DO NOT EDIT.
// versions:
//   - protoc-gen-saber-enum {{.Version}}
//   - protoc             {{.ProtocVersion}}
{{- if .IsDeprecated}}
// {{.Source}} is a deprecated file.
{{- else}}
// source: {{.Source}}
{{- end}}

package {{.Package}}

{{- range $e := .Enums}}
{{$enumName := $e.Name}}
{{if $e.MessageName}}
{{$enumName = printf "%s_%s" $e.MessageName $e.Name}}
{{end}}

// Enum value mapping for {{$enumName}}.
var (
	__{{$enumName}}Mapping_Desc = map[{{$enumName}}]string{
	{{- range $ee := $e.Values}}
		{{$ee.Number}}: "{{$ee.Mapping}}",
	{{- end}}
	}
	__{{$enumName}}Mapping_Value = map[string]{{$enumName}}{
	{{- range $ee := $e.Values}}
		"{{$ee.Mapping}}": {{$ee.Number}},
	{{- end}}
	}
)

// MappingDescriptor mapping description.
// {{$e.Comment}}
func (t {{$enumName}}) MappingDescriptor() string {
	return __{{$enumName}}Mapping_Desc[t]
}

// Get{{$enumName}}Desc mapping description.
// Deprecated: Use {{$enumName}}.MappingDescriptor instead.
func Get{{$enumName}}Desc(t {{$enumName}}) string {
	return t.MappingDescriptor()
}
// Get{{$enumName}}Value get mapping value
// {{$e.Comment}}
func Get{{$enumName}}Value(s string) int {
	return int(__{{$enumName}}Mapping_Value[s])
}
{{- end}}

