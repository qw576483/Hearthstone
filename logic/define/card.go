package define

// 卡牌的位置
type InCardsType int

const (
	InCardsTypeNone   InCardsType = iota // 不知道在哪
	InCardsTypeHand                      // 手牌中
	InCardsTypeLib                       // 牌库中
	InCardsTypeGrave                     // 坟场
	InCardsTypeBattle                    // 战场
	InCardsTypeBody                      // 身上
)

// 卡牌类型
type CardType int

const (
	CardTypeEntourage CardType = iota // 随从
	CardTypeWeapon                    // 武器
	CardTypeSorcery                   // 法术
	CardTypeSecret                    // 奥秘
	CardTypeBuff                      // buff - 不能直接使用
)

// 卡牌特质
type CardTraits int

const (
	CardTraitsOnRelease      CardTraits = iota // 战吼
	CardTraitsOnDie                            // 亡语
	CardTraitsAssault                          // 冲锋
	CardTraitsSuddenStrike                     // 突袭
	CardTraitsWindfury                         // 风怒
	CardTraitsCarom                            // 连击
	CardTraitsUnableToAttack                   // 无法攻击
	CardTraitsTaunt                            // 嘲讽
	CardTraitsHolyShield                       // 圣盾
	CardTraitsSneak                            // 潜行
	CardTraitsHighlyToxic                      // 剧毒
)

// 卡牌种族
type CardRace int

const (
	CardRaceBeast     CardRace = iota // 野兽
	CardRaceDevil                     // 恶魔
	CardRaceFish                      // 鱼人
	CardRaceMechanics                 // 机械
	CardRaceTotems                    // 图腾
)

// 卡牌系列
type CardSeries int

const (
	CardSeriesBase     CardSeries = iota // 基础
	CardSeriseClassic                    // 经典
	CardSeriseDarkmoon                   // 暗月马戏团
)

// 卡牌品质
type CardQuality int

const (
	CardQualityBase   CardQuality = iota // 基础
	CardQualityWhite                     // 普通
	CardQualityBlue                      // 稀有
	CardQualityPurple                    // 史诗
	CardQualityOrange                    // 传说
)

var BuffCardId_MyRoundEndClear = 21   // buffId ,我的回合结束清除效果
var BuffCardId_MyRoundBeginClear = 22 // buffId ,我的回合开始清除效果
