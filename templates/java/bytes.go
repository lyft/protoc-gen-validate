package java

const bytesTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Len }}
			com.lyft.pgv.BytesValidation.length("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinLen }}
			com.lyft.pgv.BytesValidation.minLength("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxLen }}
			com.lyft.pgv.BytesValidation.maxLength("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.Prefix }}
			com.lyft.pgv.BytesValidation.prefix("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetPrefix }}");
{{- end -}}
{{- if $r.Contains }}
			com.lyft.pgv.BytesValidation.contains("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetContains }}");
{{- end -}}
{{- if $r.Suffix }}
			com.lyft.pgv.BytesValidation.suffix("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetSuffix }}");
{{- end -}}
{{- if $r.GetIp }}
			com.lyft.pgv.BytesValidation.ip("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }});
{{- end -}}
{{- if $r.GetIpv4 }}
			com.lyft.pgv.BytesValidation.ipv4("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }});
{{- end -}}
{{- if $r.GetIpv6 }}
			com.lyft.pgv.BytesValidation.ipv6("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }});
{{- end -}}
`