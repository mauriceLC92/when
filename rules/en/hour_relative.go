package en

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

/*
	"10 to 8" -> 7:50
	"10 past 2" -> 2:10
	"half past 2" -> 2:30
	"quarter to 8" -> 7:45
	"quarter past 3" -> 3:15
	"6 in the morning" -> 6:00 AM
	"7 in the evening" -> 19:00
	"eleven o'clock" -> 11:00
	"5 o'clock" -> 5:00
*/

var MINUTE_WORDS = map[string]int{
	"quarter": 15,
	"half":    30,
}

// HourRelativeTo handles "X to Y" patterns like "10 to 8" -> 7:50
func HourRelativeTo(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(quarter|half|\\d{1,2})\\s+" +
			"(to|past)\\s+" +
			"(" + INTEGER_WORDS_PATTERN + "|\\d{1,2})" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && s != rules.Override {
				return false, nil
			}

			minuteStr := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			direction := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			hourStr := strings.ToLower(strings.TrimSpace(m.Captures[2]))

			var minutes int
			if min, ok := MINUTE_WORDS[minuteStr]; ok {
				minutes = min
			} else {
				var err error
				minutes, err = strconv.Atoi(minuteStr)
				if err != nil || minutes > 59 {
					return false, nil
				}
			}

			var hour int
			if h, ok := INTEGER_WORDS[hourStr]; ok {
				hour = h
			} else {
				var err error
				hour, err = strconv.Atoi(hourStr)
				if err != nil || hour > 12 {
					return false, nil
				}
			}

			if direction == "to" {
				// "10 to 8" means 7:50
				hour = hour - 1
				if hour < 0 {
					hour = 23
				}
				minutes = 60 - minutes
			}
			// "past" means the minutes are added to the hour as-is

			c.Hour = &hour
			c.Minute = &minutes
			zero := 0
			c.Second = &zero

			return true, nil
		},
	}
}

// HourInPeriod handles "X in the morning/afternoon/evening" patterns
func HourInPeriod(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + INTEGER_WORDS_PATTERN + "|\\d{1,2})\\s+" +
			"(?:in\\s+the\\s+|at\\s+)?" +
			"(morning|afternoon|evening)" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Hour != nil && s != rules.Override {
				return false, nil
			}

			hourStr := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			period := strings.ToLower(strings.TrimSpace(m.Captures[1]))

			var hour int
			if h, ok := INTEGER_WORDS[hourStr]; ok {
				hour = h
			} else {
				var err error
				hour, err = strconv.Atoi(hourStr)
				if err != nil || hour > 12 {
					return false, nil
				}
			}

			switch period {
			case "morning":
				// Morning hours are as-is (assuming 1-12 AM)
				if hour == 12 {
					hour = 0 // 12 in the morning is midnight? This is ambiguous. Let's keep as 12.
				}
			case "afternoon":
				// Afternoon: add 12 if < 12
				if hour < 12 {
					hour += 12
				}
			case "evening":
				// Evening: add 12 if < 12
				if hour < 12 {
					hour += 12
				}
			}

			c.Hour = &hour
			zero := 0
			c.Minute = &zero
			c.Second = &zero

			return true, nil
		},
	}
}

// HourOClock handles "X o'clock" patterns
func HourOClock(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + INTEGER_WORDS_PATTERN + "|\\d{1,2})\\s*" +
			"(o[''`]?clock)" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Hour != nil && s != rules.Override {
				return false, nil
			}

			hourStr := strings.ToLower(strings.TrimSpace(m.Captures[0]))

			var hour int
			if h, ok := INTEGER_WORDS[hourStr]; ok {
				hour = h
			} else {
				var err error
				hour, err = strconv.Atoi(hourStr)
				if err != nil || hour > 23 {
					return false, nil
				}
			}

			c.Hour = &hour
			zero := 0
			c.Minute = &zero
			c.Second = &zero

			return true, nil
		},
	}
}
