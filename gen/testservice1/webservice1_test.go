package testservice1_test

import (
	"fmt"
	"testing"

	"github.com/zsxm/gowsdl/gen/testservice1"
)

func TestService(t *testing.T) {
	s := testservice1.NewMessageServiceJaxbImplService()
	req := testservice1.MessageReqJaxb{}
	req.OrgNo = "3680"
	req.Version = "2.0"
	req.TypeNo = "6463"
	req.BatchId = "dslkfjlskdfjoqiwerua123f"
	req.Details = append(req.Details, testservice1.DetailsJaxb{
		Priority: "5",
		Keywords: "comName|aaaa",
		Mobile:   "15869390369",
	})
	res, err := s.OrderMsgSendJaxb(req)
	fmt.Println(err)
	fmt.Printf("%+v", res)
}
