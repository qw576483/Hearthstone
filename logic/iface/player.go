package iface

import "github.com/name5566/leaf/gate"

type IPlayer interface {
	SetRoomId(int)            // 设置房间id
	GetRoomId() int           // 获得房间id
	GetGateAgent() gate.Agent // 获得用户连接
	GetHeroId() int           // 英雄id
	GetCardIds() []int        // 获得卡牌
	SetHc(int, []int)         // 设置卡牌
}
