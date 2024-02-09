package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type InputData struct {
	Ev     string                            `json:"ev"`
	Et     string                            `json:"et"`
	Id     string                            `json:"id"`
	Uid    string                            `json:"uid"`
	Mid    string                            `json:"mid"`
	T      string                            `json:"t"`
	P      string                            `json:"p"`
	L      string                            `json:"l"`
	Sc     string                            `json:"sc"`
	Attrs  map[string]map[string]interface{} `json:"attrs"`
	UAttrs map[string]map[string]interface{} `json:"uattrs"`
}

type OutputData struct {
	Event           string                 `json:"event"`
	EventType       string                 `json:"event_type"`
	AppID           string                 `json:"app_id"`
	UserID          string                 `json:"user_id"`
	MessageID       string                 `json:"message_id"`
	PageTitle       string                 `json:"page_title"`
	PageURL         string                 `json:"page_url"`
	BrowserLanguage string                 `json:"browser_language"`
	ScreenSize      string                 `json:"screen_size"`
	Attributes      map[string]interface{} `json:"attributes"`
	Traits          map[string]interface{} `json:"traits"`
}

var webhookURL = "https://webhook.site/e7500611-e5e7-4d99-92dd-6d27f93de336"
var wg sync.WaitGroup

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var inputData InputData
	fmt.Println("Received request")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	wg.Add(1)
	go processRequest(w, inputData)
}

func processRequest(w http.ResponseWriter, inputData InputData) {
	defer wg.Done()

	outputData := transformData(inputData)

	outputJSON, err := json.Marshal(outputData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(outputJSON))
	if err != nil {
		fmt.Println("Error sending data to webhook:", err)
		http.Error(w, "Error sending data to webhook", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Webhook response:", resp.Status)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request received and processing in the background"))
}

func transformData(inputData InputData) OutputData {
	outputData := OutputData{
		Event:           inputData.Ev,
		EventType:       inputData.Et,
		AppID:           inputData.Id,
		UserID:          inputData.Uid,
		MessageID:       inputData.Mid,
		PageTitle:       inputData.T,
		PageURL:         inputData.P,
		BrowserLanguage: inputData.L,
		ScreenSize:      inputData.Sc,
		Attributes:      make(map[string]interface{}),
		Traits:          make(map[string]interface{}),
	}

	for key, value := range inputData.Attrs {
		outputData.Attributes[key] = map[string]interface{}{
			"value": value["value"],
			"type":  value["type"],
		}
	}

	for key, value := range inputData.UAttrs {
		outputData.Traits[key] = map[string]interface{}{
			"value": value["value"],
			"type":  value["type"],
		}
	}

	return outputData
}
