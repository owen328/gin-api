package vies

import (
	"bytes"
	"encoding/xml"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

const serviceUrl = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"
const envelop = `
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <Body>
        <checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
            <countryCode>{{.countryCode}}</countryCode>
            <vatNumber>{{.vatNumber}}</vatNumber>
        </checkVat>
    </Body>
</Envelope>`

type ResponseXml struct {
	XMLName xml.Name `xml:"Envelope"`
	Soap    struct {
		XMLName xml.Name `xml:"Body"`
		Soap    struct {
			XMLName     xml.Name `xml:"checkVatResponse"`
			CountryCode string   `xml:"countryCode"`
			VatNumber   string   `xml:"vatNumber"`
			RequestDate string   `xml:"requestDate"`
			Valid       bool     `xml:"valid"`
			Name        string   `xml:"name"`
			Address     string   `xml:"address"`
		}
		SoapFault struct {
			XMLName string `xml:"Fault"`
			Code    string `xml:"faultcode"`
			Message string `xml:"faultstring"`
		}
	}
}

type VatResponse struct {
	CountryCode string
	VatNumber   string
	RequestDate time.Time
	Valid       bool
	Name        string
	Address     string
}

func CheckVat(VatNumber string) (*VatResponse, error) {

	client := &http.Client{}
	requestBody, err := getRequestBody(VatNumber)
	if err != nil {
		panic(err)
	}
	reqReader := bytes.NewBufferString(requestBody)
	res, err := client.Post(serviceUrl, `text/xml; charset="utf-8"`, reqReader)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	xmlRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	respXml := &ResponseXml{}

	if err := xml.Unmarshal(xmlRes, &respXml); err != nil {
		return nil, err
	}

	if respXml.Soap.SoapFault.Message != "" {
		return nil, errors.New(respXml.Soap.SoapFault.Message)
	}
	if respXml.Soap.Soap.RequestDate == "" {
		return nil, errors.New("service returned invalid request date")
	}
	respDate, err := time.Parse("2006-01-02+15:04", respXml.Soap.Soap.RequestDate)
	if err != nil {
		return nil, err
	}
	r := &VatResponse{
		CountryCode: respXml.Soap.Soap.CountryCode,
		VatNumber:   respXml.Soap.Soap.VatNumber,
		RequestDate: respDate,
		Valid:       respXml.Soap.Soap.Valid,
		Name:        respXml.Soap.Soap.Name,
		Address:     respXml.Soap.Soap.Address,
	}
	return r, nil
}

func getRequestBody(vatNumber string) (string, error) {
	tpl, err := template.New("envelop").Parse(envelop)
	if err != nil {
		return "", err
	}
	var tplBuf bytes.Buffer
	err = tpl.Execute(&tplBuf, map[string]string{
		"countryCode": vatNumber[0:2],
		"vatNumber":   vatNumber[2:],
	})
	if err != nil {
		return "", err
	}
	return tplBuf.String(), nil
}

/*
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
<env:Header/>
<env:Body>
<ns2:checkVatResponse xmlns:ns2="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
<ns2:countryCode>FR</ns2:countryCode>
<ns2:vatNumber>29919979369</ns2:vatNumber>
<ns2:requestDate>2023-02-06+01:00</ns2:requestDate>
<ns2:valid>true</ns2:valid>
<ns2:name>PMDE TAIYUANPUTUONASHANGMAOYOUXIANGONGSI</ns2:name>
<ns2:address>
GAYE GRISHOLD-CHEZ MOKJ CONSULTING
1 PL BOIELDIEU
75002 PARIS
</ns2:address>
</ns2:checkVatResponse>
</env:Body>
</env:Envelope>
*/
