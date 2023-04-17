{{- define "user" }}
- **Email**: {{ .Credentials.Email }}
- **Github**: {{ .Credentials.Github }}
- **Status**: {{ .Status }}
{{- end }}
