package defs

const (
	ItemIDNone                    = 0x00
	ItemIDAx                      = 0x01
	ItemIDClub                    = 0x02
	ItemIDChainsaw                = 0x03
	ItemIDKnife                   = 0x04
	ItemIDProtonAx                = 0x05
	ItemIDGrenade                 = 0x06
	ItemIDPlasticExplosive        = 0x07
	ItemIDTNT                     = 0x08
	ItemIDMangler                 = 0x09
	ItemIDSabotRocket             = 0x0a
	ItemIDLAWRocket               = 0x0b
	ItemIDRPG7                    = 0x0c
	ItemIDM1911A145Pistol         = 0x0d
	ItemIDSpear                   = 0x0e
	ItemIDThrowingKnife           = 0x0f
	ItemIDVP91Z9mmPistol          = 0x10
	ItemIDFlamethrower            = 0x11
	ItemIDM17Carbine              = 0x12
	ItemIDM19Rifle                = 0x13
	ItemIDRedRyderRifle           = 0x14
	ItemIDMAC17SMG                = 0x15
	ItemIDUziMark27SMG            = 0x16
	ItemIDAK97AssaultRifle        = 0x17
	ItemIDM1989A1NATOAssaultRifle = 0x18
	ItemIDLaserPistol             = 0x19
	ItemIDIonBeamer               = 0x1a
	ItemIDLaserCarbine            = 0x1b
	ItemIDLaserRifle              = 0x1c
	ItemIDMesonCannon             = 0x1d
	ItemIDClip45                  = 0x1e
	ItemIDClip762mm               = 0x1f
	ItemIDClip9mm                 = 0x20
	ItemIDHowitzerShell           = 0x21
	ItemIDPowerPack               = 0x22
	ItemIDPowerArmor              = 0x23
	ItemIDBulletProofShirt        = 0x24
	ItemIDKevlarVest              = 0x25
	ItemIDLeatherJacket           = 0x26
	ItemIDKevlarSuit              = 0x27
	ItemIDPseudoChitinArmor       = 0x28
	ItemIDRadSuit                 = 0x29
	ItemIDRobe                    = 0x2a
	ItemIDBook                    = 0x2b
	ItemIDCanteen                 = 0x2c
	ItemIDCrowbar                 = 0x2d
	ItemIDEngine                  = 0x2e
	ItemIDGasMask                 = 0x2f
	ItemIDGeigerCounter           = 0x30
	ItemIDHandMirror              = 0x31
	ItemIDJug                     = 0x32
	ItemIDMap                     = 0x33
	ItemIDMatch                   = 0x34
	ItemIDPickAx                  = 0x35
	ItemIDRope                    = 0x36
	ItemIDShovel                  = 0x37
	ItemIDSledgeHammer            = 0x38
	ItemIDSnakeSqueezin           = 0x39
	ItemIDAndroidHead             = 0x3a
	ItemIDAntitoxin               = 0x3b
	ItemIDFinsterHead             = 0x3c
	ItemIDBlackstarKey            = 0x3d
	ItemIDBloodstaffFake          = 0x3e
	ItemIDBloodstaff              = 0x3f
	ItemIDBrokenToaster           = 0x40
	ItemIDChemical                = 0x41
	ItemIDCloneFluid              = 0x42
	ItemIDVisaCard                = 0x43
	ItemIDFusionCell              = 0x44
	ItemIDGrazerBatFetish         = 0x45
	ItemIDNovaKey                 = 0x49
	ItemIDOnyxRing                = 0x4a
	ItemIDPasskey                 = 0x4b
	ItemIDPlasmaCoupler           = 0x4c
	ItemIDPowerConverter          = 0x4d
	ItemIDPulsarKey               = 0x4e
	ItemIDQuasarKey               = 0x4f
	ItemIDRomBoard                = 0x50
	ItemIDRoomKey18               = 0x51
	ItemIDRubyRing                = 0x52
	ItemIDSecpass1                = 0x53
	ItemIDSecpass3                = 0x54
	ItemIDSecpass7                = 0x55
	ItemIDSecpassA                = 0x56
	ItemIDSecpassB                = 0x57
	ItemIDServoMotor              = 0x58
	ItemIDSonicKey                = 0x59
	ItemIDToaster                 = 0x5a
	ItemIDClayPot                 = 0x5b
	ItemIDFruit                   = 0x5c
	ItemIDJewelry                 = 0x5d
)

var ItemNames = []string{
	ItemIDNone:                    "None",
	ItemIDAx:                      "Ax",
	ItemIDClub:                    "Club",
	ItemIDChainsaw:                "Chainsaw",
	ItemIDKnife:                   "Knife",
	ItemIDProtonAx:                "ProtonAx",
	ItemIDGrenade:                 "Grenade",
	ItemIDPlasticExplosive:        "PlasticExplosive",
	ItemIDTNT:                     "TNT",
	ItemIDMangler:                 "Mangler",
	ItemIDSabotRocket:             "SabotRocket",
	ItemIDLAWRocket:               "LAWRocket",
	ItemIDRPG7:                    "RPG7",
	ItemIDM1911A145Pistol:         "M1911A145Pistol",
	ItemIDSpear:                   "Spear",
	ItemIDThrowingKnife:           "ThrowingKnife",
	ItemIDVP91Z9mmPistol:          "VP91Z9mmPistol",
	ItemIDFlamethrower:            "Flamethrower",
	ItemIDM17Carbine:              "M17Carbine",
	ItemIDM19Rifle:                "M19Rifle",
	ItemIDRedRyderRifle:           "RedRyderRifle",
	ItemIDMAC17SMG:                "MAC17SMG",
	ItemIDUziMark27SMG:            "UziMark27SMG",
	ItemIDAK97AssaultRifle:        "AK97AssaultRifle",
	ItemIDM1989A1NATOAssaultRifle: "M1989A1NATOAssaultRifle",
	ItemIDLaserPistol:             "LaserPistol",
	ItemIDIonBeamer:               "IonBeamer",
	ItemIDLaserCarbine:            "LaserCarbine",
	ItemIDLaserRifle:              "LaserRifle",
	ItemIDMesonCannon:             "MesonCannon",
	ItemIDClip45:                  "Clip45",
	ItemIDClip762mm:               "Clip762mm",
	ItemIDClip9mm:                 "Clip9mm",
	ItemIDHowitzerShell:           "HowitzerShell",
	ItemIDPowerPack:               "PowerPack",
	ItemIDPowerArmor:              "PowerArmor",
	ItemIDBulletProofShirt:        "BulletProofShirt",
	ItemIDKevlarVest:              "KevlarVest",
	ItemIDLeatherJacket:           "LeatherJacket",
	ItemIDKevlarSuit:              "KevlarSuit",
	ItemIDPseudoChitinArmor:       "PseudoChitinArmor",
	ItemIDRadSuit:                 "RadSuit",
	ItemIDRobe:                    "Robe",
	ItemIDBook:                    "Book",
	ItemIDCanteen:                 "Canteen",
	ItemIDCrowbar:                 "Crowbar",
	ItemIDEngine:                  "Engine",
	ItemIDGasMask:                 "GasMask",
	ItemIDGeigerCounter:           "GeigerCounter",
	ItemIDHandMirror:              "HandMirror",
	ItemIDJug:                     "Jug",
	ItemIDMap:                     "Map",
	ItemIDMatch:                   "Match",
	ItemIDPickAx:                  "PickAx",
	ItemIDRope:                    "Rope",
	ItemIDShovel:                  "Shovel",
	ItemIDSledgeHammer:            "SledgeHammer",
	ItemIDSnakeSqueezin:           "SnakeSqueezin",
	ItemIDAndroidHead:             "AndroidHead",
	ItemIDAntitoxin:               "Antitoxin",
	ItemIDFinsterHead:             "FinsterHead",
	ItemIDBlackstarKey:            "BlackstarKey",
	ItemIDBloodstaffFake:          "BloodstaffFake",
	ItemIDBloodstaff:              "Bloodstaff",
	ItemIDBrokenToaster:           "BrokenToaster",
	ItemIDChemical:                "Chemical",
	ItemIDCloneFluid:              "CloneFluid",
	ItemIDVisaCard:                "VisaCard",
	ItemIDFusionCell:              "FusionCell",
	ItemIDGrazerBatFetish:         "GrazerBatFetish",
	ItemIDNovaKey:                 "NovaKey",
	ItemIDOnyxRing:                "OnyxRing",
	ItemIDPasskey:                 "Passkey",
	ItemIDPlasmaCoupler:           "PlasmaCoupler",
	ItemIDPowerConverter:          "PowerConverter",
	ItemIDPulsarKey:               "PulsarKey",
	ItemIDQuasarKey:               "QuasarKey",
	ItemIDRomBoard:                "RomBoard",
	ItemIDRoomKey18:               "RoomKey18",
	ItemIDRubyRing:                "RubyRing",
	ItemIDSecpass1:                "Secpass1",
	ItemIDSecpass3:                "Secpass3",
	ItemIDSecpass7:                "Secpass7",
	ItemIDSecpassA:                "SecpassA",
	ItemIDSecpassB:                "SecpassB",
	ItemIDServoMotor:              "ServoMotor",
	ItemIDSonicKey:                "SonicKey",
	ItemIDToaster:                 "Toaster",
	ItemIDClayPot:                 "ClayPot",
	ItemIDFruit:                   "Fruit",
	ItemIDJewelry:                 "Jewelry",
}
