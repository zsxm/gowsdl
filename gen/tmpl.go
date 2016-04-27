package main

var tmplService = `
package {{.PackageName}}

import (
	"github.com/zsxm/gowsdl/webservice"
	"encoding/xml"
)

type {{.ServiceName}} struct {
	Url string
}

func New{{.ServiceName}}() *{{.ServiceName}}{
	s := {{.ServiceName}}{}
	s.Url = "{{.ServiceUrl}}"
	return &s
}{{with $s := .}}
{{range .Types}}
type {{.Name}} struct {
	XMLNamespace string {{TagDelimiter}}xml:"xmlns,attr"{{TagDelimiter}}
	{{range .Fields}}{{.Name}}	{{if StringHasValue .Type}}{{.Type}}{{end}}	{{if StringHasValue .XMLName}}{{TagDelimiter}}xml:"{{.XMLName}}"{{TagDelimiter}}{{end}}
	{{end}}
}
{{end}}{{range .Messages}}
type {{.Name}} struct {
	XMLName	xml.Name	{{TagDelimiter}}xml:"{{.XMLName}}"{{TagDelimiter}}
	Action	string		{{TagDelimiter}}xml:"-"{{TagDelimiter}}{{range .Params}}
	{{.ParamName}} {{.ParamType}} {{TagDelimiter}}xml:"{{.XMLParamName}}"{{TagDelimiter}}{{end}}
}

func (this {{.Name}}) GetAction() string {
	return this.Action
}
{{end}}{{range .Methods}}
//{{$s.ServiceName}} {{.Name}}
//Paramter {{range $i,$v:=.Params}}{{if (GT $i 0)}},{{end}}{{.ParamName}} {{.ParamType}}{{end}}
//Return {{.OutputType}}, error
func (this *{{$s.ServiceName}}) {{.Name}}({{if .HasParams}}{{range $i,$v:=.Params}}{{if (GT $i 0)}},{{end}}{{.ParamName}} {{.ParamType}}{{end}}{{end}}) (r *{{.OutputType}}, err error) {
	{{if .HasParams}}e := {{.MessageIn}}{}
	{{if StringHasValue .Action}}e.Action = "{{.Action}}"{{end}}{{range .Params}}
	e.{{.ParamName}} = {{.ParamName}}{{end}}{{end}}
	sr, err := webservice.CallService({{if .HasParams}}e{{else}}nil{{end}}, this.Url)
	if err != nil {
		return nil, err
	}
	var so {{.MessageOut}}
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}
	return &so.{{.ParamOutName}}, nil
}
{{end}}{{end}}
`
