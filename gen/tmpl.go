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
}
{{with $s := .}}
{{range .Types}}
type {{.Name}} struct {
	XMLNamespace string {{TagDelimiter}}xml:"xmlns,attr"{{TagDelimiter}}
	{{range .Fields}}{{.Name}} {{if StringHasValue .Type}}{{.Type}}{{end}} {{if StringHasValue .XMLName}}{{TagDelimiter}}xml:"{{.XMLName}}"{{TagDelimiter}}{{end}}
	{{end}}
}
{{end}}
{{range .Messages}}
type {{.Name}} struct {
	XMLName	xml.Name	{{TagDelimiter}}xml:"{{.XMLName}}"{{TagDelimiter}}
	Action	string		{{TagDelimiter}}xml:"-"{{TagDelimiter}}
	{{range .Params}}
	{{.ParamName}} {{.ParamType}} {{TagDelimiter}}xml:"{{.XMLParamName}}"{{TagDelimiter}}{{end}}
}

func (si {{.Name}}) GetAction() string {
	return si.Action
}
{{end}}{{range .Methods}}
func (s *{{$s.ServiceName}}) {{.Name}}({{if .HasParams}}p{{end}} {{.InputType}}) (r *{{.OutputType}}, err error) {
	{{if .HasParams}}si := {{.MessageIn}}{}
	si.Action = "{{.Action}}"
	si.RequestList=p{{end}}
	sr, err := webservice.CallService({{if .HasParams}}si{{else}}nil{{end}}, s.Url)
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
