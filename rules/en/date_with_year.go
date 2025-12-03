package en

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

/*
	Patterns:
	- "5th may 2017" - Day + month + year
	- "jan 3 2010" - Month day year
	- "february 14, 2004" - Month day, year (comma)
	- "february 14th, 2004" - Month ordinal, year (comma)
	- "3 jan 2000" - Day month year
	- "17 april 85" - Day month short-year
	- "October 2006" - Month + year only
	- "oct 06" - Month + short year
	- "may seventh '97" - Month + spelled ordinal + short year
*/

// parseYear converts a year string to int, handling 2-digit years
// 00-29 -> 2000-2029, 30-99 -> 1930-1999
func parseYear(yearStr string) (int, error) {
	yearStr = strings.TrimPrefix(yearStr, "'")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, err
	}
	if year < 100 {
		if year <= 29 {
			year += 2000
		} else {
			year += 1900
		}
	}
	return year, nil
}

// DayMonthYear handles "Day Month Year" patterns like "5th may 2017", "3 jan 2000"
func DayMonthYear(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + ORDINAL_WORDS_PATTERN[3:] + "\\s+" +
			"(" + MONTH_OFFSET_PATTERN[3:] + "\\s+" +
			"('?[0-9]{2,4})" +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			ordDay := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			mon := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			yearStr := strings.TrimSpace(m.Captures[2])

			day, ok := ORDINAL_WORDS[ordDay]
			if !ok {
				return false, nil
			}

			month, ok := MONTH_OFFSET[mon]
			if !ok {
				return false, nil
			}

			year, err := parseYear(yearStr)
			if err != nil {
				return false, nil
			}

			c.Year = &year
			c.Month = &month
			c.Day = &day

			return true, nil
		},
	}
}

// DayNumMonthYear handles "Day Month Year" patterns with numeric day like "3 jan 2000", "17 april 85"
func DayNumMonthYear(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"([0-9]{1,2})\\s+" +
			"(" + MONTH_OFFSET_PATTERN[3:] + "\\s+" +
			"('?[0-9]{2,4})" +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			dayStr := strings.TrimSpace(m.Captures[0])
			mon := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			yearStr := strings.TrimSpace(m.Captures[2])

			day, err := strconv.Atoi(dayStr)
			if err != nil || day < 1 || day > 31 {
				return false, nil
			}

			month, ok := MONTH_OFFSET[mon]
			if !ok {
				return false, nil
			}

			year, err := parseYear(yearStr)
			if err != nil {
				return false, nil
			}

			c.Year = &year
			c.Month = &month
			c.Day = &day

			return true, nil
		},
	}
}

// MonthDayYear handles "Month Day Year" patterns like "jan 3 2010", "february 14, 2004"
func MonthDayYear(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + MONTH_OFFSET_PATTERN[3:] + "\\s+" +
			"(?:(" + ORDINAL_WORDS_PATTERN[3:] + "|([0-9]{1,2}))\\s*,?\\s+" +
			"('?[0-9]{2,4})" +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			mon := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			ordDay := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			numDay := strings.TrimSpace(m.Captures[2])
			yearStr := strings.TrimSpace(m.Captures[3])

			month, ok := MONTH_OFFSET[mon]
			if !ok {
				return false, nil
			}

			var day int
			if ordDay != "" && numDay == "" {
				d, ok := ORDINAL_WORDS[ordDay]
				if !ok {
					return false, nil
				}
				day = d
			} else if numDay != "" {
				d, err := strconv.Atoi(numDay)
				if err != nil || d < 1 || d > 31 {
					return false, nil
				}
				day = d
			} else {
				return false, nil
			}

			year, err := parseYear(yearStr)
			if err != nil {
				return false, nil
			}

			c.Year = &year
			c.Month = &month
			c.Day = &day

			return true, nil
		},
	}
}

// MonthYear handles "Month Year" patterns like "October 2006", "oct 06"
func MonthYear(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + MONTH_OFFSET_PATTERN[3:] + "\\s+" +
			"('?[0-9]{2,4})" +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			mon := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			yearStr := strings.TrimSpace(m.Captures[1])

			month, ok := MONTH_OFFSET[mon]
			if !ok {
				return false, nil
			}

			year, err := parseYear(yearStr)
			if err != nil {
				return false, nil
			}

			c.Year = &year
			c.Month = &month
			// Day defaults to 1 when not specified
			day := 1
			c.Day = &day

			return true, nil
		},
	}
}
