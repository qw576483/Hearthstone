package bcard

import (
	"hs/logic/define"
	"hs/logic/iface"
)

// 子类方法，如果在(c *Card)中调用，需要反射调用，可以查看OnInit()
func (c *Card) OnInit()                                                      {}
func (c *Card) OnBattleBegin()                                               {}
func (c *Card) OnGet()                                                       {}
func (c *Card) OnPutToBattle(pix int)                                        {}
func (c *Card) OnOutBattle()                                                 {}
func (c *Card) OnRelease(choiceId, pidx int, rc iface.ICard, rh iface.IHero) {}
func (c *Card) OnHonorAnnihilate()                                           {}
func (c *Card) OnOverflowAnnihilate()                                        {}
func (c *Card) OnBeforeCostHp(d int) int                                     { return d }
func (c *Card) OnAfterCostHp()                                               {}
func (c *Card) OnDie(bidx int)                                               {}
func (c *Card) OnDevastate()                                                 {}
func (c *Card) OnGetMona(m int) int                                          { return m }
func (c *Card) OnGetDamage(d int) int                                        { return d }

func (c *Card) OnNRRoundBegin()                                           {}
func (c *Card) OnNRRoundEnd()                                             {}
func (c *Card) OnNROtherBeforeRelease(oc iface.ICard)                     {}
func (c *Card) OnNROtherBeforeReleaseCheckValid(oc iface.ICard) bool      { return false }
func (c *Card) OnNROtherAfterRelease(oc iface.ICard)                      {}
func (c *Card) OnNRPutToBattle(oc iface.ICard)                            {}
func (c *Card) OnNROtherDie(oc iface.ICard)                               {}
func (c *Card) OnNROtherGetMona(oc iface.ICard) int                       { return 0 }
func (c *Card) OnNROtherGetDamage(oc iface.ICard) int                     { return 0 }
func (c *Card) OnNROtherGetApDamage(oh iface.IHero) int                   { return 0 }
func (c *Card) OnNROtherGetHp(oc iface.ICard) int                         { return 0 }
func (c *Card) OnNROtherGetTraits(oc iface.ICard) []define.CardTraits     { return nil }
func (c *Card) OnNROtherHeroGetTraits(oh iface.IHero) []define.CardTraits { return nil }
