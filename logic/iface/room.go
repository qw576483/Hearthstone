package iface

type IRoom interface {
	AddToRoom(IPlayer) error
	GetBattle() IBattle
	Begin() // 房间战斗开始
	GetMembersNum() int
	GetPlayers() []IPlayer
}
