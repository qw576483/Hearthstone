package battle

import (
	"errors"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/push"
	"strconv"
	"strings"
)

// 修改预留卡牌
func (b *Battle) PlayerChangePreCards(hid int, putidxs []int /** 这个值是第几张卡*/) error {

	h := b.GetHeroByIncrId(hid)

	// 检查sign
	sign := strconv.Itoa(h.GetId()) + "changePreCards"
	if b.GetDoneSign(sign) != "" {
		return errors.New("已经选择过了初始卡牌")
	}
	b.SetDoneSign(sign, "done")

	// 更换卡牌
	h.ChangePreCrards(putidxs)

	// log
	if len(putidxs) <= 0 {
		push.PushAutoLog(h, "没有更换任何卡牌")
	} else {
		log := "【你的对手】更换了"
		for _, v := range putidxs {
			log += "第" + strconv.Itoa(v+1) + "张,"
		}
		push.PushLog(h.GetEnemy(), strings.Trim(log, ",")+"卡牌")
	}

	// 如果双方都选择完预留卡牌
	sign2 := strconv.Itoa(h.GetEnemy().GetId()) + "changePreCards"
	if b.GetDoneSign(sign2) != "" {
		b.Begin()
	}

	return nil
}

// 释放卡牌 - release
func (b *Battle) PlayerReleaseCard(hid, cid, choiceId, putidx, rcid int) error {

	h := b.GetHeroByIncrId(hid)
	c := h.GetHandCardByIncrId(cid)

	if c == nil {
		return errors.New("没有找到目标")
	}

	// 费用
	if c.GetHaveEffectMona() > h.GetMona() {
		return errors.New("法力不足")
	}

	// 拼接
	rc := h.GetCanSelectCardId(rcid)
	if rc == nil {
		rc = h.GetEnemy().GetCanSelectCardId(rcid)
	}

	if err := b.checkCanRelease(c, rcid, rc); err != nil {
		return err
	}

	// 检查是否能释放奥秘
	if c.GetConfig().Ctype == define.CardTypeSorcery &&
		c.IsHaveTraits(define.CardTraitsSecret) &&
		!h.CanReleaseSecret(c) {
		return errors.New("不能释放此奥秘")
	}

	h.CostMona(c.GetHaveEffectMona())

	// logs
	push.PushAutoLog(h, "打出了"+push.GetCardLogString(c))

	h.SetReleaseCardTimes(h.GetReleaseCardTimes() + 1)
	h.Release(c, choiceId, putidx, rc, true)

	// info
	push.PushInfoMsg(b)

	return nil
}

// 使用英雄技能 - release
func (b *Battle) PlayerUseHeroSkill(hid, choiceId, rcid int) error {

	h := b.GetHeroByIncrId(hid)
	c := h.GetHeroSkill()

	if c == nil {
		return errors.New("没找到英雄技能")
	}

	// 检查次数
	ats := c.GetAttackTimes()
	mats := c.GetMaxAttackTimes()
	if ats >= mats {
		return errors.New("最大攻击次数了")
	}

	// 费用
	if c.GetHaveEffectMona() > h.GetMona() {
		return errors.New("法力不足")
	}

	rc := h.GetCanSelectCardId(rcid)
	if rc == nil {
		rc = h.GetEnemy().GetCanSelectCardId(rcid)
	}

	if err := b.checkCanRelease(c, rcid, rc); err != nil {
		return err
	}

	h.CostMona(c.GetHaveEffectMona())

	// logs
	push.PushAutoLog(h, "使用了英雄技能")

	c.SetAttackTimes(ats + 1)
	h.Release(c, choiceId, 0, rc, true)

	// info
	push.PushInfoMsg(b)
	return nil
}

// 操作卡牌攻击 -  attack
func (b *Battle) PlayerConCardAttack(hid, cid, ecid int) error {

	if ecid == 0 {
		return errors.New("请输入敌人")
	}

	h := b.GetHeroByIncrId(hid)
	c := h.GetCanSelectCardId(cid)
	if c == nil {
		return errors.New("没有找到此卡牌")
	}

	if c.GetType() != define.CardTypeHero && c.GetType() != define.CardTypeEntourage {
		return errors.New("进攻者无效")
	}

	if c.IsHaveTraits(define.CardTraitsUnableToAttack) {
		return errors.New("此卡牌无法攻击")
	}

	// 检查次数
	ats := c.GetAttackTimes()
	mats := c.GetMaxAttackTimes()
	if ats >= mats {
		return errors.New("最大攻击次数了")
	}

	// 检查是否有嘲讽
	tids := h.GetEnemy().GetBattleCardsTraitsTauntCardIds()
	if len(tids) > 0 && (ecid == 0 || !help.InArray(ecid, tids)) {
		return errors.New("必须先攻击拥有嘲讽的卡牌")
	}

	ec := h.GetEnemy().GetCanSelectCardId(ecid)
	if err := b.checkCanAttack(ecid, ec); err != nil {
		return err
	}

	// 如果是本局攻击
	if c.GetReleaseRound() == b.GetIncrRoundId() {
		// 如果不是冲锋 ， 不是突袭
		if !c.IsHaveTraits(define.CardTraitsAssault) && !c.IsHaveTraits(define.CardTraitsSuddenStrike) {
			return errors.New("卡牌在睡眠中")
		}

		if c.IsHaveTraits(define.CardTraitsSuddenStrike) && ec.GetCardInCardsPos() == define.InCardsTypeHead {
			return errors.New("突袭不能进攻英雄")
		}
	}

	c.SetAttackTimes(ats + 1)
	h.Attack(c, ec)

	push.PushInfoMsg(b)

	return nil
}

// 结束回合
func (b *Battle) PlayerRoundEnd(hid int) error {
	b.RoundEnd()
	return nil
}
