package util

import "os"

const DebugKey = "DEBUG"

func Debugging() bool {
	return HasAnyPrefix(os.Environ(), "DEBUG=1")
}
