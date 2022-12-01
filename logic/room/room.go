package room

import (
	"errors"
	"hs/logic/battle"
	"hs/logic/heros"
	"hs/logic/iface"
)

type Room struct {
	id      int
	players []iface.IPlayer
	battle  iface.IBattle
}

// 房间
func NewRoom(id int) iface.IRoom {
	return &Room{
		id:      id,
		players: make([]iface.IPlayer, 0),
	}
}

// 添加到房间
func (r *Room) AddToRoom(p iface.IPlayer) error {

	if p.GetRoomId() != 0 {
		return errors.New("已有房间")
	}

	p.SetRoomId(r.id)
	r.players = append(r.players, p)

	r.Begin()

	return nil
}

// 房间战斗开始
func (r *Room) Begin() {

	if len(r.players) < 2 {
		return
	}

	if r.battle != nil {
		return
	}

	p1 := r.players[0]
	p2 := r.players[1]

	h1 := heros.GetHero(p1.GetHeroId())
	h2 := heros.GetHero(p2.GetHeroId())
	h1.SetGateAgent(p1.GetGateAgent())
	h2.SetGateAgent(p2.GetGateAgent())

	cards1 := iface.GetCardFact().GetCards(p1.GetCardIds())
	cards2 := iface.GetCardFact().GetCards(p2.GetCardIds())

	r.battle = battle.NewBattle(h1, h2, cards1, cards2)
}

// 获得成员数量
func (r *Room) GetMembersNum() int {
	return len(r.players)
}

// 获得战斗句柄
func (r *Room) GetBattle() iface.IBattle {
	return r.battle
}
