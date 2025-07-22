package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Send(recipients string, text string) (string, error) {
	apiKey := "e883424d-d70f-4e58-8ee3-4e21ea390ff1"
	sender := "30007546464646"

	params := url.Values{}
	params.Add("ApiKey", apiKey)
	params.Add("Text", text)
	params.Add("Sender", sender)
	params.Add("Recipients", recipients)

	baseURL := "http://api.sms-webservice.com/api/V3/Send?"
	fullURL := baseURL + params.Encode()

	resp, err := http.Get(fullURL)
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
	// مثال استفاده
	response, err := Send("09121234567", "کد تأیید شما 1234 است")
	if err != nil {
		fmt.Println("خطا:", err)
	} else {
		fmt.Println("پاسخ:", response)
	}
}
