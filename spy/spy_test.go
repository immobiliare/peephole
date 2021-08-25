package spy

import (
	"testing"
)

func TestEndpoints(t *testing.T) {
	var (
		s = Spy{endpoints: map[string]string{"1": "1"}}
		e = s.Endpoints()
	)
	if len(e) != 1 {
		t.Errorf("Wrong endpoints size")
	}
	if e[0] != "1" {
		t.Errorf("Wrong endpoint value")
	}
}
