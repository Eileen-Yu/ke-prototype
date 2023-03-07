package tetragon

import (
	"encoding/json"

	"eileenyu.io/ebpf-prototype/eventdispatcher"
)

func parseTetragonLogLine(rawData string) TetragonEvent {
	out := &TetragonLogLineType{}
	json.Unmarshal([]byte(rawData), out)
	return GetTetragonEventFromRawData(*out)
}

func normalizeTetragonEvent(event TetragonEvent, originLogLine string) eventdispatcher.Event {
	return eventdispatcher.Event{
		RawData:      originLogLine,
		ID:           event.Process.ExecutionID,
		PodName:      event.Process.Pod.Name,
		PodNamespace: event.Process.Pod.Namespace,
		Behavior:     event.Process.FunctionName,
		Time:         event.Time,
	}
}
