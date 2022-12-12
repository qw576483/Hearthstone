package iface

import (
	"hs/logic/config"

	"github.com/name5566/leaf/gate"
)

type IHero interface {
	SetGateAgent(gate.Agent)      // 设置连接
	GetGateAgent() gate.Agent     // 获得连接
	NewPoint() IHero              // 新指针
	Init(IHero, []ICard, IBattle) // 初始化

	GetId() int                                 // 获得英雄id
	GetBattle() IBattle                         // 获得战斗句柄
	GetHead() ICard                             // 获得英雄的卡牌
	IsRoundHero() bool                          // 是否是我的回合
	SetHeroSkill(ICard)                         // 设置英雄技能
	GetHeroSkill() ICard                        // 获得英雄技能
	SetConfig(*config.HeroConfig)               // 设置配置数据
	GetConfig() *config.HeroConfig              // 获得配置数据
	SetEnemy(IHero)                             // 设置敌人
	GetEnemy() IHero                            // 获得敌人
	GetIdxByCards(ICard, []ICard) int           // 获得下标
	GetPreCards() []ICard                       // 获得预存卡牌
	SetHandCards([]ICard)                       // 设置手牌
	GetHandCards() []ICard                      // 获得手牌上的卡牌
	GetHandCardByIncrId(int) ICard              // 获得手牌上的卡牌
	GetLibCards() []ICard                       // 获得牌库上的卡牌
	GetGraveCards() []ICard                     // 获得坟场上的卡牌
	GetBattleCards() []ICard                    // 获得战场上的卡牌
	GetBattleCardById(int) ICard                // 获得战场上的卡牌
	GetBattleCardsByIds([]int) []ICard          // 获得战场上的卡牌
	GetBattleCardsTraitsTauntCardIds() []int    // 获得战场上有嘲讽的卡牌ids
	GetCardById(int) ICard                      // 获得卡牌根据id
	GetCardIdx(ICard, []ICard) int              // 获得卡牌的位置
	AppendToAllCards(ICard)                     // 添加到全部卡牌
	GetAllCards() []ICard                       // 获得全部卡牌
	GetBothAllCards() []ICard                   // 获得全部卡牌
	GetApDamage() int                           // 获得法术伤害
	AddMona(int)                                // 添加法力值
	CostMona(int) bool                          // 消耗法力值
	SetMona(int)                                // 设置法力值
	GetMona() int                               // 获得法力值
	AddMonaMax(int)                             // 添加最大法力值
	GetMonaMax() int                            // 获得最大法力（当前）
	GetLockMona() int                           // 获得锁定法力值
	SetLockMona(int)                            // 设置锁定法力值
	GetLockMonaCache() int                      // 获得锁定法力值缓存
	SetLockMonaCache(int)                       // 设置锁定法力值缓存
	SetWeapon(ICard)                            // 设置武器
	GetWeapon() ICard                           // 获得当前武器
	GiveNewCardToHand(int) ICard                // 给一个新卡牌到手上
	MoveToHand(ICard)                           // 添加到手牌
	MoveOutHandOnlyHandCards(ICard)             // 撤出手牌
	MoveToBattle(ICard, int)                    // 布入战场
	MoveOutBattleOnlyBattleCards(ICard) int     // 移出战场
	CaptureCard(ICard, int)                     // 夺取卡牌
	DiscardCard(ICard)                          // 丢弃手牌
	DieCard(ICard, bool)                        // 杀死卡牌，是否立即触发亡语
	GetMaxHandCardsNum() int                    // 获得手牌上限数量
	DrawForPreBegin(int)                        // 预备开始时的抽卡
	ChangePreCrards([]int)                      // 修改预备抽卡
	DrawByTimes(int)                            // 抽卡
	SetFatigue(int)                             // 设置疲劳伤害
	GetFatigue() int                            // 获得当前疲劳伤害
	Release(ICard, int, int, ICard, bool) error // 出牌
	OnlyReleaseWeapon(ICard)                    // 仅仅出张武器卡
	Attack(ICard, ICard) error                  // 攻击
	Die()                                       // 死亡
	Push(interface{})                           // 推送数据
	RandBattleCardOrHero() ICard                // 随机战场上的卡牌或者英雄
	RandBothBattleCardOrHero() ICard            // 随机战场上的卡牌或者英雄
	RandCard([]ICard) ICard                     // 随机卡牌
	RandExcludeCard([]ICard, ICard) ICard       // 随机卡牌，排除一个卡牌
	GetReleaseCardTimes() int                   // 获得出牌次数
	SetReleaseCardTimes(int)                    // 设置出牌次数
	GetSecrets() []ICard                        // 获得奥秘
	CanReleaseSecret(ICard) bool                // 是否能释放奥秘
	OnlyReleaseSecret(ICard) bool               // 仅仅释放奥秘，返回是否释放成功
	DeleteSecret(ICard, bool)                   // 删除奥秘
	NewCountDown(int)                           // 一个新的倒计时
	CloseCountDown()                            // 关闭我的倒计时
	Henshin(ICard)                              // 变身

	PreBegin()    // 预备阶段
	RoundBegin()  // 回合开始
	RoundEnd()    // 回合结束
	FixRoundEnd() // 强制回合结束

	TrickBattleBegin()                       // 触发战斗开始事件
	TrickGetCardEvent(ICard)                 // 触发抽卡事件
	TrickRelease(ICard, int, int, ICard)     // 触发战吼
	TrickRelease2(ICard, int, int, ICard)    // 触发战吼
	TrickRoundBegin()                        // 触发回合开始事件
	TrickRoundEnd()                          // 触发回合结束事件
	TrickPutToBattleEvent(ICard, int)        // 触发步入战场事件
	TrickOutBattleEvent(ICard)               // 触发离开战场事件
	TrickAfterAttackEvent(ICard, ICard, int) // 触发攻击后事件
	TrickDieCardEvent(ICard)                 // 触发死亡事件
}
