package iface

type IRoom interface {
	AddToRoom(IPlayer, int) error
	GetBattle() IBattle
	Begin() // 房间战斗开始
	GetMembersNum() int
	GetPlayers() []IPlayer
}
