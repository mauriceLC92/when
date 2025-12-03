package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

// Reference time: January 6, 2016 (Wednesday) at 00:00:00 UTC
var specNull = time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC)

// Fixture for test cases
type SpecFixture struct {
	Text   string        // Input text to parse
	Index  int           // Expected starting index of matched phrase
	Phrase string        // Expected matched text
	Diff   time.Duration // Expected time difference from reference
}

// Helper functions matching the existing test pattern
func applySpecFixtures(t *testing.T, name string, w *when.Parser, fixt []SpecFixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, specNull)
		require.Nil(t, err, "[%s] err #%d (%s)", name, i, f.Text)
		require.NotNil(t, res, "[%s] res #%d (%s) - parser returned nil, expected match", name, i, f.Text)
		if res != nil {
			require.Equal(t, f.Index, res.Index, "[%s] index #%d (%s)", name, i, f.Text)
			require.Equal(t, f.Phrase, res.Text, "[%s] text #%d (%s)", name, i, f.Text)
			require.Equal(t, f.Diff, res.Time.Sub(specNull), "[%s] diff #%d (%s) - got %v, expected %v", name, i, f.Text, res.Time.Sub(specNull), f.Diff)
		}
	}
}

// ==============================================================================
// CATEGORY A: Today / Later Today
// ==============================================================================

func TestUpcomingSpec_TodayWithExplicitTime(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	fixt := []SpecFixture{
		// These should work - explicit times are well supported
		{"at 3pm", 0, "at 3pm", 15 * time.Hour},
		{"at 15:00", 0, "at 15:00", 15 * time.Hour},
		{"in 30 minutes", 0, "in 30 minutes", 30 * time.Minute},
		{"in 10 minutes", 0, "in 10 minutes", 10 * time.Minute},
		{"in 2 hours", 0, "in 2 hours", 2 * time.Hour},
		{"in 1 hour", 0, "in 1 hour", 1 * time.Hour},
	}

	applySpecFixtures(t, "TodayWithExplicitTime", w, fixt)
}

func TestUpcomingSpec_TodayTimeOfDay(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// Note: These tests use SPEC defaults (morning=09:00, afternoon=12:00, evening=19:00, tonight=20:00)
	// Current library defaults are different (morning=08:00, afternoon=15:00, evening=18:00, tonight=23:00)
	// So these tests will FAIL, documenting the mismatch
	fixt := []SpecFixture{
		{"this morning", 0, "this morning", 9 * time.Hour},       // Spec: 09:00, Library: 08:00
		{"this afternoon", 0, "this afternoon", 12 * time.Hour},  // Spec: 12:00, Library: 15:00
		{"this evening", 0, "this evening", 19 * time.Hour},      // Spec: 19:00, Library: 18:00
		{"tonight", 0, "tonight", 20 * time.Hour},                // Spec: 20:00, Library: 23:00
	}

	applySpecFixtures(t, "TodayTimeOfDay", w, fixt)
}

func TestUpcomingSpec_LaterToday(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: "later today" should be now + 3 hours
	// This rule does NOT exist in the library yet
	// These tests will FAIL
	fixt := []SpecFixture{
		{"later today", 0, "later today", 3 * time.Hour}, // Expected to fail - no rule exists
	}

	applySpecFixtures(t, "LaterToday", w, fixt)
}

// ==============================================================================
// CATEGORY B: Tomorrow
// ==============================================================================

func TestUpcomingSpec_TomorrowWithTime(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	fixt := []SpecFixture{
		// Tomorrow = Jan 7, 2016 (Thursday)
		{"tomorrow at 09:00", 0, "tomorrow at 09:00", (24 + 9) * time.Hour},
		{"tomorrow at 2pm", 0, "tomorrow at 2pm", (24 + 14) * time.Hour},
		{"tomorrow at 6am", 0, "tomorrow at 6am", (24 + 6) * time.Hour},
	}

	applySpecFixtures(t, "TomorrowWithTime", w, fixt)
}

func TestUpcomingSpec_TomorrowTimeOfDay(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// Note: These tests use SPEC defaults
	// Current library uses different defaults, so these will FAIL
	fixt := []SpecFixture{
		{"tomorrow morning", 0, "tomorrow morning", (24 + 9) * time.Hour},     // Spec: 09:00, Library: 08:00
		{"tomorrow afternoon", 0, "tomorrow afternoon", (24 + 12) * time.Hour}, // Spec: 12:00, Library: 15:00
		{"tomorrow evening", 0, "tomorrow evening", (24 + 19) * time.Hour},     // Spec: 19:00, Library: 18:00
	}

	applySpecFixtures(t, "TomorrowTimeOfDay", w, fixt)
}

func TestUpcomingSpec_TomorrowNoTime(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: "tomorrow" without time should default to 09:00
	// The library parses "tomorrow" but may not set a default time
	// This test will likely FAIL
	fixt := []SpecFixture{
		{"tomorrow", 0, "tomorrow", (24 + 9) * time.Hour}, // Spec: tomorrow at 09:00
	}

	applySpecFixtures(t, "TomorrowNoTime", w, fixt)
}

// ==============================================================================
// CATEGORY C: Specific Weekdays
// ==============================================================================

func TestUpcomingSpec_WeekdaysWithTime(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// Reference: Wednesday Jan 6, 2016
	fixt := []SpecFixture{
		{"on Monday at 10am", 0, "on Monday at 10am", (5*24 + 10) * time.Hour}, // Jan 11 at 10:00
		{"Friday at 4pm", 0, "Friday at 4pm", (2*24 + 16) * time.Hour},         // Jan 8 at 16:00
	}

	applySpecFixtures(t, "WeekdaysWithTime", w, fixt)
}

func TestUpcomingSpec_WeekdaysTimeOfDay(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// Note: These tests use SPEC defaults
	// These will likely FAIL due to default time mismatches
	fixt := []SpecFixture{
		{"Friday morning", 0, "Friday morning", (2*24 + 9) * time.Hour},     // Spec: 09:00, Library: 08:00
		{"Friday afternoon", 0, "Friday afternoon", (2*24 + 12) * time.Hour}, // Spec: 12:00, Library: 15:00
		{"Monday evening", 0, "Monday evening", (5*24 + 19) * time.Hour},     // Spec: 19:00, Library: 18:00
	}

	applySpecFixtures(t, "WeekdaysTimeOfDay", w, fixt)
}

func TestUpcomingSpec_WeekdaysNoTime(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: weekday without time should default to 09:00
	// The library parses weekdays but may not set a default time
	// These tests will likely FAIL
	fixt := []SpecFixture{
		{"on Wednesday", 0, "on Wednesday", (7*24 + 9) * time.Hour}, // Next Wednesday = Jan 13 at 09:00
		{"on Monday", 0, "on Monday", (5*24 + 9) * time.Hour},       // Jan 11 at 09:00
		{"on Friday", 0, "on Friday", (2*24 + 9) * time.Hour},       // Jan 8 at 09:00
	}

	applySpecFixtures(t, "WeekdaysNoTime", w, fixt)
}

// ==============================================================================
// CATEGORY D: Relative Dates
// ==============================================================================

func TestUpcomingSpec_RelativeWeekMonthYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// These patterns are supported by the library
	// But may not have the correct default times according to spec
	fixt := []SpecFixture{
		// "next week" should mean next Monday at 09:00 according to spec
		{"next week", 0, "next week", (5*24 + 9) * time.Hour}, // Jan 11 (Monday) at 09:00 - will likely FAIL (no time set)

		// "next month" should mean 1st of next month at 09:00
		// Reference: Jan 6, 2016 → Feb 1, 2016 at 09:00
		{"next month", 0, "next month", (26*24 + 9) * time.Hour}, // Feb 1 at 09:00 - will likely FAIL (no time set)
	}

	applySpecFixtures(t, "RelativeWeekMonthYear", w, fixt)
}

func TestUpcomingSpec_ThisWeekend(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: "this weekend" should mean Saturday at 10:00
	// Reference: Wednesday Jan 6 → Saturday Jan 9 at 10:00
	// This rule does NOT exist in the library
	// This test will FAIL
	fixt := []SpecFixture{
		{"this weekend", 0, "this weekend", (3*24 + 10) * time.Hour}, // Expected to fail - no rule exists
	}

	applySpecFixtures(t, "ThisWeekend", w, fixt)
}

func TestUpcomingSpec_NextQuarter(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: "next quarter" should be supported
	// This rule does NOT exist in the library
	// This test will FAIL
	fixt := []SpecFixture{
		// Q1 2016 ends March 31, so "next quarter" from Jan 6 should be April 1, 2016 at 09:00
		{"next quarter", 0, "next quarter", (86*24 + 9) * time.Hour}, // Expected to fail - no rule exists
	}

	applySpecFixtures(t, "NextQuarter", w, fixt)
}

// ==============================================================================
// CATEGORY E: Relative Durations
// ==============================================================================

func TestUpcomingSpec_RelativeDurations(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// These are well supported by the Deadline rule
	fixt := []SpecFixture{
		{"in 5 minutes", 0, "in 5 minutes", 5 * time.Minute},
		{"in 20 minutes", 0, "in 20 minutes", 20 * time.Minute},
		{"in 45 minutes", 0, "in 45 minutes", 45 * time.Minute},
		{"in 1 hour", 0, "in 1 hour", 1 * time.Hour},
		{"in 2 hours", 0, "in 2 hours", 2 * time.Hour},
		{"in 3 hours", 0, "in 3 hours", 3 * time.Hour},
		{"in 24 hours", 0, "in 24 hours", 24 * time.Hour},
	}

	applySpecFixtures(t, "RelativeDurations", w, fixt)
}

// ==============================================================================
// CATEGORY F: Recurring Reminders
// ==============================================================================

func TestUpcomingSpec_RecurringReminders(t *testing.T) {
	// NOTE: Recurring patterns are OUT OF SCOPE for this library
	// The library is designed to parse single time expressions, not schedules
	// This would require significant architecture changes

	t.Skip("Recurring patterns are out of scope for single-time parser - requires schedule support")

	// Examples that would need to be supported:
	// "every hour"
	// "every weekday at 09:00"
	// "every Monday at 9am"
	// "on the first day of every month"
	// "on the last day of every month"
}

// ==============================================================================
// CATEGORY G: Vague Phrasing
// ==============================================================================

func TestUpcomingSpec_AfterLunch(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: "after lunch" should default to 14:00 (2pm)
	// This rule does NOT exist in the library
	// This test will FAIL
	fixt := []SpecFixture{
		{"after lunch", 0, "after lunch", 14 * time.Hour}, // Expected to fail - no rule exists
	}

	applySpecFixtures(t, "AfterLunch", w, fixt)
}

func TestUpcomingSpec_AfterWork(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: "after work" should default to 18:00 (6pm)
	// This rule does NOT exist in the library
	// This test will FAIL
	fixt := []SpecFixture{
		{"after work", 0, "after work", 18 * time.Hour}, // Expected to fail - no rule exists
	}

	applySpecFixtures(t, "AfterWork", w, fixt)
}

func TestUpcomingSpec_BeforeEndOfDay(t *testing.T) {
	w := when.New(nil)
	w.Add(en.All...)

	// According to spec: "before end of day" / "before EOD" should default to 17:00 (5pm)
	// This rule does NOT exist in the library
	// These tests will FAIL
	fixt := []SpecFixture{
		{"before end of day", 0, "before end of day", 17 * time.Hour}, // Expected to fail - no rule exists
		{"before EOD", 0, "before EOD", 17 * time.Hour},               // Expected to fail - no rule exists
	}

	applySpecFixtures(t, "BeforeEndOfDay", w, fixt)
}

// ==============================================================================
// SUMMARY TEST: Default Time Behaviors
// ==============================================================================

func TestUpcomingSpec_DefaultTimeBehaviors(t *testing.T) {
	// This test documents all the default time behavior expectations from the spec
	// Many of these will FAIL due to:
	// 1. Wrong default times (morning=08:00 vs 09:00, etc.)
	// 2. No default time set at all (just date without time)

	w := when.New(nil)
	w.Add(en.All...)

	fixt := []SpecFixture{
		// Part of day defaults (according to spec)
		{"morning", 0, "morning", 9 * time.Hour},     // Spec: 09:00, Library: 08:00
		{"afternoon", 0, "afternoon", 12 * time.Hour}, // Spec: 12:00, Library: 15:00
		{"evening", 0, "evening", 19 * time.Hour},     // Spec: 19:00, Library: 18:00
		{"tonight", 0, "tonight", 20 * time.Hour},     // Spec: 20:00, Library: 23:00

		// Date without time should default to 09:00
		{"tomorrow", 0, "tomorrow", (24 + 9) * time.Hour},

		// Weekday without time should default to 09:00
		{"next Monday", 0, "next Monday", (5*24 + 9) * time.Hour},

		// "next week" should mean next Monday at 09:00
		{"next week", 0, "next week", (5*24 + 9) * time.Hour},
	}

	applySpecFixtures(t, "DefaultTimeBehaviors", w, fixt)
}
