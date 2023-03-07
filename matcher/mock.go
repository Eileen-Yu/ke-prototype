package matcher

import (
	ed "eileenyu.io/ebpf-prototype/eventdispatcher"
	pc "eileenyu.io/ebpf-prototype/policy"
)

type MockMatcher struct{}

func (mm *MockMatcher) Match(event ed.Event, policy pc.Policy) bool {
	return event.Behavior == policy.Behavior
}
