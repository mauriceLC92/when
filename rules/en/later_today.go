package en

import (
	"regexp"
	"time"

	"github.com/olebedev/when/rules"
)

func LaterToday(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)(later\s+today)(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Duration != 0 && !overwrite {
				return false, nil
			}
			c.Duration = 3 * time.Hour
			return true, nil
		},
	}
}
