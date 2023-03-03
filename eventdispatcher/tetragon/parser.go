package tetragon

import (
	"encoding/json"
	"fmt"

	"eileenyu.io/ebpf-prototype/eventdispatcher"
)

func parseTetragonLogLine(rawData string) TetragonEvent {
	// TODO
	out := &TetragonLogLineType{}
	// parse to map
	json.Unmarshal([]byte(rawData), out)
	/*
			fmt.Println(out.NodeName)
			fmt.Println(out.Time)
			fmt.Println(out.ProcessKProbe.Process.Pod.Namespace)
			fmt.Println(out.ProcessKProbe.Process.Pod.Name)

		formatStr, _ := json.MarshalIndent(event, "", "  ")
		fmt.Println(string(formatStr))
	*/

	fmt.Println(out.ProcessKProbe.FunctionName)

	return GetTetragonEventFromRawData(*out)
}

func normalizeTetragonEvent(event TetragonEvent, originLogLine string) eventdispatcher.Event {

	// TODO
	return eventdispatcher.Event{
		RawData:  originLogLine,
		Behavior: event.Process.FunctionName,
	}

}
