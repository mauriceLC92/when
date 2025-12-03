package en

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

/*
	"5 months before now" -> -5 months
	"7 days from now" -> +7 days
	"1 week hence" -> +1 week
	"a year from now" -> +1 year
	"3 hours from now" -> +3 hours
*/

func RelativeNow(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(
			"(?i)(?:\\W|^)" +
				"(" + INTEGER_WORDS_PATTERN + "|[0-9]+|an?(?:\\s*few)?|half(?:\\s*an?)?)\\s*" +
				"(seconds?|min(?:ute)?s?|hours?|days?|weeks?|months?|years?)\\s+" +
				"(before\\s+now|from\\s+now|hence)" +
				"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			numStr := strings.TrimSpace(m.Captures[0])
			unitStr := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			directionStr := strings.ToLower(strings.TrimSpace(m.Captures[2]))

			var num int
			var err error

			if n, ok := INTEGER_WORDS[numStr]; ok {
				num = n
			} else if numStr == "a" || numStr == "an" {
				num = 1
			} else if strings.Contains(numStr, "few") {
				num = 3
			} else if strings.Contains(numStr, "half") {
				// handled below
			} else {
				num, err = strconv.Atoi(numStr)
				if err != nil {
					return false, errors.Wrapf(err, "convert '%s' to int", numStr)
				}
			}

			// Determine direction: positive for "from now"/"hence", negative for "before now"
			negative := strings.Contains(directionStr, "before")

			if !strings.Contains(numStr, "half") {
				switch {
				case strings.Contains(unitStr, "second"):
					if c.Duration == 0 || overwrite {
						dur := time.Duration(num) * time.Second
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "min"):
					if c.Duration == 0 || overwrite {
						dur := time.Duration(num) * time.Minute
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "hour"):
					if c.Duration == 0 || overwrite {
						dur := time.Duration(num) * time.Hour
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "day"):
					if c.Duration == 0 || overwrite {
						dur := time.Duration(num) * 24 * time.Hour
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "week"):
					if c.Duration == 0 || overwrite {
						dur := time.Duration(num) * 7 * 24 * time.Hour
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "month"):
					if c.Month == nil || overwrite {
						month := int(ref.Month())
						if negative {
							month -= num
						} else {
							month += num
						}
						// Handle year rollover
						for month < 1 {
							month += 12
							if c.Year == nil {
								c.Year = pointer.ToInt(ref.Year() - 1)
							} else {
								c.Year = pointer.ToInt(*c.Year - 1)
							}
						}
						for month > 12 {
							month -= 12
							if c.Year == nil {
								c.Year = pointer.ToInt(ref.Year() + 1)
							} else {
								c.Year = pointer.ToInt(*c.Year + 1)
							}
						}
						c.Month = pointer.ToInt(month)
					}
				case strings.Contains(unitStr, "year"):
					if c.Year == nil || overwrite {
						year := ref.Year()
						if negative {
							year -= num
						} else {
							year += num
						}
						c.Year = pointer.ToInt(year)
					}
				}
			} else {
				// Handle "half" cases
				switch {
				case strings.Contains(unitStr, "hour"):
					if c.Duration == 0 || overwrite {
						dur := 30 * time.Minute
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "day"):
					if c.Duration == 0 || overwrite {
						dur := 12 * time.Hour
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "week"):
					if c.Duration == 0 || overwrite {
						dur := 7 * 12 * time.Hour // 3.5 days
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "month"):
					if c.Duration == 0 || overwrite {
						dur := 14 * 24 * time.Hour // ~2 weeks
						if negative {
							dur = -dur
						}
						c.Duration = dur
					}
				case strings.Contains(unitStr, "year"):
					if c.Month == nil || overwrite {
						month := int(ref.Month())
						if negative {
							month -= 6
						} else {
							month += 6
						}
						for month < 1 {
							month += 12
							if c.Year == nil {
								c.Year = pointer.ToInt(ref.Year() - 1)
							} else {
								c.Year = pointer.ToInt(*c.Year - 1)
							}
						}
						for month > 12 {
							month -= 12
							if c.Year == nil {
								c.Year = pointer.ToInt(ref.Year() + 1)
							} else {
								c.Year = pointer.ToInt(*c.Year + 1)
							}
						}
						c.Month = pointer.ToInt(month)
					}
				}
			}

			return true, nil
		},
	}
}
