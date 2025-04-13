package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DispatchSms (message string, recipient string)(bool, error){
	url := os.Getenv("ARKESEL_URL")

	payload := map[string]any{"sender": "Benevo GH", "message": message, "recipients": []string{recipient}}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return false, err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", os.Getenv("ARKESEL_API_KEY")) 
	client := &http.Client{}
	resp, err := client.Do(req)


	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	return true, nil

}