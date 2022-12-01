package heros

import (
	"hs/logic/config"
	"hs/logic/iface"
)

var heros []iface.IHero = []iface.IHero{
	0: &Hero0{},
	1: &Hero1{},
	2: &Hero2{},
}

// 获得卡牌
func GetHero(configId int) iface.IHero {

	h := heros[configId].NewPoint()
	h.SetConfig(config.GetHeroConfig(configId))

	return h
}
