package tdlib

import (
	"sync"

	"github.com/jjeejj/go-tdlib/entities"
	"github.com/jjeejj/go-tdlib/incomingevents"
)

type Handlers struct {
	rawIncomingEvent           func(event []byte)
	incomingEvent              func(event incomingevents.Event)
	onUpdateConnectionState    func(newState entities.ConnectionStateType)
	onUpdateAuthorizationState func(newState entities.AuthorizationStateType)

	eventTypeHandlerLocker sync.Mutex
	eventTypeHandlers      map[string]event
}

func NewHandlers() *Handlers {
	return &Handlers{eventTypeHandlers: map[string]event{}}
}

func NewEventHandler[T any](handler func(data *T)) Event[T] {
	return Event[T]{handler: handler}
}

// SetRawIncomingEventHandler 事件的原始数据，没有经过格式化的.
func (h *Handlers) SetRawIncomingEventHandler(fn func(eventBytes []byte)) *Handlers {
	h.rawIncomingEvent = fn
	return h
}

// SetIncomingEventHandler 根据格式化后的事件类型处理对应的事件.
func (h *Handlers) SetIncomingEventHandler(fn func(event incomingevents.Event)) *Handlers {
	h.incomingEvent = fn
	return h
}

func (h *Handlers) SetOnUpdateConnectionStateEventHandler(fn func(newState entities.ConnectionStateType)) *Handlers {
	h.onUpdateConnectionState = fn
	return h
}

func (h *Handlers) SetOnUpdateAuthorizationStateEventHandler(fn func(newState entities.AuthorizationStateType)) *Handlers {
	h.onUpdateAuthorizationState = fn
	return h
}

func (h *Handlers) AddEventTypeHandler(eventType string, e event) *Handlers {
	h.eventTypeHandlerLocker.Lock()
	defer h.eventTypeHandlerLocker.Unlock()

	h.eventTypeHandlers[eventType] = e

	return h
}
