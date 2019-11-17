package decode

import (
	"fmt"

	"github.com/badvassal/wllib/decode/action"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/msq"
)

// Meta contains an integer for every section in an MSQ block.  It is used to
// convey section sizes and offsets.
type Meta struct {
	MapData        int
	CentralDir     int
	MapInfo        int
	ActionTables   []int
	SpecialActions int
	NPCTable       int
	MonsterNames   int
	MonsterData    int
	StringsArea    int
}

// Block is a fully decoded MSQ block.
type Block struct {
	Dim     gen.Point
	Offsets Meta
	Sizes   Meta

	MapData        MapData
	CentralDir     CentralDir
	MapInfo        MapInfo
	ActionTables   action.Tables
	NPCTable       NPCTable
	SpecialActions SpecialActions
	MonsterNames   MonsterNames
	MonsterData    MonsterData
	StringsArea    StringsArea
}

// DecodeState is fully decoded saved game.
type DecodeState struct {
	Blocks [][]Block // [game-idx][block-idx]
}

// ValidateMapDim ensures the provided dimensions are valid for a block map.
// It returns an error if the dimensions are invalid.
func ValidateMapDim(dim gen.Point) error {
	if dim.X%2 != 0 {
		return fmt.Errorf("invalid map dim.x: have=%d want=even", dim.X)
	}
	return nil
}

// DecodeBlock fully decodes an MSQ block.  dim is the dimensions of the block's map.
func DecodeBlock(block msq.Block, dim gen.Point) (*Block, error) {
	cb, err := CarveBlock(block, dim)
	if err != nil {
		return nil, err
	}

	md, err := DecodeMapData(cb.MapData, dim)
	if err != nil {
		return nil, err
	}

	cd, err := DecodeCentralDir(cb.CentralDir)
	if err != nil {
		return nil, err
	}

	mi, err := DecodeMapInfo(cb.MapInfo)
	if err != nil {
		return nil, err
	}

	parseActionTable := func(idx int) (gen.Table, error) {
		if len(cb.ActionTables[idx]) == 0 {
			return gen.Table{}, nil
		} else {
			t, err := gen.ParseTable(cb.ActionTables[idx], cd.ActionTables[idx])
			if err != nil {
				return gen.Table{}, wlerr.Wrapf(err, "failed to parse action table %d", idx)
			}

			return *t, nil
		}
	}

	t0, err := parseActionTable(0)
	if err != nil {
		return nil, err
	}
	t1, err := parseActionTable(1)
	if err != nil {
		return nil, err
	}
	t2, err := parseActionTable(2)
	if err != nil {
		return nil, err
	}
	t3, err := parseActionTable(3)
	if err != nil {
		return nil, err
	}
	t4, err := parseActionTable(4)
	if err != nil {
		return nil, err
	}
	t5, err := parseActionTable(5)
	if err != nil {
		return nil, err
	}
	t6, err := parseActionTable(6)
	if err != nil {
		return nil, err
	}
	t7, err := parseActionTable(7)
	if err != nil {
		return nil, err
	}
	t8, err := parseActionTable(8)
	if err != nil {
		return nil, err
	}
	t9, err := parseActionTable(9)
	if err != nil {
		return nil, err
	}
	t11, err := parseActionTable(11)
	if err != nil {
		return nil, err
	}
	t12, err := parseActionTable(12)
	if err != nil {
		return nil, err
	}
	t13, err := parseActionTable(13)
	if err != nil {
		return nil, err
	}
	t14, err := parseActionTable(14)
	if err != nil {
		return nil, err
	}
	t15, err := parseActionTable(15)
	if err != nil {
		return nil, err
	}

	ts, err := action.DecodeTransitions(
		cb.ActionTables[action.IDTransition],
		cd.ActionTables[action.IDTransition])
	if err != nil {
		return nil, err
	}

	var sac *SpecialActions
	if cd.SpecialActions == 0 {
		sac = &SpecialActions{}
	} else {
		sac, err = DecodeSpecialActions(cb.SpecialActions, cd.SpecialActions)
		if err != nil {
			return nil, err
		}
	}

	var nt *NPCTable
	if cd.NPCTable == 0 || cb.Sizes().NPCTable < NPCTableMinLen {
		nt = &NPCTable{}
	} else {
		nt, _, err = DecodeNPCTable(cb.NPCTable, cd.NPCTable)
		if err != nil {
			return nil, err
		}
	}

	mn, err := DecodeMonsterNames(cb.MonsterNames)
	if err != nil {
		return nil, err
	}

	mo, err := DecodeMonsterData(cb.MonsterData)
	if err != nil {
		return nil, err
	}

	sa, _, err := DecodeStringsArea(cb.StringsArea)
	if err != nil {
		return nil, err
	}

	return &Block{
		Dim:     dim,
		Offsets: cb.Offsets,
		Sizes:   *cb.Sizes(),

		MapData:    *md,
		CentralDir: *cd,
		MapInfo:    *mi,
		ActionTables: action.Tables{
			T0:          t0,
			T1:          t1,
			T2:          t2,
			T3:          t3,
			T4:          t4,
			T5:          t5,
			T6:          t6,
			T7:          t7,
			T8:          t8,
			T9:          t9,
			T11:         t11,
			T12:         t12,
			T13:         t13,
			T14:         t14,
			T15:         t15,
			Transitions: ts,
		},
		SpecialActions: *sac,
		NPCTable:       *nt,
		MonsterNames:   *mn,
		MonsterData:    *mo,
		StringsArea:    *sa,
	}, nil
}
