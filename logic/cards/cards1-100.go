package cards

import (
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

func (c *Card0) OnRelease(choiceId, bidx int, rc iface.ICard) {
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

func (c *Card2) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}

	rc.ExchangeHpDamage()
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

func (c *Card3) OnRelease(choiceId, bidx int, rc iface.ICard) {
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

func (c *Card4) OnDie() {

	if len(c.GetOwner().GetBattleCards()) >= define.MaxBattleNum {
		return
	}
	dbidx := c.GetAfterDieBidx()

	nc := iface.GetCardFact().GetCard(5)
	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	nc.GetOwner().MoveToBattle(nc, dbidx)
	nc.SetReleaseRound(c.GetOwner().GetBattle().GetIncrRoundId())

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

func (c *Card6) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundBegin")
}

func (c *Card6) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundBegin")
}

func (c *Card6) OnNRRoundBegin() {

	// 在我的回合开始时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	e := h.GetEnemy()

	rc := e.RandBattleCardOrHero()

	push.PushAutoLog(h, push.GetCardLogString(c)+"对"+push.GetCardLogString(rc)+"造成了2点伤害")
	rc.CostHp(2)
}

// 铸剑师
type Card7 struct {
	bcard.Card
}

func (c *Card7) NewPoint() iface.ICard {
	return &Card7{}
}

func (c *Card7) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card7) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card7) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	tr := h.RandExcludeCard(h.GetBattleCards(), c)
	if tr == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)

	tr.AddSubCards(buff)

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

func (c *Card9) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	h := c.GetOwner()
	dmg := 1
	if c.GetOwner().GetReleaseCardTimes() > 1 {
		push.PushAutoLog(h, push.GetCardLogString(c)+"触发了连击")
		dmg = 2
	} else {
		push.PushAutoLog(h, push.GetCardLogString(c)+"未触发连击")
	}

	push.PushAutoLog(h, push.GetCardLogString(c)+"对"+push.GetCardLogString(rc)+"造成了"+strconv.Itoa(dmg)+"点伤害")
	rc.CostHp(dmg)
}

// 食腐土狼
type Card10 struct {
	bcard.Card
}

func (c *Card10) NewPoint() iface.ICard {
	return &Card10{}
}

func (c *Card10) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card10) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherDie")
}

func (c *Card10) OnNROtherDie(tc iface.ICard) {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		tc.GetOwner().GetId() != c.GetOwner().GetId() ||
		tc.GetId() == c.GetId() ||
		tc.GetConfig().Ctype != define.CardTypeEntourage ||
		!tc.IsRace(define.CardRaceBeast) {
		return
	}

	push.PushAutoLog(c.GetOwner(), "由于"+push.GetCardLogString(tc)+"死亡,"+push.GetCardLogString(c)+"获得+2/+1")

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(2)
	buff.AddHpMaxAndHp(1)

	c.AddSubCards(buff)
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

func (c *Card15) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	for i := 1; i <= 3; i++ {
		rc := h.RandBothBattleCardOrHero()

		push.PushAutoLog(h, push.GetCardLogString(c)+"的炸药桶对"+push.GetCardLogString(rc)+"造成了1点伤害")
		rc.CostHp(1)
	}
}

// 飞刀杂耍者
type Card16 struct {
	bcard.Card
}

func (c *Card16) NewPoint() iface.ICard {
	return &Card16{}
}

func (c *Card16) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRPutToBattle")
}

func (c *Card16) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRPutToBattle")
}

func (c *Card16) OnNRPutToBattle(oc iface.ICard) {
	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return
	}

	rc := h.GetEnemy().RandBattleCardOrHero()

	push.PushAutoLog(h, push.GetCardLogString(c)+"的飞刀对"+push.GetCardLogString(rc)+"造成了1点伤害")
	rc.CostHp(1)
}

// 火舌图腾
type Card17 struct {
	bcard.Card
}

func (c *Card17) NewPoint() iface.ICard {
	return &Card17{}
}

func (c *Card17) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card17) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")
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

func (c *Card18) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetMona")
}

func (c *Card18) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetMona")
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

func (c *Card19) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetHp")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card19) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetHp")
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")
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

// buff - 我的回合结束时消散
type Card21 struct {
	bcard.Card
}

func (c *Card21) NewPoint() iface.ICard {
	return &Card21{}
}

func (c *Card21) OnInit() {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card21) OnNROtherDie(oc iface.ICard) {
	if c.GetFatherCard() != nil && oc.GetId() == c.GetFatherCard().GetId() {
		c.ClearBuff()
	}
}

func (c *Card21) OnNRRoundEnd() {

	fc := c.GetFatherCard()
	if fc != nil && fc.GetNoLoopOwner().GetId() == fc.GetOwner().GetBattle().GetRoundHero().GetId() {
		c.ClearBuff()
		return
	}
}

func (c *Card21) ClearBuff() {

	fc := c.GetFatherCard()
	if fc != nil {
		fc.RemoveSubCards(c)
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherDie")
	}
}

// buff - 我的回合开始时消散
type Card22 struct {
	bcard.Card
}

func (c *Card22) NewPoint() iface.ICard {
	return &Card22{}
}

func (c *Card22) OnInit() {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundBegin")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card22) OnNROtherDie(oc iface.ICard) {
	if c.GetFatherCard() != nil && oc.GetId() == c.GetFatherCard().GetId() {
		c.ClearBuff()
	}
}

func (c *Card22) OnNRRoundBegin() {

	fc := c.GetFatherCard()
	if fc != nil && fc.GetNoLoopOwner().GetId() == fc.GetOwner().GetBattle().GetRoundHero().GetId() {
		c.ClearBuff()
		return
	}
}

func (c *Card22) ClearBuff() {
	fc := c.GetFatherCard()
	if fc != nil {
		fc.RemoveSubCards(c)
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundBegin")
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherDie")
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

func (c *Card24) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	buff.AddDamage(2)

	rc.AddSubCards(buff)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了两点攻击力")
}

// 银色保卫者
type Card25 struct {
	bcard.Card
}

func (c *Card25) NewPoint() iface.ICard {
	return &Card25{}
}

func (c *Card25) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc != nil {
		rc.AddTraits(define.CardTraitsHolyShield)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了圣盾")
	}
}

// 匕首精通
type Card26 struct {
	bcard.Card
}

func (c *Card26) NewPoint() iface.ICard {
	return &Card26{}
}

func (c *Card26) OnRelease(choiceId, bidx int, rc iface.ICard) {

	nc := iface.GetCardFact().GetCard(27)
	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	c.GetOwner().Release(nc, 0, 0, nil, false)

	push.PushAutoLog(c.GetOwner(), "装备了"+nc.GetConfig().Name)
}

// 邪恶短刀
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

func (c *Card28) OnRelease(choiceId, bidx int, rc iface.ICard) {

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

func (c *Card29) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	d := h.GetApDamage()
	d += 2

	for _, v := range h.GetEnemy().GetBattleCards() {
		v.CostHp(d)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"对"+push.GetCardLogString(v)+"造成了"+strconv.Itoa(d)+"点伤害")
	}

	h.GetEnemy().GetHead().CostHp(d)
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

func (c *Card31) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card31) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card31) OnNROtherBeforeRelease(oc, rc iface.ICard) (iface.ICard, bool) {
	if oc.GetConfig().Ctype != define.CardTypeSorcery {
		return rc, true
	}

	nc, err := oc.Copy()
	if err != nil {
		return rc, true
	}

	// 设置所属人
	nc.SetOwner(oc.GetOwner().GetEnemy())
	nc.GetNoLoopOwner().MoveToHand(nc)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"复制了"+push.GetCardLogString(nc))

	return rc, true
}

// 丛林守护者
type Card32 struct {
	bcard.Card
}

func (c *Card32) NewPoint() iface.ICard {
	return &Card32{}
}

func (c *Card32) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}
	if choiceId == 0 {
		rc.CostHp(2)
		push.PushAutoLog(c.GetOwner(), "[抉择1]"+push.GetCardLogString(c)+"对"+push.GetCardLogString(rc)+"造成了2点伤害")
	} else {

		if rc.GetCardInCardsPos() == define.InCardsTypeHead {
			return
		}
		rc.Silent()
		push.PushAutoLog(c.GetOwner(), "[抉择2]"+push.GetCardLogString(c)+"沉默了"+push.GetCardLogString(rc))
	}
}

// 年轻的酒仙
type Card33 struct {
	bcard.Card
}

func (c *Card33) NewPoint() iface.ICard {
	return &Card33{}
}

func (c *Card33) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}
	h := c.GetOwner()
	if rc.GetOwner().GetId() == h.GetId() {
		h.MoveToHand(rc)
		push.PushAutoLog(h, push.GetCardLogString(c)+"将"+push.GetCardLogString(rc)+"移动回手牌")
	}
}

// 忏悔
type Card34 struct {
	bcard.Card
}

func (c *Card34) NewPoint() iface.ICard {
	return &Card34{}
}

func (c *Card34) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherAfterRelease")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card34) OnNROtherAfterRelease(oc iface.ICard) {
	h := c.GetOwner()
	if oc.GetOwner().GetId() == h.GetEnemy().GetId() &&
		oc.GetType() == define.CardTypeEntourage && !h.IsRoundHero() {

		oc.SetHpMaxAndHp(1)
		h.DeleteSecret(c, true)
		h.GetBattle().RemoveCardFromAllEvent(c)

		push.PushAutoLog(h, c.GetConfig().Name+"(奥秘)让"+push.GetCardLogString(oc)+"生命值变为1点")
	}
}

// 狂野怒火
type Card35 struct {
	bcard.Card
}

func (c *Card35) NewPoint() iface.ICard {
	return &Card35{}
}

func (c *Card35) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}

	if !rc.IsRace(define.CardRaceBeast) {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	buff.AddDamage(2)
	buff.AddTraits(define.CardTraitsImmune)

	rc.AddSubCards(buff)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了两点攻击力和免疫")
}

// 闪电箭
type Card36 struct {
	bcard.Card
}

func (c *Card36) NewPoint() iface.ICard {
	return &Card36{}
}

func (c *Card36) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	dmg := 3
	dmg += h.GetApDamage()

	// 锁定一点法力值
	h.SetLockMonaCache(h.GetLockMonaCache() + 1)

	if rc != nil {
		push.PushAutoLog(h, "[过载+1]"+push.GetCardLogString(c)+"对"+push.GetCardLogString(rc)+"造成了"+strconv.Itoa(dmg)+"点伤害")
		rc.CostHp(dmg)
	}

}

// 古拉巴什狂暴者
type Card37 struct {
	bcard.Card
}

func (c *Card37) NewPoint() iface.ICard {
	return &Card37{}
}

func (c *Card37) OnAfterCostHp() {
	c.AddDamage(3)
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"获得了3点攻击")
}

// 熔核巨人
type Card38 struct {
	bcard.Card
}

func (c *Card38) NewPoint() iface.ICard {
	return &Card38{}
}

func (c *Card38) OnGetMona(m int) int {

	h := c.GetOwner()
	if c.GetCardInCardsPos() == define.InCardsTypeHand {
		m = m - (h.GetHead().GetHpMax() - h.GetHead().GetHp())
	}

	return m
}

// 阿曼尼狂战士
type Card39 struct {
	bcard.Card
	sub iface.ICard
}

func (c *Card39) NewPoint() iface.ICard {
	return &Card39{}
}

func (c *Card39) OnAfterHpChange() {

	if c.GetHaveEffectHp() < c.GetHaveEffectHpMax() && c.sub == nil {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
		buff.AddDamage(3)

		c.sub = buff
		c.AddSubCards(buff)

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"获得了3点攻击")

	} else if c.GetHaveEffectHp() >= c.GetHaveEffectHpMax() && c.sub != nil {

		c.RemoveSubCards(c.sub)
		c.sub = nil

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"失去了3点攻击")
	}

}

// 冰冻陷阱
type Card40 struct {
	bcard.Card
}

func (c *Card40) NewPoint() iface.ICard {
	return &Card40{}
}

func (c *Card40) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeAttack")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card40) OnNROtherBeforeAttack(oc, rc iface.ICard) iface.ICard {

	h := c.GetOwner()
	if oc.GetOwner().GetId() == h.GetEnemy().GetId() &&
		oc.GetType() == define.CardTypeEntourage && !h.IsRoundHero() {

		rc = nil

		h.DeleteSecret(c, true)
		h.GetBattle().RemoveCardFromAllEvent(c)

		h.MoveToHand(oc)
		oc.SetMona(oc.GetMona() + 2)

		push.PushAutoLog(h, push.GetCardLogString(c)+"将"+push.GetCardLogString(oc)+"移动回手牌，并+2费")
	}

	return rc
}

// 松鼠
type Card41 struct {
	bcard.Card
}

func (c *Card41) NewPoint() iface.ICard {
	return &Card41{}
}

// 魔暴龙
type Card42 struct {
	bcard.Card
}

func (c *Card42) NewPoint() iface.ICard {
	return &Card42{}
}

// 工匠大师欧沃斯巴克
type Card43 struct {
	bcard.Card
}

func (c *Card43) NewPoint() iface.ICard {
	return &Card43{}
}

func (c *Card43) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	h := c.GetOwner()
	rch := rc.GetOwner()

	rcbidx := h.GetCardIdx(rc, rch.GetBattleCards())
	rch.MoveOutBattleOnlyBattleCards(rc)

	var nc iface.ICard
	if h.GetBattle().GetRand().Intn(2) == 0 {
		nc = iface.GetCardFact().GetCard(41)
		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"变成了小松鼠...")
	} else {
		nc = iface.GetCardFact().GetCard(42)
		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"变成了魔暴龙！")
	}

	nc.Init(nc, define.InCardsTypeNone, rch, rch.GetBattle())
	nc.SetReleaseRound(rch.GetBattle().GetIncrRoundId())
	rch.MoveToBattle(nc, rcbidx)

}

// 希尔瓦娜斯·风行者
type Card44 struct {
	bcard.Card
}

func (c *Card44) NewPoint() iface.ICard {
	return &Card44{}
}

func (c *Card44) OnDie() {

	h := c.GetOwner()
	eh := h.GetEnemy()

	rc := eh.RandCard(h.GetEnemy().GetBattleCards())
	if rc == nil {
		return
	}

	dbidx := c.GetAfterDieBidx()

	h.CaptureCard(rc, dbidx)
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"死亡时，夺取了"+push.GetCardLogString(rc))
}

// 生命分流
type Card45 struct {
	bcard.Card
}

func (c *Card45) NewPoint() iface.ICard {
	return &Card45{}
}

func (c *Card45) OnRelease(choiceId, bidx int, rc iface.ICard) {

	c.GetOwner().DrawByTimes(1)
	push.PushAutoLog(c.GetOwner(), "抽了一张牌")
}

// 稳固射击
type Card46 struct {
	bcard.Card
}

func (c *Card46) NewPoint() iface.ICard {
	return &Card46{}
}

func (c *Card46) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	e := h.GetEnemy()
	e.GetHead().CostHp(2)
	push.PushAutoLog(c.GetOwner(), push.GetHeroLogString(h)+"对"+push.GetHeroLogString(e)+"造成了两点伤害")
}

// 图腾召唤
type Card47 struct {
	bcard.Card
}

func (c *Card47) NewPoint() iface.ICard {
	return &Card47{}
}

func (c *Card47) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if len(c.GetOwner().GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	h := c.GetOwner()

	randIdx := h.GetBattle().GetRand().Intn(len(define.ShamanBaseTotemsIds))
	nc := iface.GetCardFact().GetCard(define.ShamanBaseTotemsIds[randIdx])

	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	nc.GetOwner().MoveToBattle(nc, -1)
	nc.SetReleaseRound(c.GetOwner().GetBattle().GetIncrRoundId())

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 空气之怒图腾
type Card48 struct {
	bcard.Card
}

func (c *Card48) NewPoint() iface.ICard {
	return &Card48{}
}

// 灼热图腾
type Card49 struct {
	bcard.Card
}

func (c *Card49) NewPoint() iface.ICard {
	return &Card49{}
}

// 治疗图腾
type Card50 struct {
	bcard.Card
}

func (c *Card50) NewPoint() iface.ICard {
	return &Card50{}
}

func (c *Card50) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card50) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card50) OnNRRoundEnd() {

	// 在我的回合结束时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()

	for _, v := range h.GetBattleCards() {
		v.TreatmentHp(1)
	}
	push.PushAutoLog(h, push.GetCardLogString(c)+"让所有随从恢复1点生命值")
}

// 石爪图腾
type Card51 struct {
	bcard.Card
}

func (c *Card51) NewPoint() iface.ICard {
	return &Card51{}
}

// 力量图腾
type Card52 struct {
	bcard.Card
}

func (c *Card52) NewPoint() iface.ICard {
	return &Card52{}
}

func (c *Card52) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card52) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card52) OnNRRoundEnd() {

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

// 援军
type Card53 struct {
	bcard.Card
}

func (c *Card53) NewPoint() iface.ICard {
	return &Card53{}
}

func (c *Card53) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if len(c.GetOwner().GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(define.SilverHandRecruitId)

	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	nc.GetOwner().MoveToBattle(nc, -1)
	nc.SetReleaseRound(c.GetOwner().GetBattle().GetIncrRoundId())

	push.PushAutoLog(c.GetOwner(), "召唤了"+push.GetCardLogString(nc))
}

// 白银之手新兵
type Card54 struct {
	bcard.Card
}

func (c *Card54) NewPoint() iface.ICard {
	return &Card54{}
}

// 变形
type Card55 struct {
	bcard.Card
}

func (c *Card55) NewPoint() iface.ICard {
	return &Card55{}
}
func (c *Card55) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)

	h.GetHead().AddSubCards(buff)
	h.GetHead().SetShield(h.GetHead().GetShield() + 1)

	push.PushAutoLog(h, push.GetHeroLogString(h)+"获得了1点攻击力,1点护盾")
}

// 火焰冲击
type Card56 struct {
	bcard.Card
}

func (c *Card56) NewPoint() iface.ICard {
	return &Card56{}
}

func (c *Card56) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}
	h := c.GetOwner()

	push.PushAutoLog(h, push.GetHeroLogString(h)+"对"+push.GetCardLogString(rc)+"造成了1点伤害")
	rc.CostHp(1)
}

// 次级治疗术
type Card57 struct {
	bcard.Card
}

func (c *Card57) NewPoint() iface.ICard {
	return &Card57{}
}

func (c *Card57) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	h := c.GetOwner()

	push.PushAutoLog(h, push.GetHeroLogString(h)+"让"+push.GetCardLogString(rc)+"恢复了两点生命值")
	rc.TreatmentHp(2)
}

// 全副武装！
type Card58 struct {
	bcard.Card
}

func (c *Card58) NewPoint() iface.ICard {
	return &Card58{}
}

func (c *Card58) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.GetHead().SetShield(h.GetHead().GetShield() + 2)

	push.PushAutoLog(h, push.GetHeroLogString(h)+"获得了2点护盾")
}

// 山岭巨人
type Card59 struct {
	bcard.Card
}

func (c *Card59) NewPoint() iface.ICard {
	return &Card59{}
}

func (c *Card59) OnGetMona(m int) int {

	h := c.GetOwner()
	if c.GetCardInCardsPos() == define.InCardsTypeHand {
		if len(h.GetHandCards()) > 1 {
			m = m - len(h.GetHandCards()) + 1
		}
	}

	return m
}

// 海巨人
type Card60 struct {
	bcard.Card
}

func (c *Card60) NewPoint() iface.ICard {
	return &Card60{}
}

func (c *Card60) OnGetMona(m int) int {

	h := c.GetOwner()
	if c.GetCardInCardsPos() == define.InCardsTypeHand {
		m = m - len(h.GetBattleCards()) - len(h.GetEnemy().GetBattleCards())
	}

	return m
}

// 死亡之翼
type Card61 struct {
	bcard.Card
}

func (c *Card61) NewPoint() iface.ICard {
	return &Card61{}
}

func (c *Card61) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for _, v := range h.GetBattleCards() {
		h.DieCard(v, false)
		push.PushLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(v))
	}

	for _, v := range h.GetEnemy().GetBattleCards() {
		h.GetEnemy().DieCard(v, false)
		push.PushLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(v))
	}

	hcs := h.GetHandCards()
	for _, v := range hcs {
		if v.GetId() != c.GetId() {
			h.DiscardCard(v)
			push.PushLog(h, push.GetCardLogString(c)+"丢弃了"+push.GetCardLogString(v))
		}
	}
}

// 炎爆术
type Card62 struct {
	bcard.Card
}

func (c *Card62) NewPoint() iface.ICard {
	return &Card62{}
}

func (c *Card62) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}
	h := c.GetOwner()
	dmg := 10
	dmg += h.GetApDamage()

	push.PushAutoLog(h, push.GetCardLogString(c)+"对"+push.GetCardLogString(rc)+"造成了"+strconv.Itoa(dmg)+"点伤害")
	rc.CostHp(dmg)
}

// 精神控制
type Card63 struct {
	bcard.Card
}

func (c *Card63) NewPoint() iface.ICard {
	return &Card63{}
}

func (c *Card63) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc != nil {
		h.CaptureCard(rc, -1)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"夺取了"+push.GetCardLogString(rc))
	}
}

// buff - 我的回合结束时消灭宿主和自己
type Card64 struct {
	bcard.Card
}

func (c *Card64) NewPoint() iface.ICard {
	return &Card64{}
}

func (c *Card64) OnInit() {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card64) OnNROtherDie(oc iface.ICard) {
	if c.GetFatherCard() != nil && oc.GetId() == c.GetFatherCard().GetId() {
		c.ClearBuff()
	}
}

func (c *Card64) OnNRRoundEnd() {

	fc := c.GetFatherCard()
	if fc != nil && fc.GetNoLoopOwner().GetId() == fc.GetOwner().GetBattle().GetRoundHero().GetId() {
		c.ClearBuff()

		if fc.GetCardInCardsPos() == define.InCardsTypeBattle {
			fc.GetOwner().DieCard(fc, false)
		}
		return
	}
}

func (c *Card64) ClearBuff() {

	fc := c.GetFatherCard()
	if fc != nil {
		fc.RemoveSubCards(c)
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherDie")
	}
}

// buff - 我的回合开始时消灭宿主和自己
type Card65 struct {
	bcard.Card
}

func (c *Card65) NewPoint() iface.ICard {
	return &Card65{}
}

func (c *Card65) OnInit() {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundBegin")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card65) OnNROtherDie(oc iface.ICard) {
	if c.GetFatherCard() != nil && oc.GetId() == c.GetFatherCard().GetId() {
		c.ClearBuff()
	}
}

func (c *Card65) OnNRRoundBegin() {

	fc := c.GetFatherCard()
	if fc != nil && fc.GetNoLoopOwner().GetId() == fc.GetOwner().GetBattle().GetRoundHero().GetId() {
		c.ClearBuff()

		if fc.GetCardInCardsPos() == define.InCardsTypeBattle {
			fc.GetOwner().DieCard(fc, false)
		}
		return
	}
}

func (c *Card65) ClearBuff() {
	fc := c.GetFatherCard()
	if fc != nil {
		fc.RemoveSubCards(c)
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundBegin")
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherDie")
	}
}

// 加拉克苏斯大王
type Card66 struct {
	bcard.Card
}

func (c *Card66) NewPoint() iface.ICard {
	return &Card66{}
}

func (c *Card66) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	h.Henshin(c)

	nc := iface.GetCardFact().GetCard(69)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.Release(nc, 0, 0, nil, false)

	push.PushAutoLog(h, "装备了"+nc.GetConfig().Name)
}

// 地狱火！
type Card67 struct {
	bcard.Card
}

func (c *Card67) NewPoint() iface.ICard {
	return &Card67{}
}

func (c *Card67) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(68)

	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToBattle(nc, -1)
	nc.SetReleaseRound(h.GetBattle().GetIncrRoundId())

	push.PushAutoLog(h, "召唤了"+push.GetCardLogString(nc))
}

// 地狱火
type Card68 struct {
	bcard.Card
}

func (c *Card68) NewPoint() iface.ICard {
	return &Card68{}
}

// 血怒
type Card69 struct {
	bcard.Card
}

func (c *Card69) NewPoint() iface.ICard {
	return &Card69{}
}

// 暴龙王克鲁什
type Card70 struct {
	bcard.Card
}

func (c *Card70) NewPoint() iface.ICard {
	return &Card70{}
}

// 基础英雄
type Card71 struct {
	bcard.Card
}

func (c *Card71) NewPoint() iface.ICard {
	return &Card71{}
}

// 奥妮克希亚
type Card72 struct {
	bcard.Card
}

func (c *Card72) NewPoint() iface.ICard {
	return &Card72{}
}
func (c *Card72) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	left := true
	for i := 1; i <= define.MaxBattleNum; i++ {
		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return
		}

		bidx = h.GetCardIdx(c, h.GetBattleCards())

		nc := iface.GetCardFact().GetCard(define.LittleDragonId)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
		if left {
			h.MoveToBattle(nc, bidx)
		} else {
			h.MoveToBattle(nc, bidx+1)
		}
		nc.SetReleaseRound(h.GetBattle().GetIncrRoundId())

		push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))

		left = !left
	}

}

// 雏龙
type Card73 struct {
	bcard.Card
}

func (c *Card73) NewPoint() iface.ICard {
	return &Card73{}
}

// 诺兹多姆
type Card74 struct {
	bcard.Card
}

func (c *Card74) NewPoint() iface.ICard {
	return &Card74{}
}

func (c *Card74) OnPutToBattle(bidx int) {

	h := c.GetOwner()
	h.GetBattle().GetRoundHero().NewCountDown(15)
	push.PushAutoLog(h, push.GetCardLogString(c)+"将战斗时间调整到15s")

	h.GetBattle().AddCardToEvent(c, "OnNRGetBattleTime")
}

func (c *Card74) OnSilent() {
	c.OnOutBattle()
}

func (c *Card74) OnOutBattle() {
	h := c.GetOwner()
	h.GetBattle().GetRoundHero().NewCountDown(define.BattleTime)
	push.PushAutoLog(h, push.GetCardLogString(c)+"还原战斗时间为"+strconv.Itoa(define.BattleTime)+"s")

	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRGetBattleTime")
}

func (c *Card74) OnNRGetBattleTime(bt int) int {
	return 15
}

// 玛里苟斯
type Card75 struct {
	bcard.Card
}

func (c *Card75) NewPoint() iface.ICard {
	return &Card75{}
}

// 阿莱克丝塔萨
type Card76 struct {
	bcard.Card
}

func (c *Card76) NewPoint() iface.ICard {
	return &Card76{}
}

func (c *Card76) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeHero {
		return
	}

	if rc.GetHpMax() < 15 {
		rc.SetHpMax(15)
	}
	rc.SetHp(15)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"生命值变成了15")
}

// 伊瑟拉
type Card77 struct {
	bcard.Card
}

func (c *Card77) NewPoint() iface.ICard {
	return &Card77{}
}

func (c *Card77) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card77) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card77) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()

	randIdx := h.GetBattle().GetRand().Intn(len(define.YseraDreamIds))
	nc := iface.GetCardFact().GetCard(define.YseraDreamIds[randIdx])
	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())

	h.MoveToHand(nc)

	push.PushLog(h.GetEnemy(), "【你的对手】"+push.GetCardLogString(c)+"回合结束,给与一张梦境卡")
	push.PushLog(h, push.GetCardLogString(c)+"回合结束,给与一张"+push.GetCardLogString(nc))
}

// 梦魇
type Card78 struct {
	bcard.Card
}

func (c *Card78) NewPoint() iface.ICard {
	return &Card78{}
}

func (c *Card78) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundBeginFatherDie)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	buff.AddDamage(4)
	buff.SetHpMaxAndHp(4)

	rc.AddSubCards(buff)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了+4/+4")
}

// 梦境
type Card79 struct {
	bcard.Card
}

func (c *Card79) NewPoint() iface.ICard {
	return &Card79{}
}

func (c *Card79) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}
	h := c.GetOwner()
	if rc.GetOwner().GetId() != h.GetId() {
		rc.GetOwner().MoveToHand(rc)
		push.PushAutoLog(h, push.GetCardLogString(c)+"将"+push.GetCardLogString(rc)+"移动回手牌")
	}
}

// 伊瑟拉苏醒
type Card80 struct {
	bcard.Card
}

func (c *Card80) NewPoint() iface.ICard {
	return &Card80{}
}

func (c *Card80) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	d := h.GetApDamage()
	d += 5

	for _, v := range h.GetEnemy().GetBattleCards() {
		if v.GetConfig().Id == define.YseraId {
			continue
		}
		v.CostHp(d)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"对"+push.GetCardLogString(v)+"造成了"+strconv.Itoa(d)+"点伤害")
	}

	for _, v := range h.GetBattleCards() {
		if v.GetConfig().Id == define.YseraId {
			continue
		}
		v.CostHp(d)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"对"+push.GetCardLogString(v)+"造成了"+strconv.Itoa(d)+"点伤害")
	}
}

// 欢笑的姐妹
type Card81 struct {
	bcard.Card
}

func (c *Card81) NewPoint() iface.ICard {
	return &Card81{}
}

// 翡翠幼龙
type Card82 struct {
	bcard.Card
}

func (c *Card82) NewPoint() iface.ICard {
	return &Card82{}
}

// 埃隆巴克保护者
type Card83 struct {
	bcard.Card
}

func (c *Card83) NewPoint() iface.ICard {
	return &Card83{}
}

// 野性赐福
type Card84 struct {
	bcard.Card
}

func (c *Card84) NewPoint() iface.ICard {
	return &Card84{}
}

func (c *Card84) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	for _, v := range h.GetBattleCards() {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
		buff.AddDamage(2)
		buff.AddHpMaxAndHp(2)
		buff.AddTraits(define.CardTraitsTaunt)

		push.PushLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得了嘲讽并+2/+2")

		v.AddSubCards(buff)
	}
}

// 娜塔莉·塞林
type Card85 struct {
	bcard.Card
}

func (c *Card85) NewPoint() iface.ICard {
	return &Card85{}
}

func (c *Card85) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}
	addHp := rc.GetHaveEffectHp()

	rc.GetOwner().DieCard(rc, false)

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	buff.AddHpMaxAndHp(addHp)
	c.AddSubCards(buff)

	push.PushLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(rc)+"获得了并获得了"+strconv.Itoa(addHp)+"点生命值")
}

// 奥术吞噬者
type Card86 struct {
	bcard.Card
}

func (c *Card86) NewPoint() iface.ICard {
	return &Card86{}
}

func (c *Card86) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card86) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card86) OnNROtherBeforeRelease(oc, rc iface.ICard) (iface.ICard, bool) {

	if oc.GetConfig().Ctype != define.CardTypeSorcery || oc.GetOwner().GetId() != c.GetOwner().GetId() || c.GetCardInCardsPos() != define.InCardsTypeBattle {
		return rc, true
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	buff.AddDamage(2)
	buff.AddHpMaxAndHp(2)
	c.AddSubCards(buff)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"获得了+2/+2")

	return rc, true
}

// 格鲁尔
type Card87 struct {
	bcard.Card
}

func (c *Card87) NewPoint() iface.ICard {
	return &Card87{}
}

func (c *Card87) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card87) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card87) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle {
		return
	}
	h := c.GetOwner()

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)
	buff.AddHpMaxAndHp(1)

	c.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"获得了+1/+1")
}

// 塞纳留斯
type Card88 struct {
	bcard.Card
}

func (c *Card88) NewPoint() iface.ICard {
	return &Card88{}
}

func (c *Card88) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if choiceId == 0 {

		for _, v := range h.GetBattleCards() {
			if v.GetId() == c.GetId() {
				continue
			}

			buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
			buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
			buff.AddDamage(2)
			buff.AddHpMaxAndHp(2)

			v.AddSubCards(buff)

			push.PushAutoLog(c.GetOwner(), "[抉择1]"+push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得+2/+2")
		}
	} else {

		left := true
		for i := 1; i <= 2; i++ {
			if len(h.GetBattleCards()) >= define.MaxBattleNum {
				return
			}

			bidx = h.GetCardIdx(c, h.GetBattleCards())

			nc := iface.GetCardFact().GetCard(define.TreantTauntId)
			nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
			if left {
				h.MoveToBattle(nc, bidx)
			} else {
				h.MoveToBattle(nc, bidx+1)
			}
			nc.SetReleaseRound(h.GetBattle().GetIncrRoundId())

			push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))

			left = !left
		}
	}
}

// 树人
type Card89 struct {
	bcard.Card
}

func (c *Card89) NewPoint() iface.ICard {
	return &Card89{}
}

// 树人
type Card90 struct {
	bcard.Card
}

func (c *Card90) NewPoint() iface.ICard {
	return &Card90{}
}

// 提里奥·弗丁
type Card91 struct {
	bcard.Card
}

func (c *Card91) NewPoint() iface.ICard {
	return &Card91{}
}

func (c *Card91) OnDie() {

	nc := iface.GetCardFact().GetCard(92)
	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	c.GetOwner().Release(nc, 0, 0, nil, false)

	push.PushAutoLog(c.GetOwner(), "装备了"+nc.GetConfig().Name)
}

// 灰烬使者
type Card92 struct {
	bcard.Card
}

func (c *Card92) NewPoint() iface.ICard {
	return &Card92{}
}

// 圣疗术
type Card93 struct {
	bcard.Card
}

func (c *Card93) NewPoint() iface.ICard {
	return &Card93{}
}

func (c *Card93) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	rc.TreatmentHp(8)
	h.DrawByTimes(3)

	push.PushAutoLog(h, push.GetHeroLogString(h)+"抽了3张卡,并让"+push.GetCardLogString(rc)+"恢复了8点生命")
}

// 风领主奥拉基尔
type Card94 struct {
	bcard.Card
}

func (c *Card94) NewPoint() iface.ICard {
	return &Card94{}
}

// 扭曲虚空
type Card95 struct {
	bcard.Card
}

func (c *Card95) NewPoint() iface.ICard {
	return &Card95{}
}

func (c *Card95) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for _, v := range h.GetBattleCards() {
		h.DieCard(v, false)
		push.PushLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(v))
	}

	for _, v := range h.GetEnemy().GetBattleCards() {
		h.GetEnemy().DieCard(v, false)
		push.PushLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(v))
	}
}

// 格罗玛什·地狱咆哮
type Card96 struct {
	bcard.Card
	sub iface.ICard
}

func (c *Card96) NewPoint() iface.ICard {
	return &Card96{}
}

func (c *Card96) OnAfterHpChange() {

	if c.GetHaveEffectHp() < c.GetHaveEffectHpMax() && c.sub == nil {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
		buff.AddDamage(6)

		c.sub = buff
		c.AddSubCards(buff)

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"获得了6点攻击")

	} else if c.GetHaveEffectHp() >= c.GetHaveEffectHpMax() && c.sub != nil {

		c.RemoveSubCards(c.sub)
		c.sub = nil

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"失去了6点攻击")
	}

}

// 扭曲虚空
type Card97 struct {
	bcard.Card
}

func (c *Card97) NewPoint() iface.ICard {
	return &Card97{}
}

func (c *Card97) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card97) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card97) OnNRRoundEnd() {

	// 在我的回合结束时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()

	ec := h.GetEnemy().RandBattleCardOrHero()
	if ec != nil {
		ec.CostHp(8)
		push.PushAutoLog(h, push.GetCardLogString(c)+"对"+push.GetCardLogString(ec)+"造成了8点伤害")
	}
}
