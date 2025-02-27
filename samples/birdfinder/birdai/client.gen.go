// Code generated by hardc. DO NOT EDIT.

package birdai

import (
	"context"

	"github.com/be-brite/hardconversations/chat"
	"github.com/be-brite/hardconversations/samples/birdfinder/bird"
	
)


const instruction = `Given a piece of text, you are able to determine how many birds are mentioned in the text and describe each bird.
`

type Client struct {
	*chat.Client
}

func NewClient(openAIKey string) *Client {
	c := &Client{
		Client: chat.NewClient(openAIKey, instruction),
	}

	return c
}

type Thread struct {
	*chat.Thread
}

func (c *Client) NewThread() *Thread {
	return &Thread{
		Thread: c.Client.NewThread(),
	}
}



// TODO: handle different input and output types, arrays, structs, etc
func (t *Thread) CountBirds(ctx context.Context, input string) (result int, md chat.Metadata, err error) {
	const prompt = `How many birds are mentioned in the text?` // TODO initialize text embedding

	parseInstruction, err := chat.ParseInstruction(result)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	fullPrompt := parseInstruction + prompt
	inputStr, err := chat.ConvertInput(input)
	if err != nil {
		return result, chat.Metadata{}, err
	}
	fullPrompt += "\n" + inputStr
	

	output, md, err := t.Thread.ExecutePrompt(ctx, fullPrompt)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	err = chat.Parse(output, &result)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	return result, md, nil
}



// TODO: handle different input and output types, arrays, structs, etc
func (t *Thread) ParseBird(ctx context.Context) (result []bird.Bird, md chat.Metadata, err error) {
	const prompt = `Can you parse the details of each bird?` // TODO initialize text embedding

	parseInstruction, err := chat.ParseInstruction(result)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	fullPrompt := parseInstruction + prompt

	output, md, err := t.Thread.ExecutePrompt(ctx, fullPrompt)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	err = chat.Parse(output, &result)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	return result, md, nil
}



// TODO: handle different input and output types, arrays, structs, etc
func (t *Thread) DescribeBird(ctx context.Context, input bird.Bird) (result string, md chat.Metadata, err error) {
	const prompt = `Describe the bird with the given properties and add a fun fact (make it up if you have to)` // TODO initialize text embedding

	parseInstruction, err := chat.ParseInstruction(result)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	fullPrompt := parseInstruction + prompt
	inputStr, err := chat.ConvertInput(input)
	if err != nil {
		return result, chat.Metadata{}, err
	}
	fullPrompt += "\n" + inputStr
	

	output, md, err := t.Thread.ExecutePrompt(ctx, fullPrompt)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	err = chat.Parse(output, &result)
	if err != nil {
		return result, chat.Metadata{}, err
	}

	return result, md, nil
}

