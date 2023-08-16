package chat

import (
	"fmt"

	"github.com/be-brite/hardconversations/internal/tokens"
	"github.com/be-brite/hardconversations/sources"
	gogpt "github.com/sashabaranov/go-openai"
)

type Client struct {
	ai *gogpt.Client

	*Thread // global start thread; threads are spun off from this one
}

func NewClient(openAIKey string, instruction string, opt ...ConfigOption) *Client {
	openAIClient := gogpt.NewClient(openAIKey)

	systemMessage := fmt.Sprintf(baseSystemMessage, instruction)

	return &Client{
		ai: openAIClient,
		Thread: &Thread{
			ai:                  openAIClient,
			config:              NewConfig(opt...),
			systemMessage:       systemMessage,
			systemMessageTokens: tokens.MustCount(systemMessage),
			Manager:             sources.New(openAIClient),
		},
	}
}
