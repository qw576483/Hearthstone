package bhero

import (
	"hs/logic/define"
	"hs/logic/iface"
)

func (h *Hero) RobotMove() {

	if h.gateAgnet != nil {
		return
	}

	b := h.GetBattle()
	if b.GetBattleStatus() == define.BattleStatusPre {
		b.PlayerChangePreCards(h.GetId(), make([]int, 0))
		return
	}

	if b.GetBattleStatus() != define.BattleStatusRun {
		return
	}

	if b.GetRoundHero().GetId() != h.GetId() {
		return
	}

	for i := 1; i <= 2; i++ {
		h.robotCheckRelease()
		h.robotCheckAttack()
	}

	b.PlayerRoundEnd(h.GetId())
}

func (h *Hero) robotCheckRelease() {

	// 手牌
	for _, v := range h.CardsToNewInstance(h.GetHandCards()) {

		// 如果战场满了，并且v还是随从的话
		if len(h.GetBattleCards()) >= define.MaxBattleNum && v.GetType() == define.CardTypeEntourage {
			continue
		}

		if v.GetHaveEffectMona() <= h.GetMona() {
			ecid := 0
			if ec := h.robotGetReleaseTargetCard(v); ec != nil {
				ecid = ec.GetId()
			}
			h.GetBattle().PlayerReleaseCard(h.GetId(), v.GetId(), 0, -1, ecid)
		}
	}

	// 英雄技能
	if h.GetHeroSkill().GetHaveEffectMona() <= h.GetMona() {
		v := h.GetHeroSkill()
		ecid := 0
		if ec := h.robotGetReleaseTargetCard(v); ec != nil {
			ecid = ec.GetId()
		}
		h.GetBattle().PlayerReleaseCard(h.GetId(), v.GetId(), 0, -1, ecid)
	}
}

func (h *Hero) robotGetReleaseTargetCard(ic iface.ICard) iface.ICard {

	icc := ic.GetConfig()
	if icc.ReleaseFilter == define.CardReleaseFilterNone {
		return nil
	}

	e := h.GetEnemy()
	var targetCard iface.ICard

	// 优先清理嘲讽
	for _, v := range e.GetBattleCards() {
		if v.IsHaveTraits(define.CardTraitsTaunt) && !v.IsHaveTraits(define.CardTraitsSneak) {
			targetCard = v
		}
	}

	switch icc.ReleaseFilter {
	case define.CardReleaseFilterAll:
		if icc.ReleaseIncrease {
			targetCard = h.GetHead()
		}
	case define.CardReleaseFilterBothHero:
		targetCard = h.GetEnemy().GetHead()
		if icc.ReleaseIncrease {
			targetCard = h.GetHead()
		}
	case define.CardReleaseFilterMyAll:
		targetCard = h.GetEnemy().GetHead()
	case define.CardReleaseFilterEnemyAll:
		if targetCard == nil {
			targetCard = h.GetEnemy().GetHead()
		}
	case define.CardReleaseFilterMyBattle:

		for _, v := range h.GetBattleCards() {
			targetCard = v
			break
		}
	case define.CardReleaseFilterEnemyBattle:

		if targetCard != nil {
			for _, v := range h.GetEnemy().GetBattleCards() {
				targetCard = v
				break
			}
		}
	case define.CardReleaseFilterBattle:

		if icc.ReleaseIncrease {
			for _, v := range h.GetBattleCards() {
				targetCard = v
				break
			}
		}

		if targetCard == nil {
			for _, v := range h.GetEnemy().GetBattleCards() {
				targetCard = v
				break
			}
		}
	}

	return targetCard
}

func (h *Hero) robotCheckAttack() {

	// 手牌
	for _, v := range h.CardsToNewInstance(h.GetBattleCards()) {

		ats := v.GetAttackTimes()
		mats := v.GetMaxAttackTimes()
		if ats >= mats {
			continue
		}

		// 如果是本局攻击
		if v.GetReleaseRound() == h.GetBattle().GetIncrRoundId() {
			if !v.IsHaveTraits(define.CardTraitsAssault) && !v.IsHaveTraits(define.CardTraitsSuddenStrike) {
				continue
			}
		}

		if v.IsHaveTraits(define.CardTraitsFrozen) {
			continue
		}

		ecid := 0
		if ec := h.robotGetAttackTargetCard(v); ec != nil {
			ecid = ec.GetId()
		}

		if ecid != 0 {
			h.GetBattle().PlayerConCardAttack(h.GetId(), v.GetId(), ecid)
		}

	}

	// 英雄攻击
	if h.GetHead().GetHaveEffectDamage() > 0 {

		v := h.GetHead()
		ats := v.GetAttackTimes()
		mats := v.GetMaxAttackTimes()

		if ats >= mats {
			return
		}

		if v.IsHaveTraits(define.CardTraitsFrozen) {
			return
		}

		ecid := 0
		if ec := h.robotGetAttackTargetCard(v); ec != nil {
			ecid = ec.GetId()
		}

		if ecid != 0 {
			h.GetBattle().PlayerConCardAttack(h.GetId(), v.GetId(), ecid)
		}
	}
}

func (h *Hero) robotGetAttackTargetCard(ic iface.ICard) iface.ICard {

	e := h.GetEnemy()
	var targetCard iface.ICard

	// 优先清理嘲讽
	for _, v := range e.GetBattleCards() {
		if v.IsHaveTraits(define.CardTraitsTaunt) && !v.IsHaveTraits(define.CardTraitsSneak) {
			targetCard = v
		}
	}

	if targetCard == nil {
		targetCard = e.GetHead()
	}

	return targetCard
}
