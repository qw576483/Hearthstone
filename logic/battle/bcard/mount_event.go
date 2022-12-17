package bcard

import "hs/logic/iface"

func (c *Card) AddOnDie(aod iface.AddOnDie)   { c.onDieEvents = append(c.onDieEvents, aod) }
func (c *Card) GetAddOnDie() []iface.AddOnDie { return c.onDieEvents }
func (c *Card) AddOnEventClear(aoec iface.AddOnEventClear) {
	c.onEventClearEvents = append(c.onEventClearEvents, aoec)
}
func (c *Card) GetAddOnEventClear() []iface.AddOnEventClear { return c.onEventClearEvents }
