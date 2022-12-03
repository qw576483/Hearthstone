package define

// 职业
type Vocation int

const (
	VocationRobbers Vocation = iota // 盗贼
	VocationWarlock                 // 术士
	VocationHunter                  // 猎人
	VocationShaman                  // 萨满
)

var MaxBattleNum = 8
