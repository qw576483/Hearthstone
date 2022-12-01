package battle

import (
	"errors"
	"fmt"
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

// 释放卡牌
func (b *Battle) PlayerReleaseCard(hid, cid, choiceId, putidx, rcid, rhid int) error {

	h := b.GetHeroByIncrId(hid)
	card := h.GetHandCardByIncrId(cid)

	if card == nil {
		return errors.New("没有找到目标")
	}

	// 费用
	if card.GetMona() > h.GetMona() {
		return errors.New("法力不足")
	}
	h.CostMona(card.GetMona())

	rh := b.GetHeroByIncrId(rhid)
	if rhid != 0 && rh == nil {
		return errors.New("没有找到目标")
	}

	// 拼接
	rc := h.GetBattleCardById(rcid)
	if rc == nil {
		rc = h.GetEnemy().GetBattleCardById(rcid)
	}

	h.Release(card, choiceId, putidx, rc, rh, true)

	// logs
	push.PushAutoLog(h, "打出了"+push.GetCardLogString(card))
	push.PushInfoMsg(b)

	return nil
}

// 操作卡牌
func (b *Battle) PlayerConCardAttack(hid, cid, ecid, ehid int) error {

	h := b.GetHeroByIncrId(hid)
	c := h.GetBattleCardById(cid)
	if c == nil {
		return errors.New("没有找到此卡牌")
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
	ebcs := h.GetEnemy().GetBattleCards()
	var tt []int = make([]int, 0)
	for _, v := range ebcs {
		if v.IsHaveTraits(define.CardTraitsTaunt) {
			tt = append(tt, v.GetId())
		}
	}
	if len(tt) > 0 && (ecid == 0 || !help.InArray(ecid, tt)) {
		return errors.New("必须先攻击拥有嘲讽的卡牌")
	}

	// 如果是本局攻击
	if c.GetReleaseRound() == b.GetIncrRoundId() {
		// 如果不是冲锋 ， 不是突袭
		if !c.IsHaveTraits(define.CardTraitsAssault) && !c.IsHaveTraits(define.CardTraitsSuddenStrike) {
			return errors.New("卡牌在睡眠中")
		}

		// 突袭只能打英雄
		if c.IsHaveTraits(define.CardTraitsSuddenStrike) && ehid != 0 {
			return errors.New("突袭只能打英雄")
		}
	}

	var ec iface.ICard
	if ecid != 0 {
		ec = h.GetEnemy().GetBattleCardById(ecid)
		if ec == nil {
			return errors.New("没有找到此卡")
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

func (b *Battle) PlayerAttack(hid, ecid, ehid int) error {

	h := b.GetHeroByIncrId(hid)

	// 检查次数和伤害
	ats := h.GetAttackTimes()
	mats := h.GetMaxAttackTimes()
	if ats >= mats {
		return errors.New("最大攻击次数了")
	}

	fmt.Println(ats, mats)

	if h.GetDamage() <= 0 {
		return errors.New("无伤害，无法攻击")
	}

	// 检查是否有嘲讽
	ebcs := h.GetEnemy().GetBattleCards()
	var tt []int = make([]int, 0)
	for _, v := range ebcs {
		if v.IsHaveTraits(define.CardTraitsTaunt) {
			tt = append(tt, v.GetId())
		}
	}
	if len(tt) > 0 && (ecid == 0 || !help.InArray(ecid, tt)) {
		return errors.New("必须先攻击拥有嘲讽的卡牌")
	}

	var ec iface.ICard
	if ecid != 0 {
		ec = h.GetEnemy().GetBattleCardById(ecid)
		if ec == nil {
			return errors.New("没有找到此卡")
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

// 释放角色技能
func (b *Battle) PlayerHeroSkill(hid int, cs []iface.ICard, hs []iface.IHero) error {

	// h := b.GetHeroByIncrId(hid)

	return nil
}

// 结束回合
func (b *Battle) PlayerRoundEnd(hid int) error {
	b.RoundEnd()
	return nil
}
