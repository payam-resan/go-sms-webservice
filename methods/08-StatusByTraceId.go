package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ساختار داده برای ارسال
type RequestPayload struct {
	ApiKey       string   `json:"ApiKey"`
	UserTraceIds []string `json:"UserTraceIds"`
}

func StatusByUserTraceId(userTraceIds []string) (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"
	url := "http://api.sms-webservice.com/api/V3/StatusByUserTraceId"

	// ساخت بادی
	payload := RequestPayload{
		ApiKey:       apiKey,
		UserTraceIds: userTraceIds,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("خطا در ساخت JSON: %v", err)
	}

	// ارسال درخواست
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("خطا در ساخت درخواست: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("خطا در خواندن پاسخ: %v", err)
	}

	return buf.String(), nil
}

func main() {
	// نمونه استفاده
	userTraceIds := []string{"123456", "789012"}
	response, err := StatusByUserTraceId(userTraceIds)
	if err != nil {
		fmt.Println("خطا:", err)
		return
	}
	fmt.Println("پاسخ API:")
	fmt.Println(response)
}
