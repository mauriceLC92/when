package en

import (
	"regexp"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func BeforeEndOfDay(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)(before\s+(?:end\s+of\s+day|eod))(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && !overwrite {
				return false, nil
			}
			c.Hour = pointer.ToInt(17)
			c.Minute = pointer.ToInt(0)
			return true, nil
		},
	}
}
