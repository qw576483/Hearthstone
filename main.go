package main

import (
	"encoding/gob"
	"hs/logic/battle/bcard"
	"hs/logic/battle/bhero"
	"hs/logic/cards"
	"hs/net/conf"
	"hs/net/game"
	"hs/net/gate"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func init() {
	gob.Register(&bcard.Card{})
	gob.Register(&bhero.Hero{})
}

func main() {

	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	cards.InitCardFact()
	// // 获得全部三费卡牌
	// cf := iface.GetCardFact()
	// d := cf.ScreenCards(cf.GetAllCardsExcludeNotCanCarry(), iface.NewScreenCardParam(iface.SCPWithMona(3)))
	// for _, v := range d {
	// 	fmt.Println(v.GetConfig().Name)
	// }

	// // 随机一个三费机械卡牌从牌库中
	// cf := iface.GetCardFact()
	// races := []define.CardRace{define.CardRaceMechanics}
	// d := cf.RandByAllCards(rand.New(rand.NewSource(time.Now().UnixNano())), iface.NewScreenCardParam(iface.SCPWithMona(3), iface.SCPWithCardRace(races)))

	// fmt.Println(d.GetConfig().Name)

	leaf.Run(
		game.Module,
		gate.Module,
	)
}
