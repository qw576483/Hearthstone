package cards

import (
	"hs/logic/battle/bcard"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/iface"
	"hs/logic/push"
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
