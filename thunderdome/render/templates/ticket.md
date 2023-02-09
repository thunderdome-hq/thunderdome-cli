{{ define "ticket" }}
# {{ .Name }}


- **ID**:           {{ .Id }}
- **Due Date**:     {{ or .DueDate "*None*" }}
- **Claimed by**:   {{ or .User "*None*" }}
- **Repository**:   {{ or .Repository "*None*" }}
- **Pull Request**: {{ or .PullRequest "*None*" }}

## Description

{{ or .Description "*None*" }}

## Contacts

{{ range .Contacts }}
- {{ . }}
{{ else }}
*None*
{{ end }}
{{ end }}
