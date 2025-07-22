package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ساختار گیرنده پیام
type Recipient struct {
	Destination  string `json:"Destination"`
	UserTraceId  string `json:"UserTraceId"`
}

// ساختار داده ارسالی به سرور
type SendBulkRequest struct {
	ApiKey     string      `json:"ApiKey"`
	Text       string      `json:"Text"`
	Sender     string      `json:"Sender"`
	Recipients []Recipient `json:"Recipients"`
}

func SendBulk(destination, userTraceId, text string) (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"
	sender := "30007546464646"

	recipients := []Recipient{
		{
			Destination: destination,
			UserTraceId: userTraceId,
		},
	}

	requestBody := SendBulkRequest{
		ApiKey:     apiKey,
		Text:       text,
		Sender:     sender,
		Recipients: recipients,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	url := "http://api.sms-webservice.com/api/V3/SendBulk"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func main() {
	// مثال استفاده از تابع
	response, err := SendBulk("09123456789", "trace123", "پیام تستی برای ارسال انبوه")
	if err != nil {
		fmt.Println("خطا:", err)
	} else {
		fmt.Println("پاسخ سرور:", response)
	}
}
