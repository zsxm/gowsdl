package testservice

import (
	"encoding/xml"

	"code.google.com/p/wsdl-go/webservice"
)

type Entry struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type MessageServiceJaxbImplService struct {
	Url string
}

func NewMessageServiceJaxbImplService() *MessageServiceJaxbImplService {
	s := MessageServiceJaxbImplService{}
	s.Url = "http://118.244.192.173:8902/sms/services/MessageService2.0?wsdl"

	return &s
}

type MessageReqJaxb struct {
	XMLNamespace string        `xml:"xmlns,attr"`
	BatchId      string        `xml:"batchId"`
	Details      []DetailsJaxb `xml:"details"`
	OrgNo        string        `xml:"orgNo"`
	TypeNo       string        `xml:"typeNo"`
	Version      string        `xml:"version"`
}

type DetailsJaxb struct {
	XMLNamespace string `xml:"xmlns,attr"`
	Content      string `xml:"content"`
	DetailType   string `xml:"detailType"`
	Keywords     string `xml:"keywords"`
	Mobile       string `xml:"mobile"`
	Priority     string `xml:"priority"`
	Remark       string `xml:"remark"`
	SendTime     string `xml:"sendTime"`
	ShowNo       string `xml:"showNo"`
	SignInfo     string `xml:"signInfo"`
}

type MessageResJaxb struct {
	XMLNamespace string `xml:"xmlns,attr"`
	BatchId      string `xml:"batchId"`
	OrgNo        string `xml:"orgNo"`
	Remark       string `xml:"remark"`
	RetCode      string `xml:"retCode"`
	RetInfo      string `xml:"retInfo"`
	Version      string `xml:"version"`
}

type SmsMessageModel struct {
	XMLNamespace string `xml:"xmlns,attr"`
	ModContext   string `xml:"modContext"`
	ModExplain   string `xml:"modExplain"`
	ModId        string `xml:"modId"`
	ModNumber    string `xml:"modNumber"`
	ModState     string `xml:"modState"`
	ModType      string `xml:"modType"`
	OrgNo        string `xml:"orgNo"`
	Passageway   string `xml:"passageway"`
	SmsRemark1   string `xml:"smsRemark1"`
	SmsRemark2   string `xml:"smsRemark2"`
	SmsRemark3   string `xml:"smsRemark3"`
	SmsRemark4   string `xml:"smsRemark4"`
	SmsRemark5   string `xml:"smsRemark5"`
}

type OrderMsgSendJaxb struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ orderMsgSendJaxb"`
	Action  string   `xml:"-"`

	RequestList MessageReqJaxb `xml:"requestList"`
}

func (si OrderMsgSendJaxb) GetAction() string {
	return si.Action
}

type OrderMsgSendJaxbResponse struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ orderMsgSendJaxbResponse"`
	Action  string   `xml:"-"`

	Return MessageResJaxb `xml:"return"`
}

func (si OrderMsgSendJaxbResponse) GetAction() string {
	return si.Action
}

type BatchOrderMsgSendJaxb struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ batchOrderMsgSendJaxb"`
	Action  string   `xml:"-"`

	RequestList MessageReqJaxb `xml:"requestList"`
}

func (si BatchOrderMsgSendJaxb) GetAction() string {
	return si.Action
}

type BatchOrderMsgSendJaxbResponse struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ batchOrderMsgSendJaxbResponse"`
	Action  string   `xml:"-"`

	Return MessageResJaxb `xml:"return"`
}

func (si BatchOrderMsgSendJaxbResponse) GetAction() string {
	return si.Action
}

type BatchCustomMessageSendJaxb struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ batchCustomMessageSendJaxb"`
	Action  string   `xml:"-"`

	RequestList MessageReqJaxb `xml:"requestList"`
}

func (si BatchCustomMessageSendJaxb) GetAction() string {
	return si.Action
}

type BatchCustomMessageSendJaxbResponse struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ batchCustomMessageSendJaxbResponse"`
	Action  string   `xml:"-"`

	Return MessageResJaxb `xml:"return"`
}

func (si BatchCustomMessageSendJaxbResponse) GetAction() string {
	return si.Action
}

type QueryModel struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ queryModel"`
	Action  string   `xml:"-"`

	Arg0 string `xml:"arg0"`

	Arg1 string `xml:"arg1"`

	Arg2 string `xml:"arg2"`
}

func (si QueryModel) GetAction() string {
	return si.Action
}

type QueryModelResponse struct {
	XMLName xml.Name `xml:"http://ws.service.creditease.com/ queryModelResponse"`
	Action  string   `xml:"-"`

	Return SmsMessageModel `xml:"return"`
}

func (si QueryModelResponse) GetAction() string {
	return si.Action
}

func (s *MessageServiceJaxbImplService) OrderMsgSendJaxb(p MessageReqJaxb) (r *MessageResJaxb, err error) {
	si := OrderMsgSendJaxb{}
	si.Action = "1"
	si.RequestList = p

	sr, err := webservice.CallService(si, s.Url)
	if err != nil {
		return nil, err
	}

	var so OrderMsgSendJaxbResponse
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}

	return &so.Return, nil
}

func (s *MessageServiceJaxbImplService) BatchOrderMsgSendJaxb(p MessageReqJaxb) (r *MessageResJaxb, err error) {
	si := BatchOrderMsgSendJaxb{}
	si.Action = "1"
	si.RequestList = p

	sr, err := webservice.CallService(si, s.Url)
	if err != nil {
		return nil, err
	}

	var so BatchOrderMsgSendJaxbResponse
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}

	return &so.Return, nil
}

func (s *MessageServiceJaxbImplService) BatchCustomMessageSendJaxb(p MessageReqJaxb) (r *MessageResJaxb, err error) {
	si := BatchCustomMessageSendJaxb{}
	si.Action = "1"
	si.RequestList = p

	sr, err := webservice.CallService(si, s.Url)
	if err != nil {
		return nil, err
	}

	var so BatchCustomMessageSendJaxbResponse
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}

	return &so.Return, nil
}

func (s *MessageServiceJaxbImplService) QueryModel(p string) (r *SmsMessageModel, err error) {
	si := QueryModel{}
	si.Action = "1"

	sr, err := webservice.CallService(si, s.Url)
	if err != nil {
		return nil, err
	}

	var so QueryModelResponse
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}

	return &so.Return, nil
}
