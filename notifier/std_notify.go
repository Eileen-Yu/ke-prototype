package notifier

import (
	"encoding/json"
	"fmt"

	ed "eileenyu.io/ebpf-prototype/eventdispatcher"
	pc "eileenyu.io/ebpf-prototype/policy"
)

type StdNotifier struct{}

type PrettyEvent struct {
	ID           string `json:"execution_id,omitempty"`
	PodName      string `json:"pod,omitempty"`
	PodNamespace string `json:"namespace,omitempty"`
	Behavior     string `json:"kernel_func,omitempty"`
	Time         string `json:"time,omitempty"`
	Policy       string `json:"policy,omitempty"`
}

func prettify(event ed.Event, policy pc.Policy) string {
	prettyEvent := PrettyEvent{
		ID:           event.ID,
		PodName:      event.PodName,
		PodNamespace: event.PodNamespace,
		Behavior:     event.Behavior,
		Time:         event.Time,
		Policy:       policy.Name,
	}

	prettyStr, _ := json.MarshalIndent(prettyEvent, "", " ")

	return string(prettyStr)

}

func (n *StdNotifier) Notify(event ed.Event, policy pc.Policy) {
	fmt.Println(prettify(event, policy))
}
