package publisher

import (
	"context"
	"encoding/json"
	"fmt"
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

func (p *Publisher) Start() {
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()

		for event := range p.events {
			data, err := json.Marshal(event)
			if err != nil {
				fmt.Printf("failed to marshal event: %v\n", err)
				continue
			}

			fmt.Println(string(data))
		}
	}()
}

func (p *Publisher) Publish(ctx context.Context, event model.SpendEvent) error {
	select {
	case p.events <- event:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func (p *Publisher) Shutdown(ctx context.Context) error {

	close(p.events)

	done := make(chan struct{})

	go func() {
		p.wg.Wait()
		close(done)
	}()

	select {

	case <-done:
		return nil

	case <-ctx.Done():
		return ctx.Err()

	}

}
