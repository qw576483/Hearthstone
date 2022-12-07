package bcard

import (
	"hs/logic/define"
	"hs/logic/iface"
)

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

func (c *Card) OnNRRoundBegin()                                      {}               // 回合开始时
func (c *Card) OnNRRoundEnd()                                        {}               // 回合结束时
func (c *Card) OnNROtherBeforeRelease(oc iface.ICard)                {}               // 其他卡牌释放前
func (c *Card) OnNROtherBeforeReleaseCheckValid(oc iface.ICard) bool { return false } // 其他卡牌释放前，返回是否拦截
func (c *Card) OnNROtherAfterRelease(oc iface.ICard)                 {}               // 其他卡牌释放后
func (c *Card) OnNRPutToBattle(oc iface.ICard)                       {}               // 其他卡牌步入战场时
func (c *Card) OnNROtherDie(oc iface.ICard)                          {}               // 其他卡牌死亡时
func (c *Card) OnNROtherGetMona(oc iface.ICard) int                  { return 0 }     // 其他卡牌获取自己的费用时， 返回费用加成
func (c *Card) OnNROtherGetDamage(oc iface.ICard) int                { return 0 }     // 其他卡牌获取自己的攻击力时 ， 返回攻击加成
func (c *Card) OnNROtherGetApDamage(oh iface.IHero) int              { return 0 }     // 英雄获取自己的法术伤害时 ， 返回的法术伤害加成
func (c *Card) OnNROtherGetHp(oc iface.ICard) int                    { return 0 }     // 其他卡牌获取自己的血量时 ， 返回血量加成
func (c *Card) OnNROtherGetTraits(oc iface.ICard) define.CardTraits  { return -1 }    // 其他卡牌获得自己的特质时，返回特质加成
