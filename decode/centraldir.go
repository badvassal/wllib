package decode

import (
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
	CDPtrIdxActionTables   = 2
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
		p, err := gen.ReadUint16(data[off : off+2])
		if err != nil {
			panic(err.Error())
		}
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
		idx := CDPtrIdxActionTables + i
		ps[idx] = cd.ActionTables[i]
	}

	return ps
}
