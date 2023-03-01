package eventdispatcher

type EventDispatcher interface {
	ListenEvent(ch chan string)
}
