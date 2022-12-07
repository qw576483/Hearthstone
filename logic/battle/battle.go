package battle

import (
	"hs/logic/define"
	"hs/logic/iface"
	"hs/logic/push"
	"math/rand"
	"time"

	"github.com/name5566/leaf/gate"
)

type Battle struct {
	incrCardId  int                 // 自增id
	incrRoundId int                 // 自增回合id
	doneSign    map[string]string   // 完成标记
	status      define.BattleStatus // 状态
	hero        iface.IHero         // 当前回合的英雄
	heros       []iface.IHero       // 保存一下
	randSeed    int64               // 随机数种子
	rand        *rand.Rand          // 随机数句柄
}

// 初始化句柄
func NewBattle(h1, h2 iface.IHero, cs1, cs2 []iface.ICard) iface.IBattle {

	b := &Battle{
		incrCardId:  0,
		incrRoundId: 0,
		doneSign:    make(map[string]string, 0),
	}

	b.randSeed = time.Now().UnixNano()
	b.rand = rand.New(rand.NewSource(b.randSeed))

	h1.Init(cs1, b)
	h2.Init(cs2, b)
	b.heros = []iface.IHero{h1, h2}

	b.PreBegin()

	return b
}

// 获得自增的cardId
func (b *Battle) GetIncrCardId() int {
	b.incrCardId += 1
	return b.incrCardId
}

// 获得回合id
func (b *Battle) GetIncrRoundId() int {
	return b.incrRoundId
}

// 是否完成某项操作
func (b *Battle) GetDoneSign(s string) string {

	if v, ok := b.doneSign[s]; ok {
		return v
	}

	return ""
}

// 完成某项操作
func (b *Battle) SetDoneSign(s, v string) {
	b.doneSign[s] = v
}

// 获得战斗状态
func (b *Battle) GetBattleStatus() define.BattleStatus {
	return b.status
}

// 设置战斗状态
func (b *Battle) SetBattleStatus(bs define.BattleStatus) {
	b.status = bs
	push.PushLine(b)
}

// 获得当前回合的英雄
func (b *Battle) GetRoundHero() iface.IHero {
	return b.hero
}

// 获得英雄根据自增id
func (b *Battle) GetHeroByIncrId(id int) iface.IHero {
	for _, v := range b.heros {
		if v.GetId() == id {
			return v
		}
	}

	return nil
}

// 获得英雄
func (b *Battle) GetHeros() []iface.IHero {
	return b.heros
}

// 获得随机数种子
func (b *Battle) GetRandSeed() int64 {
	return b.randSeed
}

// 随机数句柄
func (b *Battle) GetRand() *rand.Rand {
	return b.rand
}

// 根据连接获得英雄
func (b *Battle) GetHeroByGateAgent(a gate.Agent) iface.IHero {

	heros := b.GetHeros()
	for _, v := range heros {
		if v.GetGateAgent() == a {
			return v
		}
	}

	return nil
}
