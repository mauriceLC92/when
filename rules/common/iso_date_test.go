package common_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/common"
	"github.com/stretchr/testify/require"
)

func TestISODate(t *testing.T) {
	w := when.New(nil)
	w.Add(common.ISODate(rules.Override))

	testCases := []struct {
		input       string
		matchText   string
		expectYear  int
		expectMonth time.Month
		expectDay   int
	}{
		{"1979-05-27", "1979-05-27", 1979, time.May, 27},
		{"2023-12-25", "2023-12-25", 2023, time.December, 25},
		{"2020-01-01", "2020-01-01", 2020, time.January, 1},
		{"deadline is 2024-06-15", "2024-06-15", 2024, time.June, 15},
		{"event on 1999-12-31 was fun", "1999-12-31", 1999, time.December, 31},
	}

	for _, tc := range testCases {
		res, err := w.Parse(tc.input, null)
		require.Nil(t, err, "error for %s", tc.input)
		require.NotNil(t, res, "result for %s", tc.input)
		require.Equal(t, tc.matchText, res.Text, "text for %s", tc.input)
		require.Equal(t, tc.expectYear, res.Time.Year(), "year for %s", tc.input)
		require.Equal(t, tc.expectMonth, res.Time.Month(), "month for %s", tc.input)
		require.Equal(t, tc.expectDay, res.Time.Day(), "day for %s", tc.input)
	}
}

func TestISODateInvalid(t *testing.T) {
	w := when.New(nil)
	w.Add(common.ISODate(rules.Override))

	invalidCases := []string{
		"2023-13-01", // Invalid month
		"2023-00-01", // Invalid month
		"2023-02-30", // Invalid day for February
		"2023-04-31", // Invalid day for April
	}

	for _, tc := range invalidCases {
		res, err := w.Parse(tc, null)
		require.Nil(t, err, "error for %s", tc)
		require.Nil(t, res, "result should be nil for %s", tc)
	}
}
