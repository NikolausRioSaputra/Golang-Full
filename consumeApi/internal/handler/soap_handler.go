package handler

import (
	"bytes"
	"consume-api/internal/domain"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConsumeSOAP handles the SOAP API call
func ConsumeSOAP(c *gin.Context) {
	var request domain.SOAPRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Buat permintaan SOAP menggunakan data dari permintaan JSON
	soapRequest := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <Add xmlns="http://tempuri.org/">
      <intA>%d</intA>
      <intB>%d</intB>
    </Add>
  </soap:Body>
</soap:Envelope>`, request.IntA, request.IntB)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var addResp domain.AddResponse
	err = xml.Unmarshal(body, &addResp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse XML"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": addResp.Body.AddResponse.Result,
	})
}
