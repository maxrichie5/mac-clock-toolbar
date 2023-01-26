package clocks

import (
	"github.com/maxrichie5/mac-clock-toolbar/internal/errpkg"
	"os"
	"sort"
)

type ZoneList []string

func (z ZoneList) Sort() {
	sort.Slice(z, func(i, j int) bool {
		return z[i] < z[j]
	})
}

func loadZones() ZoneList {
	// All zones were retrieved from /opt/homebrew/Cellar/go/1.19.1/libexec/lib/time/zoneinfo.zip
	zonesDir := "zones"
	zoneFiles, err := os.ReadDir(zonesDir)
	errpkg.CheckFatalErr(err)

	zns := make(ZoneList, 0, len(zoneFiles))
	for _, f := range zoneFiles {
		if f.IsDir() {
			// TODO nested dirs
			subzoneFiles, err := os.ReadDir(zonesDir + string(os.PathSeparator) + f.Name())
			errpkg.CheckFatalErr(err)

			for _, ff := range subzoneFiles {
				zns = append(zns, f.Name()+string(os.PathSeparator)+ff.Name())
			}
		} else {
			zns = append(zns, f.Name())
		}
	}

	zns.Sort()
	return zns
}
