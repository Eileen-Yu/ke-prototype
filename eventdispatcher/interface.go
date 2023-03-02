package eventdispatcher

type Event struct {
	// TODO: need design
	rawData  string
	resource any
	behavior string
}

type EventDispatcher interface {
	ListenEvent(ch chan Event)
}
