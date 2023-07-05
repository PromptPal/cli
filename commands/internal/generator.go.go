package internal

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed go.tpl
var tpl string

type templateGoStruct struct {
	PackageName string
	Prefix      string
	Prompts     []PromptSchema
}

func capitalizeFunc(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func GenerateGoTypes(schema []PromptSchema, cfg *ConfigurationOutputGo) ([]byte, error) {
	t := template.Must(template.New("go").Funcs(template.FuncMap{
		"capitalize": capitalizeFunc,
	}).Parse(tpl))
	var result bytes.Buffer
	t.Execute(&result, templateGoStruct{
		PackageName: cfg.PackageName,
		Prefix:      cfg.Prefix,
		Prompts:     schema,
	})
	return result.Bytes(), nil
}
