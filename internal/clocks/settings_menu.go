package clocks

import (
	"github.com/caseymrm/menuet"
	"github.com/maxrichie5/mac-clock-toolbar/internal/cfg"
)

func GetSettingsMenu() []menuet.MenuItem {
	timeFmtChildrenClickFunc := func(fmt string, setFunc func(string)) func() {
		return func() {
			setFunc(fmt)
		}
	}
	timeFmtChildrenFunc := func(setFunc func(string)) func() []menuet.MenuItem {
		return func() []menuet.MenuItem {
			return []menuet.MenuItem{
				{Text: hourZoneTimeFmt, State: , Clicked: timeFmtChildrenClickFunc(hourZoneTimeFmt, setFunc)},
				{Text: hourMinZoneTimeFmt, Clicked: timeFmtChildrenClickFunc(hourMinZoneTimeFmt, setFunc)},
				{Text: hourMinSecZoneTimeFmt, Clicked: timeFmtChildrenClickFunc(hourMinSecZoneTimeFmt, setFunc)},
			}
		}
	}

	return []menuet.MenuItem{
		{
			Text: "Settings",
			Children: func() []menuet.MenuItem {
				return []menuet.MenuItem{
					{
						Text:     "Active Clocks Format",
						Children: timeFmtChildrenFunc(cfg.SetActiveTimeZoneFmt),
					},
					{
						Text:     "Clocks Menu Time Format",
						Children: timeFmtChildrenFunc(cfg.SetTimeZoneMenuFmt),
					},
					{
						Text: "Hour Time Format",
					},
				}
			},
		},
	}
}
