package decode

import (
	"encoding/hex"
	"strings"

	"github.com/badvassal/wllib/gen/wlerr"
)

// MonsterName describes how to string-encode a single monster type.
type MonsterName struct {
	Start       string
	MidSingular string
	MidPlural   string
	End         string
}

// MonsterNames contains the name information of all monsters in an MSQ block.
// See <https://wasteland.gamepedia.com/Map_Data_Monster_Names>.
type MonsterNames struct {
	Names []MonsterName
}

// DecodeCentralDir parses a set of monster names from a sequence of bytes.
func DecodeMonsterNames(data []byte) (*MonsterNames, error) {
	var segs [][]byte

	var seg []byte
	for _, b := range data {
		if b == 0 {
			segs = append(segs, seg)
			seg = nil
		} else {
			seg = append(seg, b)
		}
	}

	if len(seg) > 1 || len(seg) == 1 && seg[0] != 0xff {
		return nil, wlerr.Errorf(
			"incomplete monster name: %s", hex.EncodeToString(seg))
	}

	names := make([]MonsterName, len(segs))

	for i, seg := range segs {
		name := MonsterName{}
		parts := strings.Split(string(seg), string([]byte{0x0a}))
		if len(parts) > 0 {
			name.Start = parts[0]
		}
		if len(parts) > 1 {
			name.MidSingular = parts[1]
		}
		if len(parts) > 2 {
			name.MidPlural = parts[2]
		}
		if len(parts) > 3 {
			name.End = parts[3]
		}

		names[i] = name
	}

	return &MonsterNames{
		Names: names,
	}, nil
}

// EncodeMonsterName encodes a single monster name to a byte sequence.
func EncodeMonsterName(n MonsterName) []byte {
	var b []byte

	b = append(b, []byte(n.Start)...)
	b = append(b, 0x0a)
	b = append(b, []byte(n.MidSingular)...)
	b = append(b, 0x0a)
	b = append(b, []byte(n.MidPlural)...)
	b = append(b, 0x0a)
	b = append(b, []byte(n.End)...)
	b = append(b, 0x00)

	return b
}

// EncodeMonsterName encodes a set of monster names to a byte sequence.
func EncodeMonsterNames(mn MonsterNames) []byte {
	var b []byte

	for _, n := range mn.Names {
		b = append(b, EncodeMonsterName(n)...)
	}

	return b
}
