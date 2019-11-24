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

// BlockModifier is used for modifying an MSQ block.
type BlockModifier struct {
	block msq.Block
	dim   gen.Point
}

func NewBlockModifier(block msq.Block, dim gen.Point) *BlockModifier {
	return &BlockModifier{
		block: block,
		dim:   dim,
	}
}

// offsetPair converts an absolute MSQ block offset to a lower level
// representation.  It returns true if the offset can be found in the encrypted
// section of the block; false otherwise.  The second return value is the
// offset within the relevant section.
func (m *BlockModifier) offsetPair(off int) (bool, int, error) {
	if off < len(m.block.EncSection) {
		return true, off, nil
	}

	plainOff := off - len(m.block.EncSection)
	if plainOff < len(m.block.PlainSection) {
		return false, plainOff, nil
	}

	return false, 0, fmt.Errorf(
		"invalid block offset: have=%d want<%d",
		off, len(m.block.EncSection)+len(m.block.PlainSection))
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

// writeBytes writes bytes to the MSQ block at the specified offset.
func (m *BlockModifier) writeBytes(data []byte, off int) error {
	isEnc, off, err := m.offsetPair(off)
	if err != nil {
		return err
	}

	var dest *[]byte
	if isEnc {
		dest = &m.block.EncSection
	} else {
		dest = &m.block.PlainSection
	}

	writeBytesToBuf(data, dest, off)

	return nil
}

// ReplaceMapInfo replaces an MSQ block's map info section with the specified
// one.
func (m *BlockModifier) ReplaceMapInfo(mi decode.MapInfo) error {
	decBlock, err := decode.DecodeBlock(m.block, m.dim)
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
	decBlock, err := decode.DecodeBlock(m.block, m.dim)
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
	db, err := decode.DecodeBlock(m.block, m.dim)
	if err != nil {
		return err
	}

	cb, err := decode.CarveBlock(m.block, m.dim)
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
	copy(m.block.EncSection[off:off+len(st)], st)

	return nil
}

// ReplaceActionTransitions replaces an MSQ block's transitions action table
// with the specified one.
func (m *BlockModifier) ReplaceActionTransitions(transitions []*action.Transition) error {
	db, err := decode.DecodeBlock(m.block, m.dim)
	if err != nil {
		return err
	}

	cb, err := decode.CarveBlock(m.block, m.dim)
	if err != nil {
		return err
	}

	db.ActionTables.Transitions = transitions

	st := serialize.SerializeActionTransitions(transitions,
		db.CentralDir.ActionTables[action.IDTransition])
	overflow := len(st) - len(cb.ActionTables[action.IDTransition])

	if overflow > 0 {
		return wlerr.Errorf(
			"failed to replace action transitions: "+
				"replacement larger than original: overflow=%d",
			overflow)
	}

	off := db.Offsets.ActionTables[action.IDTransition]
	copy(m.block.EncSection[off:off+len(st)], st)

	return nil
}

// Block returns a BlockModifier's modified block.
func (m *BlockModifier) Block() msq.Block {
	return m.block
}
