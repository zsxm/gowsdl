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
//测试用例
`func TestService(t *testing.T) {<br/>
	s := testservice1.NewMessageServiceJaxbImplService()<br/>
	req := testservice1.MessageReqJaxb{}<br/>
	req.OrgNo = "0000"<br/>
	req.Version = "2.0"<br/>
	req.TypeNo = "2000"<br/>
	req.BatchId = "dslkfjlskdfjoqiwerua123f"<br/>
	req.Details = append(req.Details, testservice1.DetailsJaxb{<br/>
		Priority: "5",<br/>
		Keywords: "comName|aaaa",<br/>
		Mobile:   "1355555555",<br/>
	})
	res, err := s.OrderMsgSendJaxb(req)<br/>
	fmt.Println(err)<br/>
	fmt.Printf("%+v", res)<br/>
}`<br/>