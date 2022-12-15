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

// 星火术
type Card114 struct {
	bcard.Card
}

func (c *Card114) NewPoint() iface.ICard {
	return &Card114{}
}

func (c *Card114) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	dmg := 5
	dmg += h.GetApDamage()

	if rc != nil {
		rc.CostHp(c, dmg)
	}

	push.PushAutoLog(h, "抽了一张牌")
	h.DrawByTimes(1)
}

// 消失
type Card115 struct {
	bcard.Card
}

func (c *Card115) NewPoint() iface.ICard {
	return &Card115{}
}

func (c *Card115) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for _, v := range h.GetBattleCards() {
		h.MoveToHand(v)
	}

	for _, v := range h.GetEnemy().GetBattleCards() {
		h.GetEnemy().MoveToHand(v)
	}
}

// 火元素
type Card116 struct {
	bcard.Card
}

func (c *Card116) NewPoint() iface.ICard {
	return &Card116{}
}

func (c *Card116) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	rc.CostHp(c, 4)
}

// 恐惧地狱火
type Card117 struct {
	bcard.Card
}

func (c *Card117) NewPoint() iface.ICard {
	return &Card117{}
}

func (c *Card117) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for _, v := range h.GetBattleCards() {
		v.CostHp(c, 1)
	}
	for _, v := range h.GetEnemy().GetBattleCards() {
		v.CostHp(c, 1)
	}
	h.GetHead().CostHp(c, 1)
	h.GetEnemy().GetHead().CostHp(c, 1)
}

// 竞技场主宰
type Card118 struct {
	bcard.Card
}

func (c *Card118) NewPoint() iface.ICard {
	return &Card118{}
}

// 鲁莽火箭兵
type Card119 struct {
	bcard.Card
}

func (c *Card119) NewPoint() iface.ICard {
	return &Card119{}
}

// 大法师
type Card120 struct {
	bcard.Card
}

func (c *Card120) NewPoint() iface.ICard {
	return &Card120{}
}

// 格尔宾·梅卡托克
type Card121 struct {
	bcard.Card
}

func (c *Card121) NewPoint() iface.ICard {
	return &Card121{}
}

func (c *Card121) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	randIdx := h.GetBattle().GetRand().Intn(len(define.GelbinMekkatorqueInventionIds))

	nc := iface.GetCardFact().GetCard(define.GelbinMekkatorqueInventionIds[randIdx])
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 壮胆机器人3000型
type Card122 struct {
	bcard.Card
}

func (c *Card122) NewPoint() iface.ICard {
	return &Card122{}
}

func (c *Card122) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card122) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card122) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	rc := h.RandBothBattleCard()
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

// 变鸡器
type Card123 struct {
	bcard.Card
}

func (c *Card123) NewPoint() iface.ICard {
	return &Card123{}
}

func (c *Card123) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundBegin")
}

func (c *Card123) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundBegin")
}

func (c *Card123) OnNRRoundBegin() {

	// 在我的回合开始时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	rc := h.RandBothBattleCard()
	if rc == nil {
		return
	}

	rch := rc.GetOwner()
	rcbidx := h.GetIdxByCards(rc, rch.GetBattleCards())
	rch.MoveOutBattleOnlyBattleCards(rc)

	nc := iface.GetCardFact().GetCard(124)
	nc.Init(nc, define.InCardsTypeNone, rch, rch.GetBattle())
	rch.MoveToBattle(nc, rcbidx)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"变成了"+push.GetCardLogString(nc))
}

// 小鸡
type Card124 struct {
	bcard.Card
}

func (c *Card124) NewPoint() iface.ICard {
	return &Card124{}
}

// 导航小鸡
type Card125 struct {
	bcard.Card
}

func (c *Card125) NewPoint() iface.ICard {
	return &Card125{}
}

func (c *Card125) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundBegin")
}

func (c *Card125) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundBegin")
}

func (c *Card125) OnNRRoundBegin() {

	// 在我的回合开始时
	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	h.DieCard(c, false)
	h.DrawByTimes(3)

	push.PushAutoLog(h, "抽了三张牌")
}

// 修理机器人
type Card126 struct {
	bcard.Card
}

func (c *Card126) NewPoint() iface.ICard {
	return &Card126{}
}

func (c *Card126) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card126) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card126) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	rc := h.RandBothInjuredBattleCardOrHero()
	if rc == nil {
		return
	}

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"恢复了六点生命值")
	rc.TreatmentHp(c, 6)
}

// 石拳食人魔
type Card127 struct {
	bcard.Card
}

func (c *Card127) NewPoint() iface.ICard {
	return &Card127{}
}

// 大检察官怀特迈恩
type Card128 struct {
	bcard.Card
}

func (c *Card128) NewPoint() iface.ICard {
	return &Card128{}
}

func (c *Card128) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	rdcs := h.GetRoundDieCards()

	i := 1
	for _, v := range rdcs {
		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return
		}

		nc := iface.GetCardFact().GetCard(v.GetConfig().Id)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
		h.MoveToBattle(nc, bidx+i)

		push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))

		i += 1
	}
}

// 长鬃草原狮
type Card129 struct {
	bcard.Card
}

func (c *Card129) NewPoint() iface.ICard {
	return &Card129{}
}

func (c *Card129) OnDie() {

	for i := 1; i <= 2; i++ {
		if len(c.GetOwner().GetBattleCards()) >= define.MaxBattleNum {
			return
		}
		dbidx := c.GetAfterDieBidx()

		nc := iface.GetCardFact().GetCard(130)
		nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
		nc.GetOwner().MoveToBattle(nc, dbidx)

		// logs
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"死亡时，召唤了"+push.GetCardLogString(nc))
	}
}

// 土狼
type Card130 struct {
	bcard.Card
}

func (c *Card130) NewPoint() iface.ICard {
	return &Card130{}
}

// 暴风雪
type Card131 struct {
	bcard.Card
}

func (c *Card131) NewPoint() iface.ICard {
	return &Card131{}
}

func (c *Card131) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := 2
	dmg += h.GetApDamage()

	for _, v := range h.GetEnemy().GetBattleCards() {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
		buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
		buff.AddTraits(define.CardTraitsFrozen)

		v.CostHp(c, dmg)
		v.AddSubCards(buff)

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得了冻结")
	}
}

// 复仇之怒
type Card132 struct {
	bcard.Card
}

func (c *Card132) NewPoint() iface.ICard {
	return &Card132{}
}

func (c *Card132) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := 8
	dmg += h.GetApDamage()

	for i := 1; i <= dmg; i++ {
		v := h.GetEnemy().RandBattleCardOrHero()
		v.CostHp(c, 1)
	}
}

// 秘教暗影祭司
type Card133 struct {
	bcard.Card
}

func (c *Card133) NewPoint() iface.ICard {
	return &Card133{}
}

func (c *Card133) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil || rc.GetOwner().GetId() == c.GetOwner().GetId() {
		return
	}

	if rc.GetHaveEffectDamage() > 2 {
		return
	}

	h.CaptureCard(rc, bidx+1)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"夺取了"+push.GetCardLogString(rc))
}

// 神圣之火
type Card134 struct {
	bcard.Card
}

func (c *Card134) NewPoint() iface.ICard {
	return &Card134{}
}

func (c *Card134) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	dmg := 5
	dmg += h.GetApDamage()

	if rc != nil {
		rc.CostHp(c, dmg)
	}

	h.GetHead().TreatmentHp(c, 5)
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(h.GetHead())+"恢复5点生命")
}

// 劫持者
type Card135 struct {
	bcard.Card
}

func (c *Card135) NewPoint() iface.ICard {
	return &Card135{}
}

func (c *Card135) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	if h.GetReleaseCardTimes() <= 1 {
		return
	}

	rc.GetOwner().MoveToHand(rc)
	push.PushAutoLog(h, push.GetCardLogString(c)+"触发了连击,将"+push.GetCardLogString(rc)+"移动回手")
}

// 银色指挥官
type Card136 struct {
	bcard.Card
}

func (c *Card136) NewPoint() iface.ICard {
	return &Card136{}
}

// 凯恩·血蹄
type Card137 struct {
	bcard.Card
}

func (c *Card137) NewPoint() iface.ICard {
	return &Card137{}
}

func (c *Card137) OnDie() {

	if len(c.GetOwner().GetBattleCards()) >= define.MaxBattleNum {
		return
	}
	dbidx := c.GetAfterDieBidx()

	nc := iface.GetCardFact().GetCard(138)
	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	nc.GetOwner().MoveToBattle(nc, dbidx)

	// logs
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"死亡时，召唤了"+push.GetCardLogString(nc))
}

// 贝恩·血蹄
type Card138 struct {
	bcard.Card
}

func (c *Card138) NewPoint() iface.ICard {
	return &Card138{}
}

// 艾露恩的女祭司
type Card139 struct {
	bcard.Card
}

func (c *Card139) NewPoint() iface.ICard {
	return &Card139{}
}

func (c *Card139) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	h.GetHead().TreatmentHp(c, 4)
	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(h.GetHead())+"恢复4点生命")
}

// 冰霜元素
type Card140 struct {
	bcard.Card
}

func (c *Card140) NewPoint() iface.ICard {
	return &Card140{}
}

func (c *Card140) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, c.GetOwner(), c.GetOwner().GetBattle())
	buff.AddTraits(define.CardTraitsFrozen)

	rc.AddSubCards(buff)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"冻结")
}

// 萨维斯
type Card141 struct {
	bcard.Card
}

func (c *Card141) NewPoint() iface.ICard {
	return &Card141{}
}

func (c *Card141) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherAfterRelease")
}

func (c *Card141) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherAfterRelease")
}

func (c *Card141) OnNROtherAfterRelease(oc iface.ICard) {

	h := c.GetOwner()
	if oc.GetOwner().GetId() != h.GetId() || oc.GetId() == c.GetId() {
		return
	}

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	bidx := h.GetIdxByCards(c, h.GetBattleCards())

	// 召唤
	nc := iface.GetCardFact().GetCard(142)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 萨维亚萨特
type Card142 struct {
	bcard.Card
}

func (c *Card142) NewPoint() iface.ICard {
	return &Card142{}
}

// 风怒鹰身人
type Card143 struct {
	bcard.Card
}

func (c *Card143) NewPoint() iface.ICard {
	return &Card143{}
}

// 霍格
type Card144 struct {
	bcard.Card
}

func (c *Card144) NewPoint() iface.ICard {
	return &Card144{}
}

func (c *Card144) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNRRoundEnd")
}

func (c *Card144) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNRRoundEnd")
}

func (c *Card144) OnNRRoundEnd() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		!c.GetOwner().IsRoundHero() {
		return
	}

	h := c.GetOwner()
	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	bidx := h.GetIdxByCards(c, h.GetBattleCards())

	// 召唤
	nc := iface.GetCardFact().GetCard(145)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 豺狼人
type Card145 struct {
	bcard.Card
}

func (c *Card145) NewPoint() iface.ICard {
	return &Card145{}
}

// 烈日行者
type Card146 struct {
	bcard.Card
}

func (c *Card146) NewPoint() iface.ICard {
	return &Card146{}
}

// 加基森拍卖师
type Card147 struct {
	bcard.Card
}

func (c *Card147) NewPoint() iface.ICard {
	return &Card147{}
}

func (c *Card147) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card147) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherBeforeRelease")
}

func (c *Card147) OnNROtherBeforeRelease(oc, rc iface.ICard) (iface.ICard, bool) {

	h := c.GetOwner()

	if oc.GetConfig().Ctype != define.CardTypeSorcery {
		return rc, true
	}

	if oc.GetType() != define.CardTypeSorcery {
		return rc, true
	}

	h.DrawByTimes(1)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(h.GetHead())+"抽一张牌")

	return rc, true
}

// 比斯巨兽
type Card148 struct {
	bcard.Card
}

func (c *Card148) NewPoint() iface.ICard {
	return &Card148{}
}

func (c *Card148) OnDie() {

	h := c.GetOwner()
	if len(h.GetEnemy().GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(149)
	nc.Init(nc, define.InCardsTypeNone, h.GetEnemy(), h.GetBattle())
	nc.GetOwner().MoveToBattle(nc, -1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"死亡时，召唤了"+push.GetCardLogString(nc)+"给了"+push.GetHeroLogString(nc.GetOwner()))
}

// 皮普·急智
type Card149 struct {
	bcard.Card
}

func (c *Card149) NewPoint() iface.ICard {
	return &Card149{}
}

// 黑骑士
type Card150 struct {
	bcard.Card
}

func (c *Card150) NewPoint() iface.ICard {
	return &Card150{}
}

func (c *Card150) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}

	if rc.IsHaveTraits(define.CardTraitsTaunt) {
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"杀死了"+push.GetCardLogString(rc))
		rc.GetOwner().DieCard(rc, false)
	}
}
