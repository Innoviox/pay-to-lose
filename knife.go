package main

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
		panic("invalid knife type")
	}
}