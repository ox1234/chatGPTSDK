package constant

type ChatModel string

const (
	GPT35TurboModel     ChatModel = "gpt-3.5-turbo"
	GPT35Turbo0301Model ChatModel = "gpt-3.5-turbo-0301"
)

type ChatRole string

const (
	SystemRole ChatRole = "system"
	UserRole   ChatRole = "user"
	AssistRole ChatRole = "assistant"
)

const OpenAIHost = "https://api.openai.com"
