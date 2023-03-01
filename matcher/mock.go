package matcher

import (
	pc "eileenyu.io/ebpf-prototype/policy"
	"strings"
)

type MockMatcher struct{}

func (mm *MockMatcher) Match(event string, policy pc.Policy) bool {
	return strings.Contains(string(policy), event)
}
