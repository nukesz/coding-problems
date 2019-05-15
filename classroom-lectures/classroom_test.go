package classroom

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	intervals := []Time{}
	assertRooms(t, 0, CountRequiredRooms(intervals))
}

func TestOneInterval(t *testing.T) {
	intervals := []Time{
		Time{0, 50},
	}
	assertRooms(t, 1, CountRequiredRooms(intervals))
}

func TestTwoIntervals(t *testing.T) {
	intervals := []Time{
		Time{0, 50},
		Time{60, 90},
	}
	assertRooms(t, 1, CountRequiredRooms(intervals))
}

func assertRooms(t *testing.T, expected, actual int) {
	if actual != 0 {
		t.Errorf("Expected %d rooms, but got: %d", expected, actual)
	}
}
