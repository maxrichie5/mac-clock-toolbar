package clocks

import (
	"fmt"
	"github.com/caseymrm/menuet"
	"log"
	"sort"
	"strings"
	"time"
)

const timeDisplayFmt = "15:04:05 MST"

var zones ZoneList

type TimeAndZone struct {
	Time time.Time
	Zone string
}

func (tz TimeAndZone) DedupString() string {
	return tz.Time.Format(timeDisplayFmt)
}

type GroupedTimes map[string][]TimeAndZone

func Start() {
	zones = loadZones()
}

func GetActiveClocks() string {
	now := time.Now()
	activeZones := GetActiveTimeZones()
	if len(activeZones) == 0 {
		return timeAtTimeZone(now, time.Local.String()).Format(timeDisplayFmt)
	}

	clks := make([]string, 0, len(activeZones))

	for i := range activeZones {
		z := activeZones[i]
		clks = append(clks, timeAtTimeZone(now, z).Format(timeDisplayFmt))
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
		menu = append(menu, groupedTimeToMenuItem(dups))
	}

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
	var childFunc func() []menuet.MenuItem
	if len(dups) > 1 {
		childFunc = func() []menuet.MenuItem {
			children := make([]menuet.MenuItem, 0, len(dups))
			for i := range dups {
				tz := dups[i]
				children = append(children, timeAndZoneChildToMenuItem(tz))
			}
			return children
		}
	}

	// TODO: single state and single clicked
	return menuet.MenuItem{
		Text:     dups[0].Time.Format(timeDisplayFmt),
		State:    false,
		Clicked:  nil,
		Children: childFunc,
	}
}

func timeAndZoneChildToMenuItem(tz TimeAndZone) menuet.MenuItem {
	return menuet.MenuItem{
		Text:  fmt.Sprintf("%s (%s)", tz.Time.Format(timeDisplayFmt), tz.Zone),
		State: TimeZomeActive(tz.Zone),
		Clicked: func() {
			if TimeZomeActive(tz.Zone) {
				RemoveActiveTimeZone(tz.Zone)
			} else {
				AddActiveTimeZone(tz.Zone)
			}
		},
	}
}
