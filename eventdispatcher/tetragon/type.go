package tetragon

type TetragonEvent struct {
	EventType string // "process_kprob", "process_exec"
	Process   TetragonProcess
}

type TetragonProcess struct {
	ExecutionID string `json:exec_id,emitempty`
	// TODO
	// ...
	Pod TetragonProcessPod
}

type TetragonProcessPod struct {
	namespace string
	name      string
}
