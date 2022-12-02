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
	hpMax        int                 // 卡牌血上限
	damage       int                 // 攻击力
	mona         int                 // 能量
	inCardsType  define.InCardsType  // 卡牌的位置
	owner        iface.IHero         // 所属人
	attackTimes  int                 // 攻击次数
	buffs        []iface.IBuff       // buff
	releaseRound int                 // 出牌回合
	initSign     bool                // 设置初始化标记
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

// 是否拥有卡牌特质
func (c *Card) IsHaveTraits(ct define.CardTraits) bool {
	return help.InArray(ct, c.GetTraits())
}

// 添加特质
func (c *Card) AddTraits(ct define.CardTraits) {

	for _, v := range c.traits {
		if v == ct {
			return
		}
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
	if c.IsHaveTraits(define.CardTraitsHolyShield) {
		num = 0
		c.RemoveTraits(define.CardTraitsHolyShield)
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"圣盾消失")
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

// 设置血上限
func (c *Card) SetHpMax(hpMax int) {
	c.hpMax = hpMax
}

// 获得卡牌最大血量
func (c *Card) GetHpMax() int {
	return c.hpMax
}

// 获得卡牌攻击力
func (c *Card) GetDamage() int {
	return c.damage
}

// 添加攻击力
func (c *Card) AddDamage(add int) {
	c.damage += add
}

// 设置攻击力
func (c *Card) SetDamage(d int) {
	c.damage = d
}

// 获得法力值
func (c *Card) GetMona() int {
	return c.mona
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
	return c.owner
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

// 获得buffs
func (c *Card) GetBuffs() []iface.IBuff {
	return c.buffs
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
	c.ctype = c.config.Ctype         // 卡牌类型
	c.race = c.config.Race           // 卡牌种族
	c.traits = c.config.Traits       // 卡牌特质
	c.hp = c.config.Hp               // 卡牌血量
	c.hpMax = c.config.Hp            // 卡牌血上限
	c.damage = c.config.Damage       // 攻击力
	c.mona = c.config.Mona           // 能量
	c.buffs = make([]iface.IBuff, 0) // buff
}

// 沉默此卡
func (c *Card) Silent() error {

	if c.GetCardInCardsPos() != define.InCardsTypeBattle {
		return errors.New("卡牌不在战场上")
	}

	// 属性，种族，buffs修正
	c.traits = make([]define.CardTraits, 0)
	c.race = make([]define.CardRace, 0)
	c.buffs = make([]iface.IBuff, 0)

	// 血量修正
	if c.hpMax > c.config.Hp {
		c.hpMax = c.config.Hp
	}
	if c.hp > c.hpMax {
		c.hp = c.hpMax
	}

	// 攻击修正
	c.damage = c.config.Damage

	return nil
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
func (c *Card) OnInit()                                                      {} // 初始化时
func (c *Card) OnBattleBegin()                                               {} // 战斗开始
func (c *Card) OnGet()                                                       {} // 获得时
func (c *Card) OnPutToBattle(pix int)                                        {} // 放置到战场时
func (c *Card) OnOutBattle()                                                 {} // 离开战场时
func (c *Card) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {} // 释放时
func (c *Card) OnHonorAnnihilate(ec iface.ICard)                             {} // 荣誉消灭
func (c *Card) OnOverflowAnnihilate(ec iface.ICard)                          {} // 超杀
func (c *Card) OnDie(bidx int)                                               {} // 卡牌死亡时（死亡后触发销毁）
func (c *Card) OnDevastate()                                                 {} // 卡牌销毁时

func (c *Card) OnNRRoundBegin()                {} // 回合开始时
func (c *Card) OnNRRoundEnd()                  {} // 回合结束时
func (c *Card) OnNRPutToBattle(oc iface.ICard) {} // 其他卡牌步入战场时
func (c *Card) OnNROtherDie(oc iface.ICard)    {} // 其他卡牌死亡时
