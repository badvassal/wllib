package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

const (
	NPCTableMinLen = 2
	NPCSize        = 256
)

// NPC represents encoded NPC data.
type NPC struct {
	Data []byte
}

// NPCTable represents an MSQ block's set of NPCs.
type NPCTable struct {
	NPCs []NPC
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
	ptrs, _, err := gen.ReadPointers(source[2:], baseOff+2)
	if err != nil {
		return nil, 0, err
	}

	ptrsSize := (1 + len(ptrs)) * 2
	dataLen := len(ptrs) * NPCSize
	totalLen := ptrsSize + dataLen
	if len(source) < totalLen {
		return nil, 0, wlerr.Errorf(
			"failed to parse NPC table: data too short: "+
				"have=%d want>=%d ptrs=%+v",
			len(source), totalLen, ptrs)
	}

	if len(ptrs) > 1 {
		prev := ptrs[0]
		for _, p := range ptrs[1:] {
			if p != prev+NPCSize {
				return nil, 0, wlerr.Errorf(
					"failed to parse NPC table: pointer span incorrect: "+
						"have=%d want=%d ptrs=%+v",
					p-prev, NPCSize, ptrs)
			}

			prev = p
		}
	}

	nt := &NPCTable{}
	for i, _ := range ptrs {
		off := ptrsSize + i*NPCSize
		end := ptrsSize + (i+1)*NPCSize
		nt.NPCs = append(nt.NPCs, NPC{
			Data: source[off:end],
		})
	}

	return nt, totalLen, nil
}

// EncodeNPCTable encodes the NPC set to a byte sequence.
func EncodeNPCTable(nt NPCTable, baseOff int) []byte {
	var b []byte

	// This table always starts with a 0 pointer for some reason.
	b = append(b, gen.WriteUint16(0)...)

	for i, _ := range nt.NPCs {
		p := baseOff + 2*(len(nt.NPCs)+1) + NPCSize*i
		b = append(b, gen.WriteUint16(uint16(p))...)
	}

	for _, n := range nt.NPCs {
		b = append(b, n.Data...)
	}

	return b
}
