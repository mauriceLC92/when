package en

import (
	"regexp"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func NextQuarter(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)(next\s+quarter)(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Month != nil || c.Day != nil) && !overwrite {
				return false, nil
			}

			// Quarters: Q1=Jan(1), Q2=Apr(4), Q3=Jul(7), Q4=Oct(10)
			currentMonth := int(ref.Month())
			currentQuarter := (currentMonth - 1) / 3
			nextQuarterStartMonth := (currentQuarter+1)*3 + 1
			nextYear := ref.Year()

			if nextQuarterStartMonth > 12 {
				nextQuarterStartMonth = 1
				nextYear++
			}

			c.Year = pointer.ToInt(nextYear)
			c.Month = pointer.ToInt(nextQuarterStartMonth)
			c.Day = pointer.ToInt(1)
			c.Hour = pointer.ToInt(9)
			c.Minute = pointer.ToInt(0)
			return true, nil
		},
	}
}
