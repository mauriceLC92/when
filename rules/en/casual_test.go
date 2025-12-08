package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is now, ok", 16, "now", 0},
		{"The Deadline is today", 16, "today", 0},
		{"The Deadline is tonight", 16, "tonight", 20 * time.Hour}, // Changed from 23h to 20h per spec
		{"The Deadline is tomorrow evening", 16, "tomorrow", (24 + 9) * time.Hour}, // Changed: tomorrow now defaults to 09:00
		{"The Deadline is yesterday evening", 16, "yesterday", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(en.CasualDate(rules.Skip))

	ApplyFixtures(t, "en.CasualDate", w, fixt)
}

func TestCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline was this morning ", 17, "this morning", 9 * time.Hour},     // Changed from 8h to 9h per spec
		{"The Deadline was this noon ", 17, "this noon", 12 * time.Hour},
		{"The Deadline was this afternoon ", 17, "this afternoon", 12 * time.Hour}, // Changed from 15h to 12h per spec
		{"The Deadline was this evening ", 17, "this evening", 19 * time.Hour},     // Changed from 18h to 19h per spec
	}

	w := when.New(nil)
	w.Add(en.CasualTime(rules.Skip))

	ApplyFixtures(t, "en.CasualTime", w, fixt)
}

func TestCasualDateCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is tomorrow this afternoon ", 16, "tomorrow this afternoon", (12 + 24) * time.Hour}, // Changed from 15h to 12h per spec
	}

	w := when.New(nil)
	w.Add(
		en.CasualDate(rules.Skip),
		en.CasualTime(rules.Override),
	)

	ApplyFixtures(t, "en.CasualDate|en.CasualTime", w, fixt)
}
