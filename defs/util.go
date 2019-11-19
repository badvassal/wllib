package defs

import (
	"strings"

	"github.com/badvassal/wllib/gen/wlerr"
)

// LocationString produces a string representation of a location.
func LocationString(loc int) string {
	s := LocationNameMap[loc]
	if s == "" {
		s = "???"
	}
	return s
}

// ParseLocation converts a string into its corresponding location code.
func ParseLocation(s string) (int, error) {
	for k, v := range LocationNameMap {
		if s == v {
			return k, nil
		}
	}

	return 0, wlerr.Errorf("invalid location string: %s", s)
}

func ParseLocationNoCase(s string) (int, error) {
	for k, v := range LocationNameMap {
		if strings.EqualFold(s, v) {
			return k, nil
		}
	}

	return 0, wlerr.Errorf("invalid location string: %s", s)
}

// LocationIsDerelict indicates whether a given location code is a derelict
// building.
func LocationIsDerelict(locID int) bool {
	return locID == LocationQuartzDerelictBuildings ||
		locID == LocationLasVegasDerelictBuildings
}

// BlockZIPToLoc retrieves the location code that is equivalent to the given
// block ZIP.
func BlockZIPToLoc(bz BlockZIP) (int, error) {
	for loc, z := range LocationBlockZIPMap {
		if bz.GameIdx == z.GameIdx &&
			bz.BlockIdx == z.BlockIdx {
			return loc, nil
		}
	}

	return 0, wlerr.Errorf("invalid block zip: %+v", bz)
}
