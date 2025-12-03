package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestRelativeWeek(t *testing.T) {
	// Note: null is January 6, 2016 (Wednesday)
	fixt := []Fixture{
		{"last week", 0, "last week", -7 * 24 * time.Hour},
		{"next week", 0, "next week", 7 * 24 * time.Hour},
		{"review next week", 7, "next week", 7 * 24 * time.Hour},
		{"did it last week", 7, "last week", -7 * 24 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.RelativeWeek(rules.Override))

	ApplyFixtures(t, "en.RelativeWeek", w, fixt)
}

func TestRelativeWeekThisSecond(t *testing.T) {
	w := when.New(nil)
	w.Add(en.RelativeWeek(rules.Override))

	// "this second" should return the reference time unchanged
	res, err := w.Parse("this second", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "this second", res.Text)
	require.Equal(t, null, res.Time) // Time should be unchanged
}

func TestRelativeMonth(t *testing.T) {
	w := when.New(nil)
	w.Add(en.RelativeWeek(rules.Override))

	// null is January 6, 2016
	res, err := w.Parse("next month", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "next month", res.Text)
	require.Equal(t, time.February, res.Time.Month())

	res, err = w.Parse("last month", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "last month", res.Text)
	require.Equal(t, time.December, res.Time.Month())
	require.Equal(t, 2015, res.Time.Year())
}

func TestRelativeYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.RelativeWeek(rules.Override))

	// null is January 6, 2016
	res, err := w.Parse("next year", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "next year", res.Text)
	require.Equal(t, 2017, res.Time.Year())

	res, err = w.Parse("last year", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "last year", res.Text)
	require.Equal(t, 2015, res.Time.Year())
}
