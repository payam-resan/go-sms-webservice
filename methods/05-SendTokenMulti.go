package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Recipient struct {
	Destination  string   `json:"Destination"`
	UserTraceId  string   `json:"UserTraceId"`
	Parameters   []string `json:"Parameters"`
}

type SendTokenMultiRequest struct {
	ApiKey      string      `json:"ApiKey"`
	TemplateKey string      `json:"TemplateKey"`
	Recipients  []Recipient `json:"Recipients"`
}

func SendTokenMulti(templateKey, destination, userTraceId string, parameters []string) (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"

	// آماده‌سازی داده‌ها
	requestBody := SendTokenMultiRequest{
		ApiKey:      apiKey,
		TemplateKey: templateKey,
		Recipients: []Recipient{
			{
				Destination: destination,
				UserTraceId: userTraceId,
				Parameters:  parameters,
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	url := "http://api.sms-webservice.com/api/V3/SendTokenMulti"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	// ارسال درخواست
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// دریافت پاسخ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// مثال تستی برای فراخوانی
	response, err := SendTokenMulti("MyTemplateKey", "09123456789", "UserID-001", []string{"پارامتر 1", "پارامتر 2"})
	if err != nil {
		fmt.Println("خطا:", err)
	} else {
		fmt.Println("پاسخ سرور:", response)
	}
}
