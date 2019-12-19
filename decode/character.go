package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

const (
	CharSkillSize = 2
	CharItemSize  = 2
	CharNumSkills = 30
	CharNumItems  = 30
	CharacterSize = 256

	SexMale   = 0
	SexFemale = 1

	NatlUS      = 0
	NatlRussian = 1
	NatlMexican = 2
	NatlIndian  = 3
	NatlChinese = 4

	MaxCharNameLen = 13
	MaxCharRankLen = 24

	CharSkillIDCount = 35
)

// CharSkill represents a single skill learned by a character.
type CharSkill struct {
	ID    int
	Level int
}

// CharItem represents a single item in a character's inventory.
type CharItem struct {
	ID     int
	Ammo   int
	Jammed bool
}

// Character represents a PC or NPC.
type Character struct {
	Name            string      // 00-0d
	Strength        int         // 0e
	IQ              int         // 0f
	Luck            int         // 10
	Speed           int         // 11
	Agility         int         // 12
	Dexterity       int         // 13
	Charisma        int         // 14
	Money           int         // 15-17
	IsFemale        bool        // 18
	Nationality     int         // 19
	AC              int         // 1a
	Maxcon          int         // 1b-1c
	Con             int         // 1d-1e
	WeaponIdx       int         // 1f (base 1)
	SkillPoints     int         // 20
	Experience      int         // 21-23
	Level           int         // 24
	ArmorIdx        int         // 25 (base 1)
	PrevCon         int         // 26-27
	Afflictions     uint        // 28
	IsNPC           bool        // 29
	RefuseItem      int         // 2b
	RefuseSkill     int         // 2c
	RefuseAttribute int         // 2d
	RefuseTrade     int         // 2e
	JoinStringIdx   int         // 30
	Obedience       int         // 31
	Rank            string      // 32-4a
	GameIsWon       bool        // 4b
	RadioAfterWon   bool        // 4c
	Skills          []CharSkill // 80-bb
	Items           []CharItem  // bd-f8
}

// DecodeCharSkill decodes a skill from a sequence of bytes.
func DecodeCharSkill(b []byte) (*CharSkill, error) {
	if len(b) < CharSkillSize {
		return nil, wlerr.Errorf(
			"failed to parse skill: data too short: have=%d want>=%d",
			len(b), CharSkillSize)
	}

	return &CharSkill{
		ID:    int(b[0]),
		Level: int(b[1]),
	}, nil
}

// DecodeCharItem decodes an inventory item from a sequence of bytes.
func DecodeCharItem(b []byte) (*CharItem, error) {
	if len(b) < CharItemSize {
		return nil, wlerr.Errorf(
			"failed to parse item: data too short: have=%d want>=%d",
			len(b), CharItemSize)
	}

	return &CharItem{
		ID:     int(b[0]),
		Ammo:   int(b[1] & 0x7f),
		Jammed: b[1]&0x80 != 0,
	}, nil
}

// DecodeCharacter decodes a PC or NPC from a sequence of bytes.
func DecodeCharacter(b []byte) (*Character, error) {
	onErr := wlerr.MakeWrapper("failed to parse character")

	if len(b) < CharacterSize {
		return nil, onErr(nil,
			"data too short: have=%d want>=%d", len(b), CharacterSize)
	}

	ch := &Character{}

	var err error

	ch.Name, err = gen.ReadString(b[:0x0e])
	if err != nil {
		return nil, onErr(err, "failed to read name")
	}

	ch.Strength = int(b[0x0e])
	ch.IQ = int(b[0x0f])
	ch.Luck = int(b[0x10])
	ch.Speed = int(b[0x11])
	ch.Agility = int(b[0x12])
	ch.Dexterity = int(b[0x13])
	ch.Charisma = int(b[0x14])

	ch.Money, err = gen.ReadUint24(b[0x15:0x18])
	if err != nil {
		return nil, onErr(err, "failed to read money")
	}

	ch.IsFemale = b[0x18] == SexFemale
	ch.Nationality = int(b[0x19])
	ch.AC = int(b[0x1a])

	ch.Maxcon, err = gen.ReadUint16(b[0x1b:0x1d])
	if err != nil {
		return nil, onErr(err, "failed to read maxcon")
	}

	ch.Con, err = gen.ReadUint16(b[0x1d:0x1f])
	if err != nil {
		return nil, onErr(err, "failed to read con")
	}

	ch.WeaponIdx = int(b[0x1f])
	ch.SkillPoints = int(b[0x20])

	ch.Experience, err = gen.ReadUint24(b[0x21:0x24])
	if err != nil {
		return nil, onErr(err, "failed to read experience")
	}

	ch.Level = int(b[0x24])
	ch.ArmorIdx = int(b[0x25])

	ch.PrevCon, err = gen.ReadUint16(b[0x26:0x28])
	if err != nil {
		return nil, onErr(err, "failed to read experience")
	}

	ch.Afflictions = uint(b[0x28])
	ch.IsNPC = b[0x29] != 0
	ch.RefuseItem = int(b[0x2b])
	ch.RefuseSkill = int(b[0x2c])
	ch.RefuseAttribute = int(b[0x2d])
	ch.RefuseTrade = int(b[0x2e])
	ch.JoinStringIdx = int(b[0x30])
	ch.Obedience = int(b[0x31])

	ch.Rank, err = gen.ReadString(b[0x32:0x4b])
	if err != nil {
		return nil, onErr(err, "failed to read rank")
	}

	ch.GameIsWon = b[0x4b] != 0
	ch.RadioAfterWon = b[0x4c] != 0

	off := 0x80
	for i := 0; i < CharNumSkills; i++ {
		end := off + CharSkillSize
		skill, err := DecodeCharSkill(b[off:end])
		if err != nil {
			return nil, err
		}
		off = end

		ch.Skills = append(ch.Skills, *skill)
	}

	off = 0xbd
	for i := 0; i < CharNumItems; i++ {
		end := off + CharItemSize
		item, err := DecodeCharItem(b[off:end])
		if err != nil {
			return nil, err
		}
		off = end

		ch.Items = append(ch.Items, *item)
	}

	return ch, nil
}

// EncodeCharSkill encodes a skill to a byte sequence.
func EncodeCharSkill(skill CharSkill) []byte {
	return []byte{
		byte(skill.ID),
		byte(skill.Level),
	}
}

// EncodeCharSkill encodes an inventory item to a byte sequence.
func EncodeCharItem(item CharItem) []byte {
	b2 := byte(item.Ammo)
	if item.Jammed {
		b2 |= 0x80
	}

	return []byte{
		byte(item.ID),
		byte(b2),
	}
}

// EncodeCharacter encodes a PC or an NPC to a byte sequence.
func EncodeCharacter(ch Character) ([]byte, error) {
	onErr := wlerr.MakeWrapper("failed to encode character")

	b := make([]byte, CharacterSize)

	if len(ch.Name) > MaxCharNameLen {
		return nil, onErr(nil,
			"name too long: have=%d want<=%d", len(ch.Name), MaxCharNameLen)
	}
	copy(b[:len(ch.Name)], []byte(ch.Name))

	b[0x0e] = byte(ch.Strength)
	b[0x0f] = byte(ch.IQ)
	b[0x10] = byte(ch.Luck)
	b[0x11] = byte(ch.Speed)
	b[0x12] = byte(ch.Agility)
	b[0x13] = byte(ch.Dexterity)
	b[0x14] = byte(ch.Charisma)

	copy(b[0x15:0x18], gen.WriteUint24(ch.Money))

	b[0x18] = gen.BoolToByte(ch.IsFemale)
	b[0x19] = byte(ch.Nationality)
	b[0x1a] = byte(ch.AC)

	copy(b[0x1b:0x1d], gen.WriteUint16(uint16(ch.Maxcon)))
	copy(b[0x1d:0x1f], gen.WriteUint16(uint16(ch.Con)))

	b[0x1f] = byte(ch.WeaponIdx)
	b[0x20] = byte(ch.SkillPoints)

	copy(b[0x21:0x24], gen.WriteUint24(ch.Experience))

	b[0x24] = byte(ch.Level)
	b[0x25] = byte(ch.ArmorIdx)

	copy(b[0x26:0x28], gen.WriteUint16(uint16(ch.PrevCon)))

	b[0x28] = byte(ch.Afflictions)
	b[0x29] = byte(gen.BoolToByte(ch.IsNPC))
	b[0x2b] = byte(ch.RefuseItem)
	b[0x2c] = byte(ch.RefuseSkill)
	b[0x2d] = byte(ch.RefuseAttribute)
	b[0x2e] = byte(ch.RefuseTrade)
	b[0x30] = byte(ch.JoinStringIdx)
	b[0x31] = byte(ch.Obedience)

	if len(ch.Rank) > MaxCharRankLen {
		return nil, onErr(nil,
			"rank too long: have=%d want<=%d", len(ch.Rank), MaxCharRankLen)
	}
	copy(b[0x32:0x32+len(ch.Rank)], []byte(ch.Rank))

	b[0x4b] = gen.BoolToByte(ch.GameIsWon)
	b[0x4c] = gen.BoolToByte(ch.RadioAfterWon)

	off := 0x80
	for _, s := range ch.Skills {
		end := off + CharSkillSize
		copy(b[off:end], EncodeCharSkill(s))

		off = end
	}

	off = 0xbd
	for _, i := range ch.Items {
		end := off + CharItemSize
		copy(b[off:end], EncodeCharItem(i))

		off = end
	}

	return b, nil
}
