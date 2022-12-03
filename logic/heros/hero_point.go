package heros

import "hs/logic/iface"

var heros []iface.IHero = []iface.IHero{
	0: &Hero0{},
	1: &Hero1{},
	2: &Hero2{},
	3: &Hero3{},
}
