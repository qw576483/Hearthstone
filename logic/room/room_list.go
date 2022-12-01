package room

import "hs/logic/iface"

type RoomList struct {
	list map[int]iface.IRoom
}

var rl *RoomList

// 获得对象
func GetRoomList() iface.IRoomList {
	if rl == nil {
		rl = &RoomList{
			list: map[int]iface.IRoom{},
		}
	}

	return rl
}

// 获得房间
func (rm *RoomList) GetRoom(rid int) iface.IRoom {
	if r, ok := rl.list[rid]; ok {
		return r
	}

	rl.list[rid] = NewRoom(rid)

	return rl.list[rid]
}
