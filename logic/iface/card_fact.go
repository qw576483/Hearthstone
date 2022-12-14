package iface

import (
	"hs/logic/define"
	"math/rand"
)

type ICardFact interface {
	GetCard(int) ICard
	GetCards([]int) []ICard
	GetAllCards() []ICard
	GetAllCardsExcludeNotCanCarry() []ICard
	RandByAllCards(*rand.Rand, *ScreenCardParam) ICard
	RandByCards([]ICard, *rand.Rand, *ScreenCardParam) ICard
	ScreenCards([]ICard, *ScreenCardParam) []ICard
}

var ICF ICardFact

func GetCardFact() ICardFact {
	return ICF
}

// 随机卡牌参数
type ScreenCardParam struct {
	Mona          int
	CardSerices   []define.CardSeries
	CardTypes     []define.CardType
	CardTraits    []define.CardTraits
	CardRaces     []define.CardRace
	CardQuality   []define.CardQuality
	CardVocations []define.Vocation
}

type RandCardOption func(*ScreenCardParam)

// 筛选费用
func SCPWithMona(max int) RandCardOption {
	return func(q *ScreenCardParam) {
		q.Mona = max
	}
}

// 筛选卡牌系列
func SCPWithCardSerices(cs []define.CardSeries) RandCardOption {
	return func(q *ScreenCardParam) {
		q.CardSerices = cs
	}
}

// 筛选卡牌类型
func SCPWithCardTypes(ct []define.CardType) RandCardOption {
	return func(q *ScreenCardParam) {
		q.CardTypes = ct
	}
}

func SCPWithCommonCardTypes() RandCardOption {
	return func(q *ScreenCardParam) {
		q.CardTypes = []define.CardType{
			define.CardTypeEntourage,
			define.CardTypeWeapon,
			define.CardTypeSorcery,
		}
	}
}

// 筛选卡牌特质
func SCPWithCardTraits(ct []define.CardTraits) RandCardOption {
	return func(q *ScreenCardParam) {
		q.CardTraits = ct
	}
}

// 筛选卡牌种族
func SCPWithCardRace(cr []define.CardRace) RandCardOption {
	return func(q *ScreenCardParam) {
		q.CardRaces = cr
	}
}

// 筛选卡牌品质
func SCPWithCardQuality(cr []define.CardQuality) RandCardOption {
	return func(q *ScreenCardParam) {
		q.CardQuality = cr
	}
}

// 筛选卡牌职业
func SCPWithCardVocations(cr []define.Vocation) RandCardOption {
	return func(q *ScreenCardParam) {
		q.CardVocations = cr
	}
}

// 创建筛选
func NewScreenCardParam(options ...RandCardOption) *ScreenCardParam {
	ScreenCardParam := &ScreenCardParam{
		Mona: -1,
	}

	for _, o := range options {
		o(ScreenCardParam)
	}

	return ScreenCardParam
}
