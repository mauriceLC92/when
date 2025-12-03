package en

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

/*
	"0800"
	"1430"
	"2359"
	"0000"

	Military time format: 4 digits without separator
*/

func MilitaryTime(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"((?:[01][0-9]|2[0-3])([0-5][0-9]))" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && s != rules.Override {
				return false, nil
			}

			timeStr := m.Captures[0]
			if len(timeStr) != 4 {
				return false, nil
			}

			hour, err := strconv.Atoi(timeStr[:2])
			if err != nil {
				return false, errors.Wrap(err, "military time rule: hour")
			}

			minute, err := strconv.Atoi(timeStr[2:])
			if err != nil {
				return false, errors.Wrap(err, "military time rule: minute")
			}

			if hour > 23 || minute > 59 {
				return false, nil
			}

			c.Hour = &hour
			c.Minute = &minute
			zero := 0
			c.Second = &zero

			return true, nil
		},
	}
}
