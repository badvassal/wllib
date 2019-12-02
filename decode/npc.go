package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

const (
	NPCTableMinLen = 2
	NPCSize        = 256
)

// NPCTable represents an MSQ block's set of NPCs.
type NPCTable struct {
	NPCs []Character
}

// DecodeNPCTable parses a set of NPCs from a sequence of bytes.  baseOff is
// the offset of the start of the NPC table relative to the start of the MSQ
// block's encoded section.  It returns the decoded NPC table and its size, in
// bytes.
func DecodeNPCTable(source []byte, baseOff int) (*NPCTable, int, error) {
	if len(source) < NPCTableMinLen {
		return nil, 0, wlerr.Errorf(
			"data too short: have=%d want>=%d", len(source), NPCTableMinLen)
	}

	// The first pointer is always 0x0000 for some reason.  Ignore it.
	t, err := gen.ParseTable(source[2:], baseOff+2)
	if err != nil {
		return nil, 0, err
	}

	nt := &NPCTable{}
	for _, e := range t.Elems {
		n, err := DecodeCharacter(e)
		if err != nil {
			return nil, 0, wlerr.Wrapf(err, "failed to decode NPC")
		}

		nt.NPCs = append(nt.NPCs, *n)
	}

	totalLen := (1+len(t.Elems))*2 + len(t.Elems)*CharacterSize

	return nt, totalLen, nil
}

// EncodeNPCTable encodes the NPC set to a byte sequence.
func EncodeNPCTable(nt NPCTable, baseOff int) ([]byte, error) {
	onErr := wlerr.MakeWrapper("failed to encode NPC table")

	var b []byte

	// This table always starts with a 0 pointer for some reason.
	b = append(b, gen.WriteUint16(0)...)
	baseOff += 2

	t := gen.Table{}
	for _, n := range nt.NPCs {
		ch, err := EncodeCharacter(n)
		if err != nil {
			return nil, onErr(err, "")
		}
		t.Elems = append(t.Elems, ch)
	}

	b = append(b, t.Encode(baseOff)...)

	return b, nil
}
