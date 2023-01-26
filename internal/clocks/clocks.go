package clocks

import (
	"github.com/caseymrm/menuet"
	"log"
	"strings"
	"time"
)

const fmtTime = "15:04:05 MST"

var (
	zones = []string{
		"HST", // Hawaii
		"AST", // Alaska
		//"UTC-08:00",
		//"PDT",
		//"PST",
		//"PT",
		"MST", // Mountain
		"CST", // Central
		"EST", // Eastern
		"UTC",
	}
)

func GetActiveClocks() string {
	activeZones := GetActiveTimeZones()
	if len(activeZones) == 0 {
		return nowAtTimeZone(time.Local.String()).Format(fmtTime)
	}

	clks := make([]string, 0, len(activeZones))

	for i := range activeZones {
		z := activeZones[i]
		clks = append(clks, nowAtTimeZone(z).Format(fmtTime))
	}

	return strings.Join(clks, " ")
}

func GetAllClocks() []menuet.MenuItem {
	clks := make([]menuet.MenuItem, 0, 10)
	for i := range zones {
		zone := zones[i]
		t := nowAtTimeZone(zone)
		if t.IsZero() {
			continue
		}

		clks = append(clks, menuet.MenuItem{
			Type:       "",
			Text:       t.Format(fmtTime),
			Image:      "",
			FontSize:   0,
			FontWeight: 0,
			State:      TimeZomeActive(zone),
			Clicked: func() {
				if TimeZomeActive(zone) {
					RemoveActiveTimeZone(zone)
				} else {
					AddActiveTimeZone(zone)
				}
			},
			Children: nil,
		})
	}
	return clks
}

func nowAtTimeZone(zone string) time.Time {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		log.Println(err)
		return time.Time{}
	}

	return time.Now().In(loc)
}
