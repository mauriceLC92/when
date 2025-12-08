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
					// "next week" means next Monday at 09:00
					// Calculate days until next Monday
					daysUntilMonday := (8 - int(ref.Weekday())) % 7
					if daysUntilMonday == 0 {
						daysUntilMonday = 7 // If today is Monday, go to next Monday
					}
					c.Duration = time.Duration(daysUntilMonday*24) * time.Hour
					c.Hour = pointer.ToInt(9)
					c.Minute = pointer.ToInt(0)
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
					c.Day = pointer.ToInt(1) // First day of the month
					// Add default time if not already set
					if c.Hour == nil && c.Minute == nil {
						c.Hour = pointer.ToInt(9)
						c.Minute = pointer.ToInt(0)
					}
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
