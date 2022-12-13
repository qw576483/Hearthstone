package config

import "hs/logic/define"

type CardConfig struct {
	Id            int                      // id
	Name          string                   // 名字
	Desc          string                   // 描述
	Hp            int                      // 血量
	Damage        int                      // 伤害
	ApDamage      int                      // 法术伤害
	Mona          int                      // 费用
	Ctype         define.CardType          // 类型
	Quality       define.CardQuality       // 品质
	Race          []define.CardRace        // 种族
	Traits        []define.CardTraits      // 特质
	Series        define.CardSeries        // 系列
	Vocation      []define.Vocation        // 职业
	ReleaseFilter define.CardReleaseFilter // 释放筛选
	CanCarry      bool                     // 是否可携带

	IntParam1 int // int参数1
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
	if configId >= len(defineCardConfig) {
		return nil
	}
	return defineCardConfig[configId]
}

// 获得全部配置
func GetAllCardConfig() []*CardConfig {
	return defineCardConfig
}

var defineCardConfig []*CardConfig = []*CardConfig{
	0: &CardConfig{
		Id:      0,
		Name:    "幸运币",
		Desc:    "在本回合中，获得一个法力水晶。",
		Mona:    0,
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeSorcery,
		Series:  define.CardSeriseClassic,
	},
	1: &CardConfig{
		Id:       1,
		Name:     "石牙野猪",
		Mona:     1,
		Damage:   2,
		Hp:       1,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceBeast),
		Traits:   MakeCardTraits(define.CardTraitsAssault),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	2: &CardConfig{
		Id:            2,
		Name:          "疯狂的炼金师",
		Desc:          "战吼：使一个随从的攻击力和生命值互换。",
		Mona:          2,
		Damage:        2,
		Hp:            2,
		Quality:       define.CardQualityBlue,
		Ctype:         define.CardTypeEntourage,
		Traits:        MakeCardTraits(define.CardTraitsOnRelease),
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterBattle,
		CanCarry:      true,
	},
	3: &CardConfig{
		Id:       3,
		Name:     "寒光智者",
		Desc:     "战吼：双方玩家抽两张牌。",
		Mona:     3,
		Damage:   2,
		Hp:       2,
		Quality:  define.CardQualityBlue,
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
		Quality:  define.CardQualityWhite,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceMechanics),
		Traits:   MakeCardTraits(define.CardTraitsOnDie),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	5: &CardConfig{
		Id:      5,
		Name:    "损坏的傀儡",
		Mona:    1,
		Damage:  2,
		Hp:      1,
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeEntourage,
		Race:    MakeCardRace(define.CardRaceMechanics),
		Series:  define.CardSeriseClassic,
	},
	6: &CardConfig{
		Id:       6,
		Name:     "攻城车",
		Desc:     "在你的回合开始时，随机对一个敌人造成2点伤害。",
		Mona:     3,
		Damage:   1,
		Hp:       4,
		Quality:  define.CardQualityBlue,
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
		Quality:  define.CardQualityBlue,
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
		Quality:  define.CardQualityWhite,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceFish),
		Traits:   MakeCardTraits(define.CardTraitsSuddenStrike, define.CardTraitsWindfury),
		Series:   define.CardSeriseDarkmoon,
		CanCarry: true,
	},
	9: &CardConfig{
		Id:            9,
		Name:          "毁灭之刃",
		Desc:          "战吼：造成1点伤害。连击：改为造成2点伤害。",
		Mona:          3,
		Damage:        2,
		Hp:            2,
		Quality:       define.CardQualityBlue,
		Ctype:         define.CardTypeWeapon,
		Traits:        MakeCardTraits(define.CardTraitsOnRelease, define.CardTraitsCarom),
		Series:        define.CardSeriseClassic,
		Vocation:      MakeCardVocation(define.VocationRobbers),
		ReleaseFilter: define.CardReleaseFilterAll,
		CanCarry:      true,
	},
	10: &CardConfig{
		Id:       10,
		Name:     "食腐土狼",
		Desc:     "每当一个友方野兽死亡，便获得+2/+1。",
		Mona:     2,
		Damage:   2,
		Hp:       2,
		Quality:  define.CardQualityWhite,
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
		Quality:  define.CardQualityBlue,
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
		Quality:  define.CardQualityWhite,
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
		Quality:  define.CardQualityWhite,
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
		Quality:  define.CardQualityPurple,
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
		Quality:  define.CardQualityWhite,
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
		Quality:  define.CardQualityBlue,
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
		Quality:  define.CardQualityBase,
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
		Quality:  define.CardQualityBlue,
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
		Quality:  define.CardQualityWhite,
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
		Quality:  define.CardQualityWhite,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	21: &CardConfig{
		Id:    21,
		Name:  "Buff",
		Desc:  "我的回合结束时消散",
		Ctype: define.CardTypeBuff,
	},
	22: &CardConfig{
		Id:    22,
		Name:  "Buff",
		Desc:  "我的回合开始时消散",
		Ctype: define.CardTypeBuff,
	},
	23: &CardConfig{
		Id:    23,
		Name:  "Buff",
		Desc:  "永久生效",
		Ctype: define.CardTypeBuff,
	},
	24: &CardConfig{
		Id:            24,
		Name:          "叫嚣的中士",
		Desc:          "战吼：在本回合中，使一个随从获得+2攻击力。",
		Mona:          1,
		Damage:        1,
		Hp:            1,
		Quality:       define.CardQualityWhite,
		Ctype:         define.CardTypeEntourage,
		Traits:        MakeCardTraits(define.CardTraitsOnRelease),
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterBattle,
		CanCarry:      true,
	},
	25: &CardConfig{
		Id:            25,
		Name:          "银色保卫者",
		Desc:          "战吼：使一个其他友方随从获得圣盾。",
		Mona:          2,
		Damage:        3,
		Hp:            2,
		Quality:       define.CardQualityWhite,
		Ctype:         define.CardTypeEntourage,
		Traits:        MakeCardTraits(define.CardTraitsOnRelease),
		Series:        define.CardSeriseClassic,
		Vocation:      MakeCardVocation(define.VocationPaladin),
		ReleaseFilter: define.CardReleaseFilterMyBattle,
		CanCarry:      true,
	},
	26: &CardConfig{
		Id:       26,
		Name:     "匕首精通",
		Desc:     "英雄技能装备一把1/2的匕首。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationRobbers),
		Series:   define.CardSeriseClassic,
	},
	27: &CardConfig{
		Id:       27,
		Name:     "邪恶短刀",
		Mona:     2,
		Damage:   1,
		Hp:       2,
		Ctype:    define.CardTypeWeapon,
		Vocation: MakeCardVocation(define.VocationRobbers),
		Series:   define.CardSeriseClassic,
	},
	28: &CardConfig{
		Id:            28,
		Name:          "铁喙猫头鹰",
		Desc:          "战吼：沉默一个随从。",
		Mona:          3,
		Damage:        2,
		Hp:            1,
		Quality:       define.CardQualityWhite,
		Ctype:         define.CardTypeEntourage,
		Race:          MakeCardRace(define.CardRaceBeast),
		Traits:        MakeCardTraits(define.CardTraitsOnRelease),
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterBattle,
		CanCarry:      true,
	},
	29: &CardConfig{
		Id:       29,
		Name:     "奉献",
		Desc:     "对所有敌人造成2点伤害。",
		Mona:     4,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeSorcery,
		Race:     MakeCardRace(define.CardRaceSacred),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationPaladin),
		CanCarry: true,
	},
	30: &CardConfig{
		Id:       30,
		Name:     "狗头人地卜师",
		Desc:     "法术伤害+1",
		Mona:     2,
		Damage:   2,
		Hp:       2,
		ApDamage: 1,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	31: &CardConfig{
		Id:       31,
		Name:     "游学者周卓",
		Desc:     "每当一个玩家施放一个法术，复制该法术，将其置入另一个玩家的手牌。",
		Mona:     2,
		Damage:   0,
		Hp:       4,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	32: &CardConfig{
		Id:            32,
		Name:          "丛林守护者",
		Desc:          "抉择：造成2点伤害；或者沉默一个随从。",
		Mona:          4,
		Damage:        2,
		Hp:            4,
		Quality:       define.CardQualityBlue,
		Ctype:         define.CardTypeEntourage,
		Series:        define.CardSeriseClassic,
		Traits:        MakeCardTraits(define.CardTraitsChoice),
		Vocation:      MakeCardVocation(define.VocationDruid),
		ReleaseFilter: define.CardReleaseFilterBattle,
		CanCarry:      true,
	},
	33: &CardConfig{
		Id:            33,
		Name:          "年轻的酒仙",
		Desc:          "战吼：使一个友方随从从战场上移回你的手牌。",
		Mona:          2,
		Damage:        3,
		Hp:            2,
		Quality:       define.CardQualityWhite,
		Ctype:         define.CardTypeEntourage,
		Series:        define.CardSeriseClassic,
		Traits:        MakeCardTraits(define.CardTraitsOnRelease),
		ReleaseFilter: define.CardReleaseFilterMyBattle,
		CanCarry:      true,
	},
	34: &CardConfig{
		Id:       34,
		Name:     "忏悔",
		Desc:     "奥秘：在你的对手使用一张随从牌后，使该随从的生命值降为1。",
		Mona:     1,
		Quality:  define.CardQualityWhite,
		Ctype:    define.CardTypeSorcery,
		Series:   define.CardSeriseClassic,
		Traits:   MakeCardTraits(define.CardTraitsSecret),
		Vocation: MakeCardVocation(define.VocationPaladin),
		CanCarry: true,
	},
	35: &CardConfig{
		Id:            35,
		Name:          "狂野怒火",
		Desc:          "在本回合中，使一个友方野兽获得+2攻击力和免疫。",
		Mona:          1,
		Quality:       define.CardQualityPurple,
		Ctype:         define.CardTypeSorcery,
		Series:        define.CardSeriseClassic,
		Vocation:      MakeCardVocation(define.VocationHunter),
		ReleaseFilter: define.CardReleaseFilterMyBattle,
		CanCarry:      true,
	},
	36: &CardConfig{
		Id:            36,
		Name:          "闪电箭",
		Desc:          "造成3点伤害，过载：（1）",
		Mona:          1,
		Quality:       define.CardQualityWhite,
		Ctype:         define.CardTypeSorcery,
		Race:          MakeCardRace(define.CardRaceNatural),
		Series:        define.CardSeriseClassic,
		Traits:        MakeCardTraits(define.CardTraitsLockMona),
		Vocation:      MakeCardVocation(define.VocationShaman),
		ReleaseFilter: define.CardReleaseFilterAll,
		CanCarry:      true,
	},
	37: &CardConfig{
		Id:       37,
		Name:     "古拉巴什狂暴者",
		Desc:     "每当该随从受到伤害，便获得+3攻击力。",
		Mona:     5,
		Damage:   2,
		Hp:       8,
		Quality:  define.CardQualityWhite,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	38: &CardConfig{
		Id:       38,
		Name:     "熔核巨人",
		Desc:     "你的英雄每受到1点伤害，本牌的法力值消耗便减少（1）点。",
		Mona:     20,
		Damage:   8,
		Hp:       8,
		Quality:  define.CardQualityPurple,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceElement),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	39: &CardConfig{
		Id:       39,
		Name:     "阿曼尼狂战士",
		Desc:     "受伤时具有+3攻击力。",
		Mona:     2,
		Damage:   2,
		Hp:       3,
		Quality:  define.CardQualityWhite,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	40: &CardConfig{
		Id:       40,
		Name:     "冰冻陷阱",
		Desc:     "奥秘：当一个敌方随从攻击时，将其移回拥有者的手牌，并且法力值消耗增加（2）点。",
		Mona:     2,
		Quality:  define.CardQualityWhite,
		Ctype:    define.CardTypeSorcery,
		Series:   define.CardSeriseClassic,
		Traits:   MakeCardTraits(define.CardTraitsSecret),
		Vocation: MakeCardVocation(define.VocationHunter),
		CanCarry: true,
	},
	41: &CardConfig{
		Id:      41,
		Name:    "松鼠",
		Mona:    1,
		Damage:  1,
		Hp:      1,
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeEntourage,
		Race:    MakeCardRace(define.CardRaceBeast),
		Series:  define.CardSeriseClassic,
	},
	42: &CardConfig{
		Id:      42,
		Name:    "魔暴龙",
		Mona:    5,
		Damage:  5,
		Hp:      5,
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeEntourage,
		Race:    MakeCardRace(define.CardRaceBeast),
		Series:  define.CardSeriseClassic,
	},
	43: &CardConfig{
		Id:            43,
		Name:          "工匠大师欧沃斯巴克",
		Desc:          "战吼：随机使另一个随从变形成为一个5/5的魔暴龙或一个1/1的松鼠。",
		Mona:          3,
		Damage:        3,
		Hp:            3,
		Quality:       define.CardQualityOrange,
		Ctype:         define.CardTypeEntourage,
		Series:        define.CardSeriseClassic,
		Traits:        MakeCardTraits(define.CardTraitsOnRelease),
		ReleaseFilter: define.CardReleaseFilterBattle,
		CanCarry:      true,
	},
	44: &CardConfig{
		Id:       44,
		Name:     "希尔瓦娜斯·风行者",
		Desc:     "亡语：随机获得一个敌方随从的控制权。",
		Mona:     6,
		Damage:   5,
		Hp:       5,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceUndead),
		Traits:   MakeCardTraits(define.CardTraitsOnDie),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	45: &CardConfig{
		Id:       45,
		Name:     "生命分流",
		Desc:     "英雄技能抽一张牌并受到2点伤害。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationWarlock),
		Series:   define.CardSeriseClassic,
	},
	46: &CardConfig{
		Id:       46,
		Name:     "稳固射击",
		Desc:     "英雄技能对敌方英雄造成2点伤害。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationHunter),
		Series:   define.CardSeriseClassic,
	},
	47: &CardConfig{
		Id:       47,
		Name:     "图腾召唤",
		Desc:     "英雄技能随机召唤一个图腾。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationShaman),
		Series:   define.CardSeriseClassic,
	},
	48: &CardConfig{
		Id:       48,
		Name:     "空气之怒图腾",
		Desc:     "法术伤害+1",
		Mona:     1,
		Damage:   0,
		Hp:       2,
		ApDamage: 1,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceTotems),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationShaman),
	},
	49: &CardConfig{
		Id:       49,
		Name:     "灼热图腾",
		Mona:     1,
		Damage:   1,
		Hp:       1,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceTotems),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationShaman),
	},
	50: &CardConfig{
		Id:       50,
		Name:     "治疗图腾",
		Desc:     "在你的回合结束时，为所有友方随从恢复1点生命值。",
		Mona:     1,
		Damage:   0,
		Hp:       2,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceTotems),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationShaman),
	},
	51: &CardConfig{
		Id:       51,
		Name:     "石爪图腾",
		Mona:     1,
		Damage:   0,
		Hp:       2,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceTotems),
		Series:   define.CardSeriseClassic,
		Traits:   MakeCardTraits(define.CardTraitsTaunt),
		Vocation: MakeCardVocation(define.VocationShaman),
	},
	52: &CardConfig{
		Id:       52,
		Name:     "力量图腾",
		Desc:     "在你的回合结束时，使另一个友方随从获得+1攻击力。",
		Mona:     1,
		Damage:   0,
		Hp:       2,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceTotems),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationShaman),
	},
	53: &CardConfig{
		Id:       53,
		Name:     "援军",
		Desc:     "英雄技能召唤一个1/1的白银之手新兵。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationPaladin),
		Series:   define.CardSeriseClassic,
	},
	54: &CardConfig{
		Id:       54,
		Name:     "白银之手新兵",
		Mona:     1,
		Damage:   1,
		Hp:       1,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationPaladin),
	},
	55: &CardConfig{
		Id:       55,
		Name:     "变形",
		Desc:     "英雄技能本回合+1攻击力。+1护甲值。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationDruid),
		Series:   define.CardSeriseClassic,
	},
	56: &CardConfig{
		Id:            56,
		Name:          "火焰冲击",
		Desc:          "英雄技能造成1点伤害。",
		Mona:          2,
		Ctype:         define.CardTypeHeroSkill,
		Vocation:      MakeCardVocation(define.VocationMaster),
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterAll,
	},
	57: &CardConfig{
		Id:            57,
		Name:          "次级治疗术",
		Desc:          "英雄技能恢复2点生命值。",
		Mona:          2,
		Ctype:         define.CardTypeHeroSkill,
		Vocation:      MakeCardVocation(define.VocationPastor),
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterAll,
	},
	58: &CardConfig{
		Id:       58,
		Name:     "全副武装！",
		Desc:     "英雄技能获得2点护甲值。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationWarrior),
		Series:   define.CardSeriseClassic,
	},
	59: &CardConfig{
		Id:       59,
		Name:     "山岭巨人",
		Desc:     "你每有一张其他手牌，本牌的法力值消耗便减少（1）点。",
		Mona:     12,
		Damage:   8,
		Hp:       8,
		Quality:  define.CardQualityPurple,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceElement),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	60: &CardConfig{
		Id:       60,
		Name:     "海巨人",
		Desc:     "战场上每有一个其他随从，本牌的法力值消耗便减少（1）点。",
		Mona:     10,
		Damage:   8,
		Hp:       8,
		Quality:  define.CardQualityPurple,
		Ctype:    define.CardTypeEntourage,
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	61: &CardConfig{
		Id:       61,
		Name:     "死亡之翼",
		Desc:     "战吼：消灭所有其他随从，并弃掉你的手牌。",
		Mona:     10,
		Damage:   12,
		Hp:       12,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceDragon),
		Traits:   MakeCardTraits(define.CardTraitsOnRelease),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	62: &CardConfig{
		Id:            62,
		Name:          "炎爆术",
		Desc:          "造成10点伤害。",
		Mona:          10,
		Quality:       define.CardQualityPurple,
		Ctype:         define.CardTypeSorcery,
		Race:          MakeCardRace(define.CardRaceFire),
		Series:        define.CardSeriseClassic,
		Vocation:      MakeCardVocation(define.VocationMaster),
		ReleaseFilter: define.CardReleaseFilterAll,
		CanCarry:      true,
	},
	63: &CardConfig{
		Id:            63,
		Name:          "精神控制",
		Desc:          "获得一个敌方随从的控制权。",
		Mona:          10,
		Quality:       define.CardQualityPurple,
		Ctype:         define.CardTypeSorcery,
		Race:          MakeCardRace(define.CardRaceShadow),
		Series:        define.CardSeriseClassic,
		Vocation:      MakeCardVocation(define.VocationPastor),
		ReleaseFilter: define.CardReleaseFilterEnemyBattle,
		CanCarry:      true,
	},
	64: &CardConfig{
		Id:    64,
		Name:  "Buff",
		Desc:  "我的回合结束时消散和消灭宿主（挂载英雄上就会消灭英雄！）",
		Ctype: define.CardTypeBuff,
	},
	65: &CardConfig{
		Id:    65,
		Name:  "Buff",
		Desc:  "我的回合开始时消散和消灭宿主（挂载英雄上就会消灭英雄！）",
		Ctype: define.CardTypeBuff,
	},
	66: &CardConfig{
		Id:       66,
		Name:     "加拉克苏斯大王",
		Desc:     "战吼：装备一把3/8的血怒。",
		Mona:     9,
		Damage:   0,
		Hp:       5,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeHeroCanRelease,
		Traits:   MakeCardTraits(define.CardTraitsOnRelease),
		Vocation: MakeCardVocation(define.VocationWarlock),
		Series:   define.CardSeriseClassic,
		CanCarry: true,

		IntParam1: 9,
	},
	67: &CardConfig{
		Id:       67,
		Name:     "地狱火！",
		Desc:     "召唤一个6/6的地狱火。",
		Mona:     2,
		Ctype:    define.CardTypeHeroSkill,
		Vocation: MakeCardVocation(define.VocationWarlock),
		Series:   define.CardSeriseClassic,
	},
	68: &CardConfig{
		Id:       68,
		Name:     "地狱火",
		Mona:     6,
		Damage:   6,
		Hp:       6,
		Quality:  define.CardQualityBase,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceDevil),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationWarlock),
	},
	69: &CardConfig{
		Id:       69,
		Name:     "血怒",
		Mona:     3,
		Damage:   3,
		Hp:       8,
		Ctype:    define.CardTypeWeapon,
		Vocation: MakeCardVocation(define.VocationWarlock),
		Series:   define.CardSeriseClassic,
	},
	70: &CardConfig{
		Id:       70,
		Name:     "暴龙王克鲁什",
		Mona:     9,
		Damage:   8,
		Hp:       8,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceBeast),
		Traits:   MakeCardTraits(define.CardTraitsAssault),
		Series:   define.CardSeriseClassic,
		Vocation: MakeCardVocation(define.VocationHunter),
		CanCarry: true,
	},
	71: &CardConfig{
		Id:    71,
		Name:  "英雄",
		Ctype: define.CardTypeHero,
	},
	72: &CardConfig{
		Id:       72,
		Name:     "奥妮克希亚",
		Desc:     "战吼：召唤数条1/1的雏龙，直到你的随从数量达到上限。",
		Mona:     9,
		Damage:   8,
		Hp:       8,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceDragon),
		Traits:   MakeCardTraits(define.CardTraitsOnRelease),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	73: &CardConfig{
		Id:      73,
		Name:    "雏龙",
		Mona:    1,
		Damage:  1,
		Hp:      1,
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeEntourage,
		Race:    MakeCardRace(define.CardRaceDragon),
	},
	74: &CardConfig{
		Id:       74,
		Name:     "诺兹多姆",
		Desc:     "玩家只有15秒的时间来进行他们的回合。",
		Mona:     9,
		Damage:   8,
		Hp:       8,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceDragon),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	75: &CardConfig{
		Id:       75,
		Name:     "玛里苟斯",
		Desc:     "法术伤害+5",
		Mona:     9,
		Damage:   4,
		Hp:       14,
		ApDamage: 5,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceDragon),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	76: &CardConfig{
		Id:            76,
		Name:          "阿莱克丝塔萨",
		Desc:          "战吼：将一方英雄的剩余生命值变为15。",
		Mona:          0,
		Damage:        8,
		Hp:            8,
		Quality:       define.CardQualityOrange,
		Ctype:         define.CardTypeEntourage,
		Race:          MakeCardRace(define.CardRaceDragon),
		Traits:        MakeCardTraits(define.CardTraitsOnRelease),
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterBothHero,
		CanCarry:      true,
	},
	77: &CardConfig{
		Id:       77,
		Name:     "伊瑟拉",
		Desc:     "在你的回合结束时，将一张梦境牌置入你的手牌。",
		Mona:     9,
		Damage:   4,
		Hp:       12,
		Quality:  define.CardQualityOrange,
		Ctype:    define.CardTypeEntourage,
		Race:     MakeCardRace(define.CardRaceDragon),
		Series:   define.CardSeriseClassic,
		CanCarry: true,
	},
	78: &CardConfig{
		Id:            78,
		Name:          "梦魇",
		Desc:          "使一个随从获得+4/+4，在你的下个回合开始时，消灭该随从。",
		Quality:       define.CardQualityBase,
		Ctype:         define.CardTypeSorcery,
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterBattle,
	},
	79: &CardConfig{
		Id:            79,
		Name:          "梦境",
		Desc:          "将一个敌方随从移回你对手的手牌。",
		Quality:       define.CardQualityBase,
		Ctype:         define.CardTypeSorcery,
		Series:        define.CardSeriseClassic,
		ReleaseFilter: define.CardReleaseFilterEnemyBattle,
	},
	80: &CardConfig{
		Id:      80,
		Name:    "伊瑟拉苏醒",
		Desc:    "对除了伊瑟拉之外的所有随从造成 5点伤害。",
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeSorcery,
		Series:  define.CardSeriseClassic,
	},
	81: &CardConfig{
		Id:      81,
		Name:    "欢笑的姐妹",
		Desc:    "无法成为法术或英雄技能的目标。",
		Mona:    2,
		Damage:  3,
		Hp:      5,
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeEntourage,
		Traits:  MakeCardTraits(define.CardTraitsMagicImmunity),
		Series:  define.CardSeriseClassic,
	},
	82: &CardConfig{
		Id:      82,
		Name:    "翡翠幼龙",
		Mona:    4,
		Damage:  7,
		Hp:      6,
		Quality: define.CardQualityBase,
		Ctype:   define.CardTypeEntourage,
		Race:    MakeCardRace(define.CardRaceDragon),
		Series:  define.CardSeriseClassic,
	},
}
