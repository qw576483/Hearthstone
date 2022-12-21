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

// 删除房间
func (rm *RoomList) DeleteByRoomId(rid int) {

	if rm, ok := rl.list[rid]; ok {
		ps := rm.GetPlayers()
		for _, v := range ps {
			v.SetRoomId(0)
		}
	}

	delete(rl.list, rid)
}
