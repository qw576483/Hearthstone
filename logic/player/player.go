package player

import (
	"hs/logic/iface"

	"github.com/name5566/leaf/gate"
)

type Player struct {
	id        int        // 玩家id
	roomId    int        // 房间id
	gateAgnet gate.Agent // 连接
	heroId    int        // 英雄id
	cardIds   []int      // 卡牌ids
}

var IncrPlayerId int

// 玩家
func NewPlayer(a gate.Agent) iface.IPlayer {
	IncrPlayerId += 1
	return &Player{
		id:        IncrPlayerId,
		roomId:    0,
		gateAgnet: a,
	}
}

// 设置房间id
func (p *Player) SetRoomId(rid int) {
	p.roomId = rid
}

// 获得房间id
func (p *Player) GetRoomId() int {
	return p.roomId
}

// 获得用户连接
func (p *Player) GetGateAgent() gate.Agent {
	return p.gateAgnet
}

// 获得用户连接
func (p *Player) GetHeroId() int {
	return p.heroId
}

// 获得卡牌
func (p *Player) GetCardIds() []int {
	return p.cardIds
}

// 设置卡牌和英雄
func (p *Player) SetHc(heroId int, cardIds []int) {
	p.heroId = heroId
	p.cardIds = cardIds
}
