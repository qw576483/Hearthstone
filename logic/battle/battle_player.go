package battle

import (
	"errors"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/iface"
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
func (b *Battle) PlayerReleaseCard(hid, cid, choiceId, putidx, rcid, rhid int) error {

	h := b.GetHeroByIncrId(hid)
	c := h.GetHandCardByIncrId(cid)

	if c == nil {
		return errors.New("没有找到目标")
	}

	// 费用
	if c.GetHaveEffectMona(c) > h.GetMona() {
		return errors.New("法力不足")
	}

	rh := b.GetHeroByIncrId(rhid)
	if rhid != 0 && rh == nil {
		return errors.New("没有找到目标")
	}

	h.CostMona(c.GetHaveEffectMona(c))

	// 拼接
	rc := h.GetBattleCardById(rcid)
	if rc == nil {
		rc = h.GetEnemy().GetBattleCardById(rcid)
	}

	if rcid != 0 {
		if rc == nil {
			return errors.New("没有找到目标")
		}

		if rc.IsHaveTraits(define.CardTraitsSneak, rc) {
			return errors.New("目标在潜行")
		}
	}

	// 检查是否能释放奥秘
	if c.GetConfig().Ctype == define.CardTypeSorcery &&
		c.IsHaveTraits(define.CardTraitsSecret, c) &&
		!h.CanReleaseSecret(c) {
		return errors.New("不能释放此奥秘")
	}

	// logs
	push.PushAutoLog(h, "打出了"+push.GetCardLogString(c))

	h.SetReleaseCardTimes(h.GetReleaseCardTimes() + 1)
	h.Release(c, choiceId, putidx, rc, rh, true)

	// info
	push.PushInfoMsg(b)

	return nil
}

// 使用英雄技能 - release
func (b *Battle) PlayerUseHeroSkill(hid, choiceId, rcid, rhid int) error {

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
	if c.GetHaveEffectMona(c) > h.GetMona() {
		return errors.New("法力不足")
	}

	rh := b.GetHeroByIncrId(rhid)
	if rhid != 0 && rh == nil {
		return errors.New("没有找到目标")
	}

	h.CostMona(c.GetHaveEffectMona(c))

	// 拼接
	rc := h.GetBattleCardById(rcid)
	if rc == nil {
		rc = h.GetEnemy().GetBattleCardById(rcid)
	}

	if rcid != 0 {
		if rc == nil {
			return errors.New("没有找到目标")
		}

		if rc.IsHaveTraits(define.CardTraitsSneak, rc) {
			return errors.New("目标在潜行")
		}
	}

	// logs
	push.PushAutoLog(h, "使用了英雄技能")

	c.SetAttackTimes(ats + 1)
	h.Release(c, choiceId, 0, rc, rh, true)

	// info
	push.PushInfoMsg(b)
	return nil
}

// 操作卡牌攻击 -  attack
func (b *Battle) PlayerConCardAttack(hid, cid, ecid, ehid int) error {

	h := b.GetHeroByIncrId(hid)
	c := h.GetBattleCardById(cid)
	if c == nil {
		return errors.New("没有找到此卡牌")
	}

	if c.IsHaveTraits(define.CardTraitsUnableToAttack, c) {
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

	// 如果是本局攻击
	if c.GetReleaseRound() == b.GetIncrRoundId() {
		// 如果不是冲锋 ， 不是突袭
		if !c.IsHaveTraits(define.CardTraitsAssault, c) && !c.IsHaveTraits(define.CardTraitsSuddenStrike, c) {
			return errors.New("卡牌在睡眠中")
		}

		// 突袭只能打英雄
		if c.IsHaveTraits(define.CardTraitsSuddenStrike, c) && ehid != 0 {
			return errors.New("突袭只能打英雄")
		}
	}

	var ec iface.ICard
	if ecid != 0 {
		ec = h.GetEnemy().GetBattleCardById(ecid)
		if ec == nil {
			return errors.New("没有找到此卡")
		}

		if ec.IsHaveTraits(define.CardTraitsSneak, ec) {
			return errors.New("目标在潜行")
		}
	}

	var eh iface.IHero
	if ehid != 0 {
		eh = b.GetHeroByIncrId(ehid)
		if eh == nil {
			return errors.New("没有找到此英雄")
		}
	}

	if eh != nil && eh.GetId() == h.GetId() {
		return errors.New("无效的敌人")
	}

	c.SetAttackTimes(ats + 1)
	h.Attack(c, ec, eh)

	push.PushInfoMsg(b)

	return nil
}

// 英雄攻击 - attack
func (b *Battle) PlayerAttack(hid, ecid, ehid int) error {

	h := b.GetHeroByIncrId(hid)

	// 检查次数和伤害
	ats := h.GetAttackTimes()
	mats := h.GetMaxAttackTimes()
	if ats >= mats {
		return errors.New("最大攻击次数了")
	}

	// fmt.Println(ats, mats)

	if h.GetDamage() <= 0 {
		return errors.New("无伤害，无法攻击")
	}

	// 检查是否有嘲讽
	tids := h.GetEnemy().GetBattleCardsTraitsTauntCardIds()
	if len(tids) > 0 && (ecid == 0 || !help.InArray(ecid, tids)) {
		return errors.New("必须先攻击拥有嘲讽的卡牌")
	}

	var ec iface.ICard
	if ecid != 0 {
		ec = h.GetEnemy().GetBattleCardById(ecid)
		if ec == nil {
			return errors.New("没有找到此卡")
		}

		if ec.IsHaveTraits(define.CardTraitsSneak, ec) {
			return errors.New("目标在潜行")
		}
	}

	var eh iface.IHero
	if ehid != 0 {
		eh = b.GetHeroByIncrId(ehid)
		if eh == nil {
			return errors.New("没有找到此英雄")
		}
	}

	h.SetAttackTimes(ats + 1)
	h.HAttack(ec, eh)

	push.PushInfoMsg(b)

	return nil
}

// 结束回合
func (b *Battle) PlayerRoundEnd(hid int) error {
	b.RoundEnd()
	return nil
}
