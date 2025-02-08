package events

import (
	"sync"
	"time"
)

type EventHandlerSpy struct {
	Event EventInterface
}

type EventSpy struct {
	GetNameResponse     string
	GetDateTimeResponse time.Time
	GetPayloadResponse  any
}

func (e EventSpy) GetName() string {
	return e.GetNameResponse
}

func (e EventSpy) GetDateTime() time.Time {
	return e.GetDateTimeResponse
}

func (e EventSpy) GetPayload() any {
	return e.GetPayloadResponse
}

func (eh EventHandlerSpy) Handle(event EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
}

// type EventInterface interface {
// 	GetName() string
// 	GetDateTime() time.Time
// 	GetPayload() any
// }

// type EventHandlerInterface interface {
// 	Handle(event EventInterface, wg *sync.WaitGroup)
// }

// type EventDispatcherInterface interface {
// 	Register(ctx context.Context, event string, handler EventHandlerInterface) error
// 	Dispatch(ctx context.Context, event EventInterface) error
// 	Remove(ctx context.Context, event string, handler EventHandlerInterface) error
// 	Has(ctx context.Context, event string, handler EventHandlerInterface) bool
// 	Clear(ctx context.Context) error
// }
