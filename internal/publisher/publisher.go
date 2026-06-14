package publisher

import (
	"context"
	"sync"

	"github.com/asif2772/campaign-spend-tracker/internal/model"
)

type Publisher struct {
	events chan model.SpendEvent
	wg     sync.WaitGroup
}

func New(bufferSize int) *Publisher { //constructor for publisher, e.x. "new Publisher()"
	return &Publisher{
		events: make(chan model.SpendEvent, bufferSize),
	}
}

func (p *Publisher) Publish(ctx context.Context, event model.SpendEvent) error {
	return nil
}

func (p *Publisher) Shutdown(ctx context.Context) error {
	return nil
}
