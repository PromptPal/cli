package internal

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed ts.tpl
var tsTpl string

func GenerateTypeScriptTypes(schema []PromptSchema, cfg *ConfigurationOutputTS) ([]byte, error) {
	if len(schema) == 0 {
		return []byte{}, nil
	}
	t := template.Must(template.New("ts").Funcs(template.FuncMap{
		"capitalize": capitalizeFunc,
	}).Parse(tsTpl))
	var result bytes.Buffer
	t.Execute(&result, templateStruct{
		PackageName: "",
		Prefix:      cfg.Prefix,
		Prompts:     schema,
	})
	return result.Bytes(), nil
}
