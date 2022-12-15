package bhero

import (
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
)

// 预备阶段
func (h *Hero) PreBegin() {
	h.NewCountDown(define.BattleTime)
	h.DrawForPreBegin(4)
}

// 回合开始
func (h *Hero) RoundBegin() {

	h.AddMonaMax(1)
	h.SetMona(h.GetMonaMax())

	// 重置攻击次数
	bcs := h.GetBattleCards()
	for _, v := range bcs {
		v.SetAttackTimes(0)
	}
	h.GetHeroSkill().SetAttackTimes(0)
	h.GetHead().SetAttackTimes(0)

	// 重置卡牌次数
	h.SetReleaseCardTimes(0)

	// 重置回合死亡
	h.roundDieCards = make([]iface.ICard, 0)

	// 锁定法力值
	h.SetLockMona(h.GetLockMonaCache())
	h.SetLockMonaCache(0)

	// 抽卡
	h.DrawByTimes(1)
	h.TrickRoundBegin()

	// 设置时间
	bt := define.BattleTime
	for _, v := range h.GetBattle().GetEventCards("OnNRGetBattleTime") {
		bt = v.OnNRGetBattleTime(bt)
	}
	h.NewCountDown(bt)
}

// 回合结束
func (h *Hero) RoundEnd() {
	h.CloseCountDown()
	h.TrickRoundEnd()

	// 检查冻结
	bcs := h.GetBattleCards()
	for _, v := range bcs {
		v.CheckFrozen()
	}
	h.GetHead().CheckFrozen()

	// 重置回合死亡
	h.roundDieCards = make([]iface.ICard, 0)
}

// 立即结束回合
func (h *Hero) FixRoundEnd() {

	push.PushAutoLog(h, "超时！强制结束回合！")
	b := h.GetBattle()

	if b.GetBattleStatus() == define.BattleStatusPre {
		b.PlayerChangePreCards(h.GetId(), make([]int, 0))
	} else if b.GetBattleStatus() == define.BattleStatusRun {
		b.PlayerRoundEnd(h.GetId())
	}
}
