package render

import (
	"bytes"
	"embed"
	"github.com/charmbracelet/glamour"
	log "github.com/sirupsen/logrus"
	"github.com/thunderdome-hq/thunderdome-api/api"
	"google.golang.org/grpc/codes"
	"path/filepath"

	"text/template"
)

const templateDir = "templates"

// Data is the embedded data directory.
//
//go:embed templates/* style.json
var Data embed.FS

func Markdown(data any, templates ...string) (string, error) {
	// Prepend mkTemplate path
	for i, t := range templates {
		templates[i] = filepath.Join(templateDir, t)
	}

	mkTemplate, err := template.ParseFS(Data, templates...)
	if err != nil {
		log.Debugln("Could not parse templates:", err)
		return "", api.Error(codes.Internal, api.CLIError, "unable to parse markdown templates %v", templates)
	}

	var buffer bytes.Buffer
	err = mkTemplate.Execute(&buffer, data)
	if err != nil {
		log.Debugln("Could not render template:", err)
		return "", api.Error(codes.Internal, api.CLIError, "unable to render markdown template")
	}

	style, err := Data.ReadFile("style.json")
	if err != nil {
		log.Debugln("Could not read style:", err)
		return "", api.Error(codes.Internal, api.CLIError, "unable to read markdown styles")
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStylesFromJSONBytes(style),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		log.Debugln("Could create renderer:", err)
		return "", api.Error(codes.Internal, api.CLIError, "unable to create markdown renderer")
	}

	out, err := renderer.Render(buffer.String())
	if err != nil {
		log.Debugln("Could render markdown:", err)
		return out, api.Error(codes.Internal, api.CLIError, "unable to render markdown")
	}

	return out, nil
}
