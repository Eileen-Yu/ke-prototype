package matcher

import (
	pc "eileenyu.io/ebpf-prototype/policy"
)

type Matcher interface {
	Match(event string, policy pc.Policy) bool
}
