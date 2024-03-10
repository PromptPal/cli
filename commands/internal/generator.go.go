package internal

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed go.tpl
var goTpl string

func GenerateGoTypes(schema []PromptSchema, cfg *ConfigurationOutputGo) ([]byte, error) {
	if len(schema) == 0 {
		return []byte{}, nil
	}
	t := template.Must(template.New("go").Funcs(template.FuncMap{
		"capitalize": capitalizeFunc,
	}).Parse(goTpl))
	var result bytes.Buffer
	t.Execute(&result, templateStruct{
		PackageName: cfg.PackageName,
		Prefix:      cfg.Prefix,
		Prompts:     normalizeSchema(schema),
	})
	return result.Bytes(), nil
}
