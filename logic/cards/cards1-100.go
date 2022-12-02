package cards

import (
	"hs/logic/battle"
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
	"strconv"
)

// 幸运币
type Card0 struct {
	battle.Card
}

func (c *Card0) NewPoint() iface.ICard {
	return &Card0{}
}

func (c *Card0) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {
	c.GetOwner().AddMona(1)
}

// 石牙野猪
type Card1 struct {
	battle.Card
}

func (c *Card1) NewPoint() iface.ICard {
	return &Card1{}
}

// 疯狂的炼金师
type Card2 struct {
	battle.Card
}

func (c *Card2) NewPoint() iface.ICard {
	return &Card2{}
}

func (c *Card2) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	if rc == nil {
		return
	}
	th := rc.GetHp()
	td := rc.GetDamage()

	rc.SetDamage(th)
	rc.SetHpMaxAndHp(td)
	rc.CostHp(0)

	// logs
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"血量和攻击互换")
}

// 寒光智者
type Card3 struct {
	battle.Card
}

func (c *Card3) NewPoint() iface.ICard {
	return &Card3{}
}

func (c *Card3) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {
	c.GetOwner().DrawByTimes(2)
	c.GetOwner().GetEnemy().DrawByTimes(2)

	// logs
	push.PushAllLog(c.GetOwner().GetBattle(), "你和你的对手都抽了两张牌")
}

// 麦田傀儡
type Card4 struct {
	battle.Card
}

func (c *Card4) NewPoint() iface.ICard {
	return &Card4{}
}

// 死亡效果
func (c *Card4) OnDie(bidx int) {

	if len(c.GetOwner().GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(5)
	nc.Init(nc, define.InCardsTypeBattle, c.GetOwner(), c.GetOwner().GetBattle())
	nc.GetOwner().MoveToBattle(nc, bidx)
	nc.SetReleaseRound(c.GetReleaseRound())

	// logs
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"死亡时，召唤了"+push.GetCardLogString(nc))
}

// 损坏的傀儡（麦田傀儡衍生物）
type Card5 struct {
	battle.Card
}

func (c *Card5) NewPoint() iface.ICard {
	return &Card5{}
}

// 攻城车
type Card6 struct {
	battle.Card
}

func (c *Card6) NewPoint() iface.ICard {
	return &Card6{}
}

func (c *Card6) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNRRoundBegin")
}

func (c *Card6) OnNRRoundBegin() {

	// 在我的回合开始时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	e := h.GetEnemy()
	rc, rh := e.RandBattleCardOrHero()

	if rc != nil {
		push.PushAutoLog(h, push.GetCardLogString(c)+"的石头对"+push.GetCardLogString(rc)+"造成了2点伤害")
		rc.CostHp(2)
	}

	if rh != nil {
		push.PushAutoLog(h, push.GetCardLogString(c)+"的石头对"+push.GetHeroLogString(rh)+"造成了2点伤害")
		rh.CostHp(2)
	}
}

func (c *Card6) OnDie(bidx int) {
	c.GetOwner().RemoveCardFromEvent(c, "OnNRRoundBegin")
}

// 铸剑师
type Card7 struct {
	battle.Card
}

func (c *Card7) NewPoint() iface.ICard {
	return &Card7{}
}

func (c *Card7) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNRRoundEnd")
}

// 在你的回合结束时，随机使另一个友方随从获得+1攻击力。
func (c *Card7) OnNRRoundEnd() {

	// 在我的回合结束时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	tr := h.RandExcludeCard(h.GetBattleCards(), c)
	if tr == nil {
		return
	}

	tr.AddDamage(1)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(tr)+"提升1点攻击力")
}

func (c *Card7) OnDie(bidx int) {
	c.GetOwner().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

// 螃蟹骑士
type Card8 struct {
	battle.Card
}

func (c *Card8) NewPoint() iface.ICard {
	return &Card8{}
}

// 毁灭之刃
type Card9 struct {
	battle.Card
}

func (c *Card9) NewPoint() iface.ICard {
	return &Card9{}
}

func (c *Card9) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	h := c.GetOwner()
	dmg := 1
	if c.GetOwner().GetReleaseCardTimes() > 1 {
		push.PushAutoLog(h, push.GetCardLogString(c)+"触发了连击")
		dmg = 2
	} else {
		push.PushAutoLog(h, push.GetCardLogString(c)+"未触发连击")
	}

	if rc != nil {
		push.PushAutoLog(h, push.GetCardLogString(c)+"对"+push.GetCardLogString(rc)+"造成了"+strconv.Itoa(dmg)+"点伤害")
		rc.CostHp(dmg)
	}

	if rh != nil {
		push.PushAutoLog(h, push.GetCardLogString(c)+"对"+push.GetHeroLogString(rh)+"造成了"+strconv.Itoa(dmg)+"点伤害")
		rh.CostHp(dmg)
	}
}

// 食腐土狼
type Card10 struct {
	battle.Card
}

func (c *Card10) NewPoint() iface.ICard {
	return &Card10{}
}

func (c *Card10) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card10) OnNROtherDie(tc iface.ICard) {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		tc.GetOwner().GetId() != c.GetOwner().GetId() ||
		tc.GetId() == c.GetId() {
		return
	}

	push.PushAutoLog(c.GetOwner(), "由于"+push.GetCardLogString(tc)+"死亡,"+push.GetCardLogString(c)+"获得+2/+1")

	c.AddDamage(2)
	c.AddHpMaxAndHp(1)
}

func (c *Card10) OnDie(bidx int) {
	c.GetOwner().RemoveCardFromEvent(c, "OnNROtherDie")
}

// 上古看守者
type Card11 struct {
	battle.Card
}

func (c *Card11) NewPoint() iface.ICard {
	return &Card11{}
}

// 持盾卫士
type Card12 struct {
	battle.Card
}

func (c *Card12) NewPoint() iface.ICard {
	return &Card12{}
}

// 银色侍从
type Card13 struct {
	battle.Card
}

func (c *Card13) NewPoint() iface.ICard {
	return &Card13{}
}

// 耐心的刺客
type Card14 struct {
	battle.Card
}

func (c *Card14) NewPoint() iface.ICard {
	return &Card14{}
}
