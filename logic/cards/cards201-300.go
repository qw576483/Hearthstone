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
	for _, v := range rc.GetOwner().GetBattleCards() {
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

func (c *Card204) OnAfterAttack(ec iface.ICard) {
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

	for _, v := range h.GetEnemy().GetBattleCards() {
		v.CostHp(c, dmg)
	}

	for _, v := range h.GetBattleCards() {
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

	for _, v := range h.GetBattleCards() {
		v.CostHp(c, dmg)
	}
	for _, v := range h.GetEnemy().GetBattleCards() {
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

	h := c.GetOwner()

	for _, v := range h.GetEnemy().GetBattleCards() {
		if v.IsRace(define.CardRaceFish) {
			dmg += 1
		}
	}

	for _, v := range h.GetBattleCards() {
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
		push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+c.GetConfig().Name+"(奥秘)")
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
	for _, v := range h.GetBattleCards() {

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

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得亡语：召唤一个2/2的树人")
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
	return who.GetOwner().GetId() == c.GetOwner().GetId()
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

	for _, v := range h.GetEnemy().GetBattleCards() {
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

	for _, v := range h.GetEnemy().GetBattleCards() {
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
	if rc == nil {
		return
	}

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
