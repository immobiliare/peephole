package util

import (
	"os"
	"syscall"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	c := make(chan bool)
	TrapSignal(syscall.SIGHUP, func() {
		c <- true
	})

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Errorf("Unable to get my PID")
	}

	err = p.Signal(syscall.SIGHUP)
	if err != nil {
		t.Errorf("Unable send signal")
	}

	go func() {
		time.Sleep(1 * time.Second)
		c <- false
	}()

	if !<-c {
		t.Errorf("Signal not trapped")
	}
}
