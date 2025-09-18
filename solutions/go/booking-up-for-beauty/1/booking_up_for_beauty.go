package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic("Couldn't parse the given string")
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	parsed, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic("Couldn't parse the given string")
	}
	return parsed.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic("Couldn't parse the given string")
	}
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointment := Schedule(date)
	formattedTime := appointment.Format("Monday, January 2, 2006, at 15:04.")
	return "You have an appointment on " + formattedTime
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
