package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

// CentralDirLen is the size, in bytes, of an MSQ block's map info.
const MapInfoLen = 50

// MapInfo contains some properties of an MSQ block's map.  See
// <https://wasteland.gamepedia.com/Map_Info>.
type MapInfo struct {
	Size          int
	EncounterFreq int
	TileSet       int
	MaxMonsters   int
	MaxEncounters int
	OobTile       int
	SecRate       int
	MinRate       int
	HealRate      int
	StringIDs     []int
}

// DecodeCentralDir parses map info from a sequence of bytes.
func DecodeMapInfo(data []byte) (*MapInfo, error) {
	if len(data) < MapInfoLen {
		return nil, wlerr.Errorf(
			"failed to decode map info: too few bytes: have=%d want>=%d",
			len(data), MapInfoLen)
	}

	off := 0
	readByte := func() int {
		b := data[off]
		off++

		return int(b)
	}
	readPtr := func() int {
		p, err := gen.ReadUint16(data[off : off+2])
		if err != nil {
			panic(err.Error())
		}
		off += 2
		return p
	}

	mi := &MapInfo{}

	off += 2 // Two bytes of padding.
	mi.Size = readByte()
	off += 2 // Two bytes of padding.
	mi.EncounterFreq = readByte()
	mi.TileSet = readByte()
	mi.MaxMonsters = readByte()
	mi.MaxEncounters = readByte()
	mi.OobTile = readByte()
	mi.SecRate = readByte()
	mi.MinRate = readByte()
	mi.HealRate = readByte()

	for i := 0; i < 18; i++ {
		mi.StringIDs = append(mi.StringIDs, readPtr())
	}

	off++ // One byte of padding (I think).

	return mi, nil
}
