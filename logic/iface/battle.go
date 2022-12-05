package iface

import (
	"hs/logic/define"
	"math/rand"

	"github.com/name5566/leaf/gate"
)

type IBattle interface {
	GetIncrRoundId() int                  // 获得自增的roundId
	GetIncrCardId() int                   // 获得自增的cardId
	GetDoneSign(string) string            // 是否完成某项操作
	SetDoneSign(string, string)           // 完成某项操作
	GetBattleStatus() define.BattleStatus // 获得战斗运行状态
	GetRoundHero() IHero                  // 获得当前回合的英雄
	GetHeroByIncrId(int) IHero            // 获得英雄，根据自增id
	GetHeros() []IHero                    // 获得两个英雄
	GetRandSeed() int64                   // 获得随机数种子
	GetRand() *rand.Rand                  // 随机数句柄
	GetHeroByGateAgent(gate.Agent) IHero  // 根据连接获得英雄

	PlayerChangePreCards(int, []int /** 这个值是第几张卡*/) error // 修改预留卡牌
	PlayerReleaseCard(int, int, int, int, int, int) error // 施放卡牌
	PlayerUseHeroSkill(int, int, int, int) error          // 释放角色技能
	PlayerConCardAttack(int, int, int, int) error         // 卡牌进攻
	PlayerAttack(int, int, int) error                     // 玩家进攻
	PlayerRoundEnd(int) error                             // 结束回合

	// 流程
	PreBegin()   // 预备战斗
	Begin()      // 开始战斗
	RoundBegin() // 回合开始
	RoundEnd()   // 回合结束
}
