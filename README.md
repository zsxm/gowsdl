# soap wsdl
web service 客户端代码生成, gen目录下 main.go配置wsdl url或文件目录，指定包名(目录名)和文件名，运行。</br>
使用方式：<br/>
var pak = "testservice1"<br/>
var wsdlFile = flag.String("w", `service4.wsdl`, "WSDL file with full path")<br/>
var xsdFile = flag.String("x", ``, "XSD file with full path")<br/>
var packageName = flag.String("p", pak, "Package name")<br/>
var outFile = flag.String("o", pak+"/service.go", "Output file")<br/>

//测试代码<br/>
package testservice1_test<br/>

import (<br/>
	"fmt"<br/>
	"testing"<br/>
<br/>
	"github.com/zsxm/gowsdl/gen/testservice1"<br/>
)<br/>
<br/>
//测试用例<br/>
`func TestService(t *testing.T) {`

`	s := testservice1.NewMessageServiceJaxbImplService()`

`	req := testservice1.MessageReqJaxb{}`

`	req.OrgNo = "0000"`

`	req.Version = "2.0"`

`	req.TypeNo = "2000"`

`	req.BatchId = "dslkfjlskdfjoqiwerua123f"`

`	req.Details = append(req.Details, testservice1.DetailsJaxb{`

`		Priority: "5",`

`		Keywords: "comName|aaaa",`

`		Mobile:   "1355555555",`

`	})`

`	res, err := s.OrderMsgSendJaxb(req)`

`	fmt.Println(err)`

`	fmt.Printf("%+v", res)`

`}`
