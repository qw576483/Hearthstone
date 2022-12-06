package cards

import (
	"fmt"
	"hs/logic/battle/bcard"
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
	"math"
	"strconv"
)

// 幸运币
type Card0 struct {
	bcard.Card
}

func (c *Card0) NewPoint() iface.ICard {
	return &Card0{}
}

func (c *Card0) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {
	c.GetOwner().AddMona(1)
}

// 石牙野猪
type Card1 struct {
	bcard.Card
}

func (c *Card1) NewPoint() iface.ICard {
	return &Card1{}
}

// 疯狂的炼金师
type Card2 struct {
	bcard.Card
}

func (c *Card2) NewPoint() iface.ICard {
	return &Card2{}
}

func (c *Card2) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	if rc == nil {
		return
	}
	rc.ExchangeHpDamage(rc)
	rc.CostHp(0)

	// logs
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"血量和攻击互换")
}

// 寒光智者
type Card3 struct {
	bcard.Card
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
	bcard.Card
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
	bcard.Card
}

func (c *Card5) NewPoint() iface.ICard {
	return &Card5{}
}

// 攻城车
type Card6 struct {
	bcard.Card
}

func (c *Card6) NewPoint() iface.ICard {
	return &Card6{}
}

func (c *Card6) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNRRoundBegin")
}

func (c *Card6) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNRRoundBegin")
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

// 铸剑师
type Card7 struct {
	bcard.Card
}

func (c *Card7) NewPoint() iface.ICard {
	return &Card7{}
}

func (c *Card7) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card7) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNRRoundEnd")
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

// 螃蟹骑士
type Card8 struct {
	bcard.Card
}

func (c *Card8) NewPoint() iface.ICard {
	return &Card8{}
}

// 毁灭之刃
type Card9 struct {
	bcard.Card
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
	bcard.Card
}

func (c *Card10) NewPoint() iface.ICard {
	return &Card10{}
}

func (c *Card10) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card10) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNROtherDie")
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

// 上古看守者
type Card11 struct {
	bcard.Card
}

func (c *Card11) NewPoint() iface.ICard {
	return &Card11{}
}

// 持盾卫士
type Card12 struct {
	bcard.Card
}

func (c *Card12) NewPoint() iface.ICard {
	return &Card12{}
}

// 银色侍从
type Card13 struct {
	bcard.Card
}

func (c *Card13) NewPoint() iface.ICard {
	return &Card13{}
}

// 耐心的刺客
type Card14 struct {
	bcard.Card
}

func (c *Card14) NewPoint() iface.ICard {
	return &Card14{}
}

// 疯狂投弹者
type Card15 struct {
	bcard.Card
}

func (c *Card15) NewPoint() iface.ICard {
	return &Card15{}
}

func (c *Card15) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	h := c.GetOwner()
	for i := 1; i <= 3; i++ {
		rc, rh := h.RandBothBattleCardOrHero()

		if rc != nil {
			push.PushAutoLog(h, push.GetCardLogString(c)+"的炸药桶对"+push.GetCardLogString(rc)+"造成了1点伤害")
			rc.CostHp(1)
		}

		if rh != nil {
			push.PushAutoLog(h, push.GetCardLogString(c)+"的炸药桶对"+push.GetHeroLogString(rh)+"造成了1点伤害")
			rh.CostHp(1)
		}
	}
}

// 飞刀杂耍者
type Card16 struct {
	bcard.Card
}

func (c *Card16) NewPoint() iface.ICard {
	return &Card16{}
}

func (c *Card16) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNRPutToBattle")
}

func (c *Card16) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNRPutToBattle")
}

func (c *Card16) OnNRPutToBattle(oc iface.ICard) {
	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return
	}

	rc, rh := h.GetEnemy().RandBothBattleCardOrHero()

	if rc != nil {
		push.PushAutoLog(h, push.GetCardLogString(c)+"的飞刀对"+push.GetCardLogString(rc)+"造成了1点伤害")
		rc.CostHp(1)
	}

	if rh != nil {
		push.PushAutoLog(h, push.GetCardLogString(c)+"的飞刀对"+push.GetHeroLogString(rh)+"造成了1点伤害")
		rh.CostHp(1)
	}
}

// 火舌图腾
type Card17 struct {
	bcard.Card
}

func (c *Card17) NewPoint() iface.ICard {
	return &Card17{}
}

func (c *Card17) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card17) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNROtherGetDamage")
}

func (c *Card17) OnNROtherGetDamage(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBattle ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	cIdx := h.GetCardIdx(c, h.GetBattleCards())
	ocIdx := h.GetCardIdx(oc, h.GetBattleCards())

	if cIdx != -1 && ocIdx != -1 && (math.Abs(float64(cIdx)-float64(ocIdx)) == 1) {
		return 2
	}

	return 0
}

// 小个子召唤师
type Card18 struct {
	bcard.Card
}

func (c *Card18) NewPoint() iface.ICard {
	return &Card18{}
}

func (c *Card18) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNROtherGetMona")
}

func (c *Card18) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNROtherGetMona")
}

func (c *Card18) OnNROtherGetMona(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeHand ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	if h.GetReleaseCardTimes() == 0 {
		return -1
	}

	return 0
}

// 暴风城勇士
type Card19 struct {
	bcard.Card
}

func (c *Card19) NewPoint() iface.ICard {
	return &Card19{}
}

func (c *Card19) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNROtherGetHp")
	c.GetOwner().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card19) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNROtherGetHp")
	c.GetOwner().RemoveCardFromEvent(c, "OnNROtherGetDamage")
}

func (c *Card19) OnNROtherGetDamage(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBattle ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	return 1
}

func (c *Card19) OnNROtherGetHp(oc iface.ICard) int {
	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBattle ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}
	return 1
}

// 小精灵
type Card20 struct {
	bcard.Card
}

func (c *Card20) NewPoint() iface.ICard {
	return &Card20{}
}

// buff - 你的回合结束时消散
type Card21 struct {
	bcard.Card
}

func (c *Card21) NewPoint() iface.ICard {
	return &Card21{}
}

func (c *Card21) OnInit() {
	c.GetOwner().AddCardToEvent(c, "OnNRRoundEnd")
	c.GetOwner().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card21) OnNROtherDie(oc iface.ICard) {
	if oc.GetId() == c.GetFatherCard().GetId() {
		c.ClearBuff()
	}
}

func (c *Card21) OnNRRoundEnd() {

	fc := c.GetFatherCard()

	if fc == nil || fc.GetNoLoopOwner().GetId() != fc.GetOwner().GetBattle().GetRoundHero().GetId() {
		return
	}

	c.ClearBuff()
}

func (c *Card21) ClearBuff() {
	fc := c.GetFatherCard()
	if fc != nil {
		fc.RemoveSubCards(c)
		c.GetOwner().RemoveCardFromEvent(c, "OnNRRoundEnd")
		c.GetOwner().RemoveCardFromEvent(c, "OnNROtherDie")
	}
}

// buff - 你的回合开始时消散
type Card22 struct {
	bcard.Card
}

func (c *Card22) NewPoint() iface.ICard {
	return &Card22{}
}

func (c *Card22) OnInit() {
	c.GetOwner().AddCardToEvent(c, "OnNRRoundBegin")
	c.GetOwner().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card22) OnNROtherDie(oc iface.ICard) {
	if oc.GetId() == c.GetFatherCard().GetId() {
		c.ClearBuff()
	}
}

func (c *Card22) OnNRRoundBegin() {

	fc := c.GetFatherCard()
	if fc == nil || fc.GetNoLoopOwner().GetId() != fc.GetOwner().GetBattle().GetRoundHero().GetId() {
		return
	}
	c.ClearBuff()
}

func (c *Card22) ClearBuff() {
	fc := c.GetFatherCard()
	if fc != nil {
		fc.RemoveSubCards(c)
		c.GetOwner().AddCardToEvent(c, "OnNRRoundBegin")
		c.GetOwner().RemoveCardFromEvent(c, "OnNROtherDie")
	}
}

// buff - 永久生效
type Card23 struct {
	bcard.Card
}

func (c *Card23) NewPoint() iface.ICard {
	return &Card23{}
}

// 叫嚣的中士
type Card24 struct {
	bcard.Card
}

func (c *Card24) NewPoint() iface.ICard {
	return &Card24{}
}

func (c *Card24) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	if rc != nil {

		nc := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
		nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
		nc.AddDamage(2)

		rc.AddSubCards(rc, nc)

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了两点攻击力")
	}

}

// 银色保卫者
type Card25 struct {
	bcard.Card
}

func (c *Card25) NewPoint() iface.ICard {
	return &Card25{}
}

func (c *Card25) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	if rc != nil {
		rc.AddTraits(define.CardTraitsHolyShield)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了圣盾")
	}
}

// 盗贼基础技能
type Card26 struct {
	bcard.Card
}

func (c *Card26) NewPoint() iface.ICard {
	return &Card26{}
}

func (c *Card26) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	nc := iface.GetCardFact().GetCard(27)
	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	c.GetOwner().Release(nc, 0, 0, nil, nil, false)

	push.PushAutoLog(c.GetOwner(), "装备了匕首")
}

// 盗贼基础技能 - 匕首
type Card27 struct {
	bcard.Card
}

func (c *Card27) NewPoint() iface.ICard {
	return &Card27{}
}

// 铁喙猫头鹰
type Card28 struct {
	bcard.Card
}

func (c *Card28) NewPoint() iface.ICard {
	return &Card28{}
}

func (c *Card28) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	if rc != nil {
		rc.Silent()
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"沉默了"+push.GetCardLogString(rc))
	}
}

// 奉献
type Card29 struct {
	bcard.Card
}

func (c *Card29) NewPoint() iface.ICard {
	return &Card29{}
}

func (c *Card29) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	h := c.GetOwner()
	d := h.GetApDamage()
	d += 2

	for _, v := range h.GetEnemy().GetBattleCards() {
		v.CostHp(d)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"对"+push.GetCardLogString(v)+"造成了"+strconv.Itoa(d)+"点伤害")
	}

	h.GetEnemy().CostHp(d)
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"对"+push.GetHeroLogString(h.GetEnemy())+"造成了"+strconv.Itoa(d)+"点伤害")
}

// 狗头人地卜师
type Card30 struct {
	bcard.Card
}

func (c *Card30) NewPoint() iface.ICard {
	return &Card30{}
}

// 游学者周卓
type Card31 struct {
	bcard.Card
}

func (c *Card31) NewPoint() iface.ICard {
	return &Card31{}
}

func (c *Card31) OnPutToBattle(pidx int) {
	c.GetOwner().AddCardToEvent(c, "OnNROtherRelease")
}

func (c *Card31) OnOutBattle() {
	c.GetOwner().RemoveCardFromEvent(c, "OnNROtherRelease")
}

func (c *Card31) OnNROtherRelease(oc iface.ICard) bool {
	if oc.GetConfig().Ctype != define.CardTypeSorcery {
		return false
	}

	nc, err := oc.Copy(oc)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 设置所属人
	nc.SetOwner(oc.GetOwner().GetEnemy())
	nc.GetNoLoopOwner().MoveToHand(nc)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"复制了"+push.GetCardLogString(nc))

	return false
}

// 丛林守护者
type Card32 struct {
	bcard.Card
}

func (c *Card32) NewPoint() iface.ICard {
	return &Card32{}
}

func (c *Card32) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	if choiceId == 0 {
		if rc != nil {
			rc.CostHp(2)
			push.PushAutoLog(c.GetOwner(), "[抉择1]"+push.GetCardLogString(c)+"对"+push.GetCardLogString(rc)+"造成了2点伤害")
		} else if rh != nil {
			rh.CostHp(2)
			push.PushAutoLog(c.GetOwner(), "[抉择1]"+push.GetCardLogString(c)+"对"+push.GetHeroLogString(rh)+"造成了2点伤害")
		}
	} else {
		if rc != nil {
			rc.Silent()
			push.PushAutoLog(c.GetOwner(), "[抉择2]"+push.GetCardLogString(c)+"沉默了"+push.GetCardLogString(rc))
		}
	}
}

// 年轻的酒仙
type Card33 struct {
	bcard.Card
}

func (c *Card33) NewPoint() iface.ICard {
	return &Card33{}
}

func (c *Card33) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {

	h := c.GetOwner()
	if rc != nil && rc.GetOwner().GetId() == h.GetId() {
		h.MoveToHand(rc)
		push.PushAutoLog(h, push.GetCardLogString(c)+"将"+push.GetCardLogString(rc)+"移动回手牌")
	}
}
