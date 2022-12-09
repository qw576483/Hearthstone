package bhero

// 预备阶段
func (h *Hero) PreBegin() {
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
	h.SetAttackTimes(0)

	// 重置卡牌次数
	h.SetReleaseCardTimes(0)

	// 锁定法力值
	h.SetLockMona(h.GetLockMonaCache())
	h.SetLockMonaCache(0)

	// 抽卡
	h.DrawByTimes(1)
	h.TrickRoundBegin()
}

// 回合结束
func (h *Hero) RoundEnd() {

	h.TrickRoundEnd()
}
