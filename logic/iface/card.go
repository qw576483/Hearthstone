package iface

import (
	"hs/logic/config"
	"hs/logic/define"
)

// 接口
type ICard interface {
	NewPoint() ICard                                // 新指针
	Init(ICard, define.InCardsType, IHero, IBattle) // 初始化

	GetId() int                            // 获得卡牌id
	SetConfig(*config.CardConfig)          // 设置配置
	GetConfig() *config.CardConfig         // 获得配置
	GetType() define.CardType              // 获得卡牌类型
	GetRace() []define.CardRace            // 获得卡牌种族
	GetTraits() []define.CardTraits        // 获得卡牌特质（冲锋，突袭，风怒...）
	IsHaveTraits(define.CardTraits) bool   // 是否拥有卡牌特质
	AddTraits(define.CardTraits)           // 添加特质
	RemoveTraits(define.CardTraits)        // 删除特质
	TreatmentHp(int)                       // 治疗血量
	AddHp(int)                             // 加血
	AddHpMaxAndHp(int)                     // 加血上限和血
	SetHpMaxAndHp(int)                     // 设置血上限和血
	CostHp(int) int                        // 扣除血量(返回值为实际消耗)
	SetHp(int)                             // 设置血量
	GetHp() int                            // 获得卡牌血量
	DeleteHpEffect()                       // 删除hp的影响数据
	GetHaveEffectHp() int                  // 获得有效果加成的卡牌血量
	SetHpMax(int)                          // 设置血上限
	GetHpMax() int                         // 获得卡牌最大血量
	GetHaveEffectHpMax() int               // 获得有效果加成的最大血量
	GetDamage() int                        // 获得卡牌攻击力
	GetHaveEffectDamage(ICard) int         // 计算有效果加成的卡牌攻击力
	AddDamage(int)                         // 添加攻击力
	SetDamage(int)                         // 设置攻击
	ExchangeHpDamage(ICard)                // 交换攻击和血
	GetMona() int                          // 获得法力值
	GetHaveEffectMona(ICard) int           // 计算有效果加成的卡牌费用
	SetCardInCardsPos(define.InCardsType)  // 设置此卡的位置
	GetCardInCardsPos() define.InCardsType // 获得此卡的位置
	GetHandPos() (int, error)              // 获得此卡在手牌中的位置
	SetOwner(IHero)                        // 设置拥有人
	GetOwner() IHero                       // 获得此卡拥有人
	GetFatherCard() ICard                  // 获得父卡牌
	GetSubCards() []ICard                  // 获得子卡牌
	AddSubCards(ICard)                     // 添加子卡牌
	RemoveSubCards(ICard)                  // 删除子卡牌
	SetAttackTimes(int)                    // 设置攻击次数
	GetAttackTimes() int                   // 获得攻击次数
	GetMaxAttackTimes() int                // 获得最大攻击次数
	Copy() (ICard, error)                  // 复制此卡
	Reset()                                // 重置此卡
	Silent(ICard)                          // 沉默此卡
	IsSilent() bool                        // 是否被沉默
	SetReleaseRound(int)                   // 设置出牌回合
	GetReleaseRound() int                  // 获得出牌回合

	// 事件 - 只需要实现接口
	OnInit()                          // 初始化时
	OnBattleBegin()                   // 战斗开始
	OnGet()                           // 获得时
	OnRelease(int, int, ICard, IHero) // 释放时
	OnPutToBattle(int)                // 步入战场时
	OnOutBattle()                     // 离开战场时
	OnHonorAnnihilate(ICard)          // 荣誉消灭
	OnOverflowAnnihilate(ICard)       // 超杀
	OnDie(int)                        // 卡牌死亡时
	OnDevastate()                     // 卡牌销毁时
	OnGetMona() int                   // 获取自己的费用时，返回费用加成
	OnGetDamage() int                 // 获取自己的攻击力时 , 返回攻击加成

	// 注册事件 - 实现前需要注册
	OnNRRoundBegin()              // 回合开始时
	OnNRRoundEnd()                // 回合结束时
	OnNRPutToBattle(ICard)        // 其他卡牌步入战场时
	OnNROtherDie(ICard)           // 其他卡牌死亡时
	OnNROtherGetMona(ICard) int   // 其他卡牌获取自己的费用时， 返回费用加成
	OnNROtherGetDamage(ICard) int // 其他卡牌获取自己的攻击力时 ， 返回攻击加成
	OnNROtherGetHp(ICard) int     // 其他卡牌获取自己的血量时 ， 返回血量加成

}
