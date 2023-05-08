package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BitlyResponse struct {
	Link string `json:"link"`
}

func ShortenURL(longURL string) (string, error) {
	// Set up a client for the Bitly API
	client := &http.Client{}
	baseURL := "https://api-ssl.bitly.com/v4/shorten"
	apiKey := "84d3b40d1b208f80ad0114d19d2a2bc654810cc1"

	authToken := fmt.Sprintf("Bearer %s", apiKey)

	// Set up the request body with the original URL
	reqBody := map[string]string{
		"long_url": longURL,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// Make a POST request to the Bitly API to shorten the URL
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the response from the Bitly API
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println(authToken)
	log.Println(string(respBodyBytes))

	var bitlyResponse BitlyResponse

	err = json.Unmarshal(respBodyBytes, &bitlyResponse)
	if err != nil {
		return "", err
	}

	return bitlyResponse.Link, nil
}
