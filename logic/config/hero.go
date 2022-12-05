package config

import "hs/logic/define"

type HeroConfig struct {
	Id          int             // id
	Vocation    define.Vocation // 职业
	Name        string          // 名字
	Hp          int             // 血量
	HpMax       int             // 最大血量
	Mona        int             // 能量
	MonaMax     int             // 最大能量
	Shield      int             // 护盾
	HeroSkillId int             // 技能
}

// 获得配置
func GetHeroConfig(configId int) *HeroConfig {
	return defineHeroConfig[configId]
}

var defineHeroConfig []*HeroConfig = []*HeroConfig{
	0: &HeroConfig{
		Id:          0,
		Vocation:    define.VocationRobbers,
		Name:        "瓦莉拉",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 26,
	},
	1: &HeroConfig{
		Id:          1,
		Vocation:    define.VocationWarlock,
		Name:        "古尔丹",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 26,
	},
	2: &HeroConfig{
		Id:          2,
		Vocation:    define.VocationHunter,
		Name:        "雷克萨",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 26,
	},
	3: &HeroConfig{
		Id:          3,
		Vocation:    define.VocationShaman,
		Name:        "萨尔",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 26,
	},
	4: &HeroConfig{
		Id:          4,
		Vocation:    define.VocationPaladin,
		Name:        "乌瑟尔",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 26,
	},
}
