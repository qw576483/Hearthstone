package cards

import (
	"hs/logic/battle/bcard"
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
	"strconv"
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

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 知识古树
type Card105 struct {
	bcard.Card
}

func (c *Card105) NewPoint() iface.ICard {
	return &Card105{}
}

func (c *Card105) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if choiceId == 0 {
		push.PushAutoLog(c.GetOwner(), "[抉择1]抽了2张牌")
		c.GetOwner().DrawByTimes(2)
	} else {
		if rc == nil {
			return
		}

		push.PushAutoLog(c.GetOwner(), "[抉择2]"+push.GetCardLogString(rc)+"恢复5点生命")
		rc.TreatmentHp(c, 5)
	}
}

// 战争古树
type Card106 struct {
	bcard.Card
}

func (c *Card106) NewPoint() iface.ICard {
	return &Card106{}
}

func (c *Card106) OnRelease(choiceId, bidx int, rc iface.ICard) {

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())

	if choiceId == 0 {
		push.PushAutoLog(c.GetOwner(), "[抉择1]"+push.GetCardLogString(c)+"获得了+5攻击力")
		buff.AddDamage(5)
	} else {
		push.PushAutoLog(c.GetOwner(), "[抉择2]"+push.GetCardLogString(c)+"+5生命值并具有嘲讽")
		buff.AddHpMaxAndHp(5)
		buff.AddTraits(define.CardTraitsTaunt)
	}

	c.AddSubCards(buff)
}

// 角斗士的长弓
type Card107 struct {
	bcard.Card
}

func (c *Card107) NewPoint() iface.ICard {
	return &Card107{}
}

func (c *Card107) OnBeforeAttack(ec iface.ICard) iface.ICard {
	c.GetOwner().GetHead().AddTraits(define.CardTraitsImmune)
	return ec
}

func (c *Card107) OnAfterAttack(ec iface.ICard) {
	c.GetOwner().GetHead().RemoveTraits(define.CardTraitsImmune)
}

// 大法师安东尼达斯
type Card108 struct {
	bcard.Card
}

func (c *Card108) NewPoint() iface.ICard {
	return &Card108{}
}

func (c *Card108) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card108) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card108) OnNROtherBeforeRelease(oc, rc iface.ICard) (iface.ICard, bool) {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeSorcery ||
		h.GetId() != oc.GetOwner().GetId() {
		return rc, true
	}

	nc := iface.GetCardFact().GetCard(109)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToHand(nc)

	push.PushAutoLog(h, push.GetHeroLogString(h)+"生成了"+push.GetCardLogString(nc))

	return rc, true
}

// 火球术
type Card109 struct {
	bcard.Card
}

func (c *Card109) NewPoint() iface.ICard {
	return &Card109{}
}

func (c *Card109) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	dmg := 6
	dmg += h.GetApDamage()

	if rc != nil {
		rc.CostHp(c, dmg)
	}
}

// 先知维伦
type Card110 struct {
	bcard.Card
}

func (c *Card110) NewPoint() iface.ICard {
	return &Card110{}
}

func (c *Card110) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeCostHp")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeTreatmentHp")
}

func (c *Card110) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherBeforeCostHp")
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherBeforeTreatmentHp")
}

func (c *Card110) OnNROtherBeforeCostHp(who iface.ICard, num int) int {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		(who.GetType() != define.CardTypeSorcery && who.GetType() != define.CardTypeHeroSkill) ||
		h.GetId() != who.GetOwner().GetId() {
		return num
	}

	if num == 0 {
		return num
	}

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(who)+"伤害翻倍"+strconv.Itoa(num)+"->"+strconv.Itoa(num*2))
	return num * 2
}

func (c *Card110) OnNROtherBeforeTreatmentHp(who iface.ICard, num int) int {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		(who.GetType() != define.CardTypeSorcery && who.GetType() != define.CardTypeHeroSkill) ||
		h.GetId() != who.GetOwner().GetId() {
		return num
	}

	if num == 0 {
		return num
	}

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(who)+"治疗翻倍"+strconv.Itoa(num)+"->"+strconv.Itoa(num*2))

	return num * 2
}

// 血吼
type Card111 struct {
	bcard.Card
	attackCard bool
}

func (c *Card111) NewPoint() iface.ICard {
	return &Card111{}
}

func (c *Card111) OnAfterAttack(ec iface.ICard) {

	if ec.GetType() == define.CardTypeEntourage {
		c.attackCard = true
	}
}

func (c *Card111) OnBeforeCostHp(num int) int {

	h := c.GetOwner()
	// 如果攻击卡牌
	if c.attackCard {

		push.PushAutoLog(h, push.GetCardLogString(c)+"攻击的是卡牌,不掉耐久,改为-1攻击")

		num -= 1
		c.attackCard = false
		c.SetDamage(c.GetDamage() - 1)

		if c.GetDamage() <= 0 {
			push.PushAutoLog(h, push.GetCardLogString(c)+"攻击<=0 , 强制破碎")
			h.DieCard(c, false)
		}
	}

	return num
}

// 拉文霍德刺客
type Card112 struct {
	bcard.Card
}

func (c *Card112) NewPoint() iface.ICard {
	return &Card112{}
}

// 迦顿男爵
type Card113 struct {
	bcard.Card
}

func (c *Card113) NewPoint() iface.ICard {
	return &Card113{}
}

func (c *Card113) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card113) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card113) OnNRRoundEnd() {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!h.IsRoundHero() {
		return
	}

	for _, v := range h.GetBattleCards() {
		if v.GetId() == c.GetId() {
			continue
		}
		v.CostHp(c, 2)
	}

	for _, v := range h.GetEnemy().GetBattleCards() {
		v.CostHp(c, 2)

	}
	h.GetEnemy().GetHead().CostHp(c, 2)
	h.GetHead().CostHp(c, 2)
}
