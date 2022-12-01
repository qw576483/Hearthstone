package iface

import "hs/logic/define"

type IBuff interface {
	AddMona(int)                       // 添加费用
	GetAddMona() int                   // 获得添加费用
	AddDamage(int)                     // 添加攻击力
	GetAddDamage() int                 // 获得添加的攻击力
	AddHp(int)                         // 添加血量
	GetAddHp() int                     // 获得添加的血量
	AddTraits(define.CardTraits)       // 添加特质
	GetAddTraits() []define.CardTraits // 获得添加的特质
	GetPCard() ICard                   // 获得挂载的卡牌
	GetPHero() IHero                   // 获得挂载的英雄（不是挂载buff卡牌的英雄，而是挂载buff的英雄）
}

type AddBuffCardOnDie func()
