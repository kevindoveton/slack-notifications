package internal

import (
	"strconv"
	"testing"
	"time"
)

func TestShouldAlert(t *testing.T) {
	tables := []struct {
		mode   Mode
		ID     string
		time   time.Time
		result bool
	}{
		// no overrides, only show
		{Mode{Order: []string{"show"}}, "@test", getDate("mon", 12, 12), true},

		// no overrides, only hide
		{Mode{Order: []string{"hide"}}, "@test", getDate("mon", 12, 12), false},

		// irelevant override, should ignore the override
		{
			Mode{Order: []string{"show", "overrides"}, Overrides: []Rule{Rule{Channel: "@test", Day: "tues", Type: HIDE, Before: "23:59", After: "00:00"}}},
			"@test",
			getDate("mon", 12, 12),
			true,
		},

		// relevant override, should respect it
		{
			Mode{Order: []string{"show", "overrides"}, Overrides: []Rule{Rule{Channel: "@test", Day: "mon", Type: HIDE, Before: "23:59", After: "00:00"}}},
			"@test",
			getDate("mon", 12, 12),
			false,
		},

		// relevant override, should respect it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "00:00"}}},
			"@test",
			getDate("mon", 12, 12),
			true,
		},

		// irelevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "00:00"}}},
			"@kevin",
			getDate("mon", 12, 12),
			false,
		},

		// irelevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "00:00"},
				Rule{Channel: "@kevin2", Day: "mon", Type: SHOW, Before: "23:59", After: "00:00"},
			}},
			"@kevin",
			getDate("mon", 12, 12),
			false,
		},

		// relevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "00:00"},
				Rule{Channel: "@kevin", Day: "mon", Type: SHOW, Before: "23:59", After: "00:00"},
			}},
			"@kevin",
			getDate("mon", 12, 12),
			true,
		},

		// time based relevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "12:13", After: "00:00"},
			}},
			"@test",
			getDate("mon", 12, 12),
			true,
		},

		// relevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "11:13", After: "00:00"},
			}},
			"@test",
			getDate("mon", 12, 12),
			false,
		},

		// relevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "12:12"},
			}},
			"@test",
			getDate("mon", 12, 12),
			true,
		},

		// relevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "12:13"},
			}},
			"@test",
			getDate("mon", 12, 12),
			false,
		},

		// relevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "11:12"},
			}},
			"@test",
			getDate("mon", 12, 12),
			true,
		},

		// relevant override, should ignore it
		{
			Mode{Order: []string{"hide", "overrides"}, Overrides: []Rule{
				Rule{Channel: "@test", Day: "mon", Type: SHOW, Before: "23:59", After: "11:13"},
			}},
			"@test",
			getDate("mon", 12, 12),
			true,
		},
	}

	for idx, table := range tables {
		alert := ShouldAlert(table.mode, table.ID, table.time)

		if alert != table.result {
			t.Error(strconv.Itoa(idx) + ": alert should be " + strconv.FormatBool(table.result))
		}
	}
}

func getDate(day string, hour int, min int) time.Time {
	dayOfWeek := strDayToInt(day)

	return time.Date(2018, 01, dayOfWeek, hour, min, 0, 0, time.Local)
}
