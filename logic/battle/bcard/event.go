package bcard

import (
	"hs/logic/define"
	"hs/logic/iface"
)

// 子类方法，如果在(c *Card)中调用，需要反射调用，可以查看OnInit()
func (c *Card) OnInit()                                                       {}
func (c *Card) OnBattleBegin()                                                {}
func (c *Card) OnGet()                                                        {}
func (c *Card) OnPutToBattle(pix int)                                         {}
func (c *Card) OnOutBattle()                                                  {}
func (c *Card) OnRelease(choiceId, bidx int, rc iface.ICard, rh iface.IHero)  {}
func (c *Card) OnRelease2(choiceId, bidx int, rc iface.ICard, rh iface.IHero) {}
func (c *Card) OnHonorAnnihilate()                                            {}
func (c *Card) OnOverflowAnnihilate()                                         {}
func (c *Card) OnBeforeCostHp(d int) int                                      { return d }
func (c *Card) OnAfterCostHp()                                                {}
func (c *Card) OnAfterHpChange()                                              {}
func (c *Card) OnDie()                                                        {}
func (c *Card) OnAfterDisCard()                                               {}
func (c *Card) OnGetMona(m int) int                                           { return m }
func (c *Card) OnGetDamage(d int) int                                         { return d }

func (C *Card) OnNRGetBattleTime(bt int) int { return bt }
func (c *Card) OnNRRoundBegin()              {}
func (c *Card) OnNRRoundEnd()                {}
func (c *Card) OnNROtherBeforeRelease(oc iface.ICard, rc iface.ICard, rh iface.IHero) (iface.ICard, iface.IHero, bool) {
	return rc, rh, true
}
func (c *Card) OnNROtherAfterRelease(oc iface.ICard) {}
func (c *Card) OnNROtherBeforeAttack(oc, ec iface.ICard, eh iface.IHero) (iface.ICard, iface.IHero) {
	return ec, eh
}
func (c *Card) OnNRPutToBattle(oc iface.ICard)                            {}
func (c *Card) OnNROtherDie(oc iface.ICard)                               {}
func (c *Card) OnNROtherGetMona(oc iface.ICard) int                       { return 0 }
func (c *Card) OnNROtherGetDamage(oc iface.ICard) int                     { return 0 }
func (c *Card) OnNROtherGetApDamage(oh iface.IHero) int                   { return 0 }
func (c *Card) OnNROtherGetHp(oc iface.ICard) int                         { return 0 }
func (c *Card) OnNROtherGetTraits(oc iface.ICard) []define.CardTraits     { return nil }
func (c *Card) OnNROtherHeroGetTraits(oh iface.IHero) []define.CardTraits { return nil }
