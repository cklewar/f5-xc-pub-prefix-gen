{{- range $val := . -}}
variable "{{$val.Geography}}_{{$val.Protocol}}" {
    type = list(string)
    default = [{{range $prefix := $val.Prefixes}}"{{$prefix}}",{{end}}]
}
{{end}}