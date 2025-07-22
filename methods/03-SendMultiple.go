package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ساختار هر گیرنده
type Recipient struct {
	Sender       string `json:"Sender"`
	Text         string `json:"Text"`
	Destination  string `json:"Destination"`
	UserTraceId  string `json:"UserTraceId"`
}

// ساختار بدنه‌ی درخواست
type SendMultipleRequest struct {
	ApiKey     string      `json:"ApiKey"`
	Recipients []Recipient `json:"Recipients"`
}

func SendMultiple(destination, userTraceId, text string) (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"
	sender := "30007546464646"

	recipients := []Recipient{
		{
			Sender:      sender,
			Text:        text,
			Destination: destination,
			UserTraceId: userTraceId,
		},
	}

	requestBody := SendMultipleRequest{
		ApiKey:     apiKey,
		Recipients: recipients,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	url := "http://api.sms-webservice.com/api/V3/SendMultiple"
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
	// تست
	response, err := SendMultiple("09123456789", "trace-id-123", "پیام تستی برای ارسال گروهی")
	if err != nil {
		fmt.Println("خطا:", err)
	} else {
		fmt.Println("پاسخ:", response)
	}
}
