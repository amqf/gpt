package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
	API_BASE_URL   = "https://api.openai.com/v1"
)

type CompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Role     string `json:"role"`
			Content  string `json:"content"`
			Created  int64  `json:"created"`
			DataType string `json:"data_type"`
		} `json:"message"`
		FinishReason  string  `json:"finish_reason"`
		Index         int     `json:"index"`
		TotalLogProbs float64 `json:"total_logprobs"`
		Tokens        int     `json:"tokens"`
	} `json:"choices"`
}

type Payload struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	messageContent := flag.String("prompt", "", "your prompt")
	temperature := flag.Float64("temperature", 1, "temperature between 0 and 1, or one of two")

	flag.Parse()

	if *messageContent == "" {
		fmt.Println("What is your prompt?")
		return
	}

	if *temperature < 0 || *temperature > 1 {
		fmt.Println("Temperature must be a value between 0 and 1, or one of two")
		return
	}

	url := fmt.Sprintf("%s/chat/completions", API_BASE_URL)

	payload := Payload{
		Model:       "gpt-3.5-turbo",
		Messages:    []Message{{Role: "user", Content: *messageContent}},
		Temperature: *temperature,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	var completionResp CompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&completionResp)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	if len(completionResp.Choices) == 0 {
		fmt.Println("Check value in OPENAI_API_KEY environment variable.")
		return
	}

	fmt.Printf("%v\n", completionResp.Choices[0].Message.Content)
}
