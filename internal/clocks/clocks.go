package clocks

import (
	"fmt"
	"github.com/caseymrm/menuet"
	"github.com/maxrichie5/mac-clock-toolbar/internal/cfg"
	"log"
	"sort"
	"strings"
	"time"
)

const hourSecMinZoneTimeFmt = "15:04:05 MST"
const zoneTimeFmt = "MST"

var zones ZoneList

type TimeAndZone struct {
	Time time.Time
	Zone string
}

func (tz TimeAndZone) DedupString() string {
	return tz.Time.Format(hourSecMinZoneTimeFmt)
}

type GroupedTimes map[string][]TimeAndZone

func Start() {
	zones = loadZones()
}

func GetActiveClocks() string {
	now := time.Now()
	activeZones := cfg.GetActiveTimeZones()
	if len(activeZones) == 0 {
		return timeAtTimeZone(now, time.Local.String()).Format(hourSecMinZoneTimeFmt)
	}

	clks := make([]string, 0, len(activeZones))

	for i := range activeZones {
		z := activeZones[i]
		clks = append(clks, timeAtTimeZone(now, z).Format(hourSecMinZoneTimeFmt))
	}

	sort.Slice(clks, func(i, j int) bool {
		return clks[i] < clks[j]
	})
	return strings.Join(clks, " ")
}

func GetClocksMenu() []menuet.MenuItem {
	now := time.Now()
	clks := make([]TimeAndZone, 0, 10)
	for i := range zones {
		zone := zones[i]
		t := timeAtTimeZone(now, zone)
		if t.IsZero() {
			continue
		}

		clks = append(clks, TimeAndZone{Time: t, Zone: zone})
	}

	dedupdClks := dedupClocks(clks)
	menu := make([]menuet.MenuItem, 0, len(dedupdClks))
	for _, dups := range dedupdClks {
		// display selected clocks first
		if cfg.TimeZomeActive(dups[0].Zone) {
			menu = append([]menuet.MenuItem{groupedTimeToMenuItem(dups)}, menu...)
			continue
		}
		menu = append(menu, groupedTimeToMenuItem(dups))
	}

	// sort by formatted zone name and if active
	sort.Slice(menu, func(i, j int) bool {
		if menu[i].State != menu[j].State {
			return menu[i].State
		}
		return strings.Split(menu[i].Text, " ")[1] < strings.Split(menu[j].Text, " ")[1]
	})
	return menu
}

func timeAtTimeZone(t time.Time, zone string) time.Time {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		log.Println(err)
		return time.Time{}
	}

	return t.In(loc)
}

func dedupClocks(clks []TimeAndZone) GroupedTimes {
	dedupd := make(GroupedTimes, len(clks))
	for _, c := range clks {
		if _, ok := dedupd[c.DedupString()]; !ok {
			dedupd[c.DedupString()] = make([]TimeAndZone, 0, 3)
		}
		dedupd[c.DedupString()] = append(dedupd[c.DedupString()], c)
	}
	return dedupd
}

func groupedTimeToMenuItem(dups []TimeAndZone) menuet.MenuItem {
	var (
		single      = len(dups) == 1
		state       = false
		first       = dups[0]
		childFunc   func() []menuet.MenuItem
		clickedFunc func()
	)

	if !single {
		childFunc = func() []menuet.MenuItem {
			children := make([]menuet.MenuItem, 0, len(dups))
			for i := range dups {
				tz := dups[i]
				// if the zone name matches the formatted zone name display it first
				if tz.Zone == tz.Time.Format(zoneTimeFmt) {
					children = append([]menuet.MenuItem{timeAndZoneChildToMenuItem(tz)}, children...)
					continue
				}
				children = append(children, timeAndZoneChildToMenuItem(tz))
			}
			return children
		}
		state = anyTimeZoneIsActive(dups)
	} else {
		clickedFunc = func() {
			if cfg.TimeZomeActive(first.Zone) {
				cfg.RemoveActiveTimeZone(first.Zone)
			} else {
				cfg.AddActiveTimeZone(first.Zone)
			}
		}
		state = cfg.TimeZomeActive(first.Zone)
	}

	return menuet.MenuItem{
		Text:     first.Time.Format(hourSecMinZoneTimeFmt),
		State:    state,
		Clicked:  clickedFunc,
		Children: childFunc,
	}
}

func timeAndZoneChildToMenuItem(tz TimeAndZone) menuet.MenuItem {
	return menuet.MenuItem{
		Text:  fmt.Sprintf("%s (%s)", tz.Time.Format(hourSecMinZoneTimeFmt), tz.Zone),
		State: cfg.TimeZomeActive(tz.Zone),
		Clicked: func() {
			if cfg.TimeZomeActive(tz.Zone) {
				cfg.RemoveActiveTimeZone(tz.Zone)
			} else {
				cfg.AddActiveTimeZone(tz.Zone)
			}
		},
	}
}

func anyTimeZoneIsActive(tzs []TimeAndZone) bool {
	for _, tz := range tzs {
		if cfg.TimeZomeActive(tz.Zone) {
			return true
		}
	}
	return false
}
