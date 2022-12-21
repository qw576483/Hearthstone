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

// 银背族长
type Card310 struct {
	bcard.Card
}

func (c *Card310) NewPoint() iface.ICard {
	return &Card310{}
}

// 达拉然法师
type Card311 struct {
	bcard.Card
}

func (c *Card311) NewPoint() iface.ICard {
	return &Card311{}
}

// 剃刀猎手
type Card312 struct {
	bcard.Card
}

func (c *Card312) NewPoint() iface.ICard {
	return &Card312{}
}

func (c *Card312) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(313)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 野猪
type Card313 struct {
	bcard.Card
}

func (c *Card313) NewPoint() iface.ICard {
	return &Card313{}
}

// 狼骑兵
type Card314 struct {
	bcard.Card
}

func (c *Card314) NewPoint() iface.ICard {
	return &Card314{}
}

// 铁炉堡火枪手
type Card315 struct {
	bcard.Card
}

func (c *Card315) NewPoint() iface.ICard {
	return &Card315{}
}

func (c *Card315) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	rc.CostHp(c, 1)
}

// 破碎残阳祭司
type Card316 struct {
	bcard.Card
}

func (c *Card316) NewPoint() iface.ICard {
	return &Card316{}
}

func (c *Card316) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)
	buff.AddHpMaxAndHp(1)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得+1/+1")
}

// 铁鬃灰熊
type Card317 struct {
	bcard.Card
}

func (c *Card317) NewPoint() iface.ICard {
	return &Card317{}
}

// 团队领袖
type Card318 struct {
	bcard.Card
}

func (c *Card318) NewPoint() iface.ICard {
	return &Card318{}
}

func (c *Card318) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card318) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")
}

func (c *Card318) OnNROtherGetDamage(oc iface.ICard) int {

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

// 岩浆暴怒者
type Card319 struct {
	bcard.Card
}

func (c *Card319) NewPoint() iface.ICard {
	return &Card319{}
}

// 光明之翼
type Card320 struct {
	bcard.Card
}

func (c *Card320) NewPoint() iface.ICard {
	return &Card320{}
}

func (c *Card320) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	ncCache := iface.GetCardFact().RandByAllCards(h.GetBattle().GetRand(), iface.NewScreenCardParam(
		iface.SCPWithCardTypes([]define.CardType{define.CardTypeEntourage}),
		iface.SCPWithCardQuality([]define.CardQuality{define.CardQualityOrange}),
	))

	if ncCache == nil {
		return
	}

	nc := iface.GetCardFact().GetCard(ncCache.GetConfig().Id)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	push.PushLog(c.GetOwner(), push.GetCardLogString(c)+"让你获得了"+push.GetCardLogString(nc))

	h.MoveToHand(nc)
}

// 腐肉食尸鬼
type Card321 struct {
	bcard.Card
}

func (c *Card321) NewPoint() iface.ICard {
	return &Card321{}
}

func (c *Card321) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherDie")
}

func (c *Card321) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherDie")
}

func (c *Card321) OnNROtherDie(tc iface.ICard) {

	h := c.GetOwner()
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		tc.GetId() == c.GetId() ||
		tc.GetConfig().Ctype != define.CardTypeEntourage {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)

	c.AddSubCards(buff)

	push.PushAutoLog(c.GetOwner(), "由于"+push.GetCardLogString(tc)+"死亡,"+push.GetCardLogString(c)+"获得+1攻击")
}

// 寒光先知
type Card322 struct {
	bcard.Card
}

func (c *Card322) NewPoint() iface.ICard {
	return &Card322{}
}

func (c *Card322) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for _, v := range h.GetBattleCards() {
		if !v.IsRace(define.CardRaceFish) {
			continue
		}

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddHpMaxAndHp(2)

		v.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得+2生命值。")
	}
}

// 奥术傀儡
type Card323 struct {
	bcard.Card
}

func (c *Card323) NewPoint() iface.ICard {
	return &Card323{}
}

func (c *Card323) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	h.GetEnemy().AddMonaMax(1)
	h.GetEnemy().AddMona(1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(h.GetEnemy().GetHead())+"获得1个法力水晶")
}

// 血色十字军战士
type Card324 struct {
	bcard.Card
}

func (c *Card324) NewPoint() iface.ICard {
	return &Card324{}
}

// 南海船长
type Card325 struct {
	bcard.Card
}

func (c *Card325) NewPoint() iface.ICard {
	return &Card325{}
}

func (c *Card325) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetHp")
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card325) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetHp")
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")
}

func (c *Card325) OnNROtherGetDamage(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBattle ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() ||
		!oc.IsRace(define.CardRacePirate) {
		return 0
	}

	return 1
}

func (c *Card325) OnNROtherGetHp(oc iface.ICard) int {
	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBattle ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() ||
		!oc.IsRace(define.CardRacePirate) {
		return 0
	}
	return 1
}

// 精神控制技师
type Card326 struct {
	bcard.Card
}

func (c *Card326) NewPoint() iface.ICard {
	return &Card326{}
}

func (c *Card326) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	ebcs := h.GetEnemy().GetBattleCards()

	if len(ebcs) < 4 {
		return
	}

	rc = ebcs[h.GetBattle().GetRand().Intn(len(ebcs))]

	h.CaptureCard(rc, bidx+1)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"夺取了"+push.GetCardLogString(rc))
}

// 血骑士
type Card327 struct {
	bcard.Card
}

func (c *Card327) NewPoint() iface.ICard {
	return &Card327{}
}

func (c *Card327) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	bcs := h.GetBattleCards()
	bcs = append(bcs, h.GetEnemy().GetBattleCards()...)

	for _, v := range bcs {
		if !v.IsHaveTraits(define.CardTraitsHolyShield) {
			continue
		}

		v.RemoveTraits(define.CardTraitsHolyShield)

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddDamage(3)
		buff.AddHpMaxAndHp(3)

		c.AddSubCards(buff)

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"吃掉了"+push.GetCardLogString(v)+"的圣盾，并获得+3/+3")
	}
}

// 任务达人
type Card328 struct {
	bcard.Card
}

func (c *Card328) NewPoint() iface.ICard {
	return &Card328{}
}

func (c *Card328) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherAfterRelease")
}

func (c *Card328) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherAfterRelease")
}

func (c *Card328) OnNROtherAfterRelease(oc iface.ICard) {

	h := c.GetOwner()
	if oc.GetOwner().GetId() != h.GetId() || oc.GetId() == c.GetId() {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(1)
	buff.AddHpMaxAndHp(1)

	c.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"获得+1/+1")
}

// 丛林猎豹
type Card329 struct {
	bcard.Card
}

func (c *Card329) NewPoint() iface.ICard {
	return &Card329{}
}

// 小鬼召唤师
type Card330 struct {
	bcard.Card
}

func (c *Card330) NewPoint() iface.ICard {
	return &Card330{}
}

func (c *Card330) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card330) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card330) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()

	bidx := h.GetIdxByCards(c, h.GetBattleCards())
	c.CostHp(c, 1)

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(331)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 小鬼
type Card331 struct {
	bcard.Card
}

func (c *Card331) NewPoint() iface.ICard {
	return &Card331{}
}

// 鱼人领军
type Card332 struct {
	bcard.Card
}

func (c *Card332) NewPoint() iface.ICard {
	return &Card332{}
}

func (c *Card332) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
}

func (c *Card332) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")
}

func (c *Card332) OnNROtherGetDamage(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBattle ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() ||
		!oc.IsRace(define.CardRaceFish) {
		return 0
	}

	return 2
}

// 帝王眼镜蛇
type Card333 struct {
	bcard.Card
}

func (c *Card333) NewPoint() iface.ICard {
	return &Card333{}
}

// 苦痛侍僧
type Card334 struct {
	bcard.Card
}

func (c *Card334) NewPoint() iface.ICard {
	return &Card334{}
}

func (c *Card334) OnAfterCostHp() {
	c.GetOwner().DrawByTimes(1)
}

// 负伤剑圣
type Card335 struct {
	bcard.Card
}

func (c *Card335) NewPoint() iface.ICard {
	return &Card335{}
}

func (c *Card335) OnRelease2(choiceId, bidx int, rc iface.ICard) {
	c.CostHp(c, 4)
}

// 大地之环先知
type Card336 struct {
	bcard.Card
}

func (c *Card336) NewPoint() iface.ICard {
	return &Card336{}
}

func (c *Card336) OnRelease(choiceId, bidx int, rc iface.ICard) {
	if rc == nil {
		return
	}

	rc.TreatmentHp(c, 3)
}

// 报警机器人
type Card337 struct {
	bcard.Card
}

func (c *Card337) NewPoint() iface.ICard {
	return &Card337{}
}

func (c *Card337) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundBegin")
}

func (c *Card337) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundBegin")
}

func (c *Card337) OnNRRoundBegin() {

	// 在我的回合开始时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	idx := h.GetIdxByCards(c, h.GetBattleCards())

	handE := []iface.ICard{}
	for _, v := range h.GetHandCards() {
		if v.GetType() == define.CardTypeEntourage {
			handE = append(handE, v)
		}
	}

	if len(handE) <= 0 {
		return
	}

	h.MoveToHand(c)

	randC := h.RandCard(handE)
	h.MoveToBattle(randC, idx)

	push.PushAutoLog(h, push.GetCardLogString(c)+"交换了"+push.GetCardLogString(randC))
}

// 穆克拉
type Card338 struct {
	bcard.Card
}

func (c *Card338) NewPoint() iface.ICard {
	return &Card338{}
}

func (c *Card338) OnRelease(choiceId, bidx int, rc iface.ICard) {
	he := c.GetOwner().GetEnemy()

	for i := 1; i <= 2; i++ {
		nc := iface.GetCardFact().GetCard(339)
		nc.Init(nc, define.InCardsTypeNone, he, he.GetBattle())
		he.MoveToHand(nc)
	}
}

// 香蕉
type Card339 struct {
	bcard.Card
}

func (c *Card339) NewPoint() iface.ICard {
	return &Card339{}
}

func (c *Card339) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, rc.GetOwner(), rc.GetOwner().GetBattle())
	buff.AddDamage(1)
	buff.AddHpMaxAndHp(1)

	rc.AddSubCards(buff)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得+1/+1")
}
