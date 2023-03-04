package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"chatgptSDK/basic"
	"chatgptSDK/constant"
)

type Chat struct {
	APIToken  string
	Model     constant.ChatModel
	ChoiceNum int
	client    *http.Client
}

func NewChat(model constant.ChatModel, apiToken string, choiceNum int) *Chat {
	if choiceNum <= 0 {
		choiceNum = 1
	}
	return &Chat{
		Model:     model,
		APIToken:  apiToken,
		ChoiceNum: choiceNum,
		client: &http.Client{
			Transport: nil,
			Timeout:   30 * time.Second,
		},
	}
}

func (c *Chat) ChatWithMessage(messages ...*basic.Message) ([]*basic.Message, int, error) {
	req := &basic.ChatRequest{
		Model:    c.Model,
		Messages: messages,
		N:        c.ChoiceNum,
	}
	resp, err := c.DoChat(req)
	if err != nil {
		return nil, 0, fmt.Errorf("do chat with openai fail: %s", err)
	}

	var allMessage []*basic.Message
	for _, msg := range resp.Choices {
		allMessage = append(allMessage, msg.Message)
	}
	return allMessage, resp.Usage.TotalTokens, nil
}

func (c *Chat) DoChat(req *basic.ChatRequest) (*basic.ChatResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal chat request fail: %s", err)
	}
	content, err := c.doRequest("POST", constant.OpenAIHost+"/v1/chat/completions", nil, b)
	if err != nil {
		return nil, fmt.Errorf("get chat repsonse from openai fail: %s", err)
	}

	var resp basic.ChatResponse
	err = json.Unmarshal(content, &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal chat response fail: %s", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("openai response no choice")
	}

	return &resp, nil
}

func (c *Chat) doRequest(method string, url string, header map[string]string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("create request fail: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIToken))
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request fail: %s", err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body fail: %s", err)
	}

	return content, nil
}

func (c *Chat) PlaySystemRole(content string) *basic.Message {
	return c.generateMessage(constant.SystemRole, content)
}

func (c *Chat) PlayAssistRole(content string) *basic.Message {
	return c.generateMessage(constant.AssistRole, content)
}

func (c *Chat) PlayUserRole(content string) *basic.Message {
	return c.generateMessage(constant.UserRole, content)
}

func (c *Chat) generateMessage(role constant.ChatRole, message string) *basic.Message {
	return &basic.Message{
		Role:    string(role),
		Content: message,
	}
}
