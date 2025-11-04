package clocks

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDedupClocks(t *testing.T) {
	aTime, err := time.Parse(time.RFC3339, "2023-01-26T10:50:40-07:00")
	require.NoError(t, err)
	aTime2 := aTime.Add(time.Hour)
	aTime3 := aTime2.Add(time.Hour)

	clocks := []TimeAndZone{
		{Time: timeAtTimeZone(aTime, "MST"), Zone: "MST"},
		{Time: timeAtTimeZone(aTime, "America/Boise"), Zone: "America/Boise"},
		{Time: timeAtTimeZone(aTime2, "EST"), Zone: "EST"},
		{Time: timeAtTimeZone(aTime2, "Jamaica"), Zone: "Jamaica"},
		{Time: timeAtTimeZone(aTime3, "UTC"), Zone: "UTC"},
	}
	result := dedupClocks(clocks)

	expected := GroupedTimes{
		clocks[0].DedupString(): {clocks[0], clocks[1]},
		clocks[2].DedupString(): {clocks[2], clocks[3]},
		clocks[4].DedupString(): {clocks[4]},
	}

	require.Equal(t, len(expected), len(result))
	for i, k := range expected {
		require.Equal(t, k, result[i])
	}
}
