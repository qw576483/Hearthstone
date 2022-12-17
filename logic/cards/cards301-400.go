package cards

import (
	"hs/logic/battle/bcard"
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
	"strconv"
)

// 闪电风暴
type Card301 struct {
	bcard.Card
}

func (c *Card301) NewPoint() iface.ICard {
	return &Card301{}
}

func (c *Card301) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	for _, v := range h.GetEnemy().GetBattleCards() {
		v.CostHp(c, dmg)
	}

	push.PushAutoLog(h, "[过载+2]")
	h.SetLockMonaCache(h.GetLockMonaCache() + 2)
}

// 视界术
type Card302 struct {
	bcard.Card
}

func (c *Card302) NewPoint() iface.ICard {
	return &Card302{}
}

func (c *Card302) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	ds := h.DrawByTimes(1)

	for _, v := range ds {
		v.SetMona(v.GetMona() - 3)
		push.PushLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"消耗-3")
	}
}

// 熔岩爆裂
type Card303 struct {
	bcard.Card
}

func (c *Card303) NewPoint() iface.ICard {
	return &Card303{}
}

func (c *Card303) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	push.PushAutoLog(h, "[过载+2]")
	h.SetLockMonaCache(h.GetLockMonaCache() + 2)

	if rc == nil {
		return
	}

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	rc.CostHp(c, dmg)
}

// 恶魔卫士
type Card304 struct {
	bcard.Card
}

func (c *Card304) NewPoint() iface.ICard {
	return &Card304{}
}

func (c *Card304) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.AddMonaMax(-1)

	push.PushAutoLog(h, "摧毁了自己的1个法力水晶")
}

// 虚空恐魔
type Card305 struct {
	bcard.Card
}

func (c *Card305) NewPoint() iface.ICard {
	return &Card305{}
}

func (c *Card305) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	bcs := h.GetBattleCards()

	var cLeft iface.ICard
	var cRight iface.ICard

	if (bidx - 1) >= 0 {
		cLeft = bcs[bidx-1]
	}
	if (bidx + 1) < len(bcs) {
		cRight = bcs[bidx+1]
	}

	lr := []iface.ICard{cLeft, cRight}

	for _, v := range lr {
		if v == nil {
			continue
		}

		ad := v.GetHaveEffectDamage()
		ahp := v.GetHaveEffectHp()

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddDamage(ad)
		buff.AddHpMaxAndHp(ahp)

		c.AddSubCards(buff)

		h.DieCard(v, false)

		push.PushAutoLog(h, push.GetCardLogString(c)+"吸收了"+push.GetCardLogString(v)+"获得了+"+strconv.Itoa(ad)+"/+"+strconv.Itoa(ahp))
	}
}

// 感知恶魔
type Card306 struct {
	bcard.Card
}

func (c *Card306) NewPoint() iface.ICard {
	return &Card306{}
}

func (c *Card306) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for i := 1; i <= 2; i++ {
		scs := iface.GetCardFact().ScreenCards(h.GetLibCards(), iface.NewScreenCardParam(
			iface.SCPWithCardRace([]define.CardRace{define.CardRaceDevil}),
		))

		randC := h.RandCard(scs)

		if randC == nil {
			nc := iface.GetCardFact().GetCard(309)
			nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
			h.MoveToHand(nc)
			continue
		}

		h.DrawByCard(randC)
	}
}

// 暴乱狂战士
type Card307 struct {
	bcard.Card
}

func (c *Card307) NewPoint() iface.ICard {
	return &Card307{}
}

func (c *Card307) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherAfterCostHp")
}

func (c *Card307) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherAfterCostHp")
}

func (c *Card307) OnNROtherAfterCostHp(who, target iface.ICard, num int) {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle || target.GetType() != define.CardTypeEntourage {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)
	c.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"获得了1点攻击力")
}

// 法力过剩
type Card308 struct {
	bcard.Card
}

func (c *Card308) NewPoint() iface.ICard {
	return &Card308{}
}

func (c *Card308) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.DrawByTimes(1)
}

// 游荡小鬼
type Card309 struct {
	bcard.Card
}

func (c *Card309) NewPoint() iface.ICard {
	return &Card309{}
}
