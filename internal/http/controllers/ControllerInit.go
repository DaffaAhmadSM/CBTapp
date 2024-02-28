package controllers

import (
	"fmt"
	"gorm.io/gorm"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Template Parse HTML templates
type Template struct {
	templates *template.Template
}

func findAndParseTemplates(rootDir string, funcMap template.FuncMap) *template.Template {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, "html") {
			if e1 != nil {
				return e1
			}

			b, e2 := os.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := strings.ReplaceAll(path[pfx:], "\\", "/")
			t := root.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})
	fmt.Printf("defined: %v\n", root.DefinedTemplates())
	if err != nil {
		panic(err)
	}
	return root
}
func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var templ = &Template{
	templates: findAndParseTemplates("./web/views/", nil),
}

type MainController struct {
	Database *gorm.DB
}
