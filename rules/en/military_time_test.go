package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

func TestMilitaryTime(t *testing.T) {
	fixt := []Fixture{
		{"meeting at 0800", 11, "0800", 8 * time.Hour},
		{"call at 1430", 8, "1430", (14 * time.Hour) + (30 * time.Minute)},
		{"standup at 0900", 11, "0900", 9 * time.Hour},
		{"0000", 0, "0000", 0},
		{"2359", 0, "2359", (23 * time.Hour) + (59 * time.Minute)},
		{"wake up 0630", 8, "0630", (6 * time.Hour) + (30 * time.Minute)},
		{"1200 lunch", 0, "1200", 12 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.MilitaryTime(rules.Override))

	ApplyFixtures(t, "en.MilitaryTime", w, fixt)
}

func TestMilitaryTimeNil(t *testing.T) {
	fixt := []Fixture{
		{"2500", 0, "", 0},  // Invalid hour
		{"0861", 0, "", 0},  // Invalid minute
		{"123", 0, "", 0},   // Only 3 digits
		{"12345", 0, "", 0}, // 5 digits
	}

	w := when.New(nil)
	w.Add(en.MilitaryTime(rules.Override))

	ApplyFixturesNil(t, "en.MilitaryTime nil", w, fixt)
}
