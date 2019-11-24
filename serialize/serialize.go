package serialize

import (
	"github.com/badvassal/wllib/decode/action"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/msq"
)

// SerializeActionTransitions encodes a set of transitions to a byte sequence.
// baseOff is the offset of the start of the transition table relative to the
// start of the encrypted section.
func SerializeActionTransitions(ats []*action.Transition, baseOff int) []byte {
	// Reserve room for the leading set of pointers.
	data := make([]byte, len(ats)*2)

	for i, at := range ats {
		// Write pointer to start.
		if at == nil {
			copy(data[i*2:i*2+2], gen.WriteUint16(uint16(0)))
		} else {
			p := baseOff + len(data)
			copy(data[i*2:i*2+2], gen.WriteUint16(uint16(p)))

			// Append transition definition to end.
			data = append(data, action.EncodeActionTransition(*at)...)
		}
	}

	return data
}

func SerializeActionLoots(loots []*action.Loot, baseOff int) []byte {
	t := gen.Table{}

	for _, loot := range loots {
		if loot != nil {
			t.Elems = append(t.Elems, action.EncodeLoot(*loot))
		}
	}

	return t.Encode(baseOff)
}

// SerializeActionTables encodes the set of action tables to a byte sequence.
// baseOff is the offset of the start of the first action table relative to the
// start of the encrypted section.
func SerializeActionTables(tables action.Tables, baseOff int) []byte {
	var b []byte

	off := baseOff

	appendBlob := func(blob []byte) {
		b = append(b, blob...)
		off += len(blob)
	}

	appendGenTable := func(t gen.Table) {
		appendBlob(t.Encode(off))
	}

	appendGenTable(tables.T0)
	appendGenTable(tables.T1)
	appendGenTable(tables.T2)
	appendGenTable(tables.T3)
	appendGenTable(tables.T4)
	appendBlob(SerializeActionLoots(tables.Loots, off))
	appendGenTable(tables.T6)
	appendGenTable(tables.T7)
	appendGenTable(tables.T8)
	appendGenTable(tables.T9)
	appendBlob(SerializeActionTransitions(tables.Transitions, off))
	appendGenTable(tables.T11)
	appendGenTable(tables.T12)
	appendGenTable(tables.T13)
	appendGenTable(tables.T14)
	appendGenTable(tables.T15)

	return b
}

// SerializeGame encodes a set of MSQ blocks to a byte sequence.  The result is
// the contents of a GAMEx file.
func SerializeGame(blocks []msq.Block) []byte {
	var out []byte

	for _, b := range blocks {
		sub := msq.EncodeMsqBlock(b)
		out = append(out, sub...)
	}

	return out
}
