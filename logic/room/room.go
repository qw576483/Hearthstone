package room

import (
	"errors"
	"hs/logic/battle"
	"hs/logic/battle/bhero"
	"hs/logic/config"
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/player"
	"hs/logic/push"
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
func (r *Room) AddToRoom(p iface.IPlayer, pve int) error {

	if p.GetRoomId() != 0 {
		return errors.New("已有房间")
	}

	p.SetRoomId(r.id)
	r.players = append(r.players, p)

	p.GetGateAgent().WriteMsg(&push.LineMsg{
		Line: 999,
	})

	if pve != 0 && len(r.players) == 1 {
		r.AddPveRobotToRoom(pve)
	}

	r.Begin()

	return nil
}

func (r *Room) AddPveRobotToRoom(pve int) {
	p := player.GetPlayerList().GetPlayer(nil)
	p.SetHc(6, []int{204, 204, 253, 253, 18, 18, 44, 62, 99, 135, 135, 161, 161, 223, 222, 312, 312, 334, 334, 109, 109, 77, 43, 30, 16, 30, 16})

	r.players = append(r.players, p)
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

	h1 := &bhero.Hero{}
	h2 := &bhero.Hero{}
	h1.SetConfig(config.GetHeroConfig(p1.GetHeroId()))
	h1.SetGateAgent(p1.GetGateAgent())
	h2.SetConfig(config.GetHeroConfig(p2.GetHeroId()))
	h2.SetGateAgent(p2.GetGateAgent())

	cards1 := iface.GetCardFact().GetCards(p1.GetCardIds())
	cards2 := iface.GetCardFact().GetCards(p2.GetCardIds())

	r.battle = battle.NewBattle(h1, h2, cards1, cards2)

	go func() {
		sc := r.battle.GetStatusChan()
		for {
			if define.BattleStatusEnd == <-sc {
				GetRoomList().DeleteByRoomId(r.id)
				break
			}
		}
	}()
}

// 获得成员数量
func (r *Room) GetMembersNum() int {
	return len(r.players)
}

// 获得战斗句柄
func (r *Room) GetBattle() iface.IBattle {
	return r.battle
}

func (r *Room) GetPlayers() []iface.IPlayer {
	return r.players
}
