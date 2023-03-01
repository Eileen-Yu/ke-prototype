package notifier

import (
	pc "eileenyu.io/ebpf-prototype/policy"
)

type Notifier interface {
	Notify(event string, policy pc.Policy)
}
