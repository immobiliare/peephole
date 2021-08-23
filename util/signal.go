package util

import (
	"os"
	"os/signal"
)

// TrapSignal provides a simple signal subscription flow
func TrapSignal(s os.Signal, f func()) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, s)
	go func() {
		for range channel {
			f()
		}
	}()
}
