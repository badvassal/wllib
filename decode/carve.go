package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/msq"
)

// CarvedBlock is an MSQ block body that has been decrypted and partitioned
// into areas.  The areas are byte slices; they have not been decoded.
type CarvedBlock struct {
	Dim     gen.Point
	Offsets Meta

	MapData        []byte
	CentralDir     []byte
	MapInfo        []byte
	ActionTables   [][]byte
	NPCTable       []byte
	SpecialActions []byte
	MonsterNames   []byte
	MonsterData    []byte
	StringsArea    []byte
}

func NewCarvedBlock() *CarvedBlock {
	return &CarvedBlock{
		Offsets: Meta{
			ActionTables: make([]int, 16),
		},
		ActionTables: make([][]byte, 16),
	}
}

// cdEntryIs0Len indicates whether an area referenced by the central directory
// has a length of 0.  An entry is "zero length" if its pointer is 0 or if the
// central directory contains a higher priority pointer with the same value.
// cdPtrIdx is the index of the pointer of the area to check.
// cdPtrs is a list of central directory pointers, sorted by priority (lowest
// to highest).
func cdEntryIs0Len(cdPtrIdx int, cdPtrs []int) bool {
	p := cdPtrs[cdPtrIdx]

	if p == 0 {
		return true
	}

	for _, next := range cdPtrs[cdPtrIdx+1:] {
		if p == next {
			return true
		}
	}

	return false
}

// carveCDEntry extracts an area pointed to by the central directory.
// encSection is the (now decrypted) leading section of the MSQ block body that
// was initially encrypted.
// cdPtrIdx is the index of the pointer of the area to carve.
// cdPtrs is a list of central directory pointers, sorted by priority (lowest
// to highest).
func carveCDEntry(encSection []byte, cdPtrIdx int, cdPtrs []int) ([]byte, error) {
	if cdPtrIdx < 0 || cdPtrIdx >= len(cdPtrs) {
		return nil, wlerr.Errorf(
			"invalid entry index: have=%d want>=0&&<%d",
			cdPtrIdx, len(cdPtrs))
	}
	p := cdPtrs[cdPtrIdx]

	if p == 0 {
		return nil, nil
	}

	if cdEntryIs0Len(cdPtrIdx, cdPtrs) {
		return nil, nil
	}

	su := gen.SortedUniqueInts(cdPtrs)
	end := gen.NextInt(p, su, len(encSection))

	return gen.ExtractBlob(encSection, p, end)
}

// CarveBlock converts an MSQ block body into a CarvedBlock.
// b is the block body to convert.
// dim is the dimensions of the block's map.
func CarveBlock(b msq.Body, dim gen.Point) (*CarvedBlock, error) {
	cb := NewCarvedBlock()

	off := 0

	var err error

	if err = ValidateMapDim(dim); err != nil {
		return nil, err
	}

	mapDataLen := MapDataLen(dim)
	cb.MapData, err = gen.ExtractBlob(b.EncSection, off, off+mapDataLen)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to carve map data")
	}
	cb.Offsets.MapData = off

	off = cb.Offsets.MapData + mapDataLen
	blob, err := gen.ExtractBlob(b.EncSection, off, -1)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to carve central directory")
	}
	cd, err := DecodeCentralDir(blob)
	if err != nil {
		return nil, err
	}
	cb.CentralDir = b.EncSection[off : off+CentralDirLen]
	cb.Offsets.CentralDir = off

	off = cb.Offsets.CentralDir + CentralDirLen
	cb.MapInfo, err = gen.ExtractBlob(b.EncSection, off, off+MapInfoLen)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to carve map info")
	}
	cb.Offsets.MapInfo = off

	cdPtrs := cd.Pointers()

	for i, at := range cd.ActionTables {
		cdIdx := ActionTablePtrPrio(i)
		cb.ActionTables[i], err = carveCDEntry(b.EncSection, cdIdx, cdPtrs)
		if err != nil {
			return nil, wlerr.Wrapf(err,
				"failed to carve action table %d", i)
		}
		cb.Offsets.ActionTables[i] = at
	}

	cb.SpecialActions, err = carveCDEntry(b.EncSection, CDPtrIdxSpecialActions,
		cdPtrs)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to carve special actions")
	}
	cb.Offsets.SpecialActions = cd.SpecialActions

	if !cdEntryIs0Len(CDPtrIdxNPCTable, cdPtrs) {
		_, ntsize, err := DecodeNPCTable(b.EncSection[cd.NPCTable:], cd.NPCTable)
		if err != nil {
			return nil, wlerr.Wrapf(err, "failed to carve NPC table")
		}
		cb.NPCTable = b.EncSection[cd.NPCTable : cd.NPCTable+ntsize]
		cb.Offsets.NPCTable = cd.NPCTable
	}

	cb.MonsterNames, err = carveCDEntry(b.EncSection, CDPtrIdxMonsterNames,
		cdPtrs)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to carve monster names")
	}
	cb.Offsets.MonsterNames = cd.MonsterNames

	cb.MonsterData, err = carveCDEntry(b.EncSection, CDPtrIdxMonsterData,
		cdPtrs)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to carve monster data")
	}
	cb.Offsets.MonsterData = cd.MonsterData

	_, sasize, err := DecodeStringsArea(b.PlainSection)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to carve strings area")
	}
	cb.StringsArea = b.PlainSection[:sasize]
	cb.Offsets.StringsArea = cd.Strings

	return cb, nil
}

// Sizes indicates the size of each area in a carved block.
func (cb *CarvedBlock) Sizes() *Meta {
	m := &Meta{
		MapData:        len(cb.MapData),
		CentralDir:     len(cb.CentralDir),
		MapInfo:        len(cb.MapInfo),
		SpecialActions: len(cb.SpecialActions),
		NPCTable:       len(cb.NPCTable),
		MonsterNames:   len(cb.MonsterNames),
		MonsterData:    len(cb.MonsterData),
		StringsArea:    len(cb.StringsArea),
	}

	for _, at := range cb.ActionTables {
		m.ActionTables = append(m.ActionTables, len(at))
	}

	return m
}
