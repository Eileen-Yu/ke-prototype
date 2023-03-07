package matcher

import (
	ed "eileenyu.io/ebpf-prototype/eventdispatcher"
	pc "eileenyu.io/ebpf-prototype/policy"
)

type Matcher interface {
	Match(event ed.Event, policy pc.Policy) bool
}
