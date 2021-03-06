package decode

import (
	"bytes"
	"encoding/binary"

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
		// Ignore error; we already verified source length.
		p, _ := gen.ReadUint16(data[off : off+2])
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

// EncodeMapInfo encodes map info to a byte sequence.
func EncodeMapInfo(mi MapInfo) []byte {
	buf := &bytes.Buffer{}

	// Two bytes of padding.
	buf.WriteByte(0)
	buf.WriteByte(0)

	buf.WriteByte(byte(mi.Size))

	// Two bytes of padding.
	buf.WriteByte(0)
	buf.WriteByte(0)

	buf.WriteByte(byte(mi.EncounterFreq))
	buf.WriteByte(byte(mi.TileSet))
	buf.WriteByte(byte(mi.MaxMonsters))
	buf.WriteByte(byte(mi.MaxEncounters))
	buf.WriteByte(byte(mi.OobTile))
	buf.WriteByte(byte(mi.SecRate))
	buf.WriteByte(byte(mi.MinRate))
	buf.WriteByte(byte(mi.HealRate))

	for _, sid := range mi.StringIDs {
		binary.Write(buf, binary.LittleEndian, uint16(sid))
	}

	// One byte of padding.
	buf.WriteByte(0)

	return buf.Bytes()
}
