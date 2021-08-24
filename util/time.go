package util

import (
	"fmt"
	"strings"
	"time"
)

func ToNow(interval string) time.Time {
	var (
		unit        = Unit(interval)
		placeholder = fmt.Sprintf("-%ss", strings.Trim(interval, string(unit)))
	)

	d, err := time.ParseDuration(placeholder)
	if err != nil {
		return time.Date(2005, 0, 0, 0, 0, 0, 0, time.UTC)
	}

	switch unit {
	case 'y': // years
		d *= 365
		d *= 24
		d *= 60
		d *= 60
	case 'M': // months
		d *= 31
		d *= 24
		d *= 60
		d *= 60
	case 'd': // days
		d *= 24
		d *= 60
		d *= 60
	case 'h': // hours
		d *= 60
		d *= 60
	case 'm': // minutes
		d *= 60
	default: // seconds
	}

	return time.Now().Add(d)
}

func Unit(interval string) rune {
	return []rune(interval[len(interval)-1:])[0]
}
