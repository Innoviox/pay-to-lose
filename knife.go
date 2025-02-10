package main

import (
	"fmt"

	common "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type KnifeType int
const (
	KBayonet KnifeType = 500
	KnifeCSS KnifeType = 503
	KnifeFlip KnifeType = 505
	KnifeGut KnifeType = 506
	KnifeKarambit KnifeType = 507
	KnifeM9Bayonet KnifeType = 508
	KnifeTactical KnifeType = 509
	KnifeFalchion KnifeType = 512
	KnifeSurvivalBowie KnifeType = 514
	KnifeButterfly KnifeType = 515
	KnifePush KnifeType = 516
	KnifeCord KnifeType = 517
	KnifeCanis KnifeType = 518
	KnifeUrsus KnifeType = 519
	KnifeGypsyJacKnife KnifeType = 520
	KnifeOutdoor KnifeType = 521
	KnifeStiletto KnifeType = 522
	KnifeWidowmaker KnifeType = 523
	KnifeSkeleton KnifeType = 525
	KnifeKukri KnifeType = 526
)

var AllKnifeTypes = [20]KnifeType{
	500, 503, 505, 506, 507, 508, 509, 512, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 525, 526,
}

func ToString(k KnifeType) string {
	switch k {
	case KBayonet:
		return "Bayonet"
	case KnifeCSS:
		return "KnifeCSS"
	case KnifeFlip:
		return "KnifeFlip"
	case KnifeGut:
		return "KnifeGut"
	case KnifeKarambit:
		return "KnifeKarambit"
	case KnifeM9Bayonet:
		return "KnifeM9Bayonet"
	case KnifeTactical:
		return "KnifeTactical"
	case KnifeFalchion:
		return "KnifeFalchion"
	case KnifeSurvivalBowie:
		return "KnifeSurvivalBowie"
	case KnifeButterfly:
		return "KnifeButterfly"
	case KnifePush:
		return "KnifePush"
	case KnifeCord:
		return "KnifeCord"
	case KnifeCanis:
		return "KnifeCanis"
	case KnifeUrsus:
		return "KnifeUrsus"
	case KnifeGypsyJacKnife:
		return "KnifeGypsyJacKnife"
	case KnifeOutdoor:
		return "KnifeOutdoor"
	case KnifeStiletto:
		return "KnifeStiletto"
	case KnifeWidowmaker:
		return "KnifeWidowmaker"
	case KnifeSkeleton:
		return "KnifeSkeleton"
	case KnifeKukri:
		return "KnifeKukri"
	default:
		return fmt.Sprintf("Unknown knife type: %d", k)
	}
}

func GetKnife(p *common.Player) int {
	if p == nil {
		return 0
	}

	for _, weapon := range p.Weapons() {
		if weapon.Type == common.EqKnife {
			return GetKnifeType(weapon)
		}
	}
	
	return 0
}

func GetKnifeType(e *common.Equipment) int {
	var knifeType = e.Entity.Property("m_iItemDefinitionIndex").Value().S2UInt64()
	if knifeType < 500 || knifeType > 526 {
		return 0
	}

	return (int)(knifeType)
}