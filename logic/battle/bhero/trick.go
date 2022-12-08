package bhero

import (
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
)

// 触发战斗开始
func (h *Hero) TrickBattleBegin() {
	for _, v := range h.GetBothAllCards() {
		v.OnBattleBegin()
	}
}

// 触发战吼
func (h *Hero) TrickRelease(c iface.ICard, choiceId, pidx int, rc iface.ICard, rh iface.IHero) {
	c.OnRelease(choiceId, pidx, rc, rh)
}

// 触发回合开始
func (h *Hero) TrickRoundBegin() {
	for _, v := range h.GetBothEventCards("OnNRRoundBegin") {
		v.OnNRRoundBegin()
	}
}

// 触发回合结束
func (h *Hero) TrickRoundEnd() {
	for _, v := range h.GetBothEventCards("OnNRRoundEnd") {
		v.OnNRRoundEnd()
	}
}

// 触发得到事件
func (h *Hero) TrickGetCardEvent(c iface.ICard) {
	c.OnGet()
}

// 触发销毁事件
func (h *Hero) TrickDevastateCardEvent(c iface.ICard) {
	c.OnDevastate()
}

// 触发攻击后事件
func (h *Hero) TrickAfterAttackEvent(c, ec iface.ICard, eh iface.IHero, trueCostHp int) {

	// 攻击者事件
	if trueCostHp > 0 {
		if ec != nil {
			if ec.GetHaveEffectHp() == 0 && trueCostHp > 0 && !c.IsSilent() {
				c.OnHonorAnnihilate()
			} else if ec.GetHaveEffectHp() < 0 && !c.IsSilent() {
				c.OnOverflowAnnihilate()
			} else if ec.GetHaveEffectHp() > 0 && c.IsHaveTraits(define.CardTraitsHighlyToxic) {
				push.PushAutoLog(h, push.GetCardLogString(c)+" 触发剧毒，"+push.GetCardLogString(ec)+"直接死亡")
				ec.GetOwner().DieCard(ec)
			}
		} else if eh != nil {
			if eh.GetHp() == 0 && !c.IsSilent() {
				c.OnHonorAnnihilate()
			} else if eh.GetHp() < 0 && !c.IsSilent() {
				c.OnOverflowAnnihilate()
			}
		}
	}

}

// 触发死亡事件
func (h *Hero) TrickDieCardEvent(c iface.ICard, bidx int) {

	if !c.IsSilent() {
		c.OnDie(bidx)
	}

	for _, v := range h.GetBothEventCards("OnNROtherDie") {
		v.OnNROtherDie(c)
	}
}

// 触发步入战场事件
func (h *Hero) TrickPutToBattleEvent(c iface.ICard, bidx int) {

	// 有可能是复制出来的卡，然后put to battle，也需要检查是否沉默
	if !c.IsSilent() {
		c.OnPutToBattle(bidx)
	}

	for _, v := range h.GetBothEventCards("OnNRPutToBattle") {
		v.OnNRPutToBattle(c)
	}
}

// 触发离开战场事件
func (h *Hero) TrickOutBattleEvent(c iface.ICard) {

	if !c.IsSilent() {
		c.OnOutBattle()
	}
}
