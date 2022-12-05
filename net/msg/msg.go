package msg

import (
	"hs/logic/push"

	"github.com/name5566/leaf/network/json"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	// 操作
	Processor.Register(&Hello{})
	Processor.Register(&GetCardsConfig{})
	Processor.Register(&JoinRoom{})
	Processor.Register(&BChangePre{})
	Processor.Register(&BEndRound{})
	Processor.Register(&BUseSkill{})
	Processor.Register(&BRelease{})
	Processor.Register(&BAttack{})
	Processor.Register(&BHAttack{})

	// push
	Processor.Register(&push.CardsConfigMsg{})
	Processor.Register(&push.LineMsg{})
	Processor.Register(&push.InitMsg{})
	Processor.Register(&push.InfoMsg{})
	Processor.Register(&push.PreCardsMsg{})
	Processor.Register(&push.ErrorMsg{})
	Processor.Register(&push.LogMsg{})
}
