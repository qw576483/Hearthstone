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

	h.GetBattle().WhileTrickCardDie()
}

// 触发得到事件
func (h *Hero) TrickGetCardEvent(c iface.ICard) {
	c.OnGet()

	h.GetBattle().WhileTrickCardDie()
}

// 触发战吼
func (h *Hero) TrickRelease(c iface.ICard, choiceId, bidx int, rc iface.ICard, rh iface.IHero) {
	c.OnRelease(choiceId, bidx, rc, rh)

	h.GetBattle().WhileTrickCardDie()
}

// 触发回合开始
func (h *Hero) TrickRoundBegin() {
	for _, v := range h.GetBattle().GetEventCards("OnNRRoundBegin") {
		v.OnNRRoundBegin()
	}

	h.GetBattle().WhileTrickCardDie()
}

// 触发回合结束
func (h *Hero) TrickRoundEnd() {
	for _, v := range h.GetBattle().GetEventCards("OnNRRoundEnd") {
		v.OnNRRoundEnd()
	}

	h.GetBattle().WhileTrickCardDie()
}

// 触发步入战场事件
func (h *Hero) TrickPutToBattleEvent(c iface.ICard, bidx int) {

	// 有可能是复制出来的卡，然后put to battle，也需要检查是否沉默
	if !c.IsSilent() {
		c.OnPutToBattle(bidx)
	}

	for _, v := range h.GetBattle().GetEventCards("OnNRPutToBattle") {
		v.OnNRPutToBattle(c)
	}

	h.GetBattle().WhileTrickCardDie()
}

// 触发离开战场事件
func (h *Hero) TrickOutBattleEvent(c iface.ICard) {

	if !c.IsSilent() {
		c.OnOutBattle()
	}
	h.GetBattle().WhileTrickCardDie()
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
				ec.GetOwner().DieCard(ec, false)
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
func (h *Hero) TrickDieCardEvent(c iface.ICard) {

	if !c.IsSilent() {
		c.OnDie()
	}

	for _, v := range h.GetBattle().GetEventCards("OnNROtherDie") {
		v.OnNROtherDie(c)
	}
}
