package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestDayMonthYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.DayMonthYear(rules.Override))

	testCases := []struct {
		input       string
		matchText   string
		expectYear  int
		expectMonth time.Month
		expectDay   int
	}{
		{"5th may 2017", "5th may 2017", 2017, time.May, 5},
		{"1st september 2020", "1st september 2020", 2020, time.September, 1},
		{"twenty-first december 2019", "twenty-first december 2019", 2019, time.December, 21},
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

func TestDayNumMonthYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.DayNumMonthYear(rules.Override))

	testCases := []struct {
		input       string
		matchText   string
		expectYear  int
		expectMonth time.Month
		expectDay   int
	}{
		{"3 jan 2000", "3 jan 2000", 2000, time.January, 3},
		{"17 april 85", "17 april 85", 1985, time.April, 17},
		{"25 dec 2023", "25 dec 2023", 2023, time.December, 25},
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

func TestMonthDayYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.MonthDayYear(rules.Override))

	testCases := []struct {
		input       string
		matchText   string
		expectYear  int
		expectMonth time.Month
		expectDay   int
	}{
		{"jan 3 2010", "jan 3 2010", 2010, time.January, 3},
		{"february 14, 2004", "february 14, 2004", 2004, time.February, 14},
		{"february 14th, 2004", "february 14th, 2004", 2004, time.February, 14},
		{"december 25, 2023", "december 25, 2023", 2023, time.December, 25},
		{"may 7 '97", "may 7 '97", 1997, time.May, 7},
		{"jan 1 '05", "jan 1 '05", 2005, time.January, 1},
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

func TestMonthYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.MonthYear(rules.Override))

	testCases := []struct {
		input       string
		matchText   string
		expectYear  int
		expectMonth time.Month
		expectDay   int
	}{
		{"October 2006", "October 2006", 2006, time.October, 1},
		{"oct 06", "oct 06", 2006, time.October, 1},
		{"january 2020", "january 2020", 2020, time.January, 1},
		{"dec 99", "dec 99", 1999, time.December, 1},
		{"february '21", "february '21", 2021, time.February, 1},
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
