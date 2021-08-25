package util

import (
	"testing"
)

type obj struct {
	Key string
}

func TestMarshaling(t *testing.T) {
	var (
		obj1 = obj{"value"}
		obj2 = new(obj)
	)

	bytes, err := Marshal(obj1)
	if err != nil {
		t.Errorf("Unable to marshal given object")
	}

	if err := Unmarshal(bytes, obj2); err != nil {
		t.Errorf("Unable to unmarshal bytes into object")
	}

	if obj1.Key != obj2.Key {
		t.Errorf("Unmarshalled object differs from original one")
	}
}
