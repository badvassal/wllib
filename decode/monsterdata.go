package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

// MonsterDataElemLen is the size, in bytes, of a single element of monster
// data (one "monster").
const MonsterDataElemLen = 8

const (
	MonsterTypeAnimal   = 1
	MonsterTypeMutant   = 2
	MonsterTypeHumanoid = 3
	MonsterTypeCyborg   = 4
	MonsterTypeRobot    = 5
)

// MonsterDataElem describes a single enemy.
type MonsterDataElem struct {
	HitPoints   int // B0,1
	HitChance   int // B2
	ExtraDamage int // B3
	Armor       int // B4b0,3
	MaxCount    int // B4b4,7
	AttackType  int // B5b0,3
	FixedDamage int // B5b4,7
	Type        int // B6
	CombatImage int // B7
}

// MonsterData describes all the enemies in an MSQ block.  See
// <https://wasteland.gamepedia.com/Map_Data_Monster_Body>.
type MonsterData struct {
	Monsters []MonsterDataElem
}

// DecodeMonsterData parses monster data from a sequence of bytes.
func DecodeMonsterData(data []byte) (*MonsterData, error) {
	off := 0

	var monsters []MonsterDataElem
	for off < len(data) {
		rem := len(data) - off
		if rem < MonsterDataElemLen {
			return nil, wlerr.Errorf(
				"failed to decode monster data: bad byte count: "+
					"have=%d want%%%d",
				len(data), MonsterDataElemLen)
		}

		m := MonsterDataElem{}
		m.HitPoints, _ = gen.ReadUint16(data[off : off+2])
		m.HitChance = int(data[off+2])
		m.ExtraDamage = int(data[off+3])
		m.Armor = int(data[off+4] & 0x0f)
		m.MaxCount = int(data[off+4] >> 4)
		m.AttackType = int(data[off+5] & 0x0f)
		m.FixedDamage = int(data[off+5] >> 4)
		m.Type = int(data[off+6])
		m.CombatImage = int(data[off+7])
		monsters = append(monsters, m)

		off += MonsterDataElemLen
	}

	return &MonsterData{
		Monsters: monsters,
	}, nil
}
