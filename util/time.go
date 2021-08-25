package util

import (
	"fmt"
	"strconv"
	"strings"
)

func RetentionSeconds(interval string) (uint32, error) {
	unit, err := Unit(interval)
	if err != nil {
		return 0, err
	}

	secs, err := strconv.ParseUint(strings.ReplaceAll(interval, string(unit), ""), 10, 32)
	if err != nil {
		return 0, err
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

	return uint32(secs), nil
}

func Unit(interval string) (rune, error) {
	if len(interval) <= 1 {
		return '0', fmt.Errorf("illegal interval")
	}

	r := rune(interval[len(interval)-1:][0])
	switch r {
	case 'y', 'M', 'd', 'h', 'm', 's':
		return r, nil
	default:
		return '0', fmt.Errorf("unsupported unit")
	}
}
