package define

// 职业
type Vocation int

const (
	VocationRobbers Vocation = iota // 盗贼
	VocationWarlock                 // 术士
	VocationHunter                  // 猎人
	VocationShaman                  // 萨满
	VocationPaladin                 // 圣骑士
	VocationDruid                   // 德鲁伊
	VocationMaster                  // 法师
	VocationPastor                  // 牧师
	VocationWarrior                 // 战士
	VocationNeutral                 // 中立
)

var MaxBattleNum = 7

var HeroId = 71
