package serialize

import (
	"bytes"
	"encoding/binary"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/decode/action"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/msq"
)

// SerializeMapData encodes map data to a byte sequence.
func SerializeMapData(md decode.MapData) []byte {
	dim := gen.Point{
		X: len(md.ActionClasses[0]),
		Y: len(md.ActionClasses),
	}

	b := make([]byte, decode.MapDataLen(dim))

	off := 0
	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			ac1 := md.ActionClasses[y][x]
			ac2 := md.ActionClasses[y][x+1]
			b[off] = byte(ac2)<<4 | byte(ac1)

			off++
		}
	}

	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			b[off] = byte(md.ActionSelectors[y][x])
			off++
		}
	}

	gen.Assert(off == len(b))

	return b
}

// SerializeCentralDir encodes a central directory to a byte sequence.
func SerializeCentralDir(cd decode.CentralDir) []byte {
	b := make([]byte, decode.CentralDirLen)

	off := 0

	writePtr := func(p int) {
		gen.Assert(off+2 <= len(b))
		copy(b[off:off+2], gen.WriteUint16(uint16(p)))
		off += 2
	}

	writePtr(cd.Strings)
	writePtr(cd.MonsterNames)
	writePtr(cd.MonsterData)
	for _, at := range cd.ActionTables {
		writePtr(at)
	}
	writePtr(cd.SpecialActions)
	writePtr(cd.NPCTable)

	gen.Assert(off == len(b))

	return b
}

// SerializeMapInfo encodes map info to a byte sequence.
func SerializeMapInfo(mi decode.MapInfo) []byte {
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

// SerializeActionTransition encodes a single transition to a byte sequence.
func SerializeActionTransition(at action.Transition) []byte {
	var buf []byte

	b0 := byte(0)
	if at.Relative {
		b0 |= 0x80
	}
	if at.Prompt {
		b0 |= 0x40
	}
	b0 |= byte(at.StringPtr)
	buf = append(buf, b0)

	buf = append(buf, byte(at.LocX))

	buf = append(buf, byte(at.LocY))

	buf = append(buf, byte(at.Location))

	buf = append(buf, byte(at.ToClass))
	if at.ToClass < action.TransitionToClassNoneMin {
		buf = append(buf, byte(at.ToSelector))
	}

	return buf
}

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
			data = append(data, SerializeActionTransition(*at)...)
		}
	}

	return data
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
	appendGenTable(tables.T5)
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

// SerializeSpecialActions encodes the special actions sequence to a byte
// sequence.
func SerializeSpecialActions(sa decode.SpecialActions) []byte {
	return sa.Actions
}

// SerializeNPCTable encodes the NPC set to a byte sequence.
func SerializeNPCTable(nt decode.NPCTable, baseOff int) []byte {
	var b []byte

	// This table always starts with a 0 pointer for some reason.
	b = append(b, gen.WriteUint16(0)...)

	for i, _ := range nt.NPCs {
		p := baseOff + 2*(len(nt.NPCs)+1) + decode.NPCSize*i
		b = append(b, gen.WriteUint16(uint16(p))...)
	}

	for _, n := range nt.NPCs {
		b = append(b, n.Data...)
	}

	return b
}

// SerializeMonsterName encodes a single monster name to a byte sequence.
func SerializeMonsterName(n decode.MonsterName) []byte {
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

// SerializeMonsterName encodes a set of monster names to a byte sequence.
func SerializeMonsterNames(mn decode.MonsterNames) []byte {
	var b []byte

	for _, n := range mn.Names {
		b = append(b, SerializeMonsterName(n)...)
	}

	return b
}

// SerializeMonsterDataElem encodes a single element of monster data to a byte
// sequence.
func SerializeMonsterDataElem(e decode.MonsterDataElem) []byte {
	b := make([]byte, decode.MonsterDataElemLen)

	copy(b[0:2], gen.WriteUint16(uint16(e.HitPoints)))
	b[2] = byte(e.HitChance)
	b[3] = byte(e.ExtraDamage)
	b[4] = byte(e.Armor) | byte(e.MaxCount)<<4
	b[5] = byte(e.AttackType) | byte(e.FixedDamage)<<4
	b[6] = byte(e.Type)
	b[7] = byte(e.CombatImage)

	return b
}

// SerializeMonsterData encodes the monster data section to a byte sequence.
func SerializeMonsterData(md decode.MonsterData) []byte {
	var b []byte

	for _, e := range md.Monsters {
		b = append(b, SerializeMonsterDataElem(e)...)
	}

	return b
}

// SerializeMsqHeader encodes an MSQ block header to a byte sequence.
func SerializeMsqHeader(hdr msq.Header) []byte {
	var out []byte

	out = append(out, []byte(msq.BlockPrefix)...)
	out = append(out, hdr.GameIdx+0x30)
	out = append(out, hdr.Xor0)
	out = append(out, hdr.Xor1)

	return out
}

// SerializeMsqBlock encodes an MSQ block to a byte sequence.
func SerializeMsqBlock(block msq.Block) []byte {
	csum := msq.CalcChecksum(block.EncSection)
	block.Hdr.Xor0 = byte(csum & 0xff)
	block.Hdr.Xor1 = byte(csum >> 8)

	hdrBytes := SerializeMsqHeader(block.Hdr)
	encBytes := msq.Encrypt(block.EncSection, block.Hdr.Xor0, block.Hdr.Xor1)

	out := append(hdrBytes, encBytes...)
	out = append(out, block.PlainSection...)

	return out
}

// SerializeGame encodes a set of MSQ blocks to a byte sequence.  The result is
// the contents of a GAMEx file.
func SerializeGame(blocks []msq.Block) []byte {
	var out []byte

	for _, b := range blocks {
		sub := SerializeMsqBlock(b)
		out = append(out, sub...)
	}

	return out
}
