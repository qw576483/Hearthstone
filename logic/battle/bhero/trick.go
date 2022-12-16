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

// 触发战吼
func (h *Hero) TrickRelease(c iface.ICard, choiceId, bidx int, rc iface.ICard) {
	c.OnRelease(choiceId, bidx, rc)

	h.GetBattle().WhileTrickCardDie()
}

// 触发战吼2
func (h *Hero) TrickRelease2(c iface.ICard, choiceId, bidx int, rc iface.ICard) {
	c.OnRelease2(choiceId, bidx, rc)

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

// 触发攻击前事件
func (h *Hero) TrickBeforeAttackEvent(c, ec iface.ICard) iface.ICard {

	if c.GetType() == define.CardTypeHero && c.GetOwner().GetWeapon() != nil && !c.GetOwner().GetWeapon().IsSilent() {
		ec = c.GetOwner().GetWeapon().OnBeforeAttack(ec)
	} else if c.GetType() == define.CardTypeEntourage && !c.IsSilent() {
		ec = c.OnBeforeAttack(ec)
	}

	for _, v := range h.GetBattle().GetEventCards("OnNROtherBeforeAttack") {
		ec = v.OnNROtherBeforeAttack(c, ec)
	}

	return ec
}

// 触发攻击后事件
func (h *Hero) TrickAfterAttackEvent(c, ec iface.ICard, trueCostHp int) {

	if c.GetType() == define.CardTypeHero && c.GetOwner().GetWeapon() != nil && !c.GetOwner().GetWeapon().IsSilent() {
		c.GetOwner().GetWeapon().OnAfterAttack(ec)
	} else if c.GetType() == define.CardTypeEntourage && !c.IsSilent() {
		c.OnAfterAttack(ec)
	}

	// 攻击者事件
	if trueCostHp > 0 {
		if ec.GetHaveEffectHp() > 0 && c.IsHaveTraits(define.CardTraitsHighlyToxic) && ec.GetType() != define.CardTypeHero {
			push.PushAutoLog(h, push.GetCardLogString(c)+" 触发剧毒，"+push.GetCardLogString(ec)+"直接死亡")
			ec.GetOwner().DieCard(ec, false)
		}
	}

}

// 触发死亡事件
func (h *Hero) TrickDieCardEvent(c iface.ICard) {

	if !c.IsSilent() {
		c.OnDie()
	}

	for _, v := range c.GetAddOnDie() {
		v()
	}

	for _, v := range c.GetSubCards() {
		for _, v2 := range v.GetAddOnDie() {
			v2()
		}
	}

	for _, v := range h.GetBattle().GetEventCards("OnNROtherDie") {
		v.OnNROtherDie(c)
	}
}
