package clocks

import (
	"github.com/maxrichie5/mac-clock-toolbar/internal/errpkg"
	"os"
	"sort"
	"strings"
)

type ZoneList []string

func (z ZoneList) Sort() {
	sort.Slice(z, func(i, j int) bool {
		return z[i] < z[j]
	})
}

func loadZones() ZoneList {
	// All zones were retrieved by running `sh get-zones.sh`
	f, err := os.ReadFile("zones.txt")
	errpkg.CheckFatalErr(err)

	fs := strings.Split(string(f), "\n")
	zns := make(ZoneList, 0, len(fs))
	for _, z := range fs {
		zns = append(zns, z)
	}

	zns.Sort() // TODO: sorting goes away on click
	return zns
}
