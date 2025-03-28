package card_creator

import "context"

type Client interface {
	CreateCards(ctx context.Context) ([]Card, error)
}

func NewClient() Client {
	return &client{}
}

type client struct{}

func (c *client) CreateCards(ctx context.Context) ([]Card, error) {
	return []Card{
		{
			Question: "What is the capital of France?",
			Answer:   "Paris",
		},
		{
			Question: "What is the capital of Germany?",
			Answer:   "Berlin",
		},
		{
			Question: "What is the capital of Italy?",
			Answer:   "Rome",
		},
	}, nil
}
