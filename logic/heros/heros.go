package heros

import (
	"hs/logic/config"
	"hs/logic/iface"
)

// 获得卡牌
func GetHero(configId int) iface.IHero {

	h := heros[configId].NewPoint()
	h.SetConfig(config.GetHeroConfig(configId))

	return h
}
