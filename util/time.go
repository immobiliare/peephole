package util

import (
	"strconv"
	"strings"
)

func RetentionSeconds(interval string) uint32 {
	unit := Unit(interval)
	secs, err := strconv.ParseUint(strings.ReplaceAll(interval, string(unit), ""), 10, 32)
	if err != nil {
		return 0
	}

	switch unit {
	case 'y': // years
		secs *= 365
		secs *= 24
		secs *= 60
		secs *= 60
	case 'M': // months
		secs *= 31
		secs *= 24
		secs *= 60
		secs *= 60
	case 'd': // days
		secs *= 24
		secs *= 60
		secs *= 60
	case 'h': // hours
		secs *= 60
		secs *= 60
	case 'm': // minutes
		secs *= 60
	default: // seconds
	}

	return uint32(secs)
}

func Unit(interval string) rune {
	return []rune(interval[len(interval)-1:])[0]
}
