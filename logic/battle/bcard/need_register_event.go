package bcard

import (
	"hs/logic/define"
	"hs/logic/iface"
)

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
func (c *Card) OnNRPutToBattle(oc iface.ICard)          {}
func (c *Card) OnNROtherDie(oc iface.ICard)             {}
func (c *Card) OnNROtherGetMona(oc iface.ICard) int     { return 0 }
func (c *Card) OnNROtherGetDamage(oc iface.ICard) int   { return 0 }
func (c *Card) OnNROtherGetApDamage(oh iface.IHero) int { return 0 }
func (c *Card) OnNROtherGetHp(oc iface.ICard) int       { return 0 }
func (c *Card) OnNROtherGetTraits(oc iface.ICard, cts []define.CardTraits) []define.CardTraits {
	return cts
}
func (c *Card) OnNROtherBeforeCostHp(who, target iface.ICard, num int) int { return num }
func (c *Card) OnNROtherAfterCostHp(who, target iface.ICard, num int)      {}
func (c *Card) OnNROtherBeforeTreatmentHp(who, target iface.ICard, num int) int {
	return num
}
func (c *Card) OnNROtherAfterTreatmentHp(who, target iface.ICard, num int) {}
