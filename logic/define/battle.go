package define

type BattleStatus int

const (
	BattleStatusPre BattleStatus = iota // 预备
	BattleStatusRun                     // 战斗
	BattleStatusEnd                     // 结束
)
