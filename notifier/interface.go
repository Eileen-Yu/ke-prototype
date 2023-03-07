package notifier

import (
	ed "eileenyu.io/ebpf-prototype/eventdispatcher"
	pc "eileenyu.io/ebpf-prototype/policy"
)

type Notifier interface {
	Notify(event ed.Event, policy pc.Policy)
}
