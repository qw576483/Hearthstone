package cards

import (
	"hs/logic/battle/bcard"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/iface"
	"hs/logic/push"
	"strconv"
)

// 横扫
type Card201 struct {
	bcard.Card
}

func (c *Card201) NewPoint() iface.ICard {
	return &Card201{}
}

func (c *Card201) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	dmg2 := 1
	dmg2 += h.GetApDamage()

	rc.CostHp(c, dmg)
	for _, v := range h.CardsToNewInstance(rc.GetOwner().GetBattleCards()) {
		if v.GetId() == rc.GetId() {
			continue
		}
		v.CostHp(c, dmg2)
	}

	if h.GetEnemy().GetHead().GetId() != rc.GetId() {
		h.GetEnemy().GetHead().CostHp(c, dmg2)
	}
}

// 驯兽师
type Card202 struct {
	bcard.Card
}

func (c *Card202) NewPoint() iface.ICard {
	return &Card202{}
}

func (c *Card202) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil || !rc.IsRace(define.CardRaceBeast) {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(2)
	buff.AddHpMaxAndHp(2)
	buff.AddTraits(define.CardTraitsTaunt)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得+2/+2和嘲讽")
}

// 多重射击
type Card203 struct {
	bcard.Card
}

func (c *Card203) NewPoint() iface.ICard {
	return &Card203{}
}

func (c *Card203) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	bcs := h.GetEnemy().GetBattleCards()

	for i := 1; i <= 2; i++ {
		randC := h.GetEnemy().RandCard(bcs)
		if randC == nil {
			return
		}

		randC.CostHp(c, dmg)
		_, bcs = help.DeleteCardFromCardsByIdx(bcs, h.GetEnemy().GetIdxByCards(randC, bcs))
	}
}

// 水元素
type Card204 struct {
	bcard.Card
}

func (c *Card204) NewPoint() iface.ICard {
	return &Card204{}
}

func (c *Card204) OnAfterCostOtherHp(ec iface.ICard) {
	h := c.GetOwner()
	ec.AddTraits(define.CardTraitsFrozen)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(ec)+"获得了冻结")
}

// 变形术
type Card205 struct {
	bcard.Card
}

func (c *Card205) NewPoint() iface.ICard {
	return &Card205{}
}

func (c *Card205) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	rch := rc.GetOwner()
	rcbidx := h.GetIdxByCards(rc, rch.GetBattleCards())
	rch.MoveOutBattleOnlyBattleCards(rc)

	nc := iface.GetCardFact().GetCard(206)
	nc.Init(nc, define.InCardsTypeNone, rch, rch.GetBattle())
	rch.MoveToBattle(nc, rcbidx)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"变成了"+push.GetCardLogString(nc))
}

// 绵羊
type Card206 struct {
	bcard.Card
}

func (c *Card206) NewPoint() iface.ICard {
	return &Card206{}
}

// 愤怒之锤
type Card207 struct {
	bcard.Card
}

func (c *Card207) NewPoint() iface.ICard {
	return &Card207{}
}

func (c *Card207) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	if rc != nil {
		rc.CostHp(c, dmg)
	}

	h.DrawByTimes(1)
}

// 王者祝福
type Card208 struct {
	bcard.Card
}

func (c *Card208) NewPoint() iface.ICard {
	return &Card208{}
}

func (c *Card208) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(4)
	buff.AddHpMaxAndHp(4)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得+4/+4")
}

// 真银圣剑
type Card209 struct {
	bcard.Card
}

func (c *Card209) NewPoint() iface.ICard {
	return &Card209{}
}

func (c *Card209) OnBeforeAttack(ec iface.ICard) iface.ICard {
	c.GetOwner().GetHead().TreatmentHp(c, 2)
	return ec
}

// 神圣新星
type Card210 struct {
	bcard.Card
}

func (c *Card210) NewPoint() iface.ICard {
	return &Card210{}
}

func (c *Card210) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	for _, v := range h.CardsToNewInstance(h.GetEnemy().GetBattleCards()) {
		v.CostHp(c, dmg)
	}

	for _, v := range h.CardsToNewInstance(h.GetBattleCards()) {
		v.TreatmentHp(c, 2)
	}
	h.GetHead().TreatmentHp(c, 2)
}

// 能量灌注
type Card211 struct {
	bcard.Card
}

func (c *Card211) NewPoint() iface.ICard {
	return &Card211{}
}

func (c *Card211) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(2)
	buff.AddHpMaxAndHp(6)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得+2/+6")
}

// 暗言术：毁
type Card212 struct {
	bcard.Card
}

func (c *Card212) NewPoint() iface.ICard {
	return &Card212{}
}

func (c *Card212) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	bcs := h.GetBattleCards()
	bcs = append(bcs, h.GetEnemy().GetBattleCards()...)

	for _, v := range bcs {
		if v.GetHaveEffectHp() >= 5 {
			v.GetOwner().DieCard(v, false)
			push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(v))
		}
	}
}

// 瘟疫使者
type Card213 struct {
	bcard.Card
}

func (c *Card213) NewPoint() iface.ICard {
	return &Card213{}
}

func (c *Card213) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddTraits(define.CardTraitsHighlyToxic)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得剧毒")
}

// 刺杀
type Card214 struct {
	bcard.Card
}

func (c *Card214) NewPoint() iface.ICard {
	return &Card214{}
}

func (c *Card214) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	rc.GetOwner().DieCard(rc, false)
	push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(rc))
}

// 刺客之刃
type Card215 struct {
	bcard.Card
}

func (c *Card215) NewPoint() iface.ICard {
	return &Card215{}
}

// 风语者
type Card216 struct {
	bcard.Card
}

func (c *Card216) NewPoint() iface.ICard {
	return &Card216{}
}

func (c *Card216) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddTraits(define.CardTraitsWindfury)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得风怒")
}

// 妖术
type Card217 struct {
	bcard.Card
}

func (c *Card217) NewPoint() iface.ICard {
	return &Card217{}
}

func (c *Card217) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	rch := rc.GetOwner()
	rcbidx := h.GetIdxByCards(rc, rch.GetBattleCards())
	rch.MoveOutBattleOnlyBattleCards(rc)

	nc := iface.GetCardFact().GetCard(218)
	nc.Init(nc, define.InCardsTypeNone, rch, rch.GetBattle())
	rch.MoveToBattle(nc, rcbidx)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"变成了"+push.GetCardLogString(nc))
}

// 青蛙
type Card218 struct {
	bcard.Card
}

func (c *Card218) NewPoint() iface.ICard {
	return &Card218{}
}

// 地狱烈焰
type Card219 struct {
	bcard.Card
}

func (c *Card219) NewPoint() iface.ICard {
	return &Card219{}
}

func (c *Card219) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	cs := h.GetEnemy().GetBattleCards()
	cs = append(cs, h.GetBattleCards()...)
	for _, v := range cs {
		v.CostHp(c, dmg)
	}

	h.GetHead().CostHp(c, dmg)
	h.GetEnemy().GetHead().CostHp(c, dmg)
}

// 库卡隆精英卫士
type Card220 struct {
	bcard.Card
}

func (c *Card220) NewPoint() iface.ICard {
	return &Card220{}
}

// 战歌侦察骑兵
type Card221 struct {
	bcard.Card
}

func (c *Card221) NewPoint() iface.ICard {
	return &Card221{}
}

// 冰风雪人
type Card222 struct {
	bcard.Card
}

func (c *Card222) NewPoint() iface.ICard {
	return &Card222{}
}

// 侏儒发明家
type Card223 struct {
	bcard.Card
}

func (c *Card223) NewPoint() iface.ICard {
	return &Card223{}
}

func (c *Card223) OnRelease(choiceId, bidx int, rc iface.ICard) {
	h := c.GetOwner()
	h.DrawByTimes(1)
}

// 机械幼龙技工
type Card224 struct {
	bcard.Card
}

func (c *Card224) NewPoint() iface.ICard {
	return &Card224{}
}

func (c *Card224) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(225)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 机械幼龙
type Card225 struct {
	bcard.Card
}

func (c *Card225) NewPoint() iface.ICard {
	return &Card225{}
}

// 暴风城骑士
type Card226 struct {
	bcard.Card
}

func (c *Card226) NewPoint() iface.ICard {
	return &Card226{}
}

// 老瞎眼
type Card227 struct {
	bcard.Card
}

func (c *Card227) NewPoint() iface.ICard {
	return &Card227{}
}

func (c *Card227) OnGetDamage(dmg int) int {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle {
		return dmg
	}

	h := c.GetOwner()

	cs := h.GetEnemy().GetBattleCards()
	cs = append(cs, h.GetBattleCards()...)
	for _, v := range cs {
		if v.IsRace(define.CardRaceFish) && v.GetId() != c.GetId() {
			dmg += 1
		}
	}

	return dmg
}

// 食人魔法师
type Card228 struct {
	bcard.Card
}

func (c *Card228) NewPoint() iface.ICard {
	return &Card228{}
}

// 食人魔法师
type Card229 struct {
	bcard.Card
}

func (c *Card229) NewPoint() iface.ICard {
	return &Card229{}
}

// 军情七处渗透者
type Card230 struct {
	bcard.Card
}

func (c *Card230) NewPoint() iface.ICard {
	return &Card230{}
}

func (c *Card230) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	ss := h.GetEnemy().GetSecrets()
	rs := h.GetEnemy().RandCard(ss)

	if rs != nil {
		h.GetEnemy().DeleteSecret(rs, false)
		push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+rs.GetConfig().Name+"(奥秘)")
	}
}

// 丛林之魂
type Card231 struct {
	bcard.Card
}

func (c *Card231) NewPoint() iface.ICard {
	return &Card231{}
}

func (c *Card231) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	for _, v := range h.CardsToNewInstance(h.GetBattleCards()) {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())

		buff.AddOnDie(func() {
			h := buff.GetOwner()
			if len(h.GetBattleCards()) >= define.MaxBattleNum {
				return
			}
			dbidx := v.GetAfterDieBidx()

			nc := iface.GetCardFact().GetCard(define.TreantId)
			nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
			nc.GetOwner().MoveToBattle(nc, dbidx)

			push.PushAutoLog(h, push.GetCardLogString(c)+"死亡时，召唤了"+push.GetCardLogString(nc))
		})

		v.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得亡语：召唤一个2/2的树人")
	}
}

// 撕咬
type Card232 struct {
	bcard.Card
}

func (c *Card232) NewPoint() iface.ICard {
	return &Card232{}
}

func (c *Card232) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(4)

	h.GetHead().AddSubCards(buff)
	h.GetHead().SetShield(h.GetHead().GetShield() + 4)

	push.PushAutoLog(h, push.GetHeroLogString(h)+"获得了4点攻击力,4点护盾")
}

// 巫师学徒
type Card233 struct {
	bcard.Card
}

func (c *Card233) NewPoint() iface.ICard {
	return &Card233{}
}

func (c *Card233) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetMona")
}

func (c *Card233) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetMona")
}

func (c *Card233) OnNROtherGetMona(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeHand ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeSorcery ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	return -1
}

// 虚灵奥术师
type Card234 struct {
	bcard.Card
}

func (c *Card234) NewPoint() iface.ICard {
	return &Card234{}
}

func (c *Card234) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card234) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card234) OnNRRoundEnd() {

	// 在我的回合结束时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()

	if len(h.GetSecrets()) > 0 {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddDamage(2)
		buff.AddHpMaxAndHp(2)

		c.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(c)+"获得+2/+2")
	}
}

// 奥金尼灵魂祭司
type Card235 struct {
	bcard.Card
}

func (c *Card235) NewPoint() iface.ICard {
	return &Card235{}
}

func (c *Card235) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherChangeTreatToCost")
}

func (c *Card235) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherChangeTreatToCost")
}

func (c *Card235) OnNROtherChangeTreatToCost(who iface.ICard) bool {

	change := who.GetOwner().GetId() == c.GetOwner().GetId()
	if change {
		push.PushAutoLog(who.GetOwner(), push.GetCardLogString(c)+"将治疗转成了伤害")
	}
	return change
}

// 控心术
type Card236 struct {
	bcard.Card
}

func (c *Card236) NewPoint() iface.ICard {
	return &Card236{}
}

func (c *Card236) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	types := []define.CardType{define.CardTypeEntourage}
	scs := iface.GetCardFact().ScreenCards(h.GetEnemy().GetLibCards(), iface.NewScreenCardParam(
		iface.SCPWithCardTypes(types),
	))

	randC := h.GetEnemy().RandCard(scs)

	if randC == nil {
		push.PushAutoLog(h, push.GetHeroLogString(h.GetEnemy())+"没有随从了")
		return
	}

	nc, err := randC.Copy()
	if err != nil {
		return
	}

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc.SetOwner(c.GetOwner())
	h.MoveToBattle(nc, -1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 群体驱散
type Card237 struct {
	bcard.Card
}

func (c *Card237) NewPoint() iface.ICard {
	return &Card237{}
}

func (c *Card237) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for _, v := range h.CardsToNewInstance(h.GetEnemy().GetBattleCards()) {
		v.Silent()
		push.PushAutoLog(h, push.GetCardLogString(c)+"沉默了"+push.GetCardLogString(v))
	}

	h.DrawByTimes(1)

}

// 伪装大师
type Card238 struct {
	bcard.Card
}

func (c *Card238) NewPoint() iface.ICard {
	return &Card238{}
}

func (c *Card238) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddTraits(define.CardTraitsSneak)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得潜行")
}

// 王牌猎人
type Card239 struct {
	bcard.Card
}

func (c *Card239) NewPoint() iface.ICard {
	return &Card239{}
}

func (c *Card239) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	if rc.GetHaveEffectDamage() < 7 {
		return
	}

	rc.GetOwner().DieCard(rc, false)
	push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(rc))
}

// 深渊领主
type Card240 struct {
	bcard.Card
}

func (c *Card240) NewPoint() iface.ICard {
	return &Card240{}
}

func (c *Card240) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.GetHead().CostHp(c, 5)
}

// 召唤传送门
type Card241 struct {
	bcard.Card
}

func (c *Card241) NewPoint() iface.ICard {
	return &Card241{}
}

func (c *Card241) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetMona")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetFinalMona")
}

func (c *Card241) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetMona")
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetFinalMona")
}

func (c *Card241) OnNROtherGetMona(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeHand ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	return -2
}

func (c *Card241) OnNROtherGetFinalMona(oc iface.ICard, m int) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeHand ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return m
	}

	if m <= 0 {
		return 1
	}

	return m
}

// 暗影烈焰
type Card242 struct {
	bcard.Card
}

func (c *Card242) NewPoint() iface.ICard {
	return &Card242{}
}

func (c *Card242) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}
	dmg := rc.GetHaveEffectDamage()
	dmg += h.GetApDamage()

	h.DieCard(rc, false)

	for _, v := range h.CardsToNewInstance(h.GetEnemy().GetBattleCards()) {
		v.CostHp(c, dmg)
	}
}

// 阿拉希武器匠
type Card243 struct {
	bcard.Card
}

func (c *Card243) NewPoint() iface.ICard {
	return &Card243{}
}

func (c *Card243) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	nc := iface.GetCardFact().GetCard(244)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.OnlyReleaseWeapon(nc)
}

// 战斧
type Card244 struct {
	bcard.Card
}

func (c *Card244) NewPoint() iface.ICard {
	return &Card244{}
}

// 致死打击
type Card245 struct {
	bcard.Card
}

func (c *Card245) NewPoint() iface.ICard {
	return &Card245{}
}

func (c *Card245) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()
	if h.GetHead().GetHaveEffectHp() <= 12 {
		dmg += 2
	}

	rc.CostHp(c, dmg)
}

// 年迈的酒仙
type Card246 struct {
	bcard.Card
}

func (c *Card246) NewPoint() iface.ICard {
	return &Card246{}
}

func (c *Card246) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}
	h := c.GetOwner()
	if rc.GetOwner().GetId() == h.GetId() {
		h.MoveToHand(rc)
		push.PushAutoLog(h, push.GetCardLogString(c)+"将"+push.GetCardLogString(rc)+"移动回手牌")
	}
}

// 黑铁矮人
type Card247 struct {
	bcard.Card
}

func (c *Card247) NewPoint() iface.ICard {
	return &Card247{}
}

func (c *Card247) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || rc.GetType() != define.CardTypeEntourage {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	buff.AddDamage(2)

	rc.AddSubCards(buff)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了两点攻击力")
}

// 魔古山守望者
type Card248 struct {
	bcard.Card
}

func (c *Card248) NewPoint() iface.ICard {
	return &Card248{}
}

// 破法者
type Card249 struct {
	bcard.Card
}

func (c *Card249) NewPoint() iface.ICard {
	return &Card249{}
}

func (c *Card249) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	rc.Silent()
	push.PushAutoLog(h, push.GetCardLogString(c)+"沉默了"+push.GetCardLogString(rc))
}

// 阿古斯防御者
type Card250 struct {
	bcard.Card
}

func (c *Card250) NewPoint() iface.ICard {
	return &Card250{}
}

func (c *Card250) OnRelease2(choiceId, bidx int, rc iface.ICard) {

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

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddDamage(1)
		buff.AddHpMaxAndHp(1)
		buff.AddTraits(define.CardTraitsTaunt)

		v.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得+1/+1和嘲讽")
	}
}

// 诅咒教派领袖
type Card251 struct {
	bcard.Card
}

func (c *Card251) NewPoint() iface.ICard {
	return &Card251{}
}

func (c *Card251) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card251) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherDie")
}

func (c *Card251) OnNROtherDie(tc iface.ICard) {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		tc.GetOwner().GetId() != c.GetOwner().GetId() ||
		tc.GetId() == c.GetId() ||
		tc.GetConfig().Ctype != define.CardTypeEntourage {
		return
	}

	push.PushAutoLog(c.GetOwner(), "由于"+push.GetCardLogString(tc)+"死亡,"+push.GetCardLogString(c)+"触发效果")
	h.DrawByTimes(1)
}

// 恐怖海盗
type Card252 struct {
	bcard.Card
}

func (c *Card252) NewPoint() iface.ICard {
	return &Card252{}
}

func (c *Card252) OnGetMona(m int) int {

	h := c.GetOwner()
	if c.GetCardInCardsPos() == define.InCardsTypeHand && h.GetWeapon() != nil {
		m = m - h.GetWeapon().GetHaveEffectDamage()
	}

	return m
}

// 年迈的法师
type Card253 struct {
	bcard.Card
}

func (c *Card253) NewPoint() iface.ICard {
	return &Card253{}
}

func (c *Card253) OnRelease2(choiceId, bidx int, rc iface.ICard) {

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

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddApDamage(1)
		v.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得法术伤害+1")
	}
}

// 紫罗兰教师
type Card254 struct {
	bcard.Card
}

func (c *Card254) NewPoint() iface.ICard {
	return &Card254{}
}

func (c *Card254) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherAfterRelease")
}

func (c *Card254) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherAfterRelease")
}

func (c *Card254) OnNROtherAfterRelease(oc iface.ICard) {

	h := c.GetOwner()
	if oc.GetOwner().GetId() != h.GetId() ||
		oc.GetId() == c.GetId() ||
		len(h.GetBattleCards()) >= define.MaxBattleNum ||
		oc.GetType() != define.CardTypeSorcery {
		return
	}

	bidx := h.GetIdxByCards(c, h.GetBattleCards())

	// 召唤
	nc := iface.GetCardFact().GetCard(255)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 紫罗兰学徒
type Card255 struct {
	bcard.Card
}

func (c *Card255) NewPoint() iface.ICard {
	return &Card255{}
}

// 暮光幼龙
type Card256 struct {
	bcard.Card
}

func (c *Card256) NewPoint() iface.ICard {
	return &Card256{}
}

func (c *Card256) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	addHp := len(h.GetHandCards())

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddHpMaxAndHp(addHp)

	c.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"获得了"+strconv.Itoa(addHp)+"生命值")

}

// 野蛮咆哮
type Card257 struct {
	bcard.Card
}

func (c *Card257) NewPoint() iface.ICard {
	return &Card257{}
}

func (c *Card257) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	bcs := h.GetBattleCards()
	bcs = append(bcs, h.GetHead())

	for _, v := range bcs {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddDamage(2)

		v.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得2点攻击")
	}
}

// 治疗之触
type Card258 struct {
	bcard.Card
}

func (c *Card258) NewPoint() iface.ICard {
	return &Card258{}
}

func (c *Card258) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	rc.TreatmentHp(c, 8)
}

// 野性成长
type Card259 struct {
	bcard.Card
}

func (c *Card259) NewPoint() iface.ICard {
	return &Card259{}
}

func (c *Card259) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if h.GetMonaMax() >= h.GetConfig().MonaMax {
		nc := iface.GetCardFact().GetCard(308)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
		h.MoveToHand(nc)

		push.PushAutoLog(h, "获得了"+push.GetCardLogString(nc))

		return
	}

	h.AddMonaMax(1)
	push.PushAutoLog(h, "获得一个空的法力水晶")
}

// 霍弗
type Card260 struct {
	bcard.Card
}

func (c *Card260) NewPoint() iface.ICard {
	return &Card260{}
}

// 雷欧克
type Card261 struct {
	bcard.Card
}

func (c *Card261) NewPoint() iface.ICard {
	return &Card261{}
}

func (c *Card261) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card261) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")
}

func (c *Card261) OnNROtherGetDamage(oc iface.ICard) int {

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

// 米莎
type Card262 struct {
	bcard.Card
}

func (c *Card262) NewPoint() iface.ICard {
	return &Card262{}
}

// 杀戮命令
type Card263 struct {
	bcard.Card
}

func (c *Card263) NewPoint() iface.ICard {
	return &Card263{}
}

func (c *Card263) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	for _, v := range h.CardsToNewInstance(h.GetBattleCards()) {
		if v.IsRace(define.CardRaceBeast) {
			dmg += 2
			break
		}
	}

	rc.CostHp(c, dmg)
}

// 动物伙伴
type Card264 struct {
	bcard.Card
}

func (c *Card264) NewPoint() iface.ICard {
	return &Card264{}
}

func (c *Card264) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	randIdx := h.GetBattle().GetRand().Intn(len(define.AnimalCompanionIds))

	nc := iface.GetCardFact().GetCard(define.AnimalCompanionIds[randIdx])
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToBattle(nc, -1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 奥术智慧
type Card265 struct {
	bcard.Card
}

func (c *Card265) NewPoint() iface.ICard {
	return &Card265{}
}

func (c *Card265) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.DrawByTimes(2)
}

// 冰霜新星
type Card266 struct {
	bcard.Card
}

func (c *Card266) NewPoint() iface.ICard {
	return &Card266{}
}

func (c *Card266) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	for _, v := range h.CardsToNewInstance(h.GetEnemy().GetBattleCards()) {

		v.AddTraits(define.CardTraitsFrozen)
		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得了冻结")
	}
}

// 刀扇
type Card267 struct {
	bcard.Card
}

func (c *Card267) NewPoint() iface.ICard {
	return &Card267{}
}

func (c *Card267) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	for _, v := range h.CardsToNewInstance(h.GetEnemy().GetBattleCards()) {
		v.CostHp(c, dmg)
	}
	h.DrawByTimes(1)
}

// 暗影箭
type Card268 struct {
	bcard.Card
}

func (c *Card268) NewPoint() iface.ICard {
	return &Card268{}
}

func (c *Card268) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	rc.CostHp(c, dmg)
}

// 吸取生命
type Card269 struct {
	bcard.Card
}

func (c *Card269) NewPoint() iface.ICard {
	return &Card269{}
}

func (c *Card269) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	rc.CostHp(c, dmg)
	h.GetHead().TreatmentHp(c, 2)
}

// 战歌指挥官
type Card270 struct {
	bcard.Card
}

func (c *Card270) NewPoint() iface.ICard {
	return &Card270{}
}

func (c *Card270) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRPutToBattle")
}

func (c *Card270) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRPutToBattle")
}

func (c *Card270) OnNRPutToBattle(oc iface.ICard) {
	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddTraits(define.CardTraitsSuddenStrike)

	oc.AddSubCards(buff)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(oc)+"获得了突袭")
}

// 冲锋
type Card271 struct {
	bcard.Card
}

func (c *Card271) NewPoint() iface.ICard {
	return &Card271{}
}

func (c *Card271) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(2)
	buff.AddTraits(define.CardTraitsAssault)

	rc.AddSubCards(buff)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"+2攻击并获得了冲锋")
}

// 盾牌格挡
type Card272 struct {
	bcard.Card
}

func (c *Card272) NewPoint() iface.ICard {
	return &Card272{}
}

func (c *Card272) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	h.DrawByTimes(1)
	h.GetHead().SetShield(h.GetHead().GetShield() + 5)
	push.PushAutoLog(h, push.GetHeroLogString(h)+"5点护盾")
}

// 战争储备箱
type Card273 struct {
	bcard.Card
}

func (c *Card273) NewPoint() iface.ICard {
	return &Card273{}
}

func (c *Card273) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	rd := []define.CardType{define.CardTypeEntourage, define.CardTypeSorcery, define.CardTypeWeapon}

	for _, v := range rd {

		ncCache := iface.GetCardFact().RandByAllCards(h.GetBattle().GetRand(), iface.NewScreenCardParam(
			iface.SCPWithCardVocations([]define.Vocation{define.VocationWarrior}), iface.SCPWithCardTypes([]define.CardType{v}),
		))

		if ncCache == nil {
			continue
		}

		nc := iface.GetCardFact().GetCard(ncCache.GetConfig().Id)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
		h.MoveToHand(nc)

		push.PushLog(h, "获得了"+push.GetCardLogString(nc))
	}
}

// 炽炎战斧
type Card274 struct {
	bcard.Card
}

func (c *Card274) NewPoint() iface.ICard {
	return &Card274{}
}

// 自然印记
type Card275 struct {
	bcard.Card
}

func (c *Card275) NewPoint() iface.ICard {
	return &Card275{}
}

func (c *Card275) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())

	if choiceId == 0 {
		push.PushAutoLog(c.GetOwner(), "[抉择1]+4攻击")
		buff.AddDamage(4)
	} else {
		push.PushAutoLog(c.GetOwner(), "[抉择2]+4生命值和嘲讽")
		buff.AddHpMaxAndHp(4)
		buff.AddTraits(define.CardTraitsTaunt)
	}

	rc.AddSubCards(buff)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得buff")
}

// 致命射击
type Card276 struct {
	bcard.Card
}

func (c *Card276) NewPoint() iface.ICard {
	return &Card276{}
}

func (c *Card276) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	he := h.GetEnemy()
	if len(he.GetBattleCards()) <= 0 {
		return
	}

	randC := he.RandCard(he.GetBattleCards())
	he.DieCard(randC, false)

	push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(randC))
}

// 关门放狗
type Card277 struct {
	bcard.Card
}

func (c *Card277) NewPoint() iface.ICard {
	return &Card277{}
}

func (c *Card277) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for i := 1; i <= len(h.GetEnemy().GetBattleCards()); i++ {
		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return
		}

		nc := iface.GetCardFact().GetCard(278)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

		h.MoveToBattle(nc, -1)
		push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
	}
}

// 猎犬
type Card278 struct {
	bcard.Card
}

func (c *Card278) NewPoint() iface.ICard {
	return &Card278{}
}

// 鹰角弓
type Card279 struct {
	bcard.Card
}

func (c *Card279) NewPoint() iface.ICard {
	return &Card279{}
}

func (c *Card279) OnWear() {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherSecretTigger")
}

func (c *Card279) OnNROtherSecretTigger(oc iface.ICard) {

	h := c.GetOwner()

	if c.GetCardInCardsPos() != define.InCardsTypeBody {
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherSecretTigger")
		return
	}

	if oc.GetOwner().GetId() != c.GetOwner().GetId() {
		return
	}

	c.AddHpMaxAndHp(1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"获得了一点耐久")
}

// 肯瑞托法师
type Card280 struct {
	bcard.Card
}

func (c *Card280) NewPoint() iface.ICard {
	return &Card280{}
}

func (c *Card280) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.GetBattle().AddCardToEvent(c, "OnNROtherGetMona")
	h.GetBattle().AddCardToEvent(c, "OnNROtherAfterRelease")
}

func (c *Card280) deleteEvent() {
	h := c.GetOwner()
	h.GetBattle().RemoveCardFromEvent(c, "OnNROtherGetMona")
	h.GetBattle().RemoveCardFromEvent(c, "OnNROtherAfterRelease")
}

func (c *Card280) OnNROtherAfterRelease(oc iface.ICard) {
	h := c.GetOwner()

	if c.GetReleaseRound() != h.GetBattle().GetIncrRoundId() {
		c.deleteEvent()
		return
	}

	if oc.GetType() != define.CardTypeSorcery || !oc.IsHaveTraits(define.CardTraitsSecret) || oc.GetOwner().GetId() != h.GetId() {
		return
	}

	c.deleteEvent()
}

func (c *Card280) OnNROtherGetMona(oc iface.ICard) int {

	h := c.GetOwner()

	if c.GetReleaseRound() != h.GetBattle().GetIncrRoundId() {
		c.deleteEvent()
		return 0
	}

	if oc.GetCardInCardsPos() != define.InCardsTypeHand ||
		oc.GetType() != define.CardTypeSorcery ||
		!oc.IsHaveTraits(define.CardTraitsSecret) ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	return -999
}

// 法术反制
type Card281 struct {
	bcard.Card
}

func (c *Card281) NewPoint() iface.ICard {
	return &Card281{}
}

func (c *Card281) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeRelease")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card281) OnNROtherBeforeRelease(oc, rc iface.ICard) (iface.ICard, bool) {
	h := c.GetOwner()
	if oc.GetOwner().GetId() == h.GetEnemy().GetId() &&
		oc.GetType() == define.CardTypeSorcery && !h.IsRoundHero() {

		h.DeleteSecret(c, true)

		push.PushAutoLog(h, c.GetConfig().Name+"(奥秘)让"+oc.GetConfig().Name+"变得无效")

		return nil, false
	}

	return rc, true
}

// 寒冰屏障
type Card282 struct {
	bcard.Card
}

func (c *Card282) NewPoint() iface.ICard {
	return &Card282{}
}

func (c *Card282) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeCostHpDie")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card282) OnNROtherBeforeCostHpDie(oc iface.ICard) {

	h := c.GetOwner()
	if !h.IsRoundHero() && oc.GetId() == c.GetOwner().GetId() {

		h.DeleteSecret(c, true)

		buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundBeginClear)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddTraits(define.CardTraitsImmune)
		oc.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(oc)+"获得免疫")
	}
}

// 镜像实体
type Card283 struct {
	bcard.Card
}

func (c *Card283) NewPoint() iface.ICard {
	return &Card283{}
}

func (c *Card283) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherAfterRelease")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card283) OnNROtherAfterRelease(oc iface.ICard) {
	h := c.GetOwner()
	if oc.GetOwner().GetId() == h.GetEnemy().GetId() &&
		oc.GetType() == define.CardTypeEntourage && !h.IsRoundHero() {

		nc, err := oc.Copy()
		if err != nil {
			return
		}

		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return
		}

		nc.SetOwner(c.GetOwner())
		h.MoveToBattle(nc, -1)

		h.DeleteSecret(c, true)

		push.PushAutoLog(h, c.GetConfig().Name+"(奥秘)召唤了"+push.GetCardLogString(nc))
	}
}

// 蒸发
type Card284 struct {
	bcard.Card
}

func (c *Card284) NewPoint() iface.ICard {
	return &Card284{}
}

func (c *Card284) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeAttack")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card284) OnNROtherBeforeAttack(oc, rc iface.ICard) iface.ICard {

	h := c.GetOwner()
	if oc.GetOwner().GetId() == h.GetEnemy().GetId() &&
		oc.GetType() == define.CardTypeEntourage && !h.IsRoundHero() {

		rc = nil

		h.DeleteSecret(c, true)
		oc.GetOwner().DieCard(oc, false)

		push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(oc))
	}

	return rc
}

// 游学者周卓
type Card285 struct {
	bcard.Card
}

func (c *Card285) NewPoint() iface.ICard {
	return &Card285{}
}

func (c *Card285) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeRelease")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card285) OnNROtherBeforeRelease(oc, rc iface.ICard) (iface.ICard, bool) {

	h := c.GetOwner()
	if oc.GetOwner().GetId() == h.GetEnemy().GetId() &&
		oc.GetType() == define.CardTypeSorcery && !h.IsRoundHero() && rc != nil {

		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return rc, true
		}

		h.DeleteSecret(c, true)

		nc := iface.GetCardFact().GetCard(286)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
		h.MoveToBattle(nc, -1)

		push.PushAutoLog(h, "召唤了"+push.GetCardLogString(nc))
		push.PushAutoLog(h, push.GetCardLogString(oc)+"目标变成了"+push.GetCardLogString(nc))

		rc = nc
	}

	return rc, true
}

// 扰咒师
type Card286 struct {
	bcard.Card
}

func (c *Card286) NewPoint() iface.ICard {
	return &Card286{}
}

// 寒冰护体
type Card287 struct {
	bcard.Card
}

func (c *Card287) NewPoint() iface.ICard {
	return &Card287{}
}

func (c *Card287) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.OnlyReleaseSecret(c) {
		c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeAttack")
		push.PushLog(h, "释放了"+c.GetConfig().Name+"(奥秘)")
	}
}

func (c *Card287) OnNROtherBeforeAttack(oc, rc iface.ICard) iface.ICard {

	h := c.GetOwner()
	if oc.GetOwner().GetId() == h.GetEnemy().GetId() && !h.IsRoundHero() {

		h.DeleteSecret(c, true)

		h.GetHead().SetShield(h.GetHead().GetShield() + 8)

		push.PushAutoLog(h, "获得了8点护甲值")
	}

	return rc
}

// 冰锥术
type Card288 struct {
	bcard.Card
}

func (c *Card288) NewPoint() iface.ICard {
	return &Card288{}
}

func (c *Card288) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	rh := rc.GetOwner()
	bcs := rh.GetBattleCards()

	rcIdx := rh.GetIdxByCards(rc, bcs)

	var cLeft iface.ICard
	var cRight iface.ICard

	if (rcIdx - 1) >= 0 {
		cLeft = bcs[rcIdx-1]
	}
	if (rcIdx + 1) < len(bcs) {
		cRight = bcs[rcIdx+1]
	}

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	rc.CostHp(c, dmg)
	rc.AddTraits(define.CardTraitsFrozen)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得了冻结")

	if cLeft != nil {
		cLeft.CostHp(c, dmg)
		cLeft.AddTraits(define.CardTraitsFrozen)
		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(cLeft)+"获得了冻结")
	}

	if cRight != nil {
		cRight.CostHp(c, dmg)
		cRight.AddTraits(define.CardTraitsFrozen)
		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(cRight)+"获得了冻结")
	}
}

// 奥尔多卫士
type Card289 struct {
	bcard.Card
}

func (c *Card289) NewPoint() iface.ICard {
	return &Card289{}
}

func (c *Card289) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	rc.SetDamage(1)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"攻击力变为1")
}

// 神恩术
type Card290 struct {
	bcard.Card
}

func (c *Card290) NewPoint() iface.ICard {
	return &Card290{}
}

func (c *Card290) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	times := len(h.GetEnemy().GetHandCards()) - len(h.GetHandCards())

	if times <= 0 {
		return
	}

	h.DrawByTimes(times)
}

// 公正之剑
type Card291 struct {
	bcard.Card
}

func (c *Card291) NewPoint() iface.ICard {
	return &Card291{}
}

func (c *Card291) OnWear() {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRPutToBattle")
}

func (c *Card291) OnNRPutToBattle(oc iface.ICard) {
	h := c.GetOwner()

	if c.GetCardInCardsPos() != define.InCardsTypeBody {
		c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRPutToBattle")
		return
	}

	if h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() ||
		oc.GetType() != define.CardTypeEntourage {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)
	buff.AddHpMaxAndHp(1)

	oc.AddSubCards(buff)
	c.CostHp(c, 1)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(oc)+"获得+1/+1")
}

// 光耀之子
type Card292 struct {
	bcard.Card
}

func (c *Card292) NewPoint() iface.ICard {
	return &Card292{}
}

func (c *Card292) OnGetDamage(d int) int {
	return c.GetHaveEffectHp()
}

// 暗影狂乱
type Card293 struct {
	bcard.Card
}

func (c *Card293) NewPoint() iface.ICard {
	return &Card293{}
}

func (c *Card293) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	if rc.GetOwner().GetId() == h.GetId() || rc.GetHaveEffectDamage() > 3 || len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}
	eh := rc.GetOwner()

	buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddOnEventClear(func(i iface.ICard, s string) {
		if i.GetId() == buff.GetId() &&
			s == "OnNRRoundEnd" &&
			rc.GetCardInCardsPos() == define.InCardsTypeBattle &&
			rc.GetOwner().GetId() == h.GetId() {

			if len(eh.GetBattleCards()) >= define.MaxBattleNum {

				push.PushAutoLog(h, "还给对方时,战场已满"+push.GetCardLogString(i)+"死亡")
				h.DieCard(i, false)
				return
			}

			push.PushAutoLog(h, "还给了对方"+push.GetCardLogString(rc))

			h.MoveOutBattleOnlyBattleCards(rc)
			rc.SetOwner(eh)

			eh.MoveToBattle(rc, -1)
		}
	})
	rc.AddSubCards(buff)

	rc.GetOwner().MoveOutBattleOnlyBattleCards(rc)
	rc.SetOwner(h)
	rc.SetAttackTimes(0)
	h.MoveToBattle(rc, -1)
}

// 艾德温·范克里夫
type Card294 struct {
	bcard.Card
}

func (c *Card294) NewPoint() iface.ICard {
	return &Card294{}
}

func (c *Card294) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	times := (h.GetReleaseCardTimes() - 1)

	if times <= 1 {
		return
	}

	add := (times - 1) * 2

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(add)
	buff.AddHpMaxAndHp(add)

	c.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"获得+"+strconv.Itoa(add)+"/+"+strconv.Itoa(add))
}

// 军情七处特工
type Card295 struct {
	bcard.Card
}

func (c *Card295) NewPoint() iface.ICard {
	return &Card295{}
}

func (c *Card295) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	if h.GetReleaseCardTimes() <= 1 {
		return
	}

	rc.CostHp(c, 2)
}

// 裂颅之击
type Card296 struct {
	bcard.Card
}

func (c *Card296) NewPoint() iface.ICard {
	return &Card296{}
}

func (c *Card296) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := c.GetConfig().Damage
	dmg += h.GetApDamage()

	h.GetEnemy().GetHead().CostHp(c, 2)

	if h.GetReleaseCardTimes() <= 1 {
		return
	}

	push.PushAutoLog(h, push.GetCardLogString(c)+"触发连击")
	h.GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card296) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeGrave {
		return
	}

	h := c.GetOwner()
	h.GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")

	nc := iface.GetCardFact().GetCard(296)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToHand(nc)

	push.PushAutoLog(h, push.GetCardLogString(c)+"移动回手牌")
}

// 法力之潮图腾
type Card297 struct {
	bcard.Card
}

func (c *Card297) NewPoint() iface.ICard {
	return &Card297{}
}

func (c *Card297) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card297) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card297) OnNRRoundEnd() {

	// 在我的回合结束时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	h.DrawByTimes(1)
}

// 无羁元素
type Card298 struct {
	bcard.Card
}

func (c *Card298) NewPoint() iface.ICard {
	return &Card298{}
}

func (c *Card298) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherAfterRelease")
}

func (c *Card298) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherAfterRelease")
}

func (c *Card298) OnNROtherAfterRelease(oc iface.ICard) {

	h := c.GetOwner()
	if oc.GetOwner().GetId() != h.GetId() || oc.GetId() == c.GetId() || !oc.IsHaveTraits(define.CardTraitsLockMona) {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)
	buff.AddHpMaxAndHp(1)
	c.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"获得+1/+1")
}

// 野性狼魂
type Card299 struct {
	bcard.Card
}

func (c *Card299) NewPoint() iface.ICard {
	return &Card299{}
}

func (c *Card299) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for i := 1; i <= 2; i++ {
		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return
		}
		nc := iface.GetCardFact().GetCard(300)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
		h.MoveToBattle(nc, -1)

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
	}
}

// 幽灵狼
type Card300 struct {
	bcard.Card
}

func (c *Card300) NewPoint() iface.ICard {
	return &Card300{}
}
