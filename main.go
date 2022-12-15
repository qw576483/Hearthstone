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

	leaf.Run(
		game.Module,
		gate.Module,
	)
}
