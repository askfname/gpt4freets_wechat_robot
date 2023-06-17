package gpt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/qingconglaixueit/wechatbot/config"
)

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
	Error   struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    interface{} `json:"code"`
	} `json:"error"`
}

type ChoiceItem struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     int    `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

type Response struct {
	Content string `json:"content"`
}

// Completions gtp文本模型回复
// curl https://api.openai.com/v1/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer your chatGPT key"
// -d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {
	//var gptResponseBody *ChatGPTResponseBody
	var resErr error
	var responseBody string
	fmt.Println(responseBody)

	responseBody, resErr = httpRequestCompletions(msg)

	if resErr != nil {
		return "", resErr
	} else {
		return responseBody, nil
	}
}

func httpRequestCompletions(msg string) (string, error) {
	cfg := config.LoadConfig()
	log.Printf("gpt request json: %s\n", msg)

	// 参数值进行编码
	msg = url.QueryEscape(msg)
	// 构建完整的 URL
	requestURL := fmt.Sprintf("%s/ask?prompt=%s&&model=%s&&site=%s", cfg.URL, msg, cfg.Model, cfg.Site)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: cfg.TimeOut * time.Second}
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("client.Do error: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll error: %v", err)
	}

	log.Printf("gpt response json: %s\n", string(body))

	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal error: %v", err)
	}

	gptResponseBody := strings.Replace(resp.Content, "\n\n", "\n \n", -1)
	return gptResponseBody, nil
}
