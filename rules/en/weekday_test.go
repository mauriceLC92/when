package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

func TestWeekday(t *testing.T) {
	// current is Wednesday (Jan 6, 2016 00:00)
	fixt := []Fixture{
		// past/last - now includes default 09:00 time per spec
		{"do it for the past Monday", 14, "past Monday", -(2*24*time.Hour - 9*time.Hour)},        // -2d + 9h = -39h
		{"past saturday", 0, "past saturday", -(4*24*time.Hour - 9*time.Hour)},                   // -4d + 9h = -87h
		{"past friday", 0, "past friday", -(5*24*time.Hour - 9*time.Hour)},                       // -5d + 9h = -111h
		{"past wednesday", 0, "past wednesday", -(7*24*time.Hour - 9*time.Hour)},                 // -7d + 9h = -159h
		{"past tuesday", 0, "past tuesday", -(24*time.Hour - 9*time.Hour)},                       // -1d + 9h = -15h
		// next - now includes default 09:00 time per spec
		{"next tuesday", 0, "next tuesday", (6*24 + 9) * time.Hour},                              // 6d + 9h = 153h
		{"drop me a line at next wednesday", 18, "next wednesday", (7*24 + 9) * time.Hour},       // 7d + 9h = 177h
		{"next saturday", 0, "next saturday", (3*24 + 9) * time.Hour},                            // 3d + 9h = 81h
		// this - now includes default 09:00 time per spec
		{"this tuesday", 0, "this tuesday", -(24*time.Hour - 9*time.Hour)},                       // -1d + 9h = -15h
		{"drop me a line at this wednesday", 18, "this wednesday", 9 * time.Hour},                // same day, 9h offset
		{"this saturday", 0, "this saturday", (3*24 + 9) * time.Hour},                            // 3d + 9h = 81h
	}

	w := when.New(nil)

	w.Add(en.Weekday(rules.Override))

	ApplyFixtures(t, "en.Weekday", w, fixt)
}
