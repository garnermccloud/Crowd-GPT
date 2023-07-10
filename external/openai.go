package external

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type Outline struct {
	Section   int    `json:"section"`
	Details   string `json:"details"`
	WordCount int    `json:"word_count"`
}

const (
	ScriptSystemPrompt = ``
)

var openaiClient *openai.Client

// GetAIClient: Create or return an OpenAI client
func GetAIClient() *openai.Client {
	if openaiClient == nil {
		openaiClient = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	}
	return openaiClient
}

// AskAIGPT35Turbo:  Send a message to OpenAI client GPT 3.5 Turbo and return response string or error
func AskAIGPT35Turbo(message string, ctx context.Context) (string, error) {
	return AskAI(message, openai.GPT3Dot5Turbo, ctx)
}

func AskAIStreamGPT35Turbo(message string, ctx context.Context) (*openai.ChatCompletionStream, error) {
	return AskAIStream(message, openai.GPT3Dot5Turbo, ctx)
}

func AskAIConversationStreamGPT35Turbo(
	messages []openai.ChatCompletionMessage,
	ctx context.Context,
) (*openai.ChatCompletionStream, error) {
	return AskAIConversationStream(messages, openai.GPT3Dot5Turbo, ctx)
}

func AskAIGPT4(message string, ctx context.Context) (string, error) {
	return AskAI(message, openai.GPT4, ctx)
}

func AskAIStreamGPT4(message string, ctx context.Context) (*openai.ChatCompletionStream, error) {
	return AskAIStream(message, openai.GPT4, ctx)
}

func AskAIConversationStreamGPT4(
	messages []openai.ChatCompletionMessage,
	ctx context.Context,
) (*openai.ChatCompletionStream, error) {
	return AskAIConversationStream(messages, openai.GPT4, ctx)
}

// AskAI:  Send a message to OpenAI client and return response string or error
func AskAI(message string, Model string, ctx context.Context) (msg string, err error) {
	preMessage := "Say OK if you understand."
	client := GetAIClient()
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: preMessage,
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "OK",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("error creating chat completion: %w", err)
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

// AskAIConversation: Send a Messages conversation to OpenAI client and return response string or error
func AskAIConversation(messages []openai.ChatCompletionMessage, Model string) (string, error) {
	client := GetAIClient()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    Model,
			Messages: messages,
		},
	)
	if err != nil {
		return "", fmt.Errorf("error creating chat completion: %w", err)
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

// AskAIConversationStream: Send a Messages conversation to OpenAI client and return response string or error
func AskAIConversationStream(
	messages []openai.ChatCompletionMessage,
	Model string,
	ctx context.Context,
) (*openai.ChatCompletionStream, error) {
	client := GetAIClient()
	stream, err := client.CreateChatCompletionStream(
		ctx,
		openai.ChatCompletionRequest{
			Model:    Model,
			Messages: messages,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error creating chat completion: %w", err)
	}
	return stream, nil
}

// AskAI:  Send a message to OpenAI client and return response string or error
func AskAIStream(message string, Model string, ctx context.Context) (*openai.ChatCompletionStream, error) {
	client := GetAIClient()
	stream, err := client.CreateChatCompletionStream(
		ctx,
		openai.ChatCompletionRequest{
			Model: Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: ScriptSystemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("error creating chat completion: %w", err)
	}

	return stream, nil
}
