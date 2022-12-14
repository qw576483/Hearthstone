package battle

import (
	"errors"
	"hs/logic/define"
	"hs/logic/iface"
)

func (b *Battle) checkCanRelease(c iface.ICard, rcid int, rc iface.ICard) error {

	if c.GetType() == define.CardTypeEntourage && len(c.GetOwner().GetBattleCards()) >= define.MaxBattleNum {
		return errors.New("随从满了")
	}

	if rcid != 0 {
		if rc == nil {
			return errors.New("没有找到目标")
		}

		if c.GetOwner().GetId() != rc.GetOwner().GetId() && rc.IsHaveTraits(define.CardTraitsSneak) {
			return errors.New("目标在潜行")
		}

		if (c.GetType() == define.CardTypeSorcery || c.GetType() == define.CardTypeHeroSkill) && rc.IsHaveTraits(define.CardTraitsMagicImmunity) {
			return errors.New("目标魔法免疫")
		}
	}

	conf := c.GetConfig()

	if conf.ReleaseFilter == define.CardReleaseFilterMyAll {
		if rc != nil && rc.GetOwner().GetId() != c.GetOwner().GetId() {
			return errors.New("必须以我方单位为目标")
		}
	} else if conf.ReleaseFilter == define.CardReleaseFilterEnemyAll {
		if rc != nil && rc.GetOwner().GetId() == c.GetOwner().GetId() {
			return errors.New("必须以敌方单位为目标")
		}
	} else if conf.ReleaseFilter == define.CardReleaseFilterMyBattle {
		if rc != nil && rc.GetOwner().GetId() != c.GetOwner().GetId() {
			return errors.New("必须以我方随从为目标")
		}

		if rc != nil && rc.GetCardInCardsPos() != define.InCardsTypeBattle {
			return errors.New("目标不在战场")
		}
	} else if conf.ReleaseFilter == define.CardReleaseFilterEnemyBattle {
		if rc != nil && rc.GetOwner().GetId() == c.GetOwner().GetId() {
			return errors.New("必须以敌方随从为目标")
		}

		if rc != nil && rc.GetCardInCardsPos() != define.InCardsTypeBattle {
			return errors.New("目标不在战场")
		}
	} else if conf.ReleaseFilter == define.CardReleaseFilterBattle {

		if rc != nil && rc.GetCardInCardsPos() != define.InCardsTypeBattle {
			return errors.New("目标不在战场")
		}
	}

	return nil
}

func (b *Battle) checkCanAttack(rcid int, rc iface.ICard) error {

	if rc == nil {
		return errors.New("没有找到目标")
	}

	if rc.GetType() != define.CardTypeEntourage && rc.GetType() != define.CardTypeHero {
		return errors.New("目标错误")
	}

	if rc.IsHaveTraits(define.CardTraitsSneak) {
		return errors.New("目标在潜行")
	}

	return nil
}
