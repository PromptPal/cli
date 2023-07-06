export enum {{ .Prefix }}Prompts {
	{{ range .Prompts }}
	{{ capitalize .Name }} = "{{ .HashID }}",
	{{ end }}
}

{{ range .Prompts }}
export type {{ $.Prefix }}Prompt{{ capitalize .Name}}Variables = {
	{{ range $v := .Variables }}
	{{ $v.Name }}: string;
	{{ end }}
}
{{ end }}