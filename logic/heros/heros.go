package heros

import (
	"hs/logic/battle"
	"hs/logic/config"
	"hs/logic/iface"
)

// 获得英雄
func GetHero(configId int) iface.IHero {

	h := heros[configId].NewPoint()
	h.SetConfig(config.GetHeroConfig(configId))

	return h
}

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

// 萨
type Hero3 struct {
	battle.Hero
}

func (h *Hero3) NewPoint() iface.IHero {
	return &Hero3{}
}

// 萨
type Hero4 struct {
	battle.Hero
}

func (h *Hero4) NewPoint() iface.IHero {
	return &Hero4{}
}
