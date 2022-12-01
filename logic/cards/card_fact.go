package cards

import (
	"hs/logic/config"
	"hs/logic/help"
	"hs/logic/iface"
	"math/rand"
	"strconv"
	"strings"
)

type CardFact struct {
	Pool map[string][]iface.ICard

	setConfigSign bool
}

func InitCardFact() {
	iface.ICF = &CardFact{
		Pool: make(map[string][]iface.ICard),
	}
}

// 获得卡牌
func (cf *CardFact) GetCard(configId int) iface.ICard {

	ac := cf.GetAllCards()
	c := ac[configId].NewPoint()
	c.SetConfig(config.GetCardConfig(configId))

	return c
}

// 获得多个卡牌
func (cf *CardFact) GetCards(configIds []int) []iface.ICard {
	var ret []iface.ICard = make([]iface.ICard, 0)

	for _, v := range configIds {
		ret = append(ret, cf.GetCard(v))
	}

	return ret
}

// 获得全部卡牌
func (cf *CardFact) GetAllCards() []iface.ICard {

	if !cf.setConfigSign {
		for k, v := range cardPoints {
			v.SetConfig(config.GetCardConfig(k))
		}

		cf.setConfigSign = true
	}

	return cardPoints
}

// 获得全部可携带的卡牌
func (cf *CardFact) GetAllCardsExcludeNotCanCarry() []iface.ICard {

	key := "acencc"
	pool, ok := cf.Pool[key]
	if !ok {
		pool = make([]iface.ICard, 0)

		for _, v := range cf.GetAllCards() {
			if v.GetConfig().CanCarry {
				pool = append(pool, v)
			}
		}

		cf.Pool[key] = pool
	}

	return pool
}

// 卡牌工厂随机个卡牌
func (cf *CardFact) RandByAllCards(r *rand.Rand, scp *iface.ScreenCardParam) iface.ICard {

	key := GetScpKey(scp)

	pool, ok := cf.Pool[key]
	if !ok {

		pool = cf.ScreenCards(cf.GetAllCardsExcludeNotCanCarry(), scp)
		cf.Pool[key] = pool
	}

	if len(pool) <= 0 {
		return nil
	}

	rk := r.Intn(len(pool))

	return pool[rk]
}

// 随机根据一些卡牌
func (cf *CardFact) RandByCards(op []iface.ICard, r *rand.Rand, scp *iface.ScreenCardParam) iface.ICard {

	pool := cf.ScreenCards(op, scp)

	if len(pool) <= 0 {
		return nil
	}

	rk := r.Intn(len(pool))

	return pool[rk]
}

// 筛选卡牌
func (cf *CardFact) ScreenCards(op []iface.ICard, scp *iface.ScreenCardParam) []iface.ICard {

	if scp == nil {
		return op
	}

	pool := make([]iface.ICard, 0)

	for _, v := range op {

		if scp.Mona != -1 && scp.Mona != v.GetConfig().Mona {
			continue
		}

		if scp.CardSerices != nil && !help.InArray(v.GetConfig().Series, scp.CardSerices) {
			continue
		}

		if scp.CardTypes != nil && !help.InArray(v.GetConfig().Ctype, scp.CardTypes) {
			continue
		}

		if scp.CardTraits != nil {

			for _, v2 := range v.GetConfig().Traits {
				if help.InArray(v2, scp.CardTraits) {
					goto endCardTraits
				}
			}

			continue
		}
	endCardTraits:

		if scp.CardRaces != nil {

			for _, v2 := range v.GetConfig().Race {
				if help.InArray(v2, scp.CardRaces) {
					goto endCardRaces
				}
			}
			continue
		}

	endCardRaces:

		pool = append(pool, v)
	}

	return pool
}

func GetScpKey(scp *iface.ScreenCardParam) string {

	if scp == nil {
		return "_"
	}

	str := strconv.Itoa(scp.Mona) + "_"

	if scp.CardSerices != nil {
		str += help.Implode(",", scp.CardSerices) + "_"
	} else {
		str += "0_"
	}

	if scp.CardTypes != nil {
		str += help.Implode(",", scp.CardTypes) + "_"
	} else {
		str += "0_"
	}

	if scp.CardTraits != nil {
		str += help.Implode(",", scp.CardTraits) + "_"
	} else {
		str += "0_"
	}

	if scp.CardTraits != nil {
		str += help.Implode(",", scp.CardTraits) + "_"
	} else {
		str += "0_"
	}

	str = strings.Trim(str, "_")

	return str
}
