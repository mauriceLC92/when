package en

import (
	"regexp"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

/*
	Ordinal patterns:
	- "3rd wednesday in november" - Nth weekday of month
	- "3rd thursday this september" - Nth weekday of month
	- "4th day last week" - Nth day of week
	- "3rd month next year" - Nth month of year
*/

// getNthWeekdayOfMonth returns the date of the Nth occurrence of a weekday in a given month/year
// For example, getNthWeekdayOfMonth(2023, 11, time.Wednesday, 3) returns the 3rd Wednesday of November 2023
func getNthWeekdayOfMonth(year, month int, weekday time.Weekday, n int) time.Time {
	// Start with the first day of the month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	// Find the first occurrence of the weekday
	daysUntilWeekday := int(weekday) - int(firstDay.Weekday())
	if daysUntilWeekday < 0 {
		daysUntilWeekday += 7
	}

	// Calculate the Nth occurrence
	day := 1 + daysUntilWeekday + (n-1)*7

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// OrdinalWeekdayInMonth handles patterns like "3rd wednesday in november"
func OrdinalWeekdayInMonth(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + ORDINAL_WORDS_PATTERN[3:] + "\\s+" +
			"(" + WEEKDAY_OFFSET_PATTERN[3:] + "\\s+" +
			"(?:in|of|this|next|last)?\\s*" +
			"(" + MONTH_OFFSET_PATTERN[3:] +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Day != nil || c.Month != nil) && !overwrite {
				return false, nil
			}

			ordStr := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			weekdayStr := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			monthStr := strings.ToLower(strings.TrimSpace(m.Captures[2]))

			n, ok := ORDINAL_WORDS[ordStr]
			if !ok {
				return false, nil
			}

			weekdayOffset, ok := WEEKDAY_OFFSET[weekdayStr]
			if !ok {
				return false, nil
			}

			month, ok := MONTH_OFFSET[monthStr]
			if !ok {
				return false, nil
			}

			// Determine year - if the month is in the past, use next year
			year := ref.Year()
			if month < int(ref.Month()) {
				year++
			}

			targetDate := getNthWeekdayOfMonth(year, month, time.Weekday(weekdayOffset), n)

			// Validate the date is in the correct month (n might be too large)
			if targetDate.Month() != time.Month(month) {
				return false, nil
			}

			day := targetDate.Day()
			c.Year = &year
			c.Month = &month
			c.Day = &day

			return true, nil
		},
	}
}

// OrdinalDayInWeek handles patterns like "4th day last week"
func OrdinalDayInWeek(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + ORDINAL_WORDS_PATTERN[3:] + "\\s+" +
			"(day)\\s+" +
			"(this|next|last|past)\\s+(week)" +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Duration != 0 && !overwrite {
				return false, nil
			}

			ordStr := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			direction := strings.ToLower(strings.TrimSpace(m.Captures[2]))

			n, ok := ORDINAL_WORDS[ordStr]
			if !ok || n < 1 || n > 7 {
				return false, nil
			}

			// Calculate the target weekday (1st day = Sunday, 2nd = Monday, etc.)
			targetWeekday := time.Weekday((n - 1) % 7) // Convert to 0-indexed

			// Calculate base date based on direction
			var baseDate time.Time
			switch direction {
			case "last", "past":
				// Go back to the start of last week (Sunday)
				daysToLastSunday := int(ref.Weekday()) + 7
				baseDate = ref.AddDate(0, 0, -daysToLastSunday)
			case "next":
				// Go forward to the start of next week (Sunday)
				daysToNextSunday := 7 - int(ref.Weekday())
				baseDate = ref.AddDate(0, 0, daysToNextSunday)
			case "this":
				// Start of this week (Sunday)
				daysToThisSunday := int(ref.Weekday())
				baseDate = ref.AddDate(0, 0, -daysToThisSunday)
			default:
				return false, nil
			}

			// Add the target weekday offset
			targetDate := baseDate.AddDate(0, 0, int(targetWeekday))
			c.Duration = targetDate.Sub(ref)

			return true, nil
		},
	}
}

// OrdinalMonthInYear handles patterns like "3rd month next year"
func OrdinalMonthInYear(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + ORDINAL_WORDS_PATTERN[3:] + "\\s+" +
			"(month)\\s+" +
			"(this|next|last|past)\\s+(year)" +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Month != nil || c.Year != nil) && !overwrite {
				return false, nil
			}

			ordStr := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			direction := strings.ToLower(strings.TrimSpace(m.Captures[2]))

			n, ok := ORDINAL_WORDS[ordStr]
			if !ok || n < 1 || n > 12 {
				return false, nil
			}

			// Determine year
			year := ref.Year()
			switch direction {
			case "last", "past":
				year--
			case "next":
				year++
			case "this":
				// Keep current year
			default:
				return false, nil
			}

			c.Year = pointer.ToInt(year)
			c.Month = pointer.ToInt(n)

			return true, nil
		},
	}
}
