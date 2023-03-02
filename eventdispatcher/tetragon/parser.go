package tetragon

import (
	"eileenyu.io/ebpf-prototype/eventdispatcher"
)

func parseTetragonLogLine(rawData string) []TetragonEvent {
	// TODO
	return []TetragonEvent{TetragonEvent{}}
}

func normalizeTetragonEvent(event TetragonEvent) eventdispatcher.Event {

	// TODO
	return eventdispatcher.Event{}

}
