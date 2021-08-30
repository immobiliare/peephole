package util

import (
	"os"
	"testing"
)

func TestDebugTrue(t *testing.T) {
	os.Setenv("DEBUG", "1")
	if !Debugging() {
		t.Errorf("Debugging should return true")
	}
}

func TestDebug(t *testing.T) {
	if !Debugging() {
		t.Errorf("Debugging should return false")
	}
}
