package chatgpt

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"log"
)

const (
	apiEndpoint = "https://api.chatanywhere.com.cn/v1/chat/completions"
)

func GptParse(info string) string {
	// Use your API KEY here
	apiKey := "sk-0JF93y6z9gGGkMJ7ax5PS5g8NiDztILwDyQ7Klkaf0RFZBzI"
	client := resty.New()

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []interface{}{map[string]interface{}{"role": "system", "content": "There are currently five tags: Happy, Angry, Complaining, Praising, and Flat. " +
				"Which label does the following belong to: " +
				info +
				"Just answer which label"}},
			"max_tokens": 50,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error while decoding JSON response:", err)
		return ""
	}

	// Extract the content from the JSON response
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	fmt.Println(content)
	return content
}
