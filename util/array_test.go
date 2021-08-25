package util

import (
	"testing"
)

func TestTrueHasAnyPrefix(t *testing.T) {
	var (
		entries = []string{"entry1", "entry2"}
		entry   = entries[len(entries)-1]
	)
	if !HasAnyPrefix(entries, entry) {
		t.Errorf("Unexpected return value")
	}
}

func TestFalseHasAnyPrefix(t *testing.T) {
	var (
		entries = []string{"entry1", "entry2"}
	)
	if HasAnyPrefix(entries, "entry3") {
		t.Errorf("Unexpected return value")
	}
}

func TestEmptyHasAnyPrefix(t *testing.T) {
	var (
		entries = []string{}
	)
	if HasAnyPrefix(entries, "entry") {
		t.Errorf("Unexpected return value")
	}
}
