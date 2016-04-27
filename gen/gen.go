package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/zsxm/gowsdl/wsdl"
	"github.com/zsxm/gowsdl/xsd"
)

type Method struct {
	Name         string
	Action       string
	HasParams    bool
	InputType    string
	OutputType   string
	MessageIn    string
	MessageOut   string
	ParamInName  string
	ParamOutName string
	Params       []MessageParamIn
}

type Message struct {
	Name    string
	XMLName string
	Action  string
	Params  []MessageParamIn

	ParamName    string
	XMLParamName string
	ParamType    string
	Input        bool
}

type MessageParamIn struct {
	ParamName    string
	XMLParamName string
	ParamType    string
	Input        bool
}

type Field struct {
	Name    string
	Type    string
	XMLName string
}

type StructType struct {
	Name   string
	Fields []Field
}

type TemplateData struct {
	PackageName string
	ServiceName string
	ServiceUrl  string
	Messages    []Message
	Methods     []Method
	Types       []StructType
}

var data TemplateData

func unmarshal(n string, i interface{}) {
	f, err := os.Open(n)
	if err != nil {
		exit(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		exit(err)
	}

	xmlUnmarshal(b, i)
}

func xmlUnmarshal(b []byte, i interface{}) {
	err := xml.Unmarshal(b, i)
	if err != nil {
		exit(err)
	}
}

func createOut(n string) (*bufio.Writer, *os.File) {
	// remove o arquivo de saida
	err := os.Remove(n)
	// verificar se houve, se houve erro e o erro não for do tipo "não existe"
	// ignora o erro se retorna o erro em questão
	if err != nil && !os.IsNotExist(err) {
		exit(err)
	}

	// cria o arquivo de saída
	f, err := os.Create(n)
	if err != nil {
		exit(err)
	}

	return bufio.NewWriter(f), f
}

// exit sai da aplicação exibindo o erro se existir
func exit(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

// create cria o arquivo com o serviço a ser consumido
func create(d *wsdl.Definitions, s *xsd.Schema, b *bufio.Writer, file *os.File) {
	funcMap := template.FuncMap{
		"StringHasValue": StringHasValue,
		"TagDelimiter":   TagDelimiter,
		"GT":             GT,
	}

	// create the template
	tmpl, err := template.New("T").Funcs(funcMap).Parse(tmplService)
	if err != nil {
		exit(err)
	}

	data.PackageName = *packageName
	data.ServiceName = d.Service.Name
	data.ServiceUrl = d.Service.Port.Address.Location
	data.Methods = make([]Method, 0)
	data.Messages = make([]Message, 0)

	for i := 0; i < len(s.ComplexTypes); i++ {
		// check if complex type not in elements
		found := false
		for j := 0; j < len(s.Elements); j++ {
			if s.ComplexTypes[i].Name == s.Elements[j].Name {
				found = true
				break
			}
		}
		var t StructType
		if !found {
			t = StructType{
				Name:   exportableSymbol(s.ComplexTypes[i].Name),
				Fields: make([]Field, 0),
			}

			if s.ComplexTypes[i].Content == nil {
				for ii := 0; ii < len(s.ComplexTypes[i].Sequence); ii++ {
					fi := Field{
						Name:    exportableSymbol(s.ComplexTypes[i].Sequence[ii].Name),
						Type:    decodeType(s.ComplexTypes[i].Sequence[ii]),
						XMLName: s.ComplexTypes[i].Sequence[ii].Name,
					}
					t.Fields = append(t.Fields, fi)
				}
			} else {
				t.Fields = append(t.Fields, Field{Name: exportableSymbol(s.ComplexTypes[i].Content.Extension.Base[4:])})
				for ii := 0; ii < len(s.ComplexTypes[i].Content.Extension.Sequence); ii++ {
					fi := Field{
						Name:    exportableSymbol(s.ComplexTypes[i].Content.Extension.Sequence[ii].Name),
						Type:    decodeType(s.ComplexTypes[i].Content.Extension.Sequence[ii]),
						XMLName: s.ComplexTypes[i].Content.Extension.Sequence[ii].Name,
					}
					t.Fields = append(t.Fields, fi)
				}
			}
			data.Types = append(data.Types, t)
		}
	}

	for i := 0; i < len(d.PortType.Operations); i++ {
		operation := d.PortType.Operations[i]
		m := Method{}
		m.Name = exportableSymbol(operation.Name)
		// TODO: get correct action in binding area
		m.Action = operation.Input.Action
		//fmt.Println("---------", m.Name, operation.Input)
		// find input parameter type
		e := findElement(s, operation.Input.Message)

		var c *xsd.ComplexType
		if e.ComplexTypes == nil {
			c = findComplexType(s, e.Name)
		} else {
			c = e.ComplexTypes
		}

		reqMsg := Message{}
		if c.Name != "" {
			reqMsg.Name = exportableSymbol(c.Name)
			reqMsg.XMLName = d.TargetNamespace + " " + c.Name
		} else {
			reqMsg.Name = exportableSymbol(e.Name)
			reqMsg.XMLName = d.TargetNamespace + " " + e.Name
		}
		reqMsg.Action = m.Action
		fmt.Println(m.Action)

		if len(c.Sequence) > 0 {
			si := strings.Index(c.Sequence[0].Type, ":")
			for _, v := range c.Sequence {
				messageParam := MessageParamIn{}
				messageParam.ParamName = exportableSymbol(v.Name)
				messageParam.XMLParamName = v.Name
				messageParam.ParamType = exportableSymbol(v.Type[si+1:])
				messageParam.Input = true

				reqMsg.Params = append(reqMsg.Params, messageParam)
				m.Params = append(m.Params, messageParam)
			}
			data.Messages = append(data.Messages, reqMsg)

			m.InputType = exportableSymbol(c.Sequence[0].Type[si+1:])
			m.MessageIn = reqMsg.Name
			m.ParamInName = reqMsg.ParamName
			m.HasParams = true
		} else {
			m.HasParams = false
		}

		// find output parameter type
		e = findElement(s, operation.Output.Message)
		if e.ComplexTypes == nil {
			c = findComplexType(s, e.Name)
		} else {
			c = e.ComplexTypes
		}
		si := strings.Index(c.Sequence[0].Type, ":")

		resMsg := Message{}

		if c.Name != "" {
			resMsg.Name = exportableSymbol(c.Name)
			resMsg.XMLName = c.Name
		} else {
			resMsg.Name = exportableSymbol(e.Name)
			resMsg.XMLName = e.Name
		}

		for _, v := range c.Sequence {
			messageParam := MessageParamIn{}
			messageParam.ParamName = exportableSymbol(v.Name)
			messageParam.XMLParamName = v.Name
			messageParam.ParamType = exportableSymbol(v.Type[si+1:])
			messageParam.Input = false
			resMsg.Params = append(resMsg.Params, messageParam)
		}

		data.Messages = append(data.Messages, resMsg)

		m.OutputType = exportableSymbol(c.Sequence[0].Type[si+1:])
		m.MessageOut = resMsg.Name
		m.ParamOutName = exportableSymbol(c.Sequence[0].Name)
		data.Methods = append(data.Methods, m)

	}

	err = tmpl.Execute(file, data)
	if err != nil {
		exit(err)
	}
}

func findElement(s *xsd.Schema, t string) *xsd.Element {
	if t[0:3] == "tns" {
		t = t[4:]
	}
	t = strings.Replace(t, "SoapIn", "", -1)
	t = strings.Replace(t, "SoapOut", "Response", -1)
	for i := 0; i < len(s.Elements); i++ {
		if s.Elements[i].Type == t || s.Elements[i].Name == t {
			return &s.Elements[i]
		}
	}
	return nil
}

func findComplexType(s *xsd.Schema, n string) *xsd.ComplexType {
	if n[0:3] == "tns" {
		n = n[4:]
	}
	n = strings.Replace(n, "SoapIn", "", -1)
	n = strings.Replace(n, "SoapOut", "Response", -1)
	for i := 0; i < len(s.ComplexTypes); i++ {
		if s.ComplexTypes[i].Name == n {
			return &s.ComplexTypes[i]
		}
	}
	return nil
}

func exportableSymbol(s string) string {
	switch s {
	case "string", "int":
		return s
	default:
		return strings.ToUpper(s[0:1]) + s[1:]
	}
}

func decodeType(e xsd.Element) string {
	t := e.Type
	// TODO(dops): tratar
	if t == "" && e.Name == "entry" {
		return "[]Entry"
	}
	if t[0:2] == "xs" {
		switch t[3:] {
		case "string":
			return "string"
		case "boolean":
			return "bool"
		case "decimal", "double":
			return "float64"
		case "int":
			return "int"
		case "long":
			return "int64"
		case "dateTime":
			return "int"
		default:
			//return "nil"
			return t[3:]
		}
	} else if t[0:2] == "s:" {
		switch t[2:] {
		case "string":
			return "string"
		case "boolean", "bool":
			return "bool"
		case "decimal", "double":
			return "float64"
		case "int":
			return "int"
		case "long":
			return "int64"
		case "dateTime":
			return "int"
		default:
			//return "nil"
			return t[3:]
		}
	} else if t[0:3] == "tns" {
		ty := exportableSymbol(t[4:])
		if e.MaxOccurs == "unbounded" {
			ty = "[]" + ty
		}
		return ty
	}
	panic("unknown type")
}

func StringHasValue(s string) bool {
	if s != "" {
		return true
	}
	return false
}

func TagDelimiter() string {
	return "`"
}

func GT(k1, k2 int) bool {
	return k1 > k2
}
