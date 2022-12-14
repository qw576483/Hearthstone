package bcard

import (
	"hs/logic/define"
	"hs/logic/iface"
)

// 子类方法，如果在(c *Card)中调用，需要反射调用，可以查看OnInit()
func (c *Card) OnInit()                                       {}
func (c *Card) OnBattleBegin()                                {}
func (c *Card) OnGet()                                        {}
func (c *Card) OnPutToBattle(pix int)                         {}
func (c *Card) OnOutBattle()                                  {}
func (c *Card) OnBeforeAttack(ec iface.ICard) iface.ICard     { return ec }
func (c *Card) OnAfterAttack(ec iface.ICard)                  {}
func (c *Card) OnRelease(choiceId, bidx int, rc iface.ICard)  {}
func (c *Card) OnRelease2(choiceId, bidx int, rc iface.ICard) {}
func (c *Card) OnHonorAnnihilate()                            {}
func (c *Card) OnOverflowAnnihilate()                         {}
func (c *Card) OnBeforeCostHp(d int) int                      { return d }
func (c *Card) OnAfterCostHp()                                {}
func (c *Card) OnAfterHpChange()                              {}
func (c *Card) OnDie()                                        {}
func (c *Card) OnAfterDisCard()                               {}
func (c *Card) OnGetMona(m int) int                           { return m }
func (c *Card) OnGetDamage(d int) int                         { return d }
func (c *Card) OnSilent()                                     {}

func (C *Card) OnNRGetBattleTime(bt int) int { return bt }
func (c *Card) OnNRRoundBegin()              {}
func (c *Card) OnNRRoundEnd()                {}
func (c *Card) OnNROtherBeforeRelease(oc iface.ICard, rc iface.ICard) (iface.ICard, bool) {
	return rc, true
}
func (c *Card) OnNROtherAfterRelease(oc iface.ICard) {}
func (c *Card) OnNROtherBeforeAttack(oc, ec iface.ICard) iface.ICard {
	return ec
}
func (c *Card) OnNRPutToBattle(oc iface.ICard)                        {}
func (c *Card) OnNROtherDie(oc iface.ICard)                           {}
func (c *Card) OnNROtherGetMona(oc iface.ICard) int                   { return 0 }
func (c *Card) OnNROtherGetDamage(oc iface.ICard) int                 { return 0 }
func (c *Card) OnNROtherGetApDamage(oh iface.IHero) int               { return 0 }
func (c *Card) OnNROtherGetHp(oc iface.ICard) int                     { return 0 }
func (c *Card) OnNROtherGetTraits(oc iface.ICard) []define.CardTraits { return nil }
