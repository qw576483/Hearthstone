package iface

import "github.com/name5566/leaf/gate"

type IPlayerList interface {
	GetPlayer(gate.Agent) IPlayer
}
