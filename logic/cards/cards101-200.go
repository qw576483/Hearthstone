package cards

import (
	"hs/logic/battle/bcard"
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
)

// 攻城恶魔
type Card101 struct {
	bcard.Card
}

func (c *Card101) NewPoint() iface.ICard {
	return &Card101{}
}

func (c *Card101) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card101) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")
}

func (c *Card101) OnNROtherGetDamage(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBattle ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	if oc.IsRace(define.CardRaceDevil) {
		return 1
	}

	return 0
}

// 作战傀儡
type Card102 struct {
	bcard.Card
}

func (c *Card102) NewPoint() iface.ICard {
	return &Card102{}
}

// 熔火恶犬
type Card103 struct {
	bcard.Card
}

func (c *Card103) NewPoint() iface.ICard {
	return &Card103{}
}

// 贫瘠之地饲养员
type Card104 struct {
	bcard.Card
}

func (c *Card104) NewPoint() iface.ICard {
	return &Card104{}
}

func (c *Card104) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	// 随机
	races := []define.CardRace{define.CardRaceBeast}
	vocations := []define.Vocation{h.GetConfig().Vocation, define.VocationNeutral}
	types := []define.CardType{define.CardTypeEntourage}
	ncCache := iface.GetCardFact().RandByAllCards(h.GetBattle().GetRand(), iface.NewScreenCardParam(
		iface.SCPWithCardRace(races), iface.SCPWithCardVocations(vocations), iface.SCPWithCardTypes(types),
	))

	if ncCache == nil {
		return
	}

	// 召唤
	nc := iface.GetCardFact().GetCard(ncCache.GetConfig().Id)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToBattle(nc, bidx+1)
	nc.SetReleaseRound(h.GetBattle().GetIncrRoundId())

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}
