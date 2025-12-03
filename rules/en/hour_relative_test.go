package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

func TestHourRelativeTo(t *testing.T) {
	fixt := []Fixture{
		{"10 to 8", 0, "10 to 8", (7 * time.Hour) + (50 * time.Minute)},
		{"10 past 2", 0, "10 past 2", (2 * time.Hour) + (10 * time.Minute)},
		{"half past 2", 0, "half past 2", (2 * time.Hour) + (30 * time.Minute)},
		{"quarter to 8", 0, "quarter to 8", (7 * time.Hour) + (45 * time.Minute)},
		{"quarter past 3", 0, "quarter past 3", (3 * time.Hour) + (15 * time.Minute)},
		{"meet at 10 to 8", 8, "10 to 8", (7 * time.Hour) + (50 * time.Minute)},
		{"5 past 11", 0, "5 past 11", (11 * time.Hour) + (5 * time.Minute)},
		{"20 to 4", 0, "20 to 4", (3 * time.Hour) + (40 * time.Minute)},
	}

	w := when.New(nil)
	w.Add(en.HourRelativeTo(rules.Override))

	ApplyFixtures(t, "en.HourRelativeTo", w, fixt)
}

func TestHourInPeriod(t *testing.T) {
	fixt := []Fixture{
		{"6 in the morning", 0, "6 in the morning", 6 * time.Hour},
		{"7 in the evening", 0, "7 in the evening", 19 * time.Hour},
		{"3 in the afternoon", 0, "3 in the afternoon", 15 * time.Hour},
		{"meet at 8 in the morning", 8, "8 in the morning", 8 * time.Hour},
		{"call at 5 in the evening", 8, "5 in the evening", 17 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.HourInPeriod(rules.Override))

	ApplyFixtures(t, "en.HourInPeriod", w, fixt)
}

func TestHourOClock(t *testing.T) {
	fixt := []Fixture{
		{"11 o'clock", 0, "11 o'clock", 11 * time.Hour},
		{"5 o'clock", 0, "5 o'clock", 5 * time.Hour},
		{"eleven o'clock", 0, "eleven o'clock", 11 * time.Hour},
		{"meet at 3 o'clock", 8, "3 o'clock", 3 * time.Hour},
		{"at twelve o'clock", 3, "twelve o'clock", 12 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.HourOClock(rules.Override))

	ApplyFixtures(t, "en.HourOClock", w, fixt)
}
