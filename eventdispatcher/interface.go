package eventdispatcher

type Event struct {
	// TODO: need design
	RawData      string `json:"raw_data,omitempty"`
	ID           string `json:"execution_id,omitempty"`
	PodName      string `json:"pod,omitempty"`
	PodNamespace string `json:"namespace,omitempty"`
	Behavior     string `json:"kernel_func,omitempty"`
	Time         string `json:"time,omitempty"`
}

type EventDispatcher interface {
	ListenEvent(ch chan Event)
}
