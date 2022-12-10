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
		HeroSkillId: 45,
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
		HeroSkillId: 46,
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
		HeroSkillId: 47,
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
		HeroSkillId: 53,
	},
	5: &HeroConfig{
		Id:          5,
		Vocation:    define.VocationDruid,
		Name:        "玛法里奥",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 55,
	},
	6: &HeroConfig{
		Id:          6,
		Vocation:    define.VocationMaster,
		Name:        "吉安娜",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 56,
	},
	7: &HeroConfig{
		Id:          7,
		Vocation:    define.VocationPastor,
		Name:        "安度因",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 57,
	},
	8: &HeroConfig{
		Id:          8,
		Vocation:    define.VocationWarrior,
		Name:        "加尔鲁什",
		Hp:          30,
		HpMax:       30,
		Mona:        0,
		MonaMax:     10,
		Shield:      0,
		HeroSkillId: 58,
	},
}
