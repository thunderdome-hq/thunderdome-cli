User:
{{- if .User}}
{{ template "user" .User }}
*None*
{{ else }}
{{ end -}}

Ticket:
{{- if .Ticket }}
{{ template "ticket" .Ticket }}
{{ else }}
*None*
{{ end -}}
