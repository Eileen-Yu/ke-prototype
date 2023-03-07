package tetragon

func GetTetragonEventFromRawData(rawData TetragonLogLineType) TetragonEvent {

	event := TetragonEvent{
		// TODO
		EventType: "process_kprobe",
		Process: TetragonProcess{
			ExecutionID: rawData.ProcessKProbe.Process.ExecutionID,
			Pod: TetragonProcessPod{
				Namespace: rawData.ProcessKProbe.Process.Pod.Namespace,
				Name:      rawData.ProcessKProbe.Process.Pod.Name,
			},
			FunctionName: rawData.ProcessKProbe.FunctionName,
		},
		Time: rawData.Time,
	}

	return event
}
