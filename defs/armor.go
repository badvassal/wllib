package defs

const (
	ArmorIDNone int = iota
	ArmorIDRobe
	ArmorIDLeatherJacket
	ArmorIDBulletProofShirt
	ArmorIDKevlarVest
	ArmorIDRadSuit
	ArmorIDKevlarSuit
	ArmorIDPseudoChitinArmor
	ArmorIDPowerArmor
)

type Armor struct {
	ItemID int
	AC     int
}

var Armors = []Armor{
	ArmorIDNone: Armor{
		ItemID: ItemIDNone,
		AC:     0,
	},
	ArmorIDRobe: Armor{
		ItemID: ItemIDRobe,
		AC:     1,
	},
	ArmorIDLeatherJacket: Armor{
		ItemID: ItemIDLeatherJacket,
		AC:     1,
	},
	ArmorIDBulletProofShirt: Armor{
		ItemID: ItemIDBulletProofShirt,
		AC:     2,
	},
	ArmorIDKevlarVest: Armor{
		ItemID: ItemIDKevlarVest,
		AC:     4,
	},
	ArmorIDRadSuit: Armor{
		ItemID: ItemIDRadSuit,
		AC:     5,
	},
	ArmorIDKevlarSuit: Armor{
		ItemID: ItemIDKevlarSuit,
		AC:     6,
	},
	ArmorIDPseudoChitinArmor: Armor{
		ItemID: ItemIDPseudoChitinArmor,
		AC:     10,
	},
	ArmorIDPowerArmor: Armor{
		ItemID: ItemIDPowerArmor,
		AC:     14,
	},
}
