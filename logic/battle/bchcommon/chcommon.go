package bchcommon

import (
	"hs/logic/help"
	"hs/logic/iface"
)

type CHCommon struct {
	Id int

	father   interface{}   // 父卡牌
	subCards []iface.ICard // 子卡牌，会拿到Hp，Damage ,ApDamage，Traits
}

func (chc *CHCommon) GetId() int {
	return chc.Id
}

// 获得父级
func (chc *CHCommon) GetFather() iface.ICHCommon {
	ichc, ok := chc.father.(iface.ICHCommon)
	if !ok {
		return nil
	}
	return ichc
}

// 获得父级card
func (chc *CHCommon) GetFatherCard() iface.ICard {
	ic, ok := chc.father.(iface.ICard)
	if !ok {
		return nil
	}
	return ic
}

// 获得父级英雄
func (chc *CHCommon) GetFatherHero() iface.IHero {
	ih, ok := chc.father.(iface.IHero)
	if !ok {
		return nil
	}
	return ih
}

// 设置父卡牌
func (chc *CHCommon) SetFather(f interface{}) {
	chc.father = f
}

// 获得子卡牌
func (chc *CHCommon) GetSubCards() []iface.ICard {
	return chc.subCards
}

// 设置子卡牌
func (chc *CHCommon) SetSubCards(scs []iface.ICard) {
	chc.subCards = scs
}

// 添加子卡牌
func (chc *CHCommon) AddSubCards(sc iface.ICard, father interface{}) {

	subCards := chc.GetSubCards()
	subCards = append(subCards, sc)
	chc.SetSubCards(subCards)

	sc.SetFather(father)
}

// 删除子卡牌
func (chc *CHCommon) RemoveSubCards(sc iface.ICard) {

	idx := -1
	for k, v := range chc.subCards {
		if v.GetId() == sc.GetId() {
			idx = k
			break
		}
	}

	if idx != -1 {
		_, chc.subCards = help.DeleteCardFromCardsByIdx(chc.subCards, idx)
	}
}
