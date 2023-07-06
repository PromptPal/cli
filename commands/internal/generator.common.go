package internal

import "strings"

type templateStruct struct {
	PackageName string
	Prefix      string
	Prompts     []PromptSchema
}

func capitalizeFunc(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
