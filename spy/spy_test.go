package spy

import (
	"testing"
)

func TestEndpoints(t *testing.T) {
	var (
		s = Spy{endpoints: map[string]*endpoint{"api": {"user", "pass", "client", "token"}}}
		e = s.Endpoints()
	)
	if len(e) != 1 {
		t.Errorf("Wrong endpoints size")
	}
	if e[0] != "api" {
		t.Errorf("Wrong endpoint value")
	}
}
