package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestOrdinalWeekdayInMonth(t *testing.T) {
	w := when.New(nil)
	w.Add(en.OrdinalWeekdayInMonth(rules.Override))

	// null is January 6, 2016 (Wednesday)
	testCases := []struct {
		input       string
		matchText   string
		expectYear  int
		expectMonth time.Month
		expectDay   int
	}{
		{"3rd wednesday in november", "3rd wednesday in november", 2016, time.November, 16},
		{"3rd thursday this september", "3rd thursday this september", 2016, time.September, 15},
		{"1st monday in march", "1st monday in march", 2016, time.March, 7},
		{"2nd friday in june", "2nd friday in june", 2016, time.June, 10},
		{"fourth sunday in december", "fourth sunday in december", 2016, time.December, 25},
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

func TestOrdinalDayInWeek(t *testing.T) {
	w := when.New(nil)
	w.Add(en.OrdinalDayInWeek(rules.Override))

	// null is January 6, 2016 (Wednesday, day 4 of the week when Sunday=1)
	testCases := []struct {
		input     string
		matchText string
		expectDay int
	}{
		{"4th day last week", "4th day last week", 30}, // 4th day (Wednesday) of last week = Dec 30, 2015
		{"1st day this week", "1st day this week", 3},  // 1st day (Sunday) of this week = Jan 3
		{"7th day this week", "7th day this week", 9},  // 7th day (Saturday) of this week = Jan 9
		{"2nd day next week", "2nd day next week", 11}, // 2nd day (Monday) of next week = Jan 11
	}

	for _, tc := range testCases {
		res, err := w.Parse(tc.input, null)
		require.Nil(t, err, "error for %s", tc.input)
		require.NotNil(t, res, "result for %s", tc.input)
		require.Equal(t, tc.matchText, res.Text, "text for %s", tc.input)
		require.Equal(t, tc.expectDay, res.Time.Day(), "day for %s: got %v", tc.input, res.Time)
	}
}

func TestOrdinalMonthInYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.OrdinalMonthInYear(rules.Override))

	// null is January 6, 2016
	testCases := []struct {
		input       string
		matchText   string
		expectYear  int
		expectMonth time.Month
	}{
		{"3rd month next year", "3rd month next year", 2017, time.March},
		{"6th month this year", "6th month this year", 2016, time.June},
		{"12th month last year", "12th month last year", 2015, time.December},
		{"first month next year", "first month next year", 2017, time.January},
	}

	for _, tc := range testCases {
		res, err := w.Parse(tc.input, null)
		require.Nil(t, err, "error for %s", tc.input)
		require.NotNil(t, res, "result for %s", tc.input)
		require.Equal(t, tc.matchText, res.Text, "text for %s", tc.input)
		require.Equal(t, tc.expectYear, res.Time.Year(), "year for %s", tc.input)
		require.Equal(t, tc.expectMonth, res.Time.Month(), "month for %s", tc.input)
	}
}
