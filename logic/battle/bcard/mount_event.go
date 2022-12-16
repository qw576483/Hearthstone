package bcard

import "hs/logic/iface"

func (c *Card) AddOnDie(aod iface.AddOnDie)   { c.OnDieEvents = append(c.OnDieEvents, aod) }
func (c *Card) GetAddOnDie() []iface.AddOnDie { return c.OnDieEvents }
