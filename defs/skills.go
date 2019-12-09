package defs

const (
	SkillIDBrawling        = 0x01
	SkillIDClimb           = 0x02
	SkillIDClipPistol      = 0x03
	SkillIDKnifeFight      = 0x04
	SkillIDPugilism        = 0x05
	SkillIDRifle           = 0x06
	SkillIDSwim            = 0x07
	SkillIDKnifeThrow      = 0x08
	SkillIDPerception      = 0x09
	SkillIDAssaultRifle    = 0x0A
	SkillIDATWeapon        = 0x0B
	SkillIDSMG             = 0x0C
	SkillIDAcrobat         = 0x0D
	SkillIDGambling        = 0x0E
	SkillIDPicklock        = 0x0F
	SkillIDSilentMove      = 0x10
	SkillIDCombatShooting  = 0x11
	SkillIDConfidence      = 0x12
	SkillIDSleightOfHand   = 0x13
	SkillIDDemolitions     = 0x14
	SkillIDForgery         = 0x15
	SkillIDAlarmDisarm     = 0x16
	SkillIDBureaucracy     = 0x17
	SkillIDBombDisarm      = 0x18
	SkillIDMedic           = 0x19
	SkillIDSafecrack       = 0x1A
	SkillIDCryptology      = 0x1B
	SkillIDMetallurgy      = 0x1C
	SkillIDHelicopterPilot = 0x1D
	SkillIDElectronics     = 0x1E
	SkillIDToasterRepair   = 0x1F
	SkillIDDoctor          = 0x20
	SkillIDCloneTech       = 0x21
	SkillIDEnergyWeapon    = 0x22
	SkillIDCyborgTech      = 0x23
	SkillIDMaxPlusOne      = 0x24
)

var SkillNames = []string{
	0:                      "???",
	SkillIDBrawling:        "Brawling",
	SkillIDClimb:           "Climb",
	SkillIDClipPistol:      "ClipPistol",
	SkillIDKnifeFight:      "KnifeFight",
	SkillIDPugilism:        "Pugilism",
	SkillIDRifle:           "Rifle",
	SkillIDSwim:            "Swim",
	SkillIDKnifeThrow:      "KnifeThrow",
	SkillIDPerception:      "Perception",
	SkillIDAssaultRifle:    "AssaultRifle",
	SkillIDATWeapon:        "ATWeapon",
	SkillIDSMG:             "SMG",
	SkillIDAcrobat:         "Acrobat",
	SkillIDGambling:        "Gambling",
	SkillIDPicklock:        "Picklock",
	SkillIDSilentMove:      "SilentMove",
	SkillIDCombatShooting:  "CombatShooting",
	SkillIDConfidence:      "Confidence",
	SkillIDSleightOfHand:   "SleightOfHand",
	SkillIDDemolitions:     "Demolitions",
	SkillIDForgery:         "Forgery",
	SkillIDAlarmDisarm:     "AlarmDisarm",
	SkillIDBureaucracy:     "Bureaucracy",
	SkillIDBombDisarm:      "BombDisarm",
	SkillIDMedic:           "Medic",
	SkillIDSafecrack:       "Safecrack",
	SkillIDCryptology:      "Cryptology",
	SkillIDMetallurgy:      "Metallurgy",
	SkillIDHelicopterPilot: "HelicopterPilot",
	SkillIDElectronics:     "Electronics",
	SkillIDToasterRepair:   "ToasterRepair",
	SkillIDDoctor:          "Doctor",
	SkillIDCloneTech:       "CloneTech",
	SkillIDEnergyWeapon:    "EnergyWeapon",
	SkillIDCyborgTech:      "CyborgTech",
}

type Skill struct {
	IQ   int
	Cost int
}

var Skills = []Skill{
	SkillIDBrawling:        Skill{3, 1},
	SkillIDClimb:           Skill{3, 1},
	SkillIDClipPistol:      Skill{3, 1},
	SkillIDKnifeFight:      Skill{3, 1},
	SkillIDPugilism:        Skill{3, 1},
	SkillIDRifle:           Skill{3, 1},
	SkillIDSwim:            Skill{3, 1},
	SkillIDKnifeThrow:      Skill{6, 1},
	SkillIDPerception:      Skill{6, 1},
	SkillIDAssaultRifle:    Skill{9, 1},
	SkillIDATWeapon:        Skill{9, 1},
	SkillIDSMG:             Skill{9, 1},
	SkillIDAcrobat:         Skill{10, 1},
	SkillIDGambling:        Skill{10, 1},
	SkillIDPicklock:        Skill{10, 1},
	SkillIDSilentMove:      Skill{10, 1},
	SkillIDCombatShooting:  Skill{10, 1},
	SkillIDConfidence:      Skill{11, 1},
	SkillIDSleightOfHand:   Skill{12, 1},
	SkillIDDemolitions:     Skill{13, 1},
	SkillIDForgery:         Skill{13, 1},
	SkillIDAlarmDisarm:     Skill{14, 1},
	SkillIDBureaucracy:     Skill{14, 1},
	SkillIDBombDisarm:      Skill{15, 2},
	SkillIDMedic:           Skill{15, 2},
	SkillIDSafecrack:       Skill{15, 2},
	SkillIDCryptology:      Skill{16, 2},
	SkillIDMetallurgy:      Skill{17, 2},
	SkillIDHelicopterPilot: Skill{19, 3},
	SkillIDElectronics:     Skill{20, 3},
	SkillIDToasterRepair:   Skill{20, 3},
	SkillIDDoctor:          Skill{21, 3},
	SkillIDCloneTech:       Skill{22, 3},
	SkillIDEnergyWeapon:    Skill{23, 3},
	SkillIDCyborgTech:      Skill{24, 3},
}
