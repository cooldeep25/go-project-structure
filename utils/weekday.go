package utils

import "time"

func GetNextComingWeekday(rday string, t time.Time) time.Time {
	for t.Weekday().String() != rday {
		t = t.Add(24 * time.Hour)
	}
	return t
}
