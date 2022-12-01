package heros

import (
	"hs/logic/battle"
	"hs/logic/iface"
)

type Hero1 struct {
	battle.Hero
}

func (h *Hero1) NewPoint() iface.IHero {
	return &Hero1{}
}
