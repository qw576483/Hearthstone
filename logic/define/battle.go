package define

type BattleStatus int

const (
	BattleStatusPre BattleStatus = iota // 预备
	BattleStatusRun                     // 战斗
	BattleStatusEnd                     // 结束
)

// 战斗时间
const BattleTime = 120
