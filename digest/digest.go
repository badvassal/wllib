package digest

import (
	"fmt"
	"strings"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/wlstrings"
)

// Monster is an element of monster data with some user friendly annotations.
type Monster struct {
	Name string
	Elem decode.MonsterDataElem
}

// MonsterNameSingular calculates the singular form of a monster name.
func MonsterNameSingular(name decode.MonsterName) string {
	return name.Start + name.MidSingular + name.End
}

// DecompressStringsArea decodes a set of compressed strings into ASCII text.
func DecompressStringsArea(sa decode.StringsArea) ([][]byte, error) {
	cgs := make([][]byte, len(sa.Pointers))
	for i, p := range sa.Pointers {
		start := p - sa.Pointers[0]

		var end int
		if i < len(sa.Pointers)-1 {
			end = sa.Pointers[i+1] - sa.Pointers[0]
		} else {
			end = len(sa.StringData)
		}

		cgs[i] = sa.StringData[start:end]
	}

	dgs := make([][]byte, len(cgs))
	for i, cg := range cgs {
		dg, err := wlstrings.DecompressStringGroup(sa.CharTable, cg)
		if err != nil {
			return nil, err
		}

		dgs[i] = dg
	}

	return dgs, nil
}

// MapDataString converts an instance of map data into a user friendly string.
func MapDataString(md decode.MapData) string {
	var lines []string

	for y := 0; y < len(md.ActionClasses); y++ {
		line := ""

		for x := 0; x < len(md.ActionClasses[0]); x++ {
			ac := md.ActionClasses[y][x]
			as := md.ActionSelectors[y][x]

			line += fmt.Sprintf("%1x%02x ", ac, as)
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
