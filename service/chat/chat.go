package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Choice struct {
	OpenAIMessage Message `json:"message"`
}
type RequestParam struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

func Chat(res http.ResponseWriter, req *http.Request) {
	log.Println("user xxx request to chat with llm")
	//
	var requestParam RequestParam
	err := json.NewDecoder(req.Body).Decode(&requestParam)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	// try send req to
	log.Printf("send message to API provider xxx")
	log.Println("user xxx end with request")
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		http.Error(res, "api key is blank", http.StatusBadRequest)
	}

	url := "https://api.openai.com/v1/chat/completions"
	//requestParam.Model = "gpt-4o-mini"
	//requestParam.Messages = []Message{
	//	{Role: "user", Content: "hello"},
	//}

	jsonData, err := json.Marshal(requestParam)

	chatReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	chatReq.Header.Set("Content-Type", "application/json")
	chatReq.Header.Set("Authorization", "Bearer "+apiKey)

	client := http.Client{}
	response, err := client.Do(chatReq)

	var jsonResponse Response = Response{}
	bo, err := io.ReadAll(response.Body)
	answer := json.Unmarshal(bo, &jsonResponse)
	log.Println(answer)
	fmt.Fprintf(res, jsonResponse.Choices[0].OpenAIMessage.Content)
}
