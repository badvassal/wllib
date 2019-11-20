package defs

import "github.com/badvassal/wllib/gen"

// BlockZIP identifies a block's address within a set of GAMEx files.
type BlockZIP struct {
	GameIdx  int
	BlockIdx int
}

// LocPair represents a transition between two locations.
type LocPair struct {
	From int
	To   int
}

const (
	Block0WorldMap                = 0
	Block0Quartz                  = 1
	Block0ScottsBar               = 2
	Block0StageCoachInn           = 3
	Block0UglysHideout            = 4
	Block0QuartzDerelictBuildings = 5
	Block0Courthouse              = 6
	Block0DesertNomads            = 7
	Block0AgCenter                = 8
	Block0Highpool                = 9
	Block0Needles                 = 10
	Block0BloodTempleTop          = 11
	Block0BloodTempleBottom       = 12
	Block0VerminCave              = 13
	Block0WastePit                = 14
	Block0NeedlesDowntownEast     = 15
	Block0NeedlesDowntownWest     = 16
	Block0PoliceStation           = 17
	Block0MineShaft               = 18
	Block0SavageVillage           = 19
	Block0NumBlocks               = 20

	Block1SleeperBaseLevel1         = 0
	Block1LasVegasDerelictBuildings = 1
	Block1LasVegas                  = 2
	Block1SleeperBaseLevel2         = 3
	Block1SleeperBaseLevel3         = 4
	Block1BaseCochiseOutside        = 5
	Block1BaseCochiseLevel1         = 6
	Block1BaseCochiseLevel3         = 7
	Block1BaseCochiseLevel2         = 8
	Block1BaseCochiseLevel4         = 9
	Block1Darwin                    = 10
	Block1DarwinBase                = 11
	Block1FinstersBrain             = 12
	Block1LasVegasSewersWest        = 13
	Block1LasVegasSewersEast        = 14
	Block1GuardianCitadelEntrance   = 15
	Block1GuardianCitadelOuter      = 16
	Block1TempleMushroom            = 17
	Block1FaranBrygos               = 18
	Block1FatFreddys                = 19
	Block1SpadesCasino              = 20
	Block1GuardianCitadelInner      = 21
	Block1NumBlocks                 = 22

	LocationWorldMap                  = 0
	LocationQuartz                    = 1
	LocationScottsBar                 = 2
	LocationStageCoachInn             = 3
	LocationUglysHideout              = 4
	LocationQuartzDerelictBuildings   = 5
	LocationCourthouse                = 6
	LocationSleeperBaseLevel1         = 7
	LocationDesertNomads              = 8
	LocationAgCenter                  = 9
	LocationHighpool                  = 10
	LocationLasVegasDerelictBuildings = 11
	LocationLasVegas                  = 12
	LocationSleeperBaseLevel2         = 13
	LocationNotDefined                = 14
	LocationSleeperBaseLevel3         = 15
	LocationBaseCochiseOutside        = 16
	LocationBaseCochiseLevel1         = 17
	LocationBaseCochiseLevel3         = 18
	LocationBaseCochiseLevel2         = 19
	LocationBaseCochiseLevel4         = 20
	LocationDarwin                    = 21
	LocationDarwinBase                = 22
	LocationFinstersBrain             = 23
	LocationLasVegasSewersWest        = 24
	LocationLasVegasSewersEast        = 25
	LocationNeedles                   = 26
	LocationBloodTempleTop            = 27
	LocationBloodTempleBottom         = 28
	LocationVerminCave                = 29
	LocationWastePit                  = 31
	LocationNeedlesDowntownEast       = 32
	LocationNeedlesDowntownWest       = 33
	LocationPoliceStation             = 34
	LocationGuardianCitadelEntrance   = 35
	LocationGuardianCitadelOuter      = 36
	LocationTempleMushroom            = 38
	LocationFaranBrygos               = 39
	LocationFatFreddys                = 40
	LocationSpadesCasino              = 41
	LocationGuardianCitadelInner      = 42
	LocationMineShaft                 = 43
	LocationSavageVillage             = 49
	LocationPrevious                  = 255
)

var MapDims = [][]gen.Point{
	0: []gen.Point{
		Block0WorldMap:                gen.Point{64, 64},
		Block0Quartz:                  gen.Point{32, 32},
		Block0ScottsBar:               gen.Point{32, 32},
		Block0StageCoachInn:           gen.Point{32, 32},
		Block0UglysHideout:            gen.Point{32, 32},
		Block0QuartzDerelictBuildings: gen.Point{32, 32},
		Block0Courthouse:              gen.Point{32, 32},
		Block0DesertNomads:            gen.Point{32, 32},
		Block0AgCenter:                gen.Point{32, 32},
		Block0Highpool:                gen.Point{32, 32},
		Block0Needles:                 gen.Point{64, 64},
		Block0BloodTempleTop:          gen.Point{32, 32},
		Block0BloodTempleBottom:       gen.Point{32, 32},
		Block0VerminCave:              gen.Point{32, 32},
		Block0WastePit:                gen.Point{32, 32},
		Block0NeedlesDowntownEast:     gen.Point{32, 32},
		Block0NeedlesDowntownWest:     gen.Point{32, 32},
		Block0PoliceStation:           gen.Point{32, 32},
		Block0MineShaft:               gen.Point{32, 32},
		Block0SavageVillage:           gen.Point{32, 32},
	},

	1: []gen.Point{
		Block1SleeperBaseLevel1:         gen.Point{32, 32},
		Block1LasVegasDerelictBuildings: gen.Point{32, 32},
		Block1LasVegas:                  gen.Point{64, 64},
		Block1SleeperBaseLevel2:         gen.Point{32, 32},
		Block1SleeperBaseLevel3:         gen.Point{32, 32},
		Block1BaseCochiseOutside:        gen.Point{32, 32},
		Block1BaseCochiseLevel1:         gen.Point{32, 32},
		Block1BaseCochiseLevel3:         gen.Point{32, 32},
		Block1BaseCochiseLevel2:         gen.Point{32, 32},
		Block1BaseCochiseLevel4:         gen.Point{32, 32},
		Block1Darwin:                    gen.Point{32, 32},
		Block1DarwinBase:                gen.Point{64, 64},
		Block1FinstersBrain:             gen.Point{32, 32},
		Block1LasVegasSewersWest:        gen.Point{32, 32},
		Block1LasVegasSewersEast:        gen.Point{32, 32},
		Block1GuardianCitadelEntrance:   gen.Point{32, 32},
		Block1GuardianCitadelOuter:      gen.Point{32, 32},
		Block1TempleMushroom:            gen.Point{32, 32},
		Block1FaranBrygos:               gen.Point{32, 32},
		Block1FatFreddys:                gen.Point{32, 32},
		Block1SpadesCasino:              gen.Point{32, 32},
		Block1GuardianCitadelInner:      gen.Point{32, 32},
	},
}

var LocationNameMap = map[int]string{
	LocationWorldMap:                  "WorldMap",
	LocationQuartz:                    "Quartz",
	LocationScottsBar:                 "ScottsBar",
	LocationStageCoachInn:             "StageCoachInn",
	LocationUglysHideout:              "UglysHideout",
	LocationQuartzDerelictBuildings:   "QuartzDerelictBuildings",
	LocationCourthouse:                "Courthouse",
	LocationSleeperBaseLevel1:         "SleeperBaseLevel1",
	LocationDesertNomads:              "DesertNomads",
	LocationAgCenter:                  "AgCenter",
	LocationHighpool:                  "Highpool",
	LocationLasVegasDerelictBuildings: "LasVegasDerelictBuildings",
	LocationLasVegas:                  "LasVegas",
	LocationSleeperBaseLevel2:         "SleeperBaseLevel2",
	LocationNotDefined:                "NotDefined",
	LocationSleeperBaseLevel3:         "SleeperBaseLevel3",
	LocationBaseCochiseOutside:        "BaseCochiseOutside",
	LocationBaseCochiseLevel1:         "BaseCochiseLevel1",
	LocationBaseCochiseLevel3:         "BaseCochiseLevel3",
	LocationBaseCochiseLevel2:         "BaseCochiseLevel2",
	LocationBaseCochiseLevel4:         "BaseCochiseLevel4",
	LocationDarwin:                    "Darwin",
	LocationDarwinBase:                "DarwinBase",
	LocationFinstersBrain:             "FinstersBrain",
	LocationLasVegasSewersWest:        "LasVegasSewersWest",
	LocationLasVegasSewersEast:        "LasVegasSewersEast",
	LocationNeedles:                   "Needles",
	LocationBloodTempleTop:            "BloodTempleTop",
	LocationBloodTempleBottom:         "BloodTempleBottom",
	LocationVerminCave:                "VerminCave",
	LocationWastePit:                  "WastePit",
	LocationNeedlesDowntownEast:       "NeedlesDowntownEast",
	LocationNeedlesDowntownWest:       "NeedlesDowntownWest",
	LocationPoliceStation:             "PoliceStation",
	LocationGuardianCitadelEntrance:   "GuardianCitadelEntrance",
	LocationGuardianCitadelOuter:      "GuardianCitadelOuter",
	LocationTempleMushroom:            "TempleMushroom",
	LocationFaranBrygos:               "FaranBrygos",
	LocationFatFreddys:                "FatFreddys",
	LocationSpadesCasino:              "SpadesCasino",
	LocationGuardianCitadelInner:      "GuardianCitadelInner",
	LocationMineShaft:                 "MineShaft",
	LocationSavageVillage:             "SavageVillage",
	LocationPrevious:                  "Previous",
}

var LocationBlockZIPMap = map[int]*BlockZIP{
	LocationWorldMap:                  &BlockZIP{0, Block0WorldMap},
	LocationQuartz:                    &BlockZIP{0, Block0Quartz},
	LocationScottsBar:                 &BlockZIP{0, Block0ScottsBar},
	LocationStageCoachInn:             &BlockZIP{0, Block0StageCoachInn},
	LocationUglysHideout:              &BlockZIP{0, Block0UglysHideout},
	LocationQuartzDerelictBuildings:   &BlockZIP{0, Block0QuartzDerelictBuildings},
	LocationCourthouse:                &BlockZIP{0, Block0Courthouse},
	LocationSleeperBaseLevel1:         &BlockZIP{1, Block1SleeperBaseLevel1},
	LocationDesertNomads:              &BlockZIP{0, Block0DesertNomads},
	LocationAgCenter:                  &BlockZIP{0, Block0AgCenter},
	LocationHighpool:                  &BlockZIP{0, Block0Highpool},
	LocationLasVegasDerelictBuildings: &BlockZIP{1, Block1LasVegasDerelictBuildings},
	LocationLasVegas:                  &BlockZIP{1, Block1LasVegas},
	LocationSleeperBaseLevel2:         &BlockZIP{1, Block1SleeperBaseLevel2},
	LocationSleeperBaseLevel3:         &BlockZIP{1, Block1SleeperBaseLevel3},
	LocationBaseCochiseOutside:        &BlockZIP{1, Block1BaseCochiseOutside},
	LocationBaseCochiseLevel1:         &BlockZIP{1, Block1BaseCochiseLevel1},
	LocationBaseCochiseLevel3:         &BlockZIP{1, Block1BaseCochiseLevel3},
	LocationBaseCochiseLevel2:         &BlockZIP{1, Block1BaseCochiseLevel2},
	LocationBaseCochiseLevel4:         &BlockZIP{1, Block1BaseCochiseLevel4},
	LocationDarwin:                    &BlockZIP{1, Block1Darwin},
	LocationDarwinBase:                &BlockZIP{1, Block1DarwinBase},
	LocationFinstersBrain:             &BlockZIP{1, Block1FinstersBrain},
	LocationLasVegasSewersWest:        &BlockZIP{1, Block1LasVegasSewersWest},
	LocationLasVegasSewersEast:        &BlockZIP{1, Block1LasVegasSewersEast},
	LocationNeedles:                   &BlockZIP{0, Block0Needles},
	LocationBloodTempleTop:            &BlockZIP{0, Block0BloodTempleTop},
	LocationBloodTempleBottom:         &BlockZIP{0, Block0BloodTempleBottom},
	LocationVerminCave:                &BlockZIP{0, Block0VerminCave},
	LocationWastePit:                  &BlockZIP{0, Block0WastePit},
	LocationNeedlesDowntownEast:       &BlockZIP{0, Block0NeedlesDowntownEast},
	LocationNeedlesDowntownWest:       &BlockZIP{0, Block0NeedlesDowntownWest},
	LocationPoliceStation:             &BlockZIP{0, Block0PoliceStation},
	LocationGuardianCitadelEntrance:   &BlockZIP{1, Block1GuardianCitadelEntrance},
	LocationGuardianCitadelOuter:      &BlockZIP{1, Block1GuardianCitadelOuter},
	LocationTempleMushroom:            &BlockZIP{1, Block1TempleMushroom},
	LocationFaranBrygos:               &BlockZIP{1, Block1FaranBrygos},
	LocationFatFreddys:                &BlockZIP{1, Block1FatFreddys},
	LocationSpadesCasino:              &BlockZIP{1, Block1SpadesCasino},
	LocationGuardianCitadelInner:      &BlockZIP{1, Block1GuardianCitadelInner},
	LocationMineShaft:                 &BlockZIP{0, Block0MineShaft},
	LocationSavageVillage:             &BlockZIP{0, Block0SavageVillage},
}
