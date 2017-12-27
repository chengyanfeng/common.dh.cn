{{range .Tables}}
### {{.TableComment}}——`{{.TableName}}` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|{{range .Fields}}
|{{.Name}}|{{.Type}}|{{.Null}}|{{.Comment}}|{{end}}
{{end}}