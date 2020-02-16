package modify

import (
	"fmt"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/decode/action"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/msq"
	"github.com/badvassal/wllib/serialize"
)

// BlockModifier is used for modifying an MSQ block body.
//
// XXX: Note: This data structure is just an intermediate solution.  Ideally,
// we would modify a block with the following sequence:
//
// 1. Decode block (create a decode.Block)
// 2. Modify decoded block.
// 3. Re-encode decoded block from scratch.
//
// Unfortunately, this sequence causes some blocks to be resized (removal of
// padding bytes).  Resizing blocks leads Wasteland to report a data corruption
// error on startup.  If we can ever make block resizing work, then we can
// remove this type (and indeed this entire file).
type BlockModifier struct {
	body msq.Body
	dim  gen.Point
}

func NewBlockModifier(body msq.Body, dim gen.Point) *BlockModifier {
	return &BlockModifier{
		body: body,
		dim:  dim,
	}
}

// offsetPair converts an absolute MSQ block body offset to a lower level
// representation.  It returns true if the offset can be found in the secure
// section of the block; false otherwise.  The second return value is the
// offset within the relevant section.
func (m *BlockModifier) offsetPair(off int) (bool, int, error) {
	if off < len(m.body.SecSection) {
		return true, off, nil
	}

	plainOff := off - len(m.body.SecSection)
	if plainOff < len(m.body.PlainSection) {
		return false, plainOff, nil
	}

	return false, 0, fmt.Errorf(
		"invalid block body offset: have=%d want<%d",
		off, len(m.body.SecSection)+len(m.body.PlainSection))
}

// writeBytesToBuf copies data from one byte slice to another at a specified
// destination offset.  If the data would extend beyond the end of the
// destination slice, the slice is extended to accommodate the source data.
func writeBytesToBuf(src []byte, dst *[]byte, off int) {
	end := off + len(src)

	extra := end - len(*dst)
	if extra > 0 {
		*dst = append(*dst, make([]byte, extra)...)
	}

	copy((*dst)[off:off+len(src)], src)
}

// writeBytes writes bytes to the MSQ block body at the specified offset.
func (m *BlockModifier) writeBytes(data []byte, off int) error {
	isEnc, off, err := m.offsetPair(off)
	if err != nil {
		return err
	}

	var dest *[]byte
	if isEnc {
		dest = &m.body.SecSection
	} else {
		dest = &m.body.PlainSection
	}

	writeBytesToBuf(data, dest, off)

	return nil
}

// ReplaceMapInfo replaces an MSQ block's map info section with the specified
// one.
func (m *BlockModifier) ReplaceMapInfo(mi decode.MapInfo) error {
	decBlock, err := decode.DecodeBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	smi := decode.EncodeMapInfo(mi)
	if len(smi) != decBlock.Sizes.MapInfo {
		return fmt.Errorf("map infos differ in size: old=%d new=%d",
			decBlock.Sizes.MapInfo, len(smi))
	}

	if err := m.writeBytes(smi, decBlock.Offsets.MapInfo); err != nil {
		return err
	}

	return nil
}

// ReplaceMonsterData replaces an MSQ block's monster data section with the
// specified one.
func (m *BlockModifier) ReplaceMonsterData(md decode.MonsterData) error {
	decBlock, err := decode.DecodeBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	smd := decode.EncodeMonsterData(md)
	if len(smd) != decBlock.Sizes.MonsterData {
		return fmt.Errorf("monster datas differ in size: old=%d new=%d",
			decBlock.Sizes.MonsterData, len(smd))
	}

	if err := m.writeBytes(smd, decBlock.Offsets.MonsterData); err != nil {
		return err
	}

	return nil
}

// ReplaceMonsterData replaces an MSQ block's loot section with the specified
// one.
func (m *BlockModifier) ReplaceLoots(loots []*action.Loot) error {
	db, err := decode.DecodeBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	cb, err := decode.CarveBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	st := serialize.SerializeActionLoots(loots, db.CentralDir.ActionTables[action.IDLoot])
	overflow := len(st) - len(cb.ActionTables[action.IDLoot])

	if overflow != 0 {
		return wlerr.Errorf(
			"failed to replace loot table: "+
				"replacement has size different from original: overflow=%d",
			overflow)
	}

	off := db.Offsets.ActionTables[action.IDLoot]
	copy(m.body.SecSection[off:off+len(st)], st)

	return nil
}

// ReplaceActionTransitions replaces an MSQ block's transitions action table
// with the specified one.
func (m *BlockModifier) ReplaceActionTransitions(transitions []*action.Transition) error {
	db, err := decode.DecodeBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	cb, err := decode.CarveBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	db.ActionTables.Transitions = transitions

	st := serialize.SerializeActionTransitions(transitions,
		db.CentralDir.ActionTables[action.IDTransition])
	overflow := len(st) - len(cb.ActionTables[action.IDTransition])

	if overflow != 0 {
		return wlerr.Errorf(
			"failed to replace action transitions: "+
				"replacement has size different from original: overflow=%d",
			overflow)
	}

	off := db.Offsets.ActionTables[action.IDTransition]
	copy(m.body.SecSection[off:off+len(st)], st)

	return nil
}

// ReplaceNPCTable replaces an MSQ block's NPC table with the specified one.
func (m *BlockModifier) ReplaceNPCTable(npcTable decode.NPCTable) error {
	if len(npcTable.NPCs) == 0 {
		return nil
	}

	db, err := decode.DecodeBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	cb, err := decode.CarveBlock(m.body, m.dim)
	if err != nil {
		return err
	}

	en, err := decode.EncodeNPCTable(npcTable, db.CentralDir.NPCTable)
	if err != nil {
		return err
	}
	overflow := len(en) - len(cb.NPCTable)

	if overflow != 0 {
		return wlerr.Errorf(
			"failed to replace NPC table: "+
				"replacement has size different from original: overflow=%d",
			overflow)
	}

	off := db.Offsets.NPCTable
	copy(m.body.SecSection[off:off+len(en)], en)

	return nil
}

// Body returns a BlockModifier's modified body.
func (m *BlockModifier) Body() msq.Body {
	return m.body
}
