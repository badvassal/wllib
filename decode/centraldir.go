package decode

import (
	"strconv"

	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
	log "github.com/sirupsen/logrus"
)

// CentralDirLen is the size, in bytes, of an MSQ block's central directory.
const CentralDirLen = 42

// These constants assign priority to the pointers contained in the central
// directory.  If two pointers have the same value, the higher priority pointer
// "wins".  In the original GAME files, some pointers in a single CD have the
// same value.  In this case, the only section that is actually present is the
// one with the highest priority pointer.
const (
	CDPtrIdxSpecialActions = 0
	CDPtrIdxNPCTable       = 1
	CDPtrIdxActionTable0   = 2
	CDPtrIdxActionTable1   = 3
	CDPtrIdxActionTable2   = 4
	CDPtrIdxActionTable3   = 5
	CDPtrIdxActionTable4   = 6
	CDPtrIdxActionTable5   = 7
	CDPtrIdxActionTable6   = 8
	CDPtrIdxActionTable7   = 9
	CDPtrIdxActionTable8   = 10
	CDPtrIdxActionTable9   = 11
	CDPtrIdxActionTable10  = 12
	CDPtrIdxActionTable11  = 13
	CDPtrIdxActionTable12  = 14
	CDPtrIdxActionTable13  = 15
	CDPtrIdxActionTable14  = 16
	CDPtrIdxActionTable15  = 17
	CDPtrIdxMonsterNames   = 18
	CDPtrIdxMonsterData    = 19
	CDPtrIdxStrings        = 20
	CDPtrCount             = 21
)

// CentralDir represents an MSQ block's central directory.  It contains
// pointers to other areas within the same block.  Each pointer is a 16-bit
// offset relative to the start of the block's encrypted section.
// (see <https://wasteland.gamepedia.com/Central_MSQ_Directory>).
type CentralDir struct {
	Strings        int
	MonsterNames   int
	MonsterData    int
	ActionTables   []int
	SpecialActions int
	NPCTable       int
}

func ActionTablePtrPrio(idx int) int {
	switch idx {
	case 0:
		return CDPtrIdxActionTable0
	case 1:
		return CDPtrIdxActionTable1
	case 2:
		return CDPtrIdxActionTable2
	case 3:
		return CDPtrIdxActionTable3
	case 4:
		return CDPtrIdxActionTable4
	case 5:
		return CDPtrIdxActionTable5
	case 6:
		return CDPtrIdxActionTable6
	case 7:
		return CDPtrIdxActionTable7
	case 8:
		return CDPtrIdxActionTable8
	case 9:
		return CDPtrIdxActionTable9
	case 10:
		return CDPtrIdxActionTable10
	case 11:
		return CDPtrIdxActionTable11
	case 12:
		return CDPtrIdxActionTable12
	case 13:
		return CDPtrIdxActionTable13
	case 14:
		return CDPtrIdxActionTable14
	case 15:
		return CDPtrIdxActionTable15
	default:
		panic("invalid action table index: " + strconv.Itoa(idx))
		return 0
	}
}

// DecodeCentralDir parses a central directory from a sequence of bytes.
func DecodeCentralDir(data []byte) (*CentralDir, error) {
	if len(data) < CentralDirLen {
		return nil, wlerr.Errorf(
			"cannot decode central directory: not enough data: "+
				"have=%d want>=%d",
			len(data), CentralDirLen)
	}

	off := 0
	readPtr := func() int {
		// Ignore error; we already verified source length.
		p, _ := gen.ReadUint16(data[off : off+2])
		off += 2
		return p
	}

	cd := &CentralDir{}

	cd.Strings = readPtr()
	cd.MonsterNames = readPtr()
	cd.MonsterData = readPtr()

	cd.ActionTables = make([]int, 16)
	for i := 0; i < len(cd.ActionTables); i++ {
		cd.ActionTables[i] = readPtr()
	}

	cd.SpecialActions = readPtr()
	cd.NPCTable = readPtr()

	log.Debugf("parsed central directory: %+v", cd)

	return cd, nil
}

// FirstActionTable retrieves the offset of the first non-zero-length action
// table.
func (cd *CentralDir) FirstActionTable() int {
	su := gen.SortedUniqueInts(cd.ActionTables)
	for _, p := range su {
		if p != 0 {
			return p
		}
	}
	return 0
}

// LastActionTableIdx retrieves the index of the last non-zero-length action
// table.  The retrieved index can be used with the central directory's
// ActionTables member to get the table's offset.
func (cd *CentralDir) LastActionTableIdx() int {
	best := 0
	for i, p := range cd.ActionTables {
		if p >= cd.ActionTables[best] {
			best = i
		}
	}

	return best
}

// Pointers returns a slice of the central directory's offsets.  The slice is
// sorted by area priority (lowest priority first).
func (cd *CentralDir) Pointers() []int {
	ps := make([]int, CDPtrCount)

	ps[CDPtrIdxStrings] = cd.Strings
	ps[CDPtrIdxMonsterNames] = cd.MonsterNames
	ps[CDPtrIdxMonsterData] = cd.MonsterData
	ps[CDPtrIdxNPCTable] = cd.NPCTable
	ps[CDPtrIdxSpecialActions] = cd.SpecialActions

	for i := 0; i < len(cd.ActionTables); i++ {
		idx := ActionTablePtrPrio(i)
		ps[idx] = cd.ActionTables[i]
	}

	return ps
}

// EncodeCentralDir encodes a central directory to a byte sequence.
func EncodeCentralDir(cd CentralDir) []byte {
	b := make([]byte, CentralDirLen)

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
