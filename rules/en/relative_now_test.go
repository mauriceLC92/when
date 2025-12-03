package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestRelativeNow(t *testing.T) {
	fixt := []Fixture{
		{"7 days from now", 0, "7 days from now", 7 * 24 * time.Hour},
		{"3 hours from now", 0, "3 hours from now", 3 * time.Hour},
		{"5 minutes from now", 0, "5 minutes from now", 5 * time.Minute},
		{"10 seconds from now", 0, "10 seconds from now", 10 * time.Second},
		{"2 weeks from now", 0, "2 weeks from now", 14 * 24 * time.Hour},
		{"a week from now", 0, "a week from now", 7 * 24 * time.Hour},
		{"an hour from now", 0, "an hour from now", time.Hour},
	}

	w := when.New(nil)
	w.Add(en.RelativeNow(rules.Override))

	ApplyFixtures(t, "en.RelativeNow from now", w, fixt)
}

func TestRelativeNowBeforeNow(t *testing.T) {
	fixt := []Fixture{
		{"7 days before now", 0, "7 days before now", -7 * 24 * time.Hour},
		{"3 hours before now", 0, "3 hours before now", -3 * time.Hour},
		{"5 minutes before now", 0, "5 minutes before now", -5 * time.Minute},
		{"a week before now", 0, "a week before now", -7 * 24 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.RelativeNow(rules.Override))

	ApplyFixtures(t, "en.RelativeNow before now", w, fixt)
}

func TestRelativeNowHence(t *testing.T) {
	fixt := []Fixture{
		{"1 week hence", 0, "1 week hence", 7 * 24 * time.Hour},
		{"3 days hence", 0, "3 days hence", 3 * 24 * time.Hour},
		{"a month hence", 0, "a month hence", 0}, // Month changes, not duration
	}

	w := when.New(nil)
	w.Add(en.RelativeNow(rules.Override))

	// Test week and days
	for i := 0; i < 2; i++ {
		f := fixt[i]
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[en.RelativeNow hence] err #%d", i)
		require.NotNil(t, res, "[en.RelativeNow hence] res #%d", i)
		require.Equal(t, f.Index, res.Index, "[en.RelativeNow hence] index #%d", i)
		require.Equal(t, f.Phrase, res.Text, "[en.RelativeNow hence] text #%d", i)
		require.Equal(t, f.Diff, res.Time.Sub(null), "[en.RelativeNow hence] diff #%d", i)
	}
}

func TestRelativeNowMonth(t *testing.T) {
	w := when.New(nil)
	w.Add(en.RelativeNow(rules.Override))

	// null is January 6, 2016
	res, err := w.Parse("5 months from now", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "5 months from now", res.Text)
	require.Equal(t, time.June, res.Time.Month())

	res, err = w.Parse("5 months before now", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "5 months before now", res.Text)
	require.Equal(t, time.August, res.Time.Month())
	require.Equal(t, 2015, res.Time.Year())
}

func TestRelativeNowYear(t *testing.T) {
	w := when.New(nil)
	w.Add(en.RelativeNow(rules.Override))

	// null is January 6, 2016
	res, err := w.Parse("2 years from now", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "2 years from now", res.Text)
	require.Equal(t, 2018, res.Time.Year())

	res, err = w.Parse("a year from now", null)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, "a year from now", res.Text)
	require.Equal(t, 2017, res.Time.Year())
}
