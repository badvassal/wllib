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
func DecodeBlock(body msq.Body, dim gen.Point) (*Block, error) {
	cb, err := CarveBlock(body, dim)
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

	tables := make([]gen.Table, 16)
	for i := 0; i < len(tables); i++ {
		tables[i], err = parseActionTable(i)
		if err != nil {
			return nil, err
		}
	}

	loots, err := action.DecodeLootTable(tables[5])
	if err != nil {
		return nil, err
	}

	ts, err := action.DecodeTransitionTable(tables[10])
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
			T0:          tables[0],
			T1:          tables[1],
			T2:          tables[2],
			T3:          tables[3],
			T4:          tables[4],
			Loots:       loots,
			T6:          tables[6],
			T7:          tables[7],
			T8:          tables[8],
			T9:          tables[9],
			Transitions: ts,
			T11:         tables[11],
			T12:         tables[12],
			T13:         tables[13],
			T14:         tables[14],
			T15:         tables[15],
		},
		SpecialActions: *sac,
		NPCTable:       *nt,
		MonsterNames:   *mn,
		MonsterData:    *mo,
		StringsArea:    *sa,
	}, nil
}
