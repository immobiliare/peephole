package util

import (
	"testing"
)

func TestYearRetention(t *testing.T) {
	secs, err := RetentionSeconds("1y")
	if err != nil {
		t.Errorf("Unable to calculate retention")
	}

	if secs != uint32(60*60*24*365) {
		t.Errorf("Incorrect retention")
	}
}

func TestMonthRetention(t *testing.T) {
	secs, err := RetentionSeconds("1M")
	if err != nil {
		t.Errorf("Unable to calculate retention")
	}

	if secs != uint32(60*60*24*31) {
		t.Errorf("Incorrect retention")
	}
}

func TestDayRetention(t *testing.T) {
	secs, err := RetentionSeconds("1d")
	if err != nil {
		t.Errorf("Unable to calculate retention")
	}

	if secs != uint32(60*60*24) {
		t.Errorf("Incorrect retention")
	}
}

func TestHourRetention(t *testing.T) {
	secs, err := RetentionSeconds("1h")
	if err != nil {
		t.Errorf("Unable to calculate retention")
	}

	if secs != uint32(60*60) {
		t.Errorf("Incorrect retention")
	}
}

func TestMinuteRetention(t *testing.T) {
	secs, err := RetentionSeconds("1m")
	if err != nil {
		t.Errorf("Unable to calculate retention")
	}

	if secs != uint32(60) {
		t.Errorf("Incorrect retention")
	}
}

func TestSecondRetention(t *testing.T) {
	secs, err := RetentionSeconds("1s")
	if err != nil {
		t.Errorf("Unable to calculate retention")
	}

	if secs != uint32(1) {
		t.Errorf("Incorrect retention")
	}
}

func TestMalformedRetention(t *testing.T) {
	if _, err := RetentionSeconds("1"); err == nil {
		t.Errorf("Malformed retention is supposed to return an error")
	}
}

func TestUknownUnit(t *testing.T) {
	if _, err := Unit("1k"); err == nil {
		t.Errorf("Uknown unit is supposed to return an error")
	}
}

func TestEmptyInterval(t *testing.T) {
	if _, err := Unit(""); err == nil {
		t.Errorf("Empty interval is supposed to return an error")
	}
}
