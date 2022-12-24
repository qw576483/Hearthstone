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
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	// m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	// log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}

func handleGetCardsConfig(args []interface{}) {
	a := args[1].(gate.Agent)

	// allConfig :=
	cacheConfig := make([]*config.CardConfig, 0)

	for _, v := range config.GetAllCardConfig() {
		cacheConfig = append(cacheConfig, v)
		if len(cacheConfig) >= 50 {
			a.WriteMsg(&push.CardsConfigMsg{
				Configs: cacheConfig,
			})

			cacheConfig = make([]*config.CardConfig, 0)
		}
	}

	if len(cacheConfig) >= 1 {
		a.WriteMsg(&push.CardsConfigMsg{
			Configs: cacheConfig,
		})
	}
}

func handleJoinRoom(args []interface{}) {
	m := args[0].(*msg.JoinRoom)
	a := args[1].(gate.Agent)

	hc := config.GetHeroConfig(m.HeroId)
	if hc == nil || !hc.CanCarry {
		a.WriteMsg(&push.ErrorMsg{
			Error: "英雄不存在:" + strconv.Itoa(m.HeroId),
		})
		return
	}

	// 玩家填写的cardIds
	var cardIds []int = make([]int, 0)
	cardIdsString := strings.Split(m.CardIds, ",")

	// 检查携带是否有效
	ac := config.GetAllCardConfig()
	for _, v := range cardIdsString {
		id, err := strconv.Atoi(v)
		if err == nil {
			if id >= len(ac) || id <= 0 {
				a.WriteMsg(&push.ErrorMsg{
					Error: "存在不能携带的卡牌:" + strconv.Itoa(id),
				})
				return
			}
			cardIds = append(cardIds, id)

			// cc := config.GetCardConfig(id)
			// if cc == nil || !cc.CanCarry {
			// 	a.WriteMsg(&push.ErrorMsg{
			// 		Error: "存在不能携带的卡牌:" + strconv.Itoa(id),
			// 	})
			// 	return
			// }

			// if len(cc.Vocation) > 0 && !help.InArray(hc.Vocation, cc.Vocation) {
			// 	a.WriteMsg(&push.ErrorMsg{
			// 		Error: "存在本职业不能携带的卡牌:" + strconv.Itoa(id),
			// 	})
			// 	return
			// }
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

	r.AddToRoom(p, m.Pve)
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

	err := b.PlayerUseHeroSkill(b.GetRoundHero().GetId(), m.ChoiceId, m.RCardId)

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
	err := b.PlayerReleaseCard(b.GetRoundHero().GetId(), m.CardId, m.ChoiceId, m.PutIdx, m.RCardId)

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

	err := b.PlayerConCardAttack(b.GetRoundHero().GetId(), m.CardId, m.ECardId)

	if err != nil {
		a.WriteMsg(&push.ErrorMsg{
			Error: err.Error(),
		})
	}
}
