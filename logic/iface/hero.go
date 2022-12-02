package iface

import (
	"hs/logic/config"

	"github.com/name5566/leaf/gate"
)

type IHero interface {
	SetGateAgent(gate.Agent)  // 设置连接
	GetGateAgent() gate.Agent // 获得连接
	NewPoint() IHero          // 新指针
	Init([]ICard, IBattle)    // 初始化

	GetBattle() IBattle                                // 获得战斗句柄
	GetId() int                                        // 获得英雄id
	IsRoundHero() bool                                 // 是否是我的回合
	SetConfig(*config.HeroConfig)                      // 设置配置数据
	GetConfig() *config.HeroConfig                     // 获得配置数据
	SetEnemy(IHero)                                    // 设置敌人
	GetEnemy() IHero                                   // 获得敌人
	GetIdxByCards(ICard, []ICard) int                  // 获得下标
	GetPreCards() []ICard                              // 获得预存卡牌
	SetHandCards([]ICard)                              // 设置手牌
	GetHandCards() []ICard                             // 获得手牌上的卡牌
	GetHandCardByIncrId(int) ICard                     // 获得手牌上的卡牌
	GetLibCards() []ICard                              // 获得牌库上的卡牌
	GetGraveCards() []ICard                            // 获得坟场上的卡牌
	GetBattleCards() []ICard                           // 获得战场上的卡牌
	GetBattleCardById(int) ICard                       // 获得战场上的卡牌
	GetBattleCardsByIds([]int) []ICard                 // 获得战场上的卡牌
	GetBattleCardsTraitsTauntCardIds() []int           // 获得战场上有嘲讽的卡牌ids
	AppendToAllCards(ICard)                            // 添加到全部卡牌
	GetAllCards() []ICard                              // 获得全部卡牌
	GetBothAllCards() []ICard                          // 获得全部卡牌
	GetDamage() int                                    // 获得攻击力
	SetAttackTimes(int)                                // 设置攻击次数
	GetAttackTimes() int                               // 获得攻击次数
	GetMaxAttackTimes() int                            // 获得最大攻击次数
	GetHp() int                                        // 获得血量
	GetHpMax() int                                     // 获得最大血量
	CostHp(int) int                                    // 扣血
	AddMona(int)                                       // 添加法力值
	CostMona(int) bool                                 // 消耗法力值
	SetMona(int)                                       // 设置法力值
	GetMona() int                                      // 获得法力值
	AddMonaMax(int)                                    // 添加最大法力值
	GetMonaMax() int                                   // 获得最大法力（当前）
	GetShield() int                                    // 获得护盾
	SetWeapon(ICard)                                   // 设置武器
	GetWeapon() ICard                                  // 获得当前武器
	GiveNewCardToHand(int) ICard                       // 给一个新卡牌到手上
	MoveToHand(ICard)                                  // 添加到手牌
	MoveOutHandOnlyHandCards(ICard)                    // 撤出手牌
	MoveToBattle(ICard, int)                           // 布入战场
	MoveOutBattleOnlyBattleCards(ICard) int            // 移出战场
	DieCard(ICard)                                     // 杀死卡牌
	GetMaxHandCardsNum() int                           // 获得手牌上限数量
	DrawForPreBegin(int)                               // 预备开始时的抽卡
	ChangePreCrards([]int)                             // 修改预备抽卡
	DrawByTimes(int)                                   // 抽卡
	SetFatigue(int)                                    // 设置疲劳伤害
	GetFatigue() int                                   // 获得当前疲劳伤害
	Release(ICard, int, int, ICard, IHero, bool) error // 出牌
	Attack(ICard, ICard, IHero) error                  // 攻击
	HAttack(ICard, IHero) error                        // 英雄攻击
	Die()                                              // 死亡
	Push(interface{})                                  // 推送数据
	RandBattleCardOrHero() (ICard, IHero)              // 随机战场上的卡牌或者英雄
	RandBothBattleCardOrHero() (ICard, IHero)          // 随机战场上的卡牌或者英雄
	RandExcludeCard([]ICard, ICard) ICard              // 随机卡牌，排除一个卡牌
	GetReleaseCardTimes() int                          // 获得出牌次数
	SetReleaseCardTimes(int)                           // 设置出牌次数
	GetEventCards(string) []ICard                      // 获得事件卡牌
	GetBothEventCards(string) []ICard                  // 获得双方事件卡牌
	AddCardToEvent(ICard, string)                      // 添加卡牌到事件
	RemoveCardFromEvent(ICard, string)                 // 从事件中删除卡牌

	PreBegin()   // 预备阶段
	RoundBegin() // 回合开始
	RoundEnd()   // 回合结束

	TrickBattleBegin()                              // 触发战斗开始事件
	TrickRelease(ICard, int, int, ICard, IHero)     // 触发战吼
	TrickRoundBegin()                               // 触发回合开始事件
	TrickRoundEnd()                                 // 触发回合结束事件
	TrickGetCardEvent(ICard)                        // 触发抽卡事件
	TrickDevastateCardEvent(ICard)                  // 触发销毁事件
	TrickAfterAttackEvent(ICard, ICard, IHero, int) // 触发攻击后事件
	TrickDieCardEvent(ICard, int)                   // 触发死亡事件
}