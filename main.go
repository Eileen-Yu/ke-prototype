package main

import (
	"fmt"

	ed "eileenyu.io/ebpf-prototype/eventdispatcher"
	tetragon "eileenyu.io/ebpf-prototype/eventdispatcher/tetragon"

	mm "eileenyu.io/ebpf-prototype/matcher"
	"eileenyu.io/ebpf-prototype/notifier"
	pc "eileenyu.io/ebpf-prototype/policy"
)

func eventHandler(ch chan ed.Event, eventReceiver ed.EventDispatcher, policyReceiver pc.PolicyCache, matcher mm.Matcher, nt notifier.Notifier) {
	// policies := policyReceiver.ListPolicies()

	go eventReceiver.ListenEvent(ch)

	for {
		event := <-ch
		fmt.Println(event)

		/*
			for _, policy := range *policies {
				if !matcher.Match(event, policy) {
					continue
				}

				nt.Notify(event, policy)
			}
		*/
	}
}

func main() {
	tetragonEventDispatcher := tetragon.TetragonEventDispatcher{}
	policyReceiver := pc.PolicySimulator{}
	matcher := mm.MockMatcher{}
	stdNotifier := notifier.StdNotifier{}

	ch := make(chan ed.Event)

	eventHandler(ch, &tetragonEventDispatcher, &policyReceiver, &matcher, &stdNotifier)
}
