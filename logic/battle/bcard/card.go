package bcard

import (
	"bytes"
	"encoding/gob"
	"hs/logic/config"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/iface"
	"hs/logic/push"
)

// 卡牌
type Card struct {
	Id           int                 // id
	ReleaseId    int                 // battle中的释放id
	realization  iface.ICard         // 我的实现
	Config       *config.CardConfig  // 配置
	Ctype        define.CardType     // 卡牌类型
	Race         []define.CardRace   // 卡牌种族
	Traits       []define.CardTraits // 卡牌特质
	Shield       int                 // 盾
	Hp           int                 // 卡牌血量
	HpEffect     map[int]int         // 卡牌影响的血量
	HpMax        int                 // 卡牌血上限
	Damage       int                 // 攻击力
	Mona         int                 // 能量
	InCardsType  define.InCardsType  // 卡牌的位置
	Owner        iface.IHero         // 所属人
	AttackTimes  int                 // 攻击次数
	ReleaseRound int                 // 出牌回合
	InitSign     bool                // 设置初始化标记
	SilentSign   bool                // 是否被沉默
	ApDamage     int                 // 法术伤害
	DbIdx        int                 // 死亡后的idx
	FatherCard   iface.ICard         // 父卡牌
	SubCards     []iface.ICard       // 子卡牌
}

// 返回新指针
func (c *Card) NewPoint() iface.ICard {
	return &Card{}
}

// 初始化卡牌 ，ic = c
func (c *Card) Init(ic iface.ICard, ict define.InCardsType, h iface.IHero, b iface.IBattle) {

	if c.InitSign {
		return
	}

	c.InitSign = true
	c.Id = b.GetIncrCardId()
	c.InCardsType = ict
	c.Owner = h
	c.AttackTimes = 0
	c.realization = ic

	h.AppendToAllCards(ic)
	c.Reset()

	ic.OnInit()
}

// 设置id
func (c *Card) SetId(id int) {
	c.Id = id
}

func (c *Card) GetId() int {
	return c.Id
}

// 获得实现
func (c *Card) GetRealization() iface.ICard {
	return c.realization
}

// 设置配置
func (c *Card) SetConfig(conf *config.CardConfig) {
	c.Config = conf
}

// 获得配置
func (c *Card) GetConfig() *config.CardConfig {
	return c.Config
}

// 获得类型
func (c *Card) GetType() define.CardType {
	return c.Ctype
}

// 获得种族
func (c *Card) GetRace() []define.CardRace {
	return c.Race
}

// 是否是某个种族
func (c *Card) IsRace(cr define.CardRace) bool {
	for _, v := range c.GetRace() {
		if v == define.CardRaceAll || v == cr {
			return true
		}
	}

	return false
}

// 获得特质
func (c *Card) GetTraits() []define.CardTraits {
	return c.Traits
}

// 获得有影响的特质 , ic = c
func (c *Card) GetHaveEffectTraits() []define.CardTraits {

	ic := c.GetRealization()

	ts := ic.GetTraits()
	for _, v := range ic.GetSubCards() {
		for _, ct := range v.GetTraits() {

			if !help.InArray(ct, ts) {
				ts = append(ts, ct)
			}
		}
	}

	// 获得光环影响
	for _, v := range ic.GetOwner().GetBattle().GetEventCards("OnNROtherGetTraits") {

		for _, v2 := range v.OnNROtherGetTraits(ic) {
			if !help.InArray(v2, ts) {
				ts = append(ts, v2)
			}
		}
	}

	return ts
}

// ic是否拥有卡牌特质
func (c *Card) IsHaveTraits(ct define.CardTraits) bool {
	ic := c.GetRealization()
	return help.InArray(ct, ic.GetHaveEffectTraits())
}

// 添加特质
func (c *Card) AddTraits(ct define.CardTraits) {

	if help.InArray(ct, c.Traits) {
		return
	}

	c.Traits = append(c.Traits, ct)
}

// 删除特质
func (c *Card) RemoveTraits(ct define.CardTraits) {
	for idx, v := range c.Traits {
		if v == ct {
			c.Traits = append(c.Traits[:idx], c.Traits[idx+1:]...)
			return
		}
	}

	for _, v := range c.GetSubCards() {
		v.RemoveTraits(ct)
	}
}

// 获得护盾
func (c *Card) GetShield() int {
	return c.Shield
}

// 设置护盾
func (c *Card) SetShield(s int) {
	c.Shield += s
}

// 治疗血量
func (c *Card) TreatmentHp(num int) {

	ic := c.GetRealization()
	oldHp := c.GetHaveEffectHp()

	c.AddHp(num)

	if c.GetHaveEffectHp() > oldHp && !c.IsSilent() {
		ic.OnAfterHpChange()
	}
}

// 加血
func (c *Card) AddHp(num int) {
	c.Hp += num
	if c.Hp > c.GetHaveEffectHpMax() {
		c.Hp = c.GetHaveEffectHpMax()
	}
}

// 加血上限和血
func (c *Card) AddHpMaxAndHp(num int) {
	c.HpMax += num
	c.AddHp(num)
}

// 设置血上限和血
func (c *Card) SetHpMaxAndHp(set int) {
	c.SetHpMax(set)
	c.SetHp(set)
}

// 扣除血量
func (c *Card) CostHp(num int) int {

	ic := c.GetRealization()

	// 是否拥有圣盾
	if num > 0 && c.IsHaveTraits(define.CardTraitsHolyShield) {
		num = 0
		c.RemoveTraits(define.CardTraitsHolyShield)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"圣盾消失")
	}

	if num > 0 && c.IsHaveTraits(define.CardTraitsImmune) {
		num = 0
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"具有免疫，伤害无效")
	}

	if num > 0 && !c.IsSilent() {
		num = ic.OnBeforeCostHp(num)
	}

	if c.Shield >= num {
		c.Shield -= num
		num = 0
	} else {
		num = num - c.Shield
		c.Shield = 0
	}

	tcNum := num
	// 扣一下光环加成的血
	if num > 0 {
		c.GetHaveEffectHp()
		for k, v := range c.HpEffect {
			if v > 0 {
				if num > v {
					num = num - v
					c.HpEffect[k] = 0
				} else {
					c.HpEffect[k] = v - num
					num = 0
				}
			}
		}
	}

	c.Hp -= num

	if tcNum > 0 && !c.IsSilent() {
		ic.OnAfterCostHp()
		ic.OnAfterHpChange()
	}

	if c.Hp <= 0 {

		h := c.GetOwner()
		var tc iface.ICard
		if c.GetCardInCardsPos() == define.InCardsTypeBody {
			tc = h.GetWeapon()
		} else if c.GetType() == define.CardTypeHero {
			// 弹出
			h.Die()
			return num
		} else {
			tc = h.GetBattleCardById(c.GetId())
		}

		h.DieCard(tc, false)
	}

	return num
}

// 设置血量
func (c *Card) SetHp(Hp int) {
	c.Hp = Hp
}

// 获得卡牌血量
func (c *Card) GetHp() int {
	return c.Hp
}

// 删除Hp影响数据
func (c *Card) DeleteHpEffect() {
	c.HpEffect = make(map[int]int, 0)
}

// 刷新影响数据
func (c *Card) flushHpEffect() {
	var cacheHpEffect = make(map[int]int, 0)

	// 获得光环的+血
	for _, v := range c.Owner.GetBattle().GetEventCards("OnNROtherGetHp") {

		eHp := v.OnNROtherGetHp(c)

		if eHp <= 0 {
			continue
		}

		cacheHpEffect[v.GetId()] = eHp

		if _, ok := c.HpEffect[v.GetId()]; !ok {
			c.HpEffect[v.GetId()] = cacheHpEffect[v.GetId()]
		}
	}

	// 获得buff的加血
	for _, v := range c.GetSubCards() {
		eHp := v.GetHp()

		if eHp <= 0 {
			continue
		}

		cacheHpEffect[v.GetId()] = eHp

		if _, ok := c.HpEffect[v.GetId()]; !ok {
			c.HpEffect[v.GetId()] = cacheHpEffect[v.GetId()]
		}
	}

	for k := range c.HpEffect {
		if _, ok := cacheHpEffect[k]; !ok {
			delete(c.HpEffect, k)
		}
	}
}

// 获得有血量影响的Hp
func (c *Card) GetHaveEffectHp() int {
	c.flushHpEffect()
	Hp := c.GetHp()
	for _, v := range c.HpEffect {
		Hp += v
	}

	return Hp
}

// 设置血上限
func (c *Card) SetHpMax(HpMax int) {
	c.HpMax = HpMax
}

// 获得卡牌最大血量
func (c *Card) GetHpMax() int {
	return c.HpMax
}

// 获得有血量影响的HpMax
func (c *Card) GetHaveEffectHpMax() int {

	HpMax := c.GetHpMax()
	for _, v := range c.Owner.GetBattle().GetEventCards("OnNROtherGetHp") {

		HpMax += v.OnNROtherGetHp(c)
	}

	for _, v := range c.GetSubCards() {
		HpMax += v.GetHpMax()
	}

	return HpMax
}

// 获得卡牌攻击力
func (c *Card) GetDamage() int {
	return c.Damage
}

// 计算ic有效果加成的卡牌攻击力
func (c *Card) GetHaveEffectDamage() int {

	ic := c.GetRealization()
	d := ic.GetDamage()

	if ic.GetType() == define.CardTypeHero {
		w := c.GetOwner().GetWeapon()
		if w != nil {
			d += w.GetHaveEffectDamage()
		}
	}

	if !ic.IsSilent() {
		d = ic.OnGetDamage(d)
	}

	for _, v := range ic.GetOwner().GetBattle().GetEventCards("OnNROtherGetDamage") {
		d += v.OnNROtherGetDamage(ic)
	}

	for _, v := range ic.GetSubCards() {
		d += v.GetDamage()
	}

	if d < 0 {
		d = 0
	}

	return d
}

// 添加攻击力
func (c *Card) AddDamage(add int) {
	c.Damage += add
}

// 设置攻击力
func (c *Card) SetDamage(d int) {
	c.Damage = d
}

// 交换ic攻击和血
func (c *Card) ExchangeHpDamage() {

	ic := c.GetRealization()
	od := ic.GetHaveEffectDamage()
	oh := ic.GetHaveEffectHp()

	ic.SetHpMaxAndHp(od)
	ic.SetDamage(oh)

	// 固化属性
	ic.DeleteHpEffect()
	scs := ic.GetSubCards()
	for _, v := range scs {
		if v.GetDamage() != 0 {
			v.SetDamage(0)
		}
		if v.GetHp() != 0 {
			v.SetHpMaxAndHp(0)
		}
	}
}

// 获得法术伤害
func (c *Card) GetApDamage() int {
	return c.ApDamage
}

// 计算ic有效果加成的卡牌法术伤害
func (c *Card) GetHaveEffectApDamage(ic iface.ICard) int {
	d := ic.GetApDamage()

	for _, v := range ic.GetSubCards() {
		d += v.GetApDamage()
	}

	if d < 0 {
		d = 0
	}

	return d
}

// 获得费用
func (c *Card) GetMona() int {
	return c.Mona
}

// 设置费用
func (c *Card) SetMona(m int) {
	c.Mona = m
}

// 计算有效果加成的卡牌费用
func (c *Card) GetHaveEffectMona() int {
	ic := c.GetRealization()
	m := ic.GetMona()

	if !ic.IsSilent() {
		m = ic.OnGetMona(m)
	}

	for _, v := range ic.GetOwner().GetBattle().GetEventCards("OnNROtherGetMona") {
		m += v.OnNROtherGetMona(ic)
	}

	if m < 0 {
		m = 0
	}

	return m
}

// 设置此卡在卡牌组中的位置
func (c *Card) SetCardInCardsPos(ict define.InCardsType) {
	c.InCardsType = ict
}

// 获得此卡在卡牌组中的位置
func (c *Card) GetCardInCardsPos() define.InCardsType {
	return c.InCardsType
}

// 设置死亡后的bidx
func (c *Card) SetAfterDieBidx(dbidx int) {
	c.DbIdx = dbidx
}

// 获得死亡后的bidx
func (c *Card) GetAfterDieBidx() int {
	return c.DbIdx
}

// 设置拥有人
func (c *Card) SetOwner(h iface.IHero) {
	c.Owner = h
}

// 获得此卡拥有人
func (c *Card) GetOwner() iface.IHero {

	if c.GetFatherCard() != nil {
		return c.GetFatherCard().GetOwner()
	}

	return c.GetNoLoopOwner()
}

func (c *Card) GetNoLoopOwner() iface.IHero {
	return c.Owner
}

// 设置攻击次数
func (c *Card) SetAttackTimes(t int) {
	c.AttackTimes = t
}

// 获得攻击次数
func (c *Card) GetAttackTimes() int {
	return c.AttackTimes
}

// 获得最大攻击次数
func (c *Card) GetMaxAttackTimes() int {

	ic := c.GetRealization()

	// 如果是英雄卡
	if ic.GetConfig().Ctype == define.CardTypeHero {

		// 需要获得到武器属性
		w := c.GetOwner().GetWeapon()
		if w != nil && help.InArray(define.CardTraitsWindfury, w.GetTraits()) {
			return 2
		}
		return 1
	}

	if help.InArray(define.CardTraitsWindfury, c.GetTraits()) {
		return 2
	}
	return 1
}

// 复制此卡 ic = c
func (c *Card) Copy() (iface.ICard, error) {

	ic := c.GetRealization()
	nc := iface.GetCardFact().GetCard(ic.GetConfig().Id)
	nc.Init(nc, define.InCardsTypeNone, c.GetOwner(), c.Owner.GetBattle())

	// 先去除owner , 结束时候还原回来
	owner := ic.GetNoLoopOwner()
	ic.SetOwner(nil)
	defer ic.SetOwner(owner)

	// deep copy
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(ic); err != nil {
		return nil, err
	}

	if err := gob.NewDecoder(&buf).Decode(nc); err != nil {
		return nil, err
	}

	// 设置owner
	nc.SetOwner(owner)
	nc.SetId(owner.GetBattle().GetIncrCardId())
	nc.SetCardInCardsPos(define.InCardsTypeNone)

	return nc, nil
}

// 重置此卡
func (c *Card) Reset() {
	c.Ctype = c.Config.Ctype          // 卡牌类型
	c.Race = c.Config.Race            // 卡牌种族
	c.Traits = c.Config.Traits        // 卡牌特质
	c.Hp = c.Config.Hp                // 卡牌血量
	c.HpMax = c.Config.Hp             // 卡牌血上限
	c.ApDamage = c.Config.ApDamage    // 法术伤害
	c.Damage = c.Config.Damage        // 攻击力
	c.Mona = c.Config.Mona            // 能量
	c.HpEffect = make(map[int]int, 0) // HpEffect
	c.SilentSign = false              // 沉默
	c.ReleaseId = 0                   // 释放id
	c.DbIdx = 0                       // 死亡后的bidx
	c.SetSubCards(make([]iface.ICard, 0))
}

// 沉默此卡
func (c *Card) Silent() {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle {
		return
	}

	// 属性，种族
	c.Traits = make([]define.CardTraits, 0)
	c.Race = make([]define.CardRace, 0)

	// 移除子卡牌和子卡牌所有事件
	for _, v := range c.GetSubCards() {
		c.GetOwner().GetBattle().RemoveCardFromAllEvent(v)
	}
	c.SetSubCards(make([]iface.ICard, 0))
	c.GetOwner().GetBattle().RemoveCardFromAllEvent(c)

	// 血量修正
	c.HpMax = c.Config.Hp
	if c.Hp > c.HpMax {
		c.Hp = c.HpMax
	}

	// 攻击修正
	c.Damage = c.Config.Damage

	// 费用修正
	c.Mona = c.Config.Mona

	// 放弃法术伤害
	c.ApDamage = 0

	c.SilentSign = true
}

// 是否被沉默
func (c *Card) IsSilent() bool {
	return c.SilentSign
}

// 设置出牌回合
func (c *Card) SetReleaseRound(r int) {
	c.ReleaseRound = r
	c.ReleaseId = c.GetOwner().GetBattle().GetIncrReleaseId()
}

// 获得出牌回合
func (c *Card) GetReleaseRound() int {
	return c.ReleaseRound
}

// 获得释放id
func (c *Card) GetReleaseId() int {
	return c.ReleaseId
}

// 设置父卡牌
func (c *Card) SetFatherCard(fc iface.ICard) {
	c.FatherCard = fc
}

// 获得父卡牌
func (c *Card) GetFatherCard() iface.ICard {
	return c.FatherCard
}

// 获得子卡牌
func (c *Card) GetSubCards() []iface.ICard {
	return c.SubCards
}

// 设置子卡牌
func (c *Card) SetSubCards(scs []iface.ICard) {
	c.SubCards = scs
}

// 添加子卡牌
func (c *Card) AddSubCards(sc iface.ICard) {

	sc.SetFatherCard(c.GetRealization())
	c.SubCards = append(c.SubCards, sc)
}

// 删除子卡牌
func (c *Card) RemoveSubCards(sc iface.ICard) {

	idx := -1
	for k, v := range c.SubCards {
		if v.GetId() == sc.GetId() {
			idx = k
			break
		}
	}

	if idx != -1 {
		_, c.SubCards = help.DeleteCardFromCardsByIdx(c.SubCards, idx)
	}
}
