package msg

type Hello struct {
	Name string
}

type GetCardsConfig struct {
}

type JoinRoom struct {
	RoomId  int
	HeroId  int
	CardIds string
}

// 修改预留卡牌
type BChangePre struct {
	Indexs string
}

// 结束回合
type BEndRound struct {
	End int
}

// 释放英雄技能
type BUseSkill struct {
	ChoiceId int
	RCardId  int
	RHeroId  int
}

// 释放卡牌
type BRelease struct {
	CardId   int
	ChoiceId int
	PutIdx   int
	RCardId  int
	RHeroId  int
}

// 卡牌进攻
type BAttack struct {
	CardId  int
	ECardId int
	EHeroId int
}

// 英雄进攻
type BHAttack struct {
	ECardId int
	EHeroId int
}
