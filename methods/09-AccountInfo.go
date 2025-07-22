package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AccountInfo() (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"
	data := map[string]string{
		"ApiKey": apiKey,
	}

	// تبدیل داده‌ها به JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// ایجاد درخواست POST
	url := "http://api.sms-webservice.com/api/V3/AccountInfo"
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

	// خواندن پاسخ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	response, err := AccountInfo()
	if err != nil {
		fmt.Println("خطا در دریافت اطلاعات:", err)
	} else {
		fmt.Println("پاسخ سرور:")
		fmt.Println(response)
	}
}
