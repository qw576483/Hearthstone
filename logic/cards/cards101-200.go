package cards

import (
	"hs/logic/battle/bcard"
	"hs/logic/define"
	"hs/logic/help"
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

func (c *Card110) OnNROtherBeforeCostHp(who, target iface.ICard, num int) int {

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

func (c *Card110) OnNROtherBeforeTreatmentHp(who, target iface.ICard, num int) int {

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

// 苔原犀牛
type Card151 struct {
	bcard.Card
}

func (c *Card151) NewPoint() iface.ICard {
	return &Card151{}
}

func (c *Card151) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetTraits")
}

func (c *Card151) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetTraits")
}

func (c *Card151) OnNROtherGetTraits(oc iface.ICard, ts []define.CardTraits) []define.CardTraits {

	if oc.GetType() == define.CardTypeEntourage && oc.IsRace(define.CardRaceBeast) && !help.InArray(define.CardTraitsAssault, ts) {
		ts = append(ts, define.CardTraitsAssault)
	}
	return ts
}

// 正义
type Card152 struct {
	bcard.Card
}

func (c *Card152) NewPoint() iface.ICard {
	return &Card152{}
}

func (c *Card152) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	for _, v := range h.GetBattleCards() {

		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得圣盾")
		v.AddTraits(define.CardTraitsHolyShield)
	}
}

// 疾跑
type Card153 struct {
	bcard.Card
}

func (c *Card153) NewPoint() iface.ICard {
	return &Card153{}
}

func (c *Card153) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.DrawByTimes(4)

	push.PushAutoLog(h, "抽了4张牌")
}

// 嗜血
type Card154 struct {
	bcard.Card
}

func (c *Card154) NewPoint() iface.ICard {
	return &Card154{}
}

func (c *Card154) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for _, v := range h.GetBattleCards() {

		buff := iface.GetCardFact().GetCard(define.BuffCardId_MyRoundEndClear)
		buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
		buff.AddDamage(3)

		v.AddSubCards(buff)
		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"获得了3点攻击力")
	}
}

// 奥金斧
type Card155 struct {
	bcard.Card
}

func (c *Card155) NewPoint() iface.ICard {
	return &Card155{}
}

// 雷矛特种兵
type Card156 struct {
	bcard.Card
}

func (c *Card156) NewPoint() iface.ICard {
	return &Card156{}
}

func (c *Card156) OnRelease(choiceId, bidx int, rc iface.ICard) {

	if rc == nil {
		return
	}
	rc.CostHp(c, 2)
}

// 霜狼督军
type Card157 struct {
	bcard.Card
}

func (c *Card157) NewPoint() iface.ICard {
	return &Card157{}
}

func (c *Card157) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	bNum := len(h.GetBattleCards())

	if bNum <= 0 {
		return
	}

	buff := iface.GetCardFact().GetCard(define.BuffCardId_Forever)
	buff.Init(buff, define.InCardsTypeNone, h, h.GetBattle())
	buff.AddDamage(bNum)
	buff.AddHpMaxAndHp(bNum)

	c.AddSubCards(buff)
	push.PushAutoLog(h, push.GetCardLogString(c)+"获得了+"+strconv.Itoa(bNum)+"/+"+strconv.Itoa(bNum))
}

// 暗鳞治愈者
type Card158 struct {
	bcard.Card
}

func (c *Card158) NewPoint() iface.ICard {
	return &Card158{}
}

func (c *Card158) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	for _, v := range h.GetBattleCards() {
		v.TreatmentHp(c, 2)
		push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(v)+"恢复了2点生命值")
	}

	h.GetHead().TreatmentHp(c, 2)
	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(h.GetHead())+"恢复了2点生命值")
}

// 夜刃刺客
type Card159 struct {
	bcard.Card
}

func (c *Card159) NewPoint() iface.ICard {
	return &Card159{}
}

func (c *Card159) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	h.GetEnemy().GetHead().CostHp(c, 3)
}

// 藏宝海湾保镖
type Card160 struct {
	bcard.Card
}

func (c *Card160) NewPoint() iface.ICard {
	return &Card160{}
}

// 精英牛头人酋长
type Card161 struct {
	bcard.Card
}

func (c *Card161) NewPoint() iface.ICard {
	return &Card161{}
}

func (c *Card161) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	push.PushAllLog(h.GetBattle(), "【你的对手】获得了一张卡牌")

	randIdx := h.GetBattle().GetRand().Intn(len(define.EliteTaurenChieftainIds))
	nc := iface.GetCardFact().GetCard(define.EliteTaurenChieftainIds[randIdx])
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToHand(nc)

	push.PushLog(h, push.GetCardLogString(c)+"让你获得了"+push.GetCardLogString(nc))

	randIdx2 := h.GetBattle().GetRand().Intn(len(define.EliteTaurenChieftainIds))
	nc2 := iface.GetCardFact().GetCard(define.EliteTaurenChieftainIds[randIdx2])
	nc2.Init(nc2, define.InCardsTypeNone, h.GetEnemy(), h.GetBattle())
	h.GetEnemy().MoveToHand(nc2)

	push.PushLog(h, push.GetCardLogString(c)+"让你获得了"+push.GetCardLogString(nc2))
}

// 我是鱼人
type Card162 struct {
	bcard.Card
}

func (c *Card162) NewPoint() iface.ICard {
	return &Card162{}
}

func (c *Card162) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	randN := h.GetBattle().GetRand().Intn(3) + 3
	for i := 1; i <= randN; i++ {

		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return
		}

		nc := iface.GetCardFact().GetCard(163)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())
		h.MoveToBattle(nc, -1)

		push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
	}
}

// 鱼人
type Card163 struct {
	bcard.Card
}

func (c *Card163) NewPoint() iface.ICard {
	return &Card163{}
}

// 潜行者的伎俩
type Card164 struct {
	bcard.Card
}

func (c *Card164) NewPoint() iface.ICard {
	return &Card164{}
}

func (c *Card164) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	dmg := 4
	dmg += h.GetApDamage()

	if rc != nil {
		rc.CostHp(c, dmg)
	}

	push.PushAutoLog(h, "抽了一张牌")
	h.DrawByTimes(1)
}

// 部落的力量
type Card165 struct {
	bcard.Card
}

func (c *Card165) NewPoint() iface.ICard {
	return &Card165{}
}

func (c *Card165) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	randIdx := h.GetBattle().GetRand().Intn(len(define.PowerOfTheHordeIds))

	nc := iface.GetCardFact().GetCard(define.PowerOfTheHordeIds[randIdx])
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToBattle(nc, -1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 霜狼步兵
type Card166 struct {
	bcard.Card
}

func (c *Card166) NewPoint() iface.ICard {
	return &Card166{}
}

// 森金持盾卫士
type Card167 struct {
	bcard.Card
}

func (c *Card167) NewPoint() iface.ICard {
	return &Card167{}
}

// 萨尔玛先知
type Card168 struct {
	bcard.Card
}

func (c *Card168) NewPoint() iface.ICard {
	return &Card168{}
}

// 银月城卫兵
type Card169 struct {
	bcard.Card
}

func (c *Card169) NewPoint() iface.ICard {
	return &Card169{}
}

// 牛头人战士
type Card170 struct {
	bcard.Card
	sub iface.ICard
}

func (c *Card170) NewPoint() iface.ICard {
	return &Card170{}
}

func (c *Card170) OnAfterHpChange() {

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

// 利爪德鲁伊
type Card171 struct {
	bcard.Card
}

func (c *Card171) NewPoint() iface.ICard {
	return &Card171{}
}

func (c *Card171) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	var nc iface.ICard
	if choiceId == 0 {
		nc = iface.GetCardFact().GetCard(173)
		push.PushAutoLog(c.GetOwner(), "[抉择1]变成了猎豹形态")
	} else {
		nc = iface.GetCardFact().GetCard(172)
		push.PushAutoLog(c.GetOwner(), "[抉择2]变成了熊形态")
	}
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveOutBattleOnlyBattleCards(c)
	h.MoveToBattle(nc, bidx)
}

// 熊形态
type Card172 struct {
	bcard.Card
}

func (c *Card172) NewPoint() iface.ICard {
	return &Card172{}
}

// 猎豹形态
type Card173 struct {
	bcard.Card
}

func (c *Card173) NewPoint() iface.ICard {
	return &Card173{}
}

// 星辰坠落
type Card174 struct {
	bcard.Card
}

func (c *Card174) NewPoint() iface.ICard {
	return &Card174{}
}

func (c *Card174) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if choiceId == 0 {
		dmg := 5
		dmg += c.GetApDamage()
		if rc == nil {
			return
		}
		push.PushAutoLog(c.GetOwner(), "[抉择1]")
		rc.CostHp(c, dmg)
	} else {

		dmg := 2
		dmg += c.GetApDamage()
		if rc == nil {
			return
		}
		push.PushAutoLog(c.GetOwner(), "[抉择2]")

		for _, v := range h.GetEnemy().GetBattleCards() {
			v.CostHp(c, dmg)
		}
	}
}

// 滋养
type Card175 struct {
	bcard.Card
}

func (c *Card175) NewPoint() iface.ICard {
	return &Card175{}
}

func (c *Card175) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if choiceId == 0 {
		push.PushAutoLog(c.GetOwner(), "[抉择1]获得两个法力水晶")
		h.AddMonaMax(2)
		h.AddMona(2)
	} else {
		push.PushAutoLog(c.GetOwner(), "[抉择2]抽三张牌")
		h.DrawByTimes(3)
	}
}

// 自然之力
type Card176 struct {
	bcard.Card
}

func (c *Card176) NewPoint() iface.ICard {
	return &Card176{}
}

func (c *Card176) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for i := 1; i <= 3; i++ {
		if len(h.GetBattleCards()) >= define.MaxBattleNum {
			return
		}

		nc := iface.GetCardFact().GetCard(define.TreantId)
		nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

		h.MoveToBattle(nc, -1)

		push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
	}
}

// 爆炸射击
type Card177 struct {
	bcard.Card
}

func (c *Card177) NewPoint() iface.ICard {
	return &Card177{}
}

func (c *Card177) OnRelease(choiceId, bidx int, rc iface.ICard) {

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

	rc.CostHp(c, 5)
	if cLeft != nil {
		cLeft.CostHp(c, 2)
	}

	if cRight != nil {
		cRight.CostHp(c, 2)
	}
}

// 神圣愤怒
type Card178 struct {
	bcard.Card
}

func (c *Card178) NewPoint() iface.ICard {
	return &Card178{}
}

func (c *Card178) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	push.PushAutoLog(h, "抽了一张牌")

	dcs := h.DrawByTimes(1)

	if len(dcs) > 0 {
		dc := dcs[0]
		rc.CostHp(c, dc.GetMona())
	}
}

// 神圣愤怒
type Card179 struct {
	bcard.Card
}

func (c *Card179) NewPoint() iface.ICard {
	return &Card179{}
}

func (c *Card179) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"攻击翻倍")
	rc.SetDamage(rc.GetHaveEffectDamage() * 2)
}

// 圣殿执行者
type Card180 struct {
	bcard.Card
}

func (c *Card180) NewPoint() iface.ICard {
	return &Card180{}
}

func (c *Card180) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if rc == nil {
		return
	}

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(rc)+"获得+3生命值。")
	rc.AddHpMaxAndHp(3)
}

// 土元素
type Card181 struct {
	bcard.Card
}

func (c *Card181) NewPoint() iface.ICard {
	return &Card181{}
}

func (c *Card181) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	push.PushAutoLog(h, "[过载+2]")
	h.SetLockMonaCache(h.GetLockMonaCache() + 2)
}

// 毁灭之锤
type Card182 struct {
	bcard.Card
}

func (c *Card182) NewPoint() iface.ICard {
	return &Card182{}
}

func (c *Card182) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	push.PushAutoLog(h, "[过载+2]")
	h.SetLockMonaCache(h.GetLockMonaCache() + 2)
}

// 末日守卫
type Card183 struct {
	bcard.Card
}

func (c *Card183) NewPoint() iface.ICard {
	return &Card183{}
}

func (c *Card183) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	for i := 1; i <= 2; i++ {
		hcs := h.GetHandCards()
		if len(hcs) == 0 {
			return
		}
		dc := h.RandCard(hcs)
		push.PushLog(h, push.GetCardLogString(c)+"丢弃了"+push.GetCardLogString(dc))
		h.DiscardCard(dc)
	}
}

// 末日灾祸
type Card184 struct {
	bcard.Card
}

func (c *Card184) NewPoint() iface.ICard {
	return &Card184{}
}

func (c *Card184) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}
	h.GetBattle().AddCardToEvent(c, "OnNROtherAfterCostHp")

	dmg := 2
	dmg += h.GetApDamage()

	rc.CostHp(c, 2)
}

func (c *Card184) OnNROtherAfterCostHp(who, target iface.ICard, num int) {
	if who.GetId() != c.GetId() {
		return
	}

	hp := target.GetHaveEffectHp()
	if hp > 0 {
		return
	}
	h := c.GetOwner()

	// 随机
	races := []define.CardRace{define.CardRaceDevil}
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
	h.MoveToBattle(nc, -1)

	push.PushAutoLog(c.GetOwner(), "由于"+push.GetCardLogString(c)+"杀死了"+push.GetCardLogString(target)+",召唤了"+push.GetCardLogString(nc))
}

// 灵魂虹吸
type Card185 struct {
	bcard.Card
}

func (c *Card185) NewPoint() iface.ICard {
	return &Card185{}
}

func (c *Card185) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if rc == nil {
		return
	}

	rc.GetOwner().DieCard(rc, false)
	h.GetHead().TreatmentHp(c, 3)

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"杀死了"+push.GetCardLogString(rc)+",并让英雄恢复3点生命")
}

// 绝命乱斗
type Card186 struct {
	bcard.Card
}

func (c *Card186) NewPoint() iface.ICard {
	return &Card186{}
}

func (c *Card186) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	bs := h.GetBattleCards()
	bs = append(bs, h.GetEnemy().GetBattleCards()...)

	if len(bs) <= 0 {
		return
	}

	randC := h.RandCard(bs)

	for _, v := range bs {
		if randC.GetId() == v.GetId() {
			continue
		}

		v.GetOwner().DieCard(v, false)

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"杀死了"+push.GetCardLogString(v))
	}
}

// 恶毒铁匠
type Card187 struct {
	bcard.Card
}

func (c *Card187) NewPoint() iface.ICard {
	return &Card187{}
}

func (c *Card187) OnAfterHpChange() {

	h := c.GetOwner()

	h.GetBattle().RemoveCardFromEvent(c, "OnNROtherGetDamage")

	if c.GetHaveEffectHp() < c.GetHaveEffectHpMax() {
		h.GetBattle().AddCardToEvent(c, "OnNROtherGetDamage")
	}
}

func (c *Card187) OnNROtherGetDamage(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeBody ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeWeapon ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	return 2
}

// 荆棘谷猛虎
type Card188 struct {
	bcard.Card
}

func (c *Card188) NewPoint() iface.ICard {
	return &Card188{}
}

// 白银之手骑士
type Card189 struct {
	bcard.Card
}

func (c *Card189) NewPoint() iface.ICard {
	return &Card189{}
}

func (c *Card189) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	nc := iface.GetCardFact().GetCard(190)
	nc.Init(nc, define.InCardsTypeNone, h, h.GetBattle())

	h.MoveToBattle(nc, bidx+1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
}

// 侍从
type Card190 struct {
	bcard.Card
}

func (c *Card190) NewPoint() iface.ICard {
	return &Card190{}
}

// 憎恶
type Card191 struct {
	bcard.Card
}

func (c *Card191) NewPoint() iface.ICard {
	return &Card191{}
}

func (c *Card191) OnDie() {

	h := c.GetOwner()
	for _, v := range h.GetBattleCards() {
		v.CostHp(c, 2)
	}

	for _, v := range h.GetEnemy().GetBattleCards() {
		v.CostHp(c, 2)
	}

	h.GetHead().CostHp(c, 2)
	h.GetEnemy().GetHead().CostHp(c, 2)
}

// 绿皮船长
type Card192 struct {
	bcard.Card
}

func (c *Card192) NewPoint() iface.ICard {
	return &Card192{}
}

func (c *Card192) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	if h.GetWeapon() == nil {
		return
	}

	h.GetWeapon().AddDamage(1)
	h.GetWeapon().AddHpMaxAndHp(1)

	push.PushAutoLog(h, push.GetCardLogString(c)+"让"+push.GetCardLogString(h.GetWeapon())+"获得+1/+1")
}

// 无面操纵者
type Card193 struct {
	bcard.Card
}

func (c *Card193) NewPoint() iface.ICard {
	return &Card193{}
}

func (c *Card193) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	if rc == nil || c.GetCardInCardsPos() != define.InCardsTypeBattle {
		return
	}

	h := c.GetOwner()
	nc, err := rc.Copy()
	if err != nil {
		return
	}

	nc.SetOwner(c.GetOwner())

	h.MoveOutBattleOnlyBattleCards(c)
	h.MoveToBattle(nc, bidx)

	push.PushAutoLog(h, push.GetCardLogString(c)+"变成了"+push.GetCardLogString(nc))
}

// 火车王里诺艾
type Card194 struct {
	bcard.Card
}

func (c *Card194) NewPoint() iface.ICard {
	return &Card194{}
}

func (c *Card194) OnRelease2(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	he := h.GetEnemy()
	for i := 1; i <= 2; i++ {
		if len(he.GetBattleCards()) >= define.MaxBattleNum {
			return
		}
		nc := iface.GetCardFact().GetCard(define.LittleDragonId)
		nc.Init(nc, define.InCardsTypeNone, he, he.GetBattle())
		he.MoveToBattle(nc, -1)

		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"召唤了"+push.GetCardLogString(nc))
	}
}

// 沼泽爬行者
type Card195 struct {
	bcard.Card
}

func (c *Card195) NewPoint() iface.ICard {
	return &Card195{}
}

// 沼泽爬行者
type Card196 struct {
	bcard.Card
}

func (c *Card196) NewPoint() iface.ICard {
	return &Card196{}
}

func (c *Card196) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()

	push.PushAutoLog(h, "抽了一张牌")
	h.DrawByTimes(1)
}

// 哈里森·琼斯
type Card197 struct {
	bcard.Card
}

func (c *Card197) NewPoint() iface.ICard {
	return &Card197{}
}

func (c *Card197) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	he := h.GetEnemy()

	if he.GetWeapon() == nil {
		return
	}

	w := he.GetWeapon()
	dtimes := w.GetHaveEffectHp()

	he.DieCard(w, false)
	h.DrawByTimes(dtimes)

	push.PushAutoLog(h, push.GetCardLogString(c)+"消灭了"+push.GetCardLogString(w)+"，抽了"+strconv.Itoa(dtimes)+"张牌")
}

// 风险投资公司雇佣兵
type Card198 struct {
	bcard.Card
}

func (c *Card198) NewPoint() iface.ICard {
	return &Card198{}
}

func (c *Card198) OnPutToBattle(bidx int) {
	c.GetOwner().GetBattle().AddCardToEvent(c, "OnNROtherGetMona")
}

func (c *Card198) OnOutBattle() {
	c.GetOwner().GetBattle().RemoveCardFromEvent(c, "OnNROtherGetMona")
}

func (c *Card198) OnNROtherGetMona(oc iface.ICard) int {

	h := c.GetOwner()
	if oc.GetCardInCardsPos() != define.InCardsTypeHand ||
		c.GetCardInCardsPos() != define.InCardsTypeBattle ||
		oc.GetType() != define.CardTypeEntourage ||
		h.GetId() != oc.GetOwner().GetId() ||
		c.GetId() == oc.GetId() {
		return 0
	}

	return 3
}

// 狂奔科多兽
type Card199 struct {
	bcard.Card
}

func (c *Card199) NewPoint() iface.ICard {
	return &Card199{}
}

func (c *Card199) OnRelease(choiceId, bidx int, rc iface.ICard) {

	h := c.GetOwner()
	bs := make([]iface.ICard, 0)

	for _, v := range h.GetEnemy().GetBattleCards() {
		if v.GetHaveEffectDamage() <= 2 {
			bs = append(bs, v)
		}
	}

	if len(bs) == 0 {
		return
	}
	dc := h.RandCard(bs)
	dc.GetOwner().DieCard(dc, false)
	push.PushAutoLog(h, push.GetCardLogString(c)+"杀死了"+push.GetCardLogString(dc))
}
