package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SendTokenSingle برای ارسال پیام توکنی با پارامترهای مشخص
func SendTokenSingle(templateKey, destination, param1, param2, param3 string) (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"

	// ساخت query string
	params := url.Values{}
	params.Add("ApiKey", apiKey)
	params.Add("TemplateKey", templateKey)
	params.Add("Destination", destination)
	params.Add("p1", param1)
	params.Add("p2", param2)
	params.Add("p3", param3)

	fullURL := "http://api.sms-webservice.com/api/V3/SendTokenSingle?" + params.Encode()

	// ارسال درخواست GET
	resp, err := http.Get(fullURL)
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
	// تست تابع
	response, err := SendTokenSingle("TemplateKey123", "09123456789", "کد1", "کد2", "کد3")
	if err != nil {
		fmt.Println("خطا در ارسال:", err)
	} else {
		fmt.Println("پاسخ سرور:", response)
	}
}
