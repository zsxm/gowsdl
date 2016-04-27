package main

import (
	"flag"
	"fmt"

	"github.com/zsxm/gowsdl/wsdl"
	"github.com/zsxm/gowsdl/xsd"
	"github.com/zsxm/scgo/chttplib"
)

var pak = "testservice2"
var wsdlFile = flag.String("w", `service4.wsdl`, "WSDL file with full path")
var xsdFile = flag.String("x", ``, "XSD file with full path")
var packageName = flag.String("p", pak, "Package name")
var outFile = flag.String("o", pak+"/service.go", "Output file")

// wsdl -w="C:\Temp\wsdl\CartaoEndpointService.wsdl" -x="C:\Temp\wsdl\CartaoEndpointService_schema1.xsd" -p="main" -o="C:\Temp\service.go"
// wsdl -w="C:\Temp\wsdl\authendpointservice.wsdl" -x="C:\Temp\wsdl\AuthEndpointService_schema1.xsd" -p="login" -o="C:\Temp\auth_service.go"
func main() {
	flag.Parse()
	if *wsdlFile == "" || *packageName == "" || *outFile == "" {
		flag.Usage()
		return
	}

	var d wsdl.Definitions
	unmarshal(*wsdlFile, &d)
	//fmt.Printf("%+v\n", d.Types.Schemas)
	var s xsd.Schema

	// se foi informado qual o arquivo de schema
	// o schema pode ser passado em um arquivo separado ou
	// poder estar dentro do proprio wsdl
	if *xsdFile != "" {
		unmarshal(*xsdFile, &s)
	} else {
		// TODO: na verdade podemos ter mais de um schema
		s = d.Types.Schemas[0]
	}

	buf, f := createOut(*outFile)
	defer buf.Flush()
	defer f.Close()
	//	fmt.Printf("Import %+v\n", s.Import)
	//	fmt.Println("-------------------------------------------------------------------------------------------", "\n")
	//	fmt.Printf("Elements %+v\n", s.Elements)
	//	fmt.Println("-------------------------------------------------------------------------------------------", "\n")
	//	fmt.Printf("ComplexTypes %+v\n", s.ComplexTypes)
	//	fmt.Println("-------------------------------------------------------------------------------------------", "\n")
	if s.Import.Namespace != "" && s.Import.SchemaLocation != "" {
		url := s.Import.SchemaLocation
		fmt.Println(url)
		res := wsdlServer(url)
		xmlUnmarshal(res, &s)
		//fmt.Printf("s %+v\n", s)
	}
	// create de service file
	create(&d, &s, buf, f, *packageName)
}

//url
func Get(wurl, packageName, outFile string) {
	var d wsdl.Definitions
	res := wsdlServer(wurl)
	xmlUnmarshal(res, &d)
	var s = d.Types.Schemas[0]
	if s.Import.Namespace != "" && s.Import.SchemaLocation != "" {
		wurl = s.Import.SchemaLocation
		fmt.Println(wurl)
		res := wsdlServer(wurl)
		xmlUnmarshal(res, &s)
	}
	gen(d, s, outFile, packageName)
}

//file
func File(wsdlFile, xsdFile, packageName, outFile string) {
	if wsdlFile == "" || packageName == "" || outFile == "" {
		exit(fmt.Errorf("%s", "wsdlFile or packageName or outFile is null"))
	}
	var d wsdl.Definitions
	unmarshal(wsdlFile, &d)
	var s xsd.Schema
	if xsdFile != "" {
		unmarshal(xsdFile, &s)
	} else {
		s = d.Types.Schemas[0]
	}
	if s.Import.Namespace != "" && s.Import.SchemaLocation != "" {
		url := s.Import.SchemaLocation
		fmt.Println(url)
		res := wsdlServer(url)
		xmlUnmarshal(res, &s)
	}
	gen(d, s, outFile, packageName)
}

func gen(d wsdl.Definitions, s xsd.Schema, packageName, outFile string) {
	buf, f := createOut(outFile)
	defer buf.Flush()
	defer f.Close()
	create(&d, &s, buf, f, packageName)
}

func wsdlServer(wurl string) []byte {
	req := chttplib.Get(wurl)
	res, err := req.Bytes()
	if err != nil {
		exit(err)
	}
	return res
}
