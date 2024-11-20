package refy

type EventBus[T any] struct {
	Handlers map[string][]T
}

func NewEventBus[T any]() *EventBus[T] {
	return &EventBus[T]{
		Handlers: map[string][]T{},
	}
}

func (e *EventBus[T]) On(id string, f T) {
	e.Handlers[id] = append(e.Handlers[id], f)
}

func (e *EventBus[T]) Trigger(id string, args ...any) {
	for _, v := range e.Handlers[id] {

	}
}
