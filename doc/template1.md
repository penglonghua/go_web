
# readme

模版， 根元素 `{{.}}`


# Control Structures

{{ /* a comment */ }}
{{ . }}  根元素
{{ .Title }} 渲染 Title元素
{{ if .Done }} {{ else }} {{ end }}
{{ range .TODOs }} {{ . }}  {{ end}}
{{ block "content" .}} {{ end}}  定义一个 content ?



