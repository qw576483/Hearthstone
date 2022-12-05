package internal

import (
	"hs/logic/config"
	"hs/logic/define"
	"hs/logic/player"
	"hs/logic/push"
	"hs/logic/room"
	"hs/net/msg"
	"reflect"
	"strconv"
	"strings"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Hello{}, handleHello)
	handler(&msg.GetCardsConfig{}, handleGetCardsConfig)
	handler(&msg.JoinRoom{}, handleJoinRoom)

	handler(&msg.BChangePre{}, handleBChangePre)
	handler(&msg.BEndRound{}, handleBEndRound)
	handler(&msg.BUseSkill{}, handleBUseSkill)
	handler(&msg.BRelease{}, handleBRelease)
	handler(&msg.BAttack{}, handleBAttack)
	handler(&msg.BHAttack{}, handleBHAttack)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}

func handleGetCardsConfig(args []interface{}) {
	a := args[1].(gate.Agent)
	a.WriteMsg(&push.CardsConfigMsg{
		Configs: config.GetAllCardConfig(),
	})
}

func handleJoinRoom(args []interface{}) {
	m := args[0].(*msg.JoinRoom)
	a := args[1].(gate.Agent)

	// 玩家填写的cardIds
	var cardIds []int = make([]int, 0)
	cardIdsString := strings.Split(m.CardIds, ",")

	for _, v := range cardIdsString {
		id, err := strconv.Atoi(v)
		if err == nil {
			cardIds = append(cardIds, id)
		}
	}

	p := player.GetPlayerList().GetPlayer(a)
	p.SetHc(m.HeroId, cardIds)

	if p.GetRoomId() != 0 {
		a.WriteMsg(&push.ErrorMsg{
			Error: "我已经加入房间了",
		})
		return
	}

	r := room.GetRoomList().GetRoom(m.RoomId)

	if r.GetMembersNum() >= 2 {
		a.WriteMsg(&push.ErrorMsg{
			Error: "房间已满员",
		})
		return
	}

	r.AddToRoom(p)
}

func handleBChangePre(args []interface{}) {
	m := args[0].(*msg.BChangePre)
	a := args[1].(gate.Agent)

	var indexs []int = make([]int, 0)
	indexsString := strings.Split(m.Indexs, ",")

	for _, v := range indexsString {
		id, err := strconv.Atoi(v)
		if err == nil {
			if (id-1 < 0) || (id-1 > 3) {
				continue
			}
			indexs = append(indexs, id-1)
		}
	}

	p := player.GetPlayerList().GetPlayer(a)
	r := room.GetRoomList().GetRoom(p.GetRoomId())
	b := r.GetBattle()

	if b == nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: "战斗无效",
		})
		return
	}

	if b.GetBattleStatus() != define.BattleStatusPre {
		a.WriteMsg(&push.ErrorMsg{
			Error: "当前不是换牌环节",
		})
		return
	}

	h := b.GetHeroByGateAgent(a)

	err := b.PlayerChangePreCards(h.GetId(), indexs)

	if err != nil {
		push.PushError(h, err.Error())
	}
}

func handleBEndRound(args []interface{}) {
	// m := args[0].(*msg.BChangePre)
	a := args[1].(gate.Agent)

	p := player.GetPlayerList().GetPlayer(a)
	r := room.GetRoomList().GetRoom(p.GetRoomId())
	b := r.GetBattle()

	if b == nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: "战斗无效",
		})
		return
	}

	if b.GetBattleStatus() != define.BattleStatusRun {
		a.WriteMsg(&push.ErrorMsg{
			Error: "当前不是战斗环节",
		})
		return
	}

	if b.GetHeroByGateAgent(a) != b.GetRoundHero() {
		a.WriteMsg(&push.ErrorMsg{
			Error: "不是我的出手回合",
		})
		return
	}
	b.PlayerRoundEnd(b.GetHeroByGateAgent(a).GetId())
}

func handleBUseSkill(args []interface{}) {
	m := args[0].(*msg.BUseSkill)
	a := args[1].(gate.Agent)

	p := player.GetPlayerList().GetPlayer(a)
	r := room.GetRoomList().GetRoom(p.GetRoomId())
	b := r.GetBattle()

	if b == nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: "战斗无效",
		})
		return
	}

	if b.GetBattleStatus() != define.BattleStatusRun {
		a.WriteMsg(&push.ErrorMsg{
			Error: "当前不是战斗环节",
		})
		return
	}

	if b.GetHeroByGateAgent(a) != b.GetRoundHero() {
		a.WriteMsg(&push.ErrorMsg{
			Error: "不是我的出手回合",
		})
		return
	}

	err := b.PlayerUseHeroSkill(b.GetRoundHero().GetId(), m.ChoiceId, m.RCardId, m.RHeroId)

	if err != nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: err.Error(),
		})
	}
}

func handleBRelease(args []interface{}) {
	m := args[0].(*msg.BRelease)
	a := args[1].(gate.Agent)

	p := player.GetPlayerList().GetPlayer(a)
	r := room.GetRoomList().GetRoom(p.GetRoomId())
	b := r.GetBattle()

	if b == nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: "战斗无效",
		})
		return
	}

	if b.GetBattleStatus() != define.BattleStatusRun {
		a.WriteMsg(&push.ErrorMsg{
			Error: "当前不是战斗环节",
		})
		return
	}

	if b.GetHeroByGateAgent(a) != b.GetRoundHero() {
		a.WriteMsg(&push.ErrorMsg{
			Error: "不是我的出手回合",
		})
		return
	}

	m.PutIdx -= 1
	err := b.PlayerReleaseCard(b.GetRoundHero().GetId(), m.CardId, m.ChoiceId, m.PutIdx, m.RCardId, m.RHeroId)

	if err != nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: err.Error(),
		})
	}

}

func handleBAttack(args []interface{}) {
	m := args[0].(*msg.BAttack)
	a := args[1].(gate.Agent)

	p := player.GetPlayerList().GetPlayer(a)
	r := room.GetRoomList().GetRoom(p.GetRoomId())
	b := r.GetBattle()

	if b == nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: "战斗无效",
		})
		return
	}

	if b.GetBattleStatus() != define.BattleStatusRun {
		a.WriteMsg(&push.ErrorMsg{
			Error: "当前不是战斗环节",
		})
		return
	}

	if b.GetHeroByGateAgent(a) != b.GetRoundHero() {
		a.WriteMsg(&push.ErrorMsg{
			Error: "不是我的出手回合",
		})
		return
	}

	err := b.PlayerConCardAttack(b.GetRoundHero().GetId(), m.CardId, m.ECardId, m.EHeroId)

	if err != nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: err.Error(),
		})
	}
}

func handleBHAttack(args []interface{}) {
	m := args[0].(*msg.BHAttack)
	a := args[1].(gate.Agent)

	p := player.GetPlayerList().GetPlayer(a)
	r := room.GetRoomList().GetRoom(p.GetRoomId())
	b := r.GetBattle()

	if b == nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: "战斗无效",
		})
		return
	}

	if b.GetBattleStatus() != define.BattleStatusRun {
		a.WriteMsg(&push.ErrorMsg{
			Error: "当前不是战斗环节",
		})
		return
	}

	if b.GetHeroByGateAgent(a) != b.GetRoundHero() {
		a.WriteMsg(&push.ErrorMsg{
			Error: "不是我的出手回合",
		})
		return
	}

	err := b.PlayerAttack(b.GetRoundHero().GetId(), m.ECardId, m.EHeroId)

	if err != nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: err.Error(),
		})
	}
}
