package iface

type IRoomList interface {
	GetRoom(int) IRoom
	DeleteByRoomId(int)
}
