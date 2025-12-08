package en

import (
	"regexp"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func ThisWeekend(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)(this\s+weekend)(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Duration != 0 && !overwrite {
				return false, nil
			}

			// Saturday is weekday 6
			daysUntilSaturday := (6 - int(ref.Weekday()) + 7) % 7
			if daysUntilSaturday == 0 && ref.Hour() >= 10 {
				// If it's Saturday after 10am, use next Saturday
				daysUntilSaturday = 7
			}

			c.Duration = time.Duration(daysUntilSaturday*24) * time.Hour
			c.Hour = pointer.ToInt(10)
			c.Minute = pointer.ToInt(0)
			return true, nil
		},
	}
}
