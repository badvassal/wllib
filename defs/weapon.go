package defs

const (
	WeaponIDNone int = iota
	WeaponIDAx
	WeaponIDClub
	WeaponIDChainsaw
	WeaponIDCrowbar
	WeaponIDProtonAx
	WeaponIDKnife
	WeaponIDThrowingKnife
	WeaponIDM1911A145Pistol
	WeaponIDVP91Z9mmPistol
	WeaponIDMAC17SMG
	WeaponIDUziMark27SMG
	WeaponIDM17Carbine
	WeaponIDM19Rifle
	WeaponIDRedRyderRifle
	WeaponIDAK97AssaultRifle
	WeaponIDM1989A1NATOAssaultRifle
	WeaponIDFlamethrower
	WeaponIDLaserPistol
	WeaponIDIonBeamer
	WeaponIDLaserCarbine
	WeaponIDLaserRifle
	WeaponIDMesonCannon
	WeaponIDGrenade
	WeaponIDTNT
	WeaponIDPlasticExplosive
	WeaponIDSpear
	WeaponIDMangler
	WeaponIDSabotRocket
	WeaponIDLAWRocket
	WeaponIDRPG7
	WeaponIDMaxPlusOne
)

type Weapon struct {
	ItemID       int
	SkillID      int
	Cost         int
	Reusable     bool
	ClipItemID   int
	AmmoCapacity int
}

var Weapons = []Weapon{
	WeaponIDNone: Weapon{
		ItemID:       ItemIDNone,
		SkillID:      SkillIDPugilism,
		Cost:         0,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDAx: Weapon{
		ItemID:       ItemIDAx,
		SkillID:      SkillIDBrawling,
		Cost:         25,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDClub: Weapon{
		ItemID:       ItemIDClub,
		SkillID:      SkillIDBrawling,
		Cost:         15,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDChainsaw: Weapon{
		ItemID:       ItemIDChainsaw,
		SkillID:      SkillIDBrawling,
		Cost:         500,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDCrowbar: Weapon{
		ItemID:       ItemIDCrowbar,
		SkillID:      0,
		Cost:         15,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDProtonAx: Weapon{
		ItemID:       ItemIDProtonAx,
		SkillID:      SkillIDBrawling,
		Cost:         10000,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDKnife: Weapon{
		ItemID:       ItemIDKnife,
		SkillID:      SkillIDKnifeFight,
		Cost:         20,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDThrowingKnife: Weapon{
		ItemID:       ItemIDThrowingKnife,
		SkillID:      SkillIDKnifeThrow,
		Cost:         20,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDM1911A145Pistol: Weapon{
		ItemID:       ItemIDM1911A145Pistol,
		SkillID:      SkillIDClipPistol,
		Cost:         150,
		Reusable:     true,
		ClipItemID:   ItemIDClip45,
		AmmoCapacity: 7,
	},
	WeaponIDVP91Z9mmPistol: Weapon{
		ItemID:       ItemIDVP91Z9mmPistol,
		SkillID:      SkillIDClipPistol,
		Cost:         150,
		Reusable:     true,
		ClipItemID:   ItemIDClip9mm,
		AmmoCapacity: 18,
	},
	WeaponIDMAC17SMG: Weapon{
		ItemID:       ItemIDMAC17SMG,
		SkillID:      SkillIDSMG,
		Cost:         500,
		Reusable:     true,
		ClipItemID:   ItemIDClip45,
		AmmoCapacity: 30,
	},
	WeaponIDUziMark27SMG: Weapon{
		ItemID:       ItemIDUziMark27SMG,
		SkillID:      SkillIDSMG,
		Cost:         750,
		Reusable:     true,
		ClipItemID:   ItemIDClip9mm,
		AmmoCapacity: 40,
	},
	WeaponIDM17Carbine: Weapon{
		ItemID:       ItemIDM17Carbine,
		SkillID:      SkillIDRifle,
		Cost:         250,
		Reusable:     true,
		ClipItemID:   ItemIDClip762mm,
		AmmoCapacity: 10,
	},
	WeaponIDM19Rifle: Weapon{
		ItemID:       ItemIDM19Rifle,
		SkillID:      SkillIDRifle,
		Cost:         350,
		Reusable:     true,
		ClipItemID:   ItemIDClip762mm,
		AmmoCapacity: 8,
	},
	WeaponIDRedRyderRifle: Weapon{
		ItemID:       ItemIDRedRyderRifle,
		SkillID:      SkillIDRifle,
		Cost:         12345,
		Reusable:     true,
		ClipItemID:   ItemIDClip762mm,
		AmmoCapacity: 62,
	},
	WeaponIDAK97AssaultRifle: Weapon{
		ItemID:       ItemIDAK97AssaultRifle,
		SkillID:      SkillIDAssaultRifle,
		Cost:         1300,
		Reusable:     true,
		ClipItemID:   ItemIDClip762mm,
		AmmoCapacity: 30,
	},
	WeaponIDM1989A1NATOAssaultRifle: Weapon{
		ItemID:       ItemIDM1989A1NATOAssaultRifle,
		SkillID:      SkillIDAssaultRifle,
		Cost:         1500,
		Reusable:     true,
		ClipItemID:   ItemIDClip762mm,
		AmmoCapacity: 35,
	},
	WeaponIDFlamethrower: Weapon{
		ItemID:       ItemIDFlamethrower,
		SkillID:      0,
		Cost:         3000,
		Reusable:     true,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 60,
	},
	WeaponIDLaserPistol: Weapon{
		ItemID:       ItemIDLaserPistol,
		SkillID:      SkillIDEnergyWeapon,
		Cost:         8000,
		Reusable:     true,
		ClipItemID:   ItemIDPowerPack,
		AmmoCapacity: 40,
	},
	WeaponIDIonBeamer: Weapon{
		ItemID:       ItemIDIonBeamer,
		SkillID:      SkillIDEnergyWeapon,
		Cost:         17000,
		Reusable:     true,
		ClipItemID:   ItemIDPowerPack,
		AmmoCapacity: 20,
	},
	WeaponIDLaserCarbine: Weapon{
		ItemID:       ItemIDLaserCarbine,
		SkillID:      SkillIDEnergyWeapon,
		Cost:         11500,
		Reusable:     true,
		ClipItemID:   ItemIDPowerPack,
		AmmoCapacity: 30,
	},
	WeaponIDLaserRifle: Weapon{
		ItemID:       ItemIDLaserRifle,
		SkillID:      SkillIDEnergyWeapon,
		Cost:         13000,
		Reusable:     true,
		ClipItemID:   ItemIDPowerPack,
		AmmoCapacity: 20,
	},
	WeaponIDMesonCannon: Weapon{
		ItemID:       ItemIDMesonCannon,
		SkillID:      SkillIDEnergyWeapon,
		Cost:         18000,
		Reusable:     true,
		ClipItemID:   ItemIDPowerPack,
		AmmoCapacity: 10,
	},
	WeaponIDGrenade: Weapon{
		ItemID:       ItemIDGrenade,
		SkillID:      SkillIDDemolitions,
		Cost:         150,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDTNT: Weapon{
		ItemID:       ItemIDTNT,
		SkillID:      SkillIDDemolitions,
		Cost:         50,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDPlasticExplosive: Weapon{
		ItemID:       ItemIDPlasticExplosive,
		SkillID:      SkillIDDemolitions,
		Cost:         300,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDSpear: Weapon{
		ItemID:       ItemIDSpear,
		SkillID:      SkillIDBrawling,
		Cost:         35,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDMangler: Weapon{
		ItemID:       ItemIDMangler,
		SkillID:      SkillIDATWeapon,
		Cost:         500,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDSabotRocket: Weapon{
		ItemID:       ItemIDSabotRocket,
		SkillID:      SkillIDATWeapon,
		Cost:         1100,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDLAWRocket: Weapon{
		ItemID:       ItemIDLAWRocket,
		SkillID:      SkillIDATWeapon,
		Cost:         2500,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
	WeaponIDRPG7: Weapon{
		ItemID:       ItemIDRPG7,
		SkillID:      SkillIDATWeapon,
		Cost:         5000,
		Reusable:     false,
		ClipItemID:   ItemIDNone,
		AmmoCapacity: 0,
	},
}
