package eventdispatcher

import "time"

type TetragonEventDispatcher struct{}

func (ed *TetragonEventDispatcher) ListenEvent(ch chan string) {
	// TODO: Read tetragon log lines
	for {
		time.Sleep(100 * time.Millisecond)
		val := "😉"
		ch <- val
	}
}
