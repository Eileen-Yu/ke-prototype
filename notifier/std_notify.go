package notifier

import (
	pc "eileenyu.io/ebpf-prototype/policy"
	"fmt"
)

type StdNotifier struct{}

func (n *StdNotifier) Notify(event string, policy pc.Policy) {

	fmt.Printf("Event: %s\nPolicy: %s\n", event, policy)

}
