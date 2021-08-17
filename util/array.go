package util

import "strings"

func HasAnyPrefix(array []string, prefix string) bool {
	for _, e := range array {
		if strings.HasPrefix(e, prefix) {
			return true
		}
	}
	return false
}
