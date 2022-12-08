package push

import (
	"hs/logic/config"
	"hs/logic/define"
	"hs/logic/iface"
	"strconv"
)

// 推送卡牌配置
type CardsConfigMsg struct {
	Configs []*config.CardConfig
}

// 推送赛程
type LineMsg struct {
	Line define.BattleStatus
}

func PushLine(b iface.IBattle) {
	hs := b.GetHeros()
	bs := b.GetBattleStatus()
	for _, v := range hs {
		v.Push(&LineMsg{
			Line: bs,
		})
	}

	if bs == define.BattleStatusPre {
		PushAllLog(b, "==== 战斗准备，进入预留卡牌环节")
	} else if bs == define.BattleStatusRun {
		PushAllLog(b, "==== 战斗开始")
	}
}

// 推送初始数据
type InitMsg struct {
	Em  *EnemyMsg
	Mm  *MyMsg
	Mpm []*CardMsg
}
type EnemyMsg struct {
	Id           int
	Name         string
	Hp           int
	HpMax        int
	Shield       int
	Mona         int
	MonaMax      int
	HandCardsNum int
	LibCardsNum  int
	Weapon       *CardMsg
	Secret       []define.Vocation
}
type MyMsg struct {
	Id          int
	Name        string
	Hp          int
	HpMax       int
	Shield      int
	Mona        int
	MonaMax     int
	HandCards   []*CardMsg
	LibCardsNum int
	Weapon      *CardMsg
	Secret      []*CardMsg
}

type CardMsg struct {
	Id     int
	Name   string
	Mona   int
	Damage int
	Hp     int
}

func BuildEnemyMsg(h iface.IHero) *EnemyMsg {
	return &EnemyMsg{
		Id:           h.GetId(),
		Name:         h.GetConfig().Name,
		Hp:           h.GetHp(),
		HpMax:        h.GetHpMax(),
		Shield:       h.GetShield(),
		Mona:         h.GetMona(),
		MonaMax:      h.GetMonaMax(),
		HandCardsNum: len(h.GetHandCards()),
		LibCardsNum:  len(h.GetLibCards()),
		Weapon:       BuildWeaponMsg(h.GetWeapon()),
		Secret:       BuildEnemySecret(h.GetSecrets()),
	}
}
func BuildMyMsg(h iface.IHero) *MyMsg {
	return &MyMsg{
		Id:          h.GetId(),
		Name:        h.GetConfig().Name,
		Hp:          h.GetHp(),
		HpMax:       h.GetHpMax(),
		Shield:      h.GetShield(),
		Mona:        h.GetMona(),
		MonaMax:     h.GetMonaMax(),
		HandCards:   BuildCardsMsg(h.GetHandCards()),
		LibCardsNum: len(h.GetLibCards()),
		Weapon:      BuildWeaponMsg(h.GetWeapon()),
		Secret:      BuildCardsMsg(h.GetSecrets()),
	}
}

func BuildWeaponMsg(w iface.ICard) *CardMsg {
	if w == nil {
		return nil
	}
	return &CardMsg{
		Id:     w.GetId(),
		Name:   w.GetConfig().Name,
		Damage: w.GetHaveEffectDamage(),
		Hp:     w.GetHaveEffectHp(),
	}
}

func BuildCardsMsg(cs []iface.ICard) []*CardMsg {
	var cm []*CardMsg
	for _, v := range cs {
		cm = append(cm, &CardMsg{
			Id:     v.GetId(),
			Name:   v.GetConfig().Name,
			Mona:   v.GetHaveEffectMona(),
			Damage: v.GetHaveEffectDamage(),
			Hp:     v.GetHaveEffectHp(),
		})
	}
	return cm
}

func BuildEnemySecret(s []iface.ICard) []define.Vocation {

	esm := make([]define.Vocation, 0)
	for _, v := range s {
		cv := v.GetConfig().Vocation
		for _, v2 := range cv {
			esm = append(esm, v2)
			break
		}
	}
	return esm
}

func PushInit(b iface.IBattle) {
	h := b.GetRoundHero()
	e := h.GetEnemy()

	h.Push(
		&InitMsg{
			Em:  BuildEnemyMsg(e),
			Mm:  BuildMyMsg(h),
			Mpm: BuildCardsMsg(h.GetPreCards()),
		},
	)

	e.Push(
		&InitMsg{
			Em:  BuildEnemyMsg(h),
			Mm:  BuildMyMsg(e),
			Mpm: BuildCardsMsg(e.GetPreCards()),
		},
	)
}

// 推送全部数据数据
type InfoMsg struct {
	Em  *EnemyMsg
	Mm  *MyMsg
	Ebm []*CardMsg
	Mbm []*CardMsg
	Mhm []*CardMsg
}

func PushInfoMsg(b iface.IBattle) {
	h := b.GetRoundHero()
	e := h.GetEnemy()

	h.Push(
		&InfoMsg{
			Em:  BuildEnemyMsg(e),
			Mm:  BuildMyMsg(h),
			Ebm: BuildCardsMsg(e.GetBattleCards()),
			Mbm: BuildCardsMsg(h.GetBattleCards()),
			Mhm: BuildCardsMsg(h.GetHandCards()),
		},
	)

	e.Push(
		&InfoMsg{
			Em:  BuildEnemyMsg(h),
			Mm:  BuildMyMsg(e),
			Ebm: BuildCardsMsg(h.GetBattleCards()),
			Mbm: BuildCardsMsg(e.GetBattleCards()),
			Mhm: BuildCardsMsg(e.GetHandCards()),
		},
	)
}

// 推送预选卡牌
type PreCardsMsg struct {
	Mpm []*CardMsg
}

func PushMpm(h iface.IHero) {
	h.Push(
		&PreCardsMsg{
			Mpm: BuildCardsMsg(h.GetPreCards()),
		},
	)
}

// 推送error
type ErrorMsg struct {
	Error string
}

func PushError(h iface.IHero, e string) {
	h.Push(
		&ErrorMsg{
			Error: e,
		},
	)
}

// 推送log
type LogMsg struct {
	Log string
}

func PushLog(h iface.IHero, l string) {
	h.Push(
		&LogMsg{
			Log: l,
		},
	)
}

func PushAllLog(b iface.IBattle, l string) {
	hs := b.GetHeros()
	for _, v := range hs {
		PushLog(v, l)
	}
}

func PushAutoLog(h iface.IHero, l string) {
	PushLog(h, l)
	PushLog(h.GetEnemy(), "【你的对手】"+l)
}

func GetCardLogString(c iface.ICard) string {

	if c.GetType() == define.CardTypeSorcery {
		if c.IsHaveTraits(define.CardTraitsSecret) {
			return "奥秘"
		}
		return c.GetConfig().Name
	}

	return c.GetConfig().Name + "(" + strconv.Itoa(c.GetId()) + ")" + strconv.Itoa(c.GetHaveEffectMona()) + "-" + strconv.Itoa(c.GetHaveEffectDamage()) + "-" + strconv.Itoa(c.GetHaveEffectHp())
}

func GetHeroLogString(h iface.IHero) string {
	return h.GetConfig().Name + "(" + strconv.Itoa(h.GetId()) + ")"
}
