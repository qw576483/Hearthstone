package iface

import (
	"hs/logic/config"
	"hs/logic/define"
)

// 接口
type ICard interface {
	NewPoint() ICard                                // 新指针
	Init(ICard, define.InCardsType, IHero, IBattle) // 初始化

	SetId(int)                                // 设置id
	GetId() int                               // 获得id
	GetReleaseId() int                        // 获得releaseid
	SetConfig(*config.CardConfig)             // 设置配置
	GetConfig() *config.CardConfig            // 获得配置
	GetType() define.CardType                 // 获得卡牌类型
	GetRace() []define.CardRace               // 获得卡牌种族
	IsRace(define.CardRace) bool              // 是否是某个种族
	GetTraits() []define.CardTraits           // 获得卡牌特质（冲锋，突袭，风怒...）
	GetHaveEffectTraits() []define.CardTraits // 获得有效果加成的卡牌特质
	IsHaveTraits(define.CardTraits) bool      // 是否拥有卡牌特质
	AddTraits(define.CardTraits)              // 添加特质
	RemoveTraits(define.CardTraits)           // 删除特质
	GetShield() int                           // 获得护盾
	SetShield(int)                            // 设置护盾
	TreatmentHp(ICard, int)                   // 治疗血量
	AddHp(int)                                // 加血
	AddHpMaxAndHp(int)                        // 加血上限和血
	SetHpMaxAndHp(int)                        // 设置血上限和血
	CostHp(ICard, int) int                    // 扣除血量(返回值为实际消耗)
	SetHp(int)                                // 设置血量
	GetHp() int                               // 获得卡牌血量
	DeleteHpEffect()                          // 删除hp的影响数据
	GetHaveEffectHp() int                     // 获得有效果加成的卡牌血量
	SetHpMax(int)                             // 设置血上限
	GetHpMax() int                            // 获得卡牌最大血量
	GetHaveEffectHpMax() int                  // 获得有效果加成的最大血量
	GetDamage() int                           // 获得卡牌攻击力
	GetHaveEffectDamage() int                 // 计算有效果加成的卡牌攻击力
	AddDamage(int)                            // 添加攻击力
	SetDamage(int)                            // 设置攻击
	ExchangeHpDamage()                        // 交换攻击和血
	GetApDamage() int                         // 获得法术伤害
	GetMona() int                             // 获得费用
	SetMona(int)                              // 设置费用
	GetHaveEffectMona() int                   // 计算有效果加成的卡牌费用
	SetCardInCardsPos(define.InCardsType)     // 设置此卡的位置
	GetCardInCardsPos() define.InCardsType    // 获得此卡的位置
	SetAfterDieBidx(int)                      // 设置死亡时的idx
	GetAfterDieBidx() int                     // 获得死亡时的idx
	SetOwner(IHero)                           // 设置拥有人
	GetOwner() IHero                          // 获得此卡拥有人
	GetNoLoopOwner() IHero                    // 获得不循环的拥有人，一般用于buff
	SetAttackTimes(int)                       // 设置攻击次数
	GetAttackTimes() int                      // 获得攻击次数
	GetMaxAttackTimes() int                   // 获得最大攻击次数
	Copy() (ICard, error)                     // 复制此卡
	Reset()                                   // 重置此卡
	Silent()                                  // 沉默此卡
	IsSilent() bool                           // 是否被沉默
	SetReleaseRound(int)                      // 设置出牌回合
	GetReleaseRound() int                     // 获得出牌回合
	SetFatherCard(ICard)                      // 设置父卡牌
	GetFatherCard() ICard                     // 获得父卡牌
	GetSubCards() []ICard                     // 获得子卡牌
	SetSubCards([]ICard)                      // 设置子卡牌
	AddSubCards(ICard)                        // 添加子卡牌
	RemoveSubCards(ICard)                     // 删除子卡牌

	// 事件 - 只需要实现接口
	OnInit()                    // 初始化时
	OnBattleBegin()             // 战斗开始
	OnGet()                     // 获得时
	OnRelease(int, int, ICard)  // 释放时 ， 输入抉择id(0,1)，站位，战吼目标
	OnRelease2(int, int, ICard) // 释放时 ， 输入抉择id(0,1)，站位，战吼目标
	OnPutToBattle(int)          // 步入战场时 ， 输入站位
	OnOutBattle()               // 离开战场时
	OnBeforeAttack(ICard) ICard // 攻击前
	OnAfterAttack(ICard)        // 攻击后
	OnHonorAnnihilate()         // 荣誉消灭
	OnOverflowAnnihilate()      // 超杀
	OnBeforeCostHp(int) int     // 受伤前，输入damage，输出新damage
	OnAfterCostHp()             // 受伤后
	OnAfterHpChange()           // 生命值改变后
	OnDie()                     // 卡牌死亡时
	OnAfterDisCard()            // 卡牌丢弃后
	OnGetMona(int) int          // 获取自己的费用时，输入mona ,输出新mona
	OnGetDamage(int) int        // 获取自己的攻击力时，输入damage ,输出新damage
	OnSilent()                  // 被沉默后

	// 注册事件 - 实现前需要注册
	OnNRGetBattleTime(int) int                         // 获得战斗时间 ， 输入战斗时间，返回新的战斗时间
	OnNRRoundBegin()                                   // 回合开始时
	OnNRRoundEnd()                                     // 回合结束时
	OnNROtherBeforeRelease(ICard, ICard) (ICard, bool) // 其他卡牌释放前，输入其他卡牌，攻击目标。输出攻击目标，是否生效。
	OnNROtherAfterRelease(ICard)                       // 其他卡牌释放前，输入其他卡牌
	OnNROtherBeforeAttack(ICard, ICard) ICard          // 其他卡牌攻击前，输入其他卡牌，攻击目标。输出攻击目标。
	OnNRPutToBattle(ICard)                             // 其他卡牌步入战场时，输入其他卡牌
	OnNROtherDie(ICard)                                // 其他卡牌死亡时，输入其他卡牌
	OnNROtherGetMona(ICard) int                        // 其他卡牌获取自己的费用时，输入其他卡牌， 输出费用加成
	OnNROtherGetDamage(ICard) int                      // 其他卡牌获取自己的攻击力时 ，输入其他卡牌， 输出攻击加成
	OnNROtherGetApDamage(IHero) int                    // 英雄获取自己的法术伤害时 ，输入其他卡牌， 输出的法术伤害加成
	OnNROtherGetHp(ICard) int                          // 其他卡牌获取自己的血量时 ，输入其他卡牌， 输出血量加成
	OnNROtherGetTraits(ICard) []define.CardTraits      // 其他卡牌获取自己的特质时 ，输入其他卡牌， 输出特质加成
	OnNROtherBeforeCostHp(ICard, int) int              // 受伤前，输入攻击者,num，输出新num
	OnNROtherAfterCostHp(ICard, int)                   // 受伤后，输入攻击者,num
	OnNROtherBeforeTreatmentHp(ICard, int) int         // 治疗前，输入治疗者,num，输出新num
	OnNROtherAfterTreatmentHp(ICard, int)              // 治疗后，输入治疗者,num
}
