package battle

import (
	"hs/logic/define"
	"hs/logic/iface"
)

// 感觉像鱼人杀手蟹这种 ，消灭鱼人+2 +2，这不是buff，感觉像是一个效果影响
// 叫嚣的中士，回合内+2这种感觉是buff
type Buff struct {
	damage int                 // 增加伤害
	mona   int                 // 增加费用
	hp     int                 // 增加的hp
	traits []define.CardTraits // 增加的特质
	card   iface.ICard         // 挂载卡
	hero   iface.IHero         // 挂载英雄
}

func NewBuff(c iface.ICard, h iface.IHero) iface.IBuff {
	return &Buff{
		traits: make([]define.CardTraits, 0),
	}
}

// 添加费用，可以为负数
func (bf *Buff) AddMona(d int) {
	bf.mona += d
}

// 获得费用
func (bf *Buff) GetAddMona() int {
	return bf.mona
}

// 添加伤害，可以为负数
func (bf *Buff) AddDamage(d int) {
	bf.damage += d
}

// 获得伤害
func (bf *Buff) GetAddDamage() int {
	return bf.damage
}

// 添加血量，可以为负数
func (bf *Buff) AddHp(d int) {
	bf.hp += d
}

// 获得血量
func (bf *Buff) GetAddHp() int {
	return bf.hp
}

// 添加特质
func (bf *Buff) AddTraits(t define.CardTraits) {
	bf.traits = append(bf.traits, t)
}

// 获得特质
func (bf *Buff) GetAddTraits() []define.CardTraits {
	return bf.traits
}

// 获得挂载的卡牌
func (bf *Buff) GetPCard() iface.ICard {
	return bf.card
}

// 获得挂载的英雄（不是挂载buff卡牌的英雄，而是挂载buff的英雄）
func (bf *Buff) GetPHero() iface.IHero {
	return bf.hero
}
