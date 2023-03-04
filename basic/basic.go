package basic

import (
	"chatgptSDK/constant"
)

type ChatRequest struct {
	Model    constant.ChatModel `json:"model"`
	Messages []*Message         `json:"messages"`
	N        int                `json:"n"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Id      string              `json:"id"`
	Object  string              `json:"object"`
	Created int                 `json:"created"`
	Choices []*MessageResponse  `json:"choices"`
	Usage   *TokenUsageResponse `json:"usage"`
}

type MessageResponse struct {
	Index        int `json:"index"`
	Message      *Message
	FinishReason string `json:"finish_reason"`
}

type TokenUsageResponse struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
