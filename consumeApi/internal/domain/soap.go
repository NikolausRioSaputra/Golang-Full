package domain

import "encoding/xml"

// xml adalah format data yg di gunakan untuk mengirim dan menerima data dengan struktur jelas
// Struct untuk menyimpan respons SOAP
type AddResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AddResponse AddResult `xml:"AddResponse"`
}

type AddResult struct {
	Result int `xml:"AddResult"`
}
