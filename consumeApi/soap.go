package main

// // tujuan program ini ialah yang menggunakan framework gin di go untuk mengirim permintaan soap kelayanan external kemudian respon nya mengembalikan json
// import (
// 	"bytes"
// 	"encoding/xml"
// 	"io"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // xml adalah format data yg di gunakan untuk mengirim dan menerima data dengan struktur jelas
// // Struct untuk menyimpan respons SOAP
// type AddResponse struct {
// 	XMLName xml.Name `xml:"Envelope"`
// 	Body    Body     `xml:"Body"`
// }

// type Body struct {
// 	AddResponse AddResult `xml:"AddResponse"`
// }

// type AddResult struct {
// 	Result int `xml:"AddResult"`
// }

// func consumeSOAP(c *gin.Context) {
// 	soapRequest := `<?xml version="1.0" encoding="utf-8"?>
// <soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
//   <soap:Body>
//     <Add xmlns="http://tempuri.org/">
//       <intA>12</intA>
//       <intB>12</intB>
//     </Add>
//   </soap:Body>
// </soap:Envelope>`

// 	//Membuat sebuah instance baru dari http.Client. http.Client digunakan untuk melakukan permintaan HTTP. Ini menyediakan metode Do untuk mengirimkan permintaan dan mendapatkan respons.
// 	client := &http.Client{}

// 	// untuk membuat permintaan HTTP baru
// 	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer([]byte(soapRequest)))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	req.Header.Set("Content-Type", "text/xml")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim permintaan"})
// 		return
// 	}
// 	defer resp.Body.Close() // Baris ini memastikan bahwa body respons akan ditutup setelah selesai diproses untuk membebaskan sumber daya.

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
// 		return
// 	}

// 	var addResp AddResponse
// 	//Fungsi ini digunakan untuk memparsing byte array body yang berisi respons SOAP ke dalam struct AddResponse.
// 	err = xml.Unmarshal(body, &addResp)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memparsing XML"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		//Setelah parsing, kita dapat mengakses hasil penjumlahan
// 		"result": addResp.Body.AddResponse.Result,
// 	})
// }
