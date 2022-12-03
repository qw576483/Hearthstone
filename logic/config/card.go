package config

import "hs/logic/define"

type CardConfig struct {
	Id       int                 // id
	Name     string              // 名字
	Desc     string              // 描述
	Hp       int                 // 血量
	Damage   int                 // 伤害
	Mona     int                 // 费用
	Ctype    define.CardType     // 类型
	Race     []define.CardRace   // 种族
	Traits   []define.CardTraits // 特质
	Series   define.CardSeries   // 系列
	Vocation []define.Vocation   // 职业
	CanCarry bool                // 是否可携带
}

// 创建种族
func MakeCardRace(races ...define.CardRace) []define.CardRace {
	ret := make([]define.CardRace, 0)
	return append(ret, races...)
}

// 创建特质
func MakeCardTraits(traits ...define.CardTraits) []define.CardTraits {
	ret := make([]define.CardTraits, 0)
	return append(ret, traits...)
}

// 创建职业
func MakeCardVocation(vocations ...define.Vocation) []define.Vocation {
	ret := make([]define.Vocation, 0)
	return append(ret, vocations...)
}

// 获得配置
func GetCardConfig(configId int) *CardConfig {
	return defineCardConfig[configId]
}

// 获得全部配置
func GetAllCardConfig() []*CardConfig {
	return defineCardConfig
}

var defineCardConfig []*CardConfig = []*CardConfig{
	0: &CardConfig{
		Id:       0,
		Name:     "幸运币",
		Desc:     "在本回合中，获得一个法力水晶。",
		Mona:     0,
		Ctype:    define.CardTypeSorcery,
		Series:   define.CardSeriesBase,
		CanCarry: false,
	},
	1: &CardConfig{
		Id:       1,
		Name:     "石牙野猪",
		Mona:     1,
		Damage:   2,
		Hp:       1,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceBeast),
		Traits:   MakeCardTraits(define.CardTraitsAssault),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	2: &CardConfig{
		Id:       2,
		Name:     "疯狂的炼金师",
		Desc:     "战吼：使一个随从的攻击力和生命值互换。",
		Mona:     2,
		Damage:   2,
		Hp:       2,
		Ctype:    define.CardTypeEntourage,
		Traits:   MakeCardTraits(define.CardTraitsOnRelease),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	3: &CardConfig{
		Id:       3,
		Name:     "寒光智者",
		Desc:     "战吼：双方每个玩家抽两张牌。",
		Mona:     3,
		Damage:   2,
		Hp:       2,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceFish),
		Traits:   MakeCardTraits(define.CardTraitsOnRelease),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	4: &CardConfig{
		Id:       4,
		Name:     "麦田傀儡",
		Desc:     "亡语：召唤一个2/1的损坏的傀儡。",
		Mona:     3,
		Damage:   2,
		Hp:       3,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceMechanics),
		Traits:   MakeCardTraits(define.CardTraitsOnDie),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	5: &CardConfig{
		Id:       5,
		Name:     "损坏的傀儡[麦田傀儡衍生物]",
		Mona:     1,
		Damage:   2,
		Hp:       1,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceMechanics),
		Series:   define.CardSeriseClassic,
		CanCarry: false,
	},
	6: &CardConfig{
		Id:       6,
		Name:     "攻城车",
		Desc:     "在你的回合开始时，随机对一个敌人造成2点伤害。",
		Mona:     3,
		Damage:   1,
		Hp:       4,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceMechanics),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	7: &CardConfig{
		Id:       7,
		Name:     "铸剑师",
		Desc:     "在你的回合结束时，随机使另一个友方随从获得+1攻击力。",
		Mona:     2,
		Damage:   1,
		Hp:       3,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	8: &CardConfig{
		Id:       8,
		Name:     "螃蟹骑士",
		Mona:     2,
		Damage:   1,
		Hp:       4,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceFish),
		Traits:   MakeCardTraits(define.CardTraitsSuddenStrike, define.CardTraitsWindfury),
		Series:   define.CardSeriseDarkmoon,
		CanCarry: true,
	},
	9: &CardConfig{
		Id:       9,
		Name:     "毁灭之刃",
		Desc:     "战吼：造成1点伤害。连击：改为造成2点伤害。",
		Mona:     3,
		Damage:   2,
		Hp:       2,
		Ctype:    define.CardTypeWeapon,
		Traits:   MakeCardTraits(define.CardTraitsOnRelease, define.CardTraitsCarom),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationRobbers),
		CanCarry: true,
	},
	10: &CardConfig{
		Id:       10,
		Name:     "食腐土狼",
		Desc:     "每当一个友方野兽死亡，便获得+2/+1。",
		Mona:     2,
		Damage:   2,
		Hp:       2,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceBeast),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationHunter),
		CanCarry: true,
	},
	11: &CardConfig{
		Id:       11,
		Name:     "上古看守者",
		Mona:     2,
		Damage:   4,
		Hp:       5,
		Ctype:    define.CardTypeEntourage,
		Traits:   MakeCardTraits(define.CardTraitsUnableToAttack),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	12: &CardConfig{
		Id:       12,
		Name:     "持盾卫士",
		Mona:     0,
		Damage:   0,
		Hp:       4,
		Ctype:    define.CardTypeEntourage,
		Traits:   MakeCardTraits(define.CardTraitsTaunt),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	13: &CardConfig{
		Id:       13,
		Name:     "银色侍从",
		Mona:     1,
		Damage:   1,
		Hp:       1,
		Ctype:    define.CardTypeEntourage,
		Traits:   MakeCardTraits(define.CardTraitsHolyShield),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	14: &CardConfig{
		Id:       14,
		Name:     "耐心的刺客",
		Mona:     2,
		Damage:   1,
		Hp:       1,
		Ctype:    define.CardTypeEntourage,
		Traits:   MakeCardTraits(define.CardTraitsSneak, define.CardTraitsHighlyToxic),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationRobbers),
		CanCarry: true,
	},
	15: &CardConfig{
		Id:       15,
		Name:     "疯狂投弹者",
		Desc:     "战吼：造成3点伤害，随机分配到所有其他角色身上。",
		Mona:     2,
		Damage:   3,
		Hp:       2,
		Ctype:    define.CardTypeEntourage,
		Traits:   MakeCardTraits(define.CardTraitsOnRelease),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	16: &CardConfig{
		Id:       16,
		Name:     "飞刀杂耍者",
		Desc:     "在你召唤一个随从后，随机对一个敌人造成1点伤害。",
		Mona:     2,
		Damage:   3,
		Hp:       2,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	17: &CardConfig{
		Id:       17,
		Name:     "火舌图腾",
		Desc:     "相邻的随从获得+2攻击力。",
		Mona:     2,
		Damage:   0,
		Hp:       3,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceTotems),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationShaman),
		CanCarry: true,
	},
	18: &CardConfig{
		Id:       18,
		Name:     "小个子召唤师",
		Desc:     "你每个回合使用的第一张随从牌的法力值消耗减少（1）点。",
		Mona:     2,
		Damage:   2,
		Hp:       2,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	19: &CardConfig{
		Id:       19,
		Name:     "暴风城勇士",
		Desc:     "你的其他随从获得+1/+1",
		Mona:     7,
		Damage:   6,
		Hp:       6,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	20: &CardConfig{
		Id:       20,
		Name:     "小精灵",
		Mona:     0,
		Damage:   1,
		Hp:       1,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
}
