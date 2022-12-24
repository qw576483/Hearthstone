package player

import (
	"hs/logic/iface"

	"github.com/name5566/leaf/gate"
)

type PlayerList struct {
	list map[gate.Agent]iface.IPlayer
}

var pl *PlayerList

// 获得对象
func GetPlayerList() iface.IPlayerList {
	if pl == nil {
		pl = &PlayerList{
			list: map[gate.Agent]iface.IPlayer{},
		}
	}

	return pl
}

// 获得房间
func (pl *PlayerList) GetPlayer(a gate.Agent) iface.IPlayer {

	if a == nil {
		return NewPlayer(nil)
	}

	if p, ok := pl.list[a]; ok {
		return p
	}

	pl.list[a] = NewPlayer(a)

	return pl.list[a]
}

// 删除玩家
func (pl *PlayerList) DeletePlayer(a gate.Agent) {
	delete(pl.list, a)
}
