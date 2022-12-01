package gate

import (
	"hs/net/game"
	"hs/net/msg"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GetCardsConfig{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.JoinRoom{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.BChangePre{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.BEndRound{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.BRelease{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.BAttack{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.BHAttack{}, game.ChanRPC)
}
