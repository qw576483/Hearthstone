package heros

import (
	"hs/logic/battle"
	"hs/logic/iface"
)

// 贼
type Hero0 struct {
	battle.Hero
}

func (h *Hero0) NewPoint() iface.IHero {
	return &Hero0{}
}

// 术
type Hero1 struct {
	battle.Hero
}

func (h *Hero1) NewPoint() iface.IHero {
	return &Hero1{}
}

// 猎
type Hero2 struct {
	battle.Hero
}

func (h *Hero2) NewPoint() iface.IHero {
	return &Hero2{}
}
