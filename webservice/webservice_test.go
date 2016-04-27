package webservice

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type ObterUsuariosSoapIn struct {
	XMLName xml.Name        `xml:"http://ws.service.creditease.com/ orderMsgSendJaxb"`
	Action  string          `xml:"-"`
	Param   UsuarioParamDTO `xml:"parameters"`
}

func (si ObterUsuariosSoapIn) GetAction() string {
	return si.Action
}

type ObterUsuariosSoapOut struct {
	XMLName xml.Name          `xml:"http://ws.service.creditease.com/ orderMsgSendJaxbResponse"`
	Return  UsuarioResultList `xml:"return"`
}

func NewObterUsuariosSoapIn() ObterUsuariosSoapIn {
	si := ObterUsuariosSoapIn{}
	si.Action = ""

	return si
}

type UsuarioParamDTO struct {
	XMLNamespace string `xml:"xmlns,attr"`
	Version      string `xml:"version"`
	BatchId      string `xml:"batchId"`
	OrgNo        string `xml:"orgNo"`
	TypeNo       string `xml:"typeNo"`
}

type UsuarioResultList struct {
	CodErro   string `xml:"codErro"`
	MsgErro   string `xml:"msgErro"`
	TipoConta string `xml:"tipoConta"`
}

type AuthEndpointService struct {
	Url string
}

func NewAuthEndpointService() *AuthEndpointService {
	s := &AuthEndpointService{}
	s.Url = "http://118.244.192.173:8902/sms/services/MessageService2.0"

	return s
}

// função de chamada de exemplo
func (s AuthEndpointService) ObterUsuarios(param UsuarioParamDTO) (r *UsuarioResultList, err error) {
	si := NewObterUsuariosSoapIn()
	si.Param = param

	// chama o serviço apontando para determina url
	sr, err := CallService(si, s.Url)
	if err != nil {
		return nil, err
	}

	// monta a estrutura de retorno
	var so ObterUsuariosSoapOut
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}

	return &so.Return, nil
}

func TestAuthEndpointService(t *testing.T) {
	s := NewAuthEndpointService()

	p := UsuarioParamDTO{}
	p.Version = "2.0"
	p.BatchId = "11112222333344445555666677778888"
	p.OrgNo = "0324"
	p.TypeNo = "0126"

	l, err := s.ObterUsuarios(p)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(l)

}
