package battle

import (
	"errors"
	"hs/logic/config"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/iface"
	"hs/logic/push"
)

// 卡牌
type Card struct {
	id           int                 // battle中id
	config       *config.CardConfig  // 配置
	ctype        define.CardType     // 卡牌类型
	race         []define.CardRace   // 卡牌种族
	traits       []define.CardTraits // 卡牌特质
	hp           int                 // 卡牌血量
	hpEffect     map[int]int         // 卡牌影响的血量
	hpMax        int                 // 卡牌血上限
	damage       int                 // 攻击力
	mona         int                 // 能量
	inCardsType  define.InCardsType  // 卡牌的位置
	owner        iface.IHero         // 所属人
	attackTimes  int                 // 攻击次数
	fatherCard   iface.ICard         // 父卡牌
	subCards     []iface.ICard       // 子卡牌，会拿到hp，damage，traits
	releaseRound int                 // 出牌回合
	initSign     bool                // 设置初始化标记
	silent       bool                // 是否被沉默
}

// 返回新指针
func (c *Card) NewPoint() iface.ICard {
	return &Card{}
}

// 初始化卡牌 ，实际上ic和c是同一个东西，防止断言错指针才这么传
func (c *Card) Init(ic iface.ICard, ict define.InCardsType, h iface.IHero, b iface.IBattle) {

	if c.initSign {
		return
	}

	c.initSign = true
	c.id = b.GetIncrCardId()
	c.inCardsType = ict
	c.owner = h
	c.attackTimes = 0

	// 实际上ic和c是同一个东西，句柄不一样
	h.AppendToAllCards(ic)
	ic.OnInit()

	c.Reset()
}

// 获得id
func (c *Card) GetId() int {
	return c.id
}

// 设置配置
func (c *Card) SetConfig(conf *config.CardConfig) {
	c.config = conf
}

// 获得配置
func (c *Card) GetConfig() *config.CardConfig {
	return c.config
}

// 获得类型
func (c *Card) GetType() define.CardType {
	return c.ctype
}

// 获得种族
func (c *Card) GetRace() []define.CardRace {
	return c.race
}

// 获得特质
func (c *Card) GetTraits() []define.CardTraits {
	return c.traits
}

// 获得有影响的特质
func (c *Card) GetHaveEffectTraits(oc iface.ICard) []define.CardTraits {

	ts := c.traits
	for _, v := range c.subCards {
		for _, ct := range v.GetTraits() {

			if !help.InArray(ct, c.traits) {
				ts = append(ts, ct)
			}
		}
	}

	// 获得光环影响
	for _, v := range c.owner.GetBothEventCards("OnNROtherGetTraits") {

		nt := v.OnNROtherGetTraits(oc)

		if nt != -1 {
			if !help.InArray(nt, c.traits) {
				ts = append(ts, nt)
			}
		}
	}

	return ts
}

// 是否拥有卡牌特质
func (c *Card) IsHaveTraits(ct define.CardTraits, oc iface.ICard) bool {
	return help.InArray(ct, c.GetHaveEffectTraits(oc))
}

// 添加特质
func (c *Card) AddTraits(ct define.CardTraits) {

	if help.InArray(ct, c.traits) {
		return
	}

	c.traits = append(c.traits, ct)
}

// 删除特质
func (c *Card) RemoveTraits(ct define.CardTraits) {
	for idx, v := range c.traits {
		if v == ct {
			c.traits = append(c.traits[:idx], c.traits[idx+1:]...)
			return
		}
	}
}

// 治疗血量
func (c *Card) TreatmentHp(num int) {
	c.AddHp(num)
}

// 加血
func (c *Card) AddHp(num int) {
	c.hp += num
	if c.hp > c.hpMax {
		c.hp = c.hpMax
	}
}

// 加血上限和血
func (c *Card) AddHpMaxAndHp(num int) {
	c.hpMax += num
	c.AddHp(num)
}

// 设置血上限和血
func (c *Card) SetHpMaxAndHp(set int) {
	c.SetHpMax(set)
	c.SetHp(set)
}

// 扣除血量
func (c *Card) CostHp(num int) int {

	// 是否拥有圣盾
	if c.IsHaveTraits(define.CardTraitsHolyShield, c) {
		num = 0
		c.RemoveTraits(define.CardTraitsHolyShield)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"圣盾消失")
	}

	// 扣一下光环加成的血
	if num > 0 {
		c.GetHaveEffectHp()
		for k, v := range c.hpEffect {
			if v > 0 {
				if num > v {
					num = num - v
					c.hpEffect[k] = 0
				} else {
					c.hpEffect[k] = v - num
					num = 0
				}
			}
		}
	}

	c.hp -= num
	if c.hp <= 0 {

		// logs
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"死亡")

		h := c.GetOwner()
		var tc iface.ICard
		if c.GetCardInCardsPos() == define.InCardsTypeBody {
			tc = h.GetWeapon()
		} else {
			tc = h.GetBattleCardById(c.GetId())
		}

		h.DieCard(tc)
	}

	return num
}

// 设置血量
func (c *Card) SetHp(hp int) {
	c.hp = hp
}

// 获得卡牌血量
func (c *Card) GetHp() int {
	return c.hp
}

// 删除hp影响数据
func (c *Card) DeleteHpEffect() {
	c.hpEffect = make(map[int]int, 0)
}

// 刷新影响数据
func (c *Card) flushHpEffect() {
	var cacheHpEffect = make(map[int]int, 0)

	// 获得光环的+血
	for _, v := range c.owner.GetBothEventCards("OnNROtherGetHp") {

		eHp := v.OnNROtherGetHp(c)

		if eHp <= 0 {
			continue
		}

		cacheHpEffect[v.GetId()] = eHp

		if _, ok := c.hpEffect[v.GetId()]; !ok {
			c.hpEffect[v.GetId()] = cacheHpEffect[v.GetId()]
		}
	}

	// 获得buff的加血
	for _, v := range c.subCards {
		eHp := v.GetHp()

		if eHp <= 0 {
			continue
		}

		cacheHpEffect[v.GetId()] = eHp

		if _, ok := c.hpEffect[v.GetId()]; !ok {
			c.hpEffect[v.GetId()] = cacheHpEffect[v.GetId()]
		}
	}

	for k := range c.hpEffect {
		if _, ok := cacheHpEffect[k]; !ok {
			delete(c.hpEffect, k)
		}
	}
}

// 获得有血量影响的hp
func (c *Card) GetHaveEffectHp() int {
	c.flushHpEffect()
	hp := c.GetHp()
	for _, v := range c.hpEffect {
		hp += v
	}

	return hp
}

// 设置血上限
func (c *Card) SetHpMax(hpMax int) {
	c.hpMax = hpMax
}

// 获得卡牌最大血量
func (c *Card) GetHpMax() int {
	return c.hpMax
}

// 获得有血量影响的hpMax
func (c *Card) GetHaveEffectHpMax() int {

	hpMax := c.GetHpMax()
	for _, v := range c.owner.GetBothEventCards("OnNROtherGetHp") {

		hpMax += v.OnNROtherGetHp(c)
	}

	return hpMax
}

// 获得卡牌攻击力
func (c *Card) GetDamage() int {
	return c.damage
}

// 计算有效果加成的卡牌攻击力
func (c *Card) GetHaveEffectDamage(tc iface.ICard) int {
	d := c.GetDamage()
	d += tc.OnGetDamage()

	for _, v := range c.owner.GetBothEventCards("OnNROtherGetDamage") {
		d += v.OnNROtherGetDamage(tc)
	}

	for _, v := range c.subCards {
		d += v.GetDamage()
	}

	if d < 0 {
		d = 0
	}

	return d
}

// 添加攻击力
func (c *Card) AddDamage(add int) {
	c.damage += add
}

// 设置攻击力
func (c *Card) SetDamage(d int) {
	c.damage = d
}

// 交换攻击和血
func (c *Card) ExchangeHpDamage(oc iface.ICard) {

	od := oc.GetHaveEffectDamage(oc)
	oh := oc.GetHaveEffectHp()

	oc.SetHpMaxAndHp(od)
	oc.SetDamage(oh)

	// 固化属性
	oc.DeleteHpEffect()
	scs := oc.GetSubCards()
	for _, v := range scs {
		if v.GetDamage() != 0 {
			v.SetDamage(0)
		}
		if v.GetHp() != 0 {
			v.SetHpMaxAndHp(0)
		}
	}
}

// 获得费用
func (c *Card) GetMona() int {
	return c.mona
}

// 计算有效果加成的卡牌费用
func (c *Card) GetHaveEffectMona(tc iface.ICard) int {
	d := c.GetMona()
	d += tc.OnGetMona()

	for _, v := range c.GetOwner().GetBothEventCards("OnNROtherGetMona") {
		d += v.OnNROtherGetMona(tc)
	}

	if d < 0 {
		d = 0
	}

	return d
}

// 设置此卡在卡牌组中的位置
func (c *Card) SetCardInCardsPos(ict define.InCardsType) {
	c.inCardsType = ict
}

// 获得此卡在卡牌组中的位置
func (c *Card) GetCardInCardsPos() define.InCardsType {
	return c.inCardsType
}

// 获得此卡在手牌中的位置
func (c *Card) GetHandPos() (int, error) {

	handCards := c.GetOwner().GetHandCards()

	for k, v := range handCards {
		if v.GetId() == c.id {
			return k, nil
		}
	}

	return 0, errors.New("not found this card")
}

// 设置拥有人
func (c *Card) SetOwner(h iface.IHero) {
	c.owner = h
}

// 获得此卡拥有人
func (c *Card) GetOwner() iface.IHero {
	if c.GetFatherCard() != nil {
		return c.GetFatherCard().GetOwner()
	}
	return c.owner
}

// 获得父级card
func (c *Card) GetFatherCard() iface.ICard {
	return c.fatherCard
}

// 获得子卡牌
func (c *Card) GetSubCards() []iface.ICard {
	return c.subCards
}

// 添加子卡牌
func (c *Card) AddSubCards(sc iface.ICard) {
	c.subCards = append(c.subCards, sc)
}

// 删除子卡牌
func (c *Card) RemoveSubCards(sc iface.ICard) {

	idx := -1
	for k, v := range c.subCards {
		if v.GetId() == sc.GetId() {
			idx = k
			break
		}
	}

	if idx != -1 {
		_, c.subCards = help.DeleteCardFromCardsByIdx(c.subCards, idx)
	}
}

// 设置攻击次数
func (c *Card) SetAttackTimes(t int) {
	c.attackTimes = t
}

// 获得攻击次数
func (c *Card) GetAttackTimes() int {
	return c.attackTimes
}

// 获得最大攻击次数
func (c *Card) GetMaxAttackTimes() int {
	if help.InArray(define.CardTraitsWindfury, c.GetTraits()) {
		return 2
	}
	return 1
}

// 复制此卡
func (c *Card) Copy() (iface.ICard, error) {

	// 检查是否满了

	// 复制
	nc := c

	return nc, nil
}

// 重置此卡
func (c *Card) Reset() {
	c.ctype = c.config.Ctype            // 卡牌类型
	c.race = c.config.Race              // 卡牌种族
	c.traits = c.config.Traits          // 卡牌特质
	c.hp = c.config.Hp                  // 卡牌血量
	c.hpMax = c.config.Hp               // 卡牌血上限
	c.damage = c.config.Damage          // 攻击力
	c.mona = c.config.Mona              // 能量
	c.hpEffect = make(map[int]int, 0)   // hpEffect
	c.subCards = make([]iface.ICard, 0) // 子卡牌
	c.silent = false                    // 沉默
}

// 沉默此卡
func (c *Card) Silent(c2 iface.ICard) {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle {
		return
	}

	// 属性，种族
	c.traits = make([]define.CardTraits, 0)
	c.race = make([]define.CardRace, 0)

	// 移除子卡牌和子卡牌所有事件
	for _, v := range c.subCards {
		c.GetOwner().RemoveCardFromBothEvent(v)
	}
	c.subCards = make([]iface.ICard, 0)
	c.GetOwner().RemoveCardFromBothEvent(c)

	// 血量修正
	c.hpMax = c.config.Hp
	if c.hp > c.hpMax {
		c.hp = c.hpMax
	}

	// 攻击修正
	c.damage = c.config.Damage
	c.silent = true
}

// 是否被沉默
func (c *Card) IsSilent() bool {
	return c.silent
}

// 设置出牌回合
func (c *Card) SetReleaseRound(r int) {
	c.releaseRound = r
}

// 获得出牌回合
func (c *Card) GetReleaseRound() int {
	return c.releaseRound
}

// 子类方法，如果在(c *Card)中调用，需要反射调用，可以查看OnInit()
func (c *Card) OnInit()                                                      {}           // 初始化时
func (c *Card) OnBattleBegin()                                               {}           // 战斗开始
func (c *Card) OnGet()                                                       {}           // 获得时
func (c *Card) OnPutToBattle(pix int)                                        {}           // 放置到战场时
func (c *Card) OnOutBattle()                                                 {}           // 离开战场时
func (c *Card) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {}           // 释放时
func (c *Card) OnHonorAnnihilate(ec iface.ICard)                             {}           // 荣誉消灭
func (c *Card) OnOverflowAnnihilate(ec iface.ICard)                          {}           // 超杀
func (c *Card) OnDie(bidx int)                                               {}           // 卡牌死亡时（死亡后触发销毁）
func (c *Card) OnDevastate()                                                 {}           // 卡牌销毁时
func (c *Card) OnGetMona() int                                               { return 0 } // 获取自己的费用时，返回费用加成
func (c *Card) OnGetDamage() int                                             { return 0 } // 获取自己的攻击力时 , 返回攻击加成

func (c *Card) OnNRRoundBegin()                                     {}               // 回合开始时
func (c *Card) OnNRRoundEnd()                                       {}               // 回合结束时
func (c *Card) OnNROtherRelease(oc iface.ICard) bool                { return false } // 其他卡牌释放时，返回是否拦截
func (c *Card) OnNRPutToBattle(oc iface.ICard)                      {}               // 其他卡牌步入战场时
func (c *Card) OnNROtherDie(oc iface.ICard)                         {}               // 其他卡牌死亡时
func (c *Card) OnNROtherGetMona(oc iface.ICard) int                 { return 0 }     // 其他卡牌获取自己的费用时， 返回费用加成
func (c *Card) OnNROtherGetDamage(oc iface.ICard) int               { return 0 }     // 其他卡牌获取自己的攻击力时 ， 返回攻击加成
func (c *Card) OnNROtherGetHp(oc iface.ICard) int                   { return 0 }     // 其他卡牌获取自己的血量时 ， 返回血量加成
func (c *Card) OnNROtherGetTraits(oc iface.ICard) define.CardTraits { return -1 }    // 其他卡牌获得自己的特质时，返回特质加成
