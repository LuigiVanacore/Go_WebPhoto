package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir   string = "static/layouts/"
	TemplateExt string = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	var templates []string
	templates = append(templates, "static/index.gohtml")
	templates = append(templates, layoutFiles()...)
	templates = append(templates, files...)
	t, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
