package common

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
)

/*
ISO 8601 date format: YYYY-MM-DD
Examples:
- 1979-05-27
- 2023-12-25
- 2020-01-01
*/

func ISODate(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"((?:19|20)[0-9]{2})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Day != nil || c.Month != nil || c.Year != nil) && s != rules.Override {
				return false, nil
			}

			year, err := strconv.Atoi(m.Captures[0])
			if err != nil {
				return false, nil
			}

			month, err := strconv.Atoi(m.Captures[1])
			if err != nil {
				return false, nil
			}

			day, err := strconv.Atoi(m.Captures[2])
			if err != nil {
				return false, nil
			}

			// Validate day for the given month
			if getDays(year, month) < day {
				return false, nil
			}

			c.Year = &year
			c.Month = &month
			c.Day = &day

			return true, nil
		},
	}
}
