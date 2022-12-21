package internal

import (
	"hs/logic/define"
	"hs/logic/player"
	"hs/logic/push"
	"hs/logic/room"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	// 如果玩家断链
	p := player.GetPlayerList().GetPlayer(a)
	roomId := p.GetRoomId()
	if roomId != 0 {

		r := room.GetRoomList().GetRoom(roomId)
		b := r.GetBattle()
		if b != nil {
			push.PushAllLog(b, push.GetHeroLogString(b.GetHeroByGateAgent(a))+"离开了游戏")
			b.SetBattleStatus(define.BattleStatusEnd)
		} else {
			room.GetRoomList().DeleteByRoomId(roomId)
		}
	}
	player.GetPlayerList().DeletePlayer(a)
}
