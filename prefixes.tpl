{{- range $val := . -}}
variable "f5xc_ip_ranges_{{$val.Geography}}_{{$val.Protocol}}" {
    type = list(string)
    default = [{{range $prefix := $val.Prefixes}}"{{$prefix}}",{{end}}]
}
{{end}}