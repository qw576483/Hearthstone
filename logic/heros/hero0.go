package heros

import (
	"hs/logic/battle"
	"hs/logic/iface"
)

type Hero0 struct {
	battle.Hero
}

func (h *Hero0) NewPoint() iface.IHero {
	return &Hero0{}
}
