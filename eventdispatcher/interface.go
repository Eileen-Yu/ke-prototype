package eventdispatcher

type Event struct {
	// TODO: need design
	RawData  string
	Resource any
	Behavior string
}

type EventDispatcher interface {
	ListenEvent(ch chan Event)
}
