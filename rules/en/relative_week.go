package en

import (
	"regexp"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

/*
	"last week" -> -7 days
	"next week" -> +7 days
	"this week" -> current week
	"last month" -> previous month
	"next month" -> next month
	"this second" -> now
	"last year" -> previous year
	"next year" -> next year
*/

func RelativeWeek(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(this|last|past|next)\\s+" +
			"(second|week|month|year)" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			direction := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			unit := strings.ToLower(strings.TrimSpace(m.Captures[1]))

			switch unit {
			case "second":
				// "this second" means now - no change needed
				return true, nil

			case "week":
				if c.Duration != 0 && s != rules.Override {
					return false, nil
				}
				switch direction {
				case "last", "past":
					c.Duration = -7 * 24 * time.Hour
				case "next":
					c.Duration = 7 * 24 * time.Hour
				case "this":
					// this week - no change (current week)
				}

			case "month":
				if c.Month != nil && s != rules.Override {
					return false, nil
				}
				switch direction {
				case "last", "past":
					month := int(ref.Month()) - 1
					if month < 1 {
						month = 12
						c.Year = pointer.ToInt(ref.Year() - 1)
					}
					c.Month = pointer.ToInt(month)
				case "next":
					month := int(ref.Month()) + 1
					if month > 12 {
						month = 1
						c.Year = pointer.ToInt(ref.Year() + 1)
					}
					c.Month = pointer.ToInt(month)
				case "this":
					c.Month = pointer.ToInt(int(ref.Month()))
				}

			case "year":
				if c.Year != nil && s != rules.Override {
					return false, nil
				}
				switch direction {
				case "last", "past":
					c.Year = pointer.ToInt(ref.Year() - 1)
				case "next":
					c.Year = pointer.ToInt(ref.Year() + 1)
				case "this":
					c.Year = pointer.ToInt(ref.Year())
				}
			}

			return true, nil
		},
	}
}
