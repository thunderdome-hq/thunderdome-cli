User:
{{ template "user" .User }}

Ticket:
{{- if .Ticket.Id }}
{{ template "ticket" .Ticket }}
{{ else }}
*None*
{{ end -}}
