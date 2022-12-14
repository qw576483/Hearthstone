package iface

import (
	"hs/logic/define"
	"math/rand"

	"github.com/name5566/leaf/gate"
)

type IBattle interface {
	GetIncrRoundId() int                     // 获得自增的roundId
	GetIncrCardId() int                      // 获得自增的cardId
	GetIncrReleaseId() int                   // 获得自增的releaseId
	GetDoneSign(string) string               // 是否完成某项操作
	SetDoneSign(string, string)              // 完成某项操作
	GetBattleStatus() define.BattleStatus    // 获得战斗运行状态
	SetBattleStatus(define.BattleStatus)     // 设置战斗运行状态
	GetStatusChan() chan define.BattleStatus // 获得状态管道
	GetRoundHero() IHero                     // 获得当前回合的英雄
	GetHeroByIncrId(int) IHero               // 获得英雄，根据自增id
	GetHeros() []IHero                       // 获得两个英雄
	GetRandSeed() int64                      // 获得随机数种子
	GetRand() *rand.Rand                     // 随机数句柄
	GetHeroByGateAgent(gate.Agent) IHero     // 根据连接获得英雄

	PlayerChangePreCards(int, []int /** 这个值是第几张卡*/) error // 修改预留卡牌
	PlayerReleaseCard(int, int, int, int, int) error      // 施放卡牌
	PlayerUseHeroSkill(int, int, int) error               // 释放角色技能
	PlayerConCardAttack(int, int, int) error              // 卡牌进攻
	PlayerRoundEnd(int) error                             // 结束回合

	// 流程
	PreBegin()   // 预备战斗
	Begin()      // 开始战斗
	RoundBegin() // 回合开始
	RoundEnd()   // 回合结束

	// 事件
	GetEvent() map[string][]ICard      // 获得所有事件
	GetEventCards(string) []ICard      // 获得事件卡牌
	AddCardToEvent(ICard, string)      // 添加卡牌到事件
	RemoveCardFromEvent(ICard, string) // 从事件中删除卡牌
	RemoveCardFromAllEvent(ICard)      // 从事件中删除卡牌

	// 收集亡语
	RecordCardDie(ICard) // 收集真实死亡
	TrickCardDie()       // 触发死亡
	WhileTrickCardDie()  // 循环的触发亡语
}
