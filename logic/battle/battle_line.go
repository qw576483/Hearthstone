package battle

import (
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
	"strconv"
)

// 预备战斗
func (b *Battle) PreBegin() {

	// 分配先后手
	randIndex := b.rand.Intn(2)
	b.hero = b.heros[randIndex]
	var enemy iface.IHero
	if randIndex == 0 {
		enemy = b.heros[1]
	} else {
		enemy = b.heros[0]
	}
	b.hero.SetEnemy(enemy)
	enemy.SetEnemy(b.hero)

	// 设置状态
	b.status = define.BattleStatusPre

	for _, v := range b.heros {
		v.PreBegin()
	}

	push.PushLine(b)
	push.PushInit(b)
}

// 开始战斗
func (b *Battle) Begin() {

	// 将预抽保存到手牌
	b.hero.SetHandCards(b.hero.GetPreCards())
	b.hero.GetEnemy().SetHandCards(b.hero.GetEnemy().GetPreCards())
	// 添加一个幸运币
	b.hero.GetEnemy().GiveNewCardToHand(0)

	b.status = define.BattleStatusRun

	// 战斗开始
	b.hero.TrickBattleBegin()
	b.hero.GetEnemy().TrickBattleBegin()

	push.PushLine(b)

	// 补卡
	b.RoundBegin()
}

// 回合开始
func (b *Battle) RoundBegin() {
	b.incrRoundId += 1

	push.PushLog(b.hero.GetEnemy(), "==== 【你的对手】回合开始("+strconv.Itoa(b.incrRoundId)+")")
	push.PushLog(b.hero, "==== 你的回合开始("+strconv.Itoa(b.incrRoundId)+")")

	b.hero.RoundBegin()

	push.PushInfoMsg(b)
}

// 回合结束
func (b *Battle) RoundEnd() {

	push.PushLog(b.hero.GetEnemy(), "==== 【你的对手】回合结束("+strconv.Itoa(b.incrRoundId)+")")
	push.PushLog(b.hero, "==== 你的回合结束("+strconv.Itoa(b.incrRoundId)+")")

	b.hero.RoundEnd()

	// 切换出手
	b.hero = b.hero.GetEnemy()
	b.RoundBegin()
}
