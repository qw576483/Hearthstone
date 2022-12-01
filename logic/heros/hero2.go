package heros

import (
	"hs/logic/battle"
	"hs/logic/iface"
)

type Hero2 struct {
	battle.Hero
}

func (h *Hero2) NewPoint() iface.IHero {
	return &Hero2{}
}
