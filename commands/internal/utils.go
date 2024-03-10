package internal

import "unicode"

// TODO: need sync this logic to website to make sure it's correct.
func convertString(str string) string {
	var result string
	isPreviousUnderscore := false
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			if isPreviousUnderscore {
				result += string(unicode.ToUpper(char))
				isPreviousUnderscore = false
			} else {
				result += string(char)
			}
		} else {
			isPreviousUnderscore = true
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
