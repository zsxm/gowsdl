package testservice_test

import (
	"fmt"
	"testing"

	"github.com/zsxm/gowsdl/gen/testservice"
)

func TestService(t *testing.T) {
	s := testservice.NewMessageServiceJaxbImplService()
	req := testservice.MessageReqJaxb{}
	req.OrgNo = "3680"
	req.Version = "2.0"
	req.TypeNo = "6463"
	req.BatchId = "dslkfjlskdfjoqiwerua123f"
	req.Details = append(req.Details, testservice.DetailsJaxb{
		Priority: "5",
		Keywords: "comName|aaaa",
		Mobile:   "15869390369",
	})
	res, err := s.OrderMsgSendJaxb(req)
	fmt.Println(err)
	fmt.Printf("%+v", res)
}
