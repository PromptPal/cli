package internal

import "unicode"

func convertString(str string) string {
	var result string
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || char == '_' {
			result += string(char)
		} else {
			result += "_"
		}
	}
	return result
}

func normalizeSchema(schema []PromptSchema) []PromptSchema {
	for i := 0; i < len(schema); i++ {
		s := &schema[i]
		s.Name = convertString(s.Name)
		for j := 0; j < len(s.Variables); j++ {
			s.Variables[j].Name = convertString(s.Variables[j].Name)
		}
	}
	return schema
}
