package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func StatusById(ids []string) (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"

	// ساختار داده برای JSON
	requestData := map[string]interface{}{
		"ApiKey": apiKey,
		"Ids":    ids,
	}

	// تبدیل به JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	// ارسال درخواست POST
	url := "http://api.sms-webservice.com/api/V3/StatusById"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// نمونه‌ای از استفاده
	ids := []string{"123", "456", "789"}
	response, err := StatusById(ids)
	if err != nil {
		fmt.Println("خطا:", err)
	} else {
		fmt.Println("پاسخ:", response)
	}
}
