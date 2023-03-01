package main

import (
	ed "eileenyu.io/ebpf-prototype/eventdispatcher"
	mm "eileenyu.io/ebpf-prototype/matcher"
	"eileenyu.io/ebpf-prototype/notifier"
	pc "eileenyu.io/ebpf-prototype/policy"
)

func eventHandler(ch chan string, eventReceiver ed.EventDispatcher, policyReceiver pc.PolicyCache, matcher mm.Matcher, nt notifier.Notifier) {
	policies := policyReceiver.ListPolicies()

	go eventReceiver.ListenEvent(ch)

	for {
		event := <-ch

		for _, policy := range *policies {
			if !matcher.Match(event, policy) {
				continue
			}

			nt.Notify(event, policy)
		}
	}
}

func main() {
	tetragonEventDispatcher := ed.TetragonEventDispatcher{}
	policyReceiver := pc.PolicySimulator{}
	matcher := mm.MockMatcher{}
	stdNotifier := notifier.StdNotifier{}

	ch := make(chan string)

	eventHandler(ch, &tetragonEventDispatcher, &policyReceiver, &matcher, &stdNotifier)
}
