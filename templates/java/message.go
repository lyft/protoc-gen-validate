package java

const messageTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{- if $r.GetSkip }}
			// skipping validation for {{ $f.Name }}
	{{- else -}}
		{{- template "required" . }}
		{{- if (isOfMessageType $f) }}
			// Validate {{ $f.Name }}
			com.lyft.pgv.GeneratedValidatorIndex.validatorFor({{ accessor . }}).assertValid({{ accessor . }});
		{{- end -}}
	{{- end -}}
`
