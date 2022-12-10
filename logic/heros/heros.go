package heros

import (
	"hs/logic/battle/bhero"
	"hs/logic/config"
	"hs/logic/iface"
)

// 获得英雄
func GetHero(configId int) iface.IHero {

	h := heros[configId].NewPoint()
	h.SetConfig(config.GetHeroConfig(configId))

	return h
}

// 盗贼
type Hero0 struct {
	bhero.Hero
}

func (h *Hero0) NewPoint() iface.IHero {
	return &Hero0{}
}

// 术士
type Hero1 struct {
	bhero.Hero
}

func (h *Hero1) NewPoint() iface.IHero {
	return &Hero1{}
}

// 猎人
type Hero2 struct {
	bhero.Hero
}

func (h *Hero2) NewPoint() iface.IHero {
	return &Hero2{}
}

// 萨满
type Hero3 struct {
	bhero.Hero
}

func (h *Hero3) NewPoint() iface.IHero {
	return &Hero3{}
}

// 圣骑士
type Hero4 struct {
	bhero.Hero
}

func (h *Hero4) NewPoint() iface.IHero {
	return &Hero4{}
}

// 德鲁伊
type Hero5 struct {
	bhero.Hero
}

func (h *Hero5) NewPoint() iface.IHero {
	return &Hero5{}
}

// 法师
type Hero6 struct {
	bhero.Hero
}

func (h *Hero6) NewPoint() iface.IHero {
	return &Hero6{}
}

// 牧师
type Hero7 struct {
	bhero.Hero
}

func (h *Hero7) NewPoint() iface.IHero {
	return &Hero7{}
}

// 战士
type Hero8 struct {
	bhero.Hero
}

func (h *Hero8) NewPoint() iface.IHero {
	return &Hero8{}
}
