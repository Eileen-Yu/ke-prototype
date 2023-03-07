package tetragon

type TetragonEvent struct {
	EventType string // "process_kprobe", "process_exec"
	Process   TetragonProcess
	Time      string
}

type TetragonProcess struct {
	ExecutionID  string
	Pod          TetragonProcessPod
	FunctionName string
}

type TetragonProcessPod struct {
	Namespace string
	Name      string
}

// for json decode
type TetragonLogLineType struct {
	NodeName      string            `json:"node_name,omitempty"`
	Time          string            `json:"time,omitempty"`
	ProcessKProbe ProcessKProbeType `json:"process_kprobe,omitempty"`

	// ProcessExec  ProcessExecType `json:"process_exec,omitempty"`
}

type ProcessKProbeType struct {
	Process      ProcessType `json:"process,omitempty"`
	FunctionName string      `json:"function_name,omitempty"`
}

type ProcessExecType struct {
	Process ProcessType `json:"process,omitempty"`
}

type ProcessType struct {
	ExecutionID string  `json:"exec_id,omitempty"`
	Pod         PodType `json:"pod,omitempty"`
}

type PodType struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}
