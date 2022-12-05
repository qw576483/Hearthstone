package battle

import (
	"hs/logic/config"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/iface"
	"hs/logic/push"
	"sort"
	"strconv"
	"strings"

	"github.com/name5566/leaf/gate"
)

type Hero struct {
	gateAgnet        gate.Agent               // 连接
	battle           iface.IBattle            // 战斗句柄
	id               int                      // 唯一id
	config           *config.HeroConfig       // 配置数据
	enemy            iface.IHero              // 敌人
	preCards         []iface.ICard            // 预存卡牌
	handCards        []iface.ICard            // 手牌
	libCards         []iface.ICard            // 牌库
	graveCards       []iface.ICard            // 坟场
	battleCards      []iface.ICard            // 战场
	allCards         []iface.ICard            // 全部卡牌
	secretCards      []iface.ICard            // 奥秘
	events           map[string][]iface.ICard // 事件
	damage           int                      // 攻击力
	attackTimes      int                      // 攻击次数
	hp               int                      // 血量
	hpMax            int                      // 最大血量
	mona             int                      // 法力值
	monaMax          int                      // 最大法力
	shield           int                      // 护盾
	weapon           iface.ICard              // 武器
	maxHandCardsNum  int                      // 手牌上限数量
	fatigue          int                      // 疲劳伤害
	releaseCardTimes int                      // 本回合出牌次数
}

func (h *Hero) NewPoint() iface.IHero {
	return &Hero{}
}

// 设置链接
func (h *Hero) SetGateAgent(a gate.Agent) {
	h.gateAgnet = a
}

// 获得用户连接
func (h *Hero) GetGateAgent() gate.Agent {
	return h.gateAgnet
}

// 初始化
func (h *Hero) Init(cards []iface.ICard, b iface.IBattle) {

	h.battle = b
	h.id = b.GetIncrCardId()
	h.enemy = nil
	h.preCards = make([]iface.ICard, 0)
	h.handCards = make([]iface.ICard, 0)
	h.libCards = cards
	h.graveCards = make([]iface.ICard, 0)
	h.allCards = make([]iface.ICard, 0)
	h.secretCards = make([]iface.ICard, 0)
	h.events = make(map[string][]iface.ICard, 0)
	h.hp = h.config.Hp
	h.hpMax = h.config.HpMax
	h.mona = h.config.Mona
	h.monaMax = h.config.Mona
	h.shield = h.config.Shield
	h.weapon = nil
	h.maxHandCardsNum = 10

	for _, v := range h.libCards {
		v.Init(v, define.InCardsTypeLib, h, b)
	}
}

// 获得战斗句柄
func (h *Hero) GetBattle() iface.IBattle {
	return h.battle
}

// 获得英雄id
func (h *Hero) GetId() int {
	return h.id
}

// 是否是我的回合
func (h *Hero) IsRoundHero() bool {
	return h.GetBattle().GetRoundHero().GetId() == h.GetId()
}

// 设置配置数据
func (h *Hero) SetConfig(conf *config.HeroConfig) {
	h.config = conf
}

// 获得配置数据
func (h *Hero) GetConfig() *config.HeroConfig {
	return h.config
}

// 设置敌人
func (h *Hero) SetEnemy(eh iface.IHero) {
	h.enemy = eh
}

// 获得敌人
func (h *Hero) GetEnemy() iface.IHero {
	return h.enemy
}

// 获得索引
func (h *Hero) GetIdxByCards(c iface.ICard, cs []iface.ICard) int {

	for k, v := range cs {
		if v.GetId() == c.GetId() {
			return k
		}
	}

	return -1
}

// 获得预存卡牌
func (h *Hero) GetPreCards() []iface.ICard {
	return h.preCards
}

// 设置手牌卡牌
func (h *Hero) SetHandCards(cs []iface.ICard) {
	for _, v := range cs {
		v.SetCardInCardsPos(define.InCardsTypeHand)
	}
	h.handCards = cs
}

// 获得手牌上的卡牌
func (h *Hero) GetHandCards() []iface.ICard {
	return h.handCards
}

// 获得手牌上的卡牌
func (h *Hero) GetHandCardByIncrId(id int) iface.ICard {

	hcs := h.GetHandCards()
	for _, v := range hcs {
		if v.GetId() == id {
			return v
		}
	}

	return nil
}

// 获得牌库中的卡牌
func (h *Hero) GetLibCards() []iface.ICard {
	return h.libCards
}

// 获得坟场卡牌
func (h *Hero) GetGraveCards() []iface.ICard {
	return h.graveCards
}

// 获得战场上的卡牌
func (h *Hero) GetBattleCards() []iface.ICard {
	return h.battleCards
}

// 获得战场上的嘲讽卡牌
func (h *Hero) GetBattleCardsTraitsTauntCardIds() []int {
	tsid := make([]int, 0)

	// 有嘲讽特质，没有潜行特质
	for _, v := range h.GetBattleCards() {
		if v.IsHaveTraits(define.CardTraitsTaunt, v) && !v.IsHaveTraits(define.CardTraitsSneak, v) {
			tsid = append(tsid, v.GetId())
		}
	}
	return tsid
}

// 获得卡牌的位置
func (h *Hero) GetCardIdx(c iface.ICard, cs []iface.ICard) int {

	for k, v := range cs {
		if v.GetId() == c.GetId() {
			return k
		}
	}

	return -1
}

// 添加到全部卡牌
func (h *Hero) AppendToAllCards(c iface.ICard) {
	for _, v := range h.allCards {
		if v.GetId() == c.GetId() {
			return
		}
	}
	h.allCards = append(h.allCards, c)
}

// 获得全部卡牌
func (h *Hero) GetAllCards() []iface.ICard {
	return h.allCards
}

// 获得全部卡牌
func (h *Hero) GetBothAllCards() []iface.ICard {

	ecs := h.GetEnemy().GetAllCards()
	return append(h.allCards, ecs...)
}

// 获得战场上的卡牌
func (h *Hero) GetBattleCardById(id int) iface.ICard {

	bcs := h.GetBattleCards()

	for _, v := range bcs {
		if v.GetId() == id {
			return v
		}
	}

	return nil
}

// 获得战场上的卡牌
func (h *Hero) GetBattleCardsByIds(ids []int) []iface.ICard {

	var cs []iface.ICard = make([]iface.ICard, 0)
	bcs := h.GetBattleCards()

	for _, v := range bcs {
		if help.InArray(v.GetId(), ids) {
			cs = append(cs, v)
		}
	}

	return cs
}

// 获得英雄攻击力
func (h *Hero) GetDamage() int {

	w := h.GetWeapon()
	if w == nil {
		return h.damage
	}

	return h.damage + w.GetHaveEffectDamage(w)
}

// 设置攻击次数
func (h *Hero) SetAttackTimes(t int) {
	h.attackTimes = t
}

// 获得攻击次数
func (h *Hero) GetAttackTimes() int {
	return h.attackTimes
}

// 获得最大攻击次数
func (h *Hero) GetMaxAttackTimes() int {

	w := h.GetWeapon()

	if w != nil && help.InArray(define.CardTraitsWindfury, w.GetTraits()) {
		return 2
	}

	return 1
}

// 获得血量
func (h *Hero) GetHp() int {
	return h.hp
}

// 消耗血量
func (h *Hero) CostHp(num int) int {

	if h.shield >= num {
		h.shield -= num
	} else {
		num = num - h.shield
		h.shield = 0
		h.hp -= num
	}

	if h.hp <= 0 {
		h.Die()
	}

	return num
}

// 获得最大血量
func (h *Hero) GetHpMax() int {
	return h.hpMax
}

// 添加法力值
func (h *Hero) AddMona(add int) {
	h.mona += add
	if h.mona > h.config.MonaMax {
		h.mona = h.config.MonaMax
	}
}

// 消耗法力值
func (h *Hero) CostMona(cost int) bool {
	if h.mona < cost {
		return false
	}
	h.mona -= cost
	return true
}

// 设置法力值
func (h *Hero) SetMona(set int) {
	h.mona = set
	if h.mona > h.config.MonaMax {
		h.mona = h.config.MonaMax
	}
}

// 获得法力值
func (h *Hero) GetMona() int {
	return h.mona
}

// 添加最大法力值
func (h *Hero) AddMonaMax(add int) {
	h.monaMax += add
	if h.monaMax > h.config.MonaMax {
		h.monaMax = h.config.MonaMax
	}
}

// 获得法力值
func (h *Hero) GetMonaMax() int {
	return h.monaMax
}

// 获得护盾
func (h *Hero) GetShield() int {
	return h.shield
}

// 设置武器
func (h *Hero) SetWeapon(c iface.ICard) {
	h.weapon = c
}

// 获得当前武器
func (h *Hero) GetWeapon() iface.ICard {
	return h.weapon
}

// 给一个新卡牌到手上
func (h *Hero) GiveNewCardToHand(configId int) iface.ICard {

	card := iface.GetCardFact().GetCard(configId)
	card.Init(card, define.InCardsTypeNone, h, h.GetBattle())
	h.MoveToHand(card)

	return card
}

// 添加到手牌
func (h *Hero) MoveToHand(c iface.ICard) {

	if len(h.handCards) >= h.GetMaxHandCardsNum() {
		h.DieCard(c)
		return
	}

	// 添加到手牌
	h.handCards = append(h.handCards, c)
	c.SetCardInCardsPos(define.InCardsTypeHand)

	// 触发得到事件
	h.TrickGetCardEvent(c)
}

// 移除手牌
func (h *Hero) MoveOutHandOnlyHandCards(card iface.ICard) {

	idx := -1
	for k, v := range h.GetHandCards() {
		if v == card {
			idx = k
			break
		}
	}

	if idx != -1 {
		_, h.handCards = help.DeleteCardFromCardsByIdx(h.GetHandCards(), idx)
	}

}

// 步入战场
func (h *Hero) MoveToBattle(c iface.ICard, pidx int) {

	h.MoveOutHandOnlyHandCards(c)

	// 满员了
	if len(h.GetBattleCards()) >= define.MaxBattleNum {
		return
	}

	// 位置修正
	if pidx < 0 || pidx > len(h.GetBattleCards()) {
		pidx = len(h.GetBattleCards())
	}

	// 添加到战场
	h.battleCards = help.AddCardToCardsByIdx(h.GetBattleCards(), pidx, c)
	c.SetCardInCardsPos(define.InCardsTypeBattle)

	// 触发效果
	h.TrickPutToBattleEvent(c, pidx)
}

// 移出战场
func (h *Hero) MoveOutBattleOnlyBattleCards(c iface.ICard) int {
	idx := -1
	for k, v := range h.GetBattleCards() {
		if v.GetId() == c.GetId() {
			idx = k
			break
		}
	}

	if idx != -1 {
		_, h.battleCards = help.DeleteCardFromCardsByIdx(h.GetBattleCards(), idx)

		h.TrickOutBattleEvent(c)
	}

	return idx
}

// 卡牌死亡
func (h *Hero) DieCard(c iface.ICard) {

	// 如果在身上或者在场上触发死亡效果
	if c.GetCardInCardsPos() == define.InCardsTypeBattle || c.GetCardInCardsPos() == define.InCardsTypeBody {

		// 如果在场上，则移出场

		battleIdx := h.MoveOutBattleOnlyBattleCards(c)

		// 如果在身上就卸下
		if c.GetCardInCardsPos() == define.InCardsTypeBody {
			c.GetOwner().SetWeapon(nil)
		}

		h.TrickDieCardEvent(c, battleIdx)
	}

	// 触发销毁事件
	h.TrickDevastateCardEvent(c)

	// 设置位置为坟场
	c.SetCardInCardsPos(define.InCardsTypeGrave)
}

// 获得最大手牌上限
func (h *Hero) GetMaxHandCardsNum() int {
	return h.maxHandCardsNum
}

// 预备开始抽卡
func (h *Hero) DrawForPreBegin(t int) {

	for i := 1; i <= t; i++ {
		lcn := len(h.libCards)
		if lcn <= 0 {
			return
		}

		// 随机一张卡
		var card iface.ICard
		idx := h.GetBattle().GetRand().Intn(lcn)

		card, h.libCards = help.DeleteCardFromCardsByIdx(h.GetLibCards(), idx)
		h.preCards = append(h.preCards, card)
	}
}

// 修改预备抽卡
func (h *Hero) ChangePreCrards(putidxs []int) {

	if len(putidxs) <= 0 {
		return
	}

	// 排序一下
	sort.Slice(putidxs, func(i, j int) bool {
		return j > i
	})

	log := "你更换了"

	// 放一张抽一张，保证切片不乱
	for _, v := range putidxs {
		var c iface.ICard
		c, h.preCards = help.DeleteCardFromCardsByIdx(h.preCards, v)

		log += "[第" + strconv.Itoa(v+1) + "张卡牌:" + push.GetCardLogString(c) + " -> "

		lcn := len(h.libCards)
		if lcn <= 0 {
			return
		}

		// 随机一张卡
		var card iface.ICard
		idx := h.GetBattle().GetRand().Intn(lcn)

		card, h.libCards = help.DeleteCardFromCardsByIdx(h.GetLibCards(), idx)
		h.preCards = help.AddCardToCardsByIdx(h.preCards, v, card)

		log += push.GetCardLogString(card) + "],"
	}

	push.PushLog(h, strings.Trim(log, ","))
	push.PushMpm(h)
}

// 抽卡根据次数
func (h *Hero) DrawByTimes(t int) {

	for i := 1; i <= t; i++ {
		lcn := len(h.libCards)
		if lcn <= 0 {

			// 增加疲劳伤害
			f := h.GetFatigue()
			h.SetFatigue(f + 1)

			// 扣血
			h.CostHp(f + 1)
			return
		}

		// 随机一张卡
		var card iface.ICard
		idx := h.GetBattle().GetRand().Intn(lcn)
		card, h.libCards = help.DeleteCardFromCardsByIdx(h.GetLibCards(), idx)

		h.MoveToHand(card)
	}
}

// 获得当前疲劳伤害
func (h *Hero) GetFatigue() int {
	return h.fatigue
}

// 设置当前疲劳伤害
func (h *Hero) SetFatigue(f int) {
	h.fatigue = f
}

// 出牌
func (h *Hero) Release(c iface.ICard, choiceId, putidx int, rc iface.ICard, rh iface.IHero, trickRelease bool) error {

	// 增加出牌次数
	h.releaseCardTimes += 1

	cType := c.GetType()

	if putidx == -1 {
		putidx = len(h.GetBattleCards())
	}

	// 战吼优先触发
	if trickRelease {
		h.TrickRelease(c, choiceId, putidx, rc, rh)
	}

	if cType == define.CardTypeEntourage { // 随从
		h.MoveToBattle(c, putidx)
	} else if cType == define.CardTypeWeapon { // 武器
		h.OnlyReleaseWeapon(c)
	} else if cType == define.CardTypeSorcery {
		h.MoveOutHandOnlyHandCards(c)
	}

	c.SetReleaseRound(c.GetOwner().GetBattle().GetIncrRoundId())

	// 使用完成后销毁法术
	if cType == define.CardTypeSorcery {
		c.SetCardInCardsPos(define.InCardsTypeGrave)
		h.DieCard(c)
	}

	return nil
}

// 仅仅是装备武器
func (h *Hero) OnlyReleaseWeapon(c iface.ICard) {
	w := h.GetWeapon()
	if w != nil {
		h.DieCard(w)
	}
	h.SetWeapon(c)
	h.MoveOutHandOnlyHandCards(c)
	c.SetCardInCardsPos(define.InCardsTypeBody)
}

// 进攻 ， 这里不减次数， 放在battle那边
func (h *Hero) Attack(c, ec iface.ICard, eh iface.IHero) error {

	dmg := c.GetHaveEffectDamage(c)

	var trueCostHp int // 实际伤血
	if ec != nil {     // 如果对手是卡牌

		dmg2 := ec.GetHaveEffectDamage(ec)

		// logs
		push.PushAutoLog(h, push.GetCardLogString(c)+" 对"+push.GetCardLogString(ec)+"造成了"+strconv.Itoa(dmg)+"点伤害")
		push.PushAutoLog(h.GetEnemy(), push.GetCardLogString(ec)+" 对"+push.GetCardLogString(c)+"反击"+strconv.Itoa(dmg2)+"点伤害")

		trueCostHp = ec.CostHp(dmg)
		c.CostHp(dmg2)

	} else if eh != nil { // 如果对手是英雄

		// logs
		push.PushAutoLog(h, push.GetCardLogString(c)+" 对"+push.GetHeroLogString(eh)+"造成了"+strconv.Itoa(dmg)+"伤害")

		trueCostHp = eh.CostHp(dmg)
	}

	h.TrickAfterAttackEvent(c, ec, eh, trueCostHp)

	return nil
}

func (h *Hero) HAttack(ec iface.ICard, eh iface.IHero) error {

	dmg := h.GetDamage()

	var trueCostHp int // 实际伤血
	if ec != nil {     // 如果对手是卡牌

		dmg2 := ec.GetHaveEffectDamage(ec)

		// logs
		push.PushAutoLog(h, push.GetHeroLogString(h)+"对"+push.GetCardLogString(ec)+"造成了"+strconv.Itoa(dmg)+"点伤害")
		push.PushAutoLog(h.GetEnemy(), push.GetCardLogString(ec)+"对"+push.GetHeroLogString(h)+"反击"+strconv.Itoa(dmg2)+"点伤害")

		trueCostHp = ec.CostHp(dmg)
		h.CostHp(dmg2)

	} else if eh != nil { // 如果对手是英雄

		// logs
		push.PushAutoLog(h, push.GetHeroLogString(h)+"对"+push.GetHeroLogString(eh)+"造成了"+strconv.Itoa(dmg)+"伤害")

		trueCostHp = eh.CostHp(dmg)
	}

	if h.GetWeapon() != nil {
		c := h.GetWeapon()
		h.TrickAfterAttackEvent(c, ec, eh, trueCostHp)
	}

	if h.GetWeapon() != nil {
		h.GetWeapon().CostHp(1)
	}

	return nil
}

// 死亡
func (h *Hero) Die() {
}

// 推送数据
func (h *Hero) Push(data interface{}) {
	a := h.GetGateAgent()
	if a != nil {
		a.WriteMsg(data)
	}
}

// 随机战场上的卡牌或者英雄
func (h *Hero) RandBattleCardOrHero() (iface.ICard, iface.IHero) {

	r := h.GetBattle().GetRand()
	bs := h.GetBattleCards()
	rn := r.Intn(len(bs) + 1)

	if rn >= len(bs) {
		return nil, h
	}

	return bs[rn], nil
}

// 随机战场上的卡牌或者英雄
func (h *Hero) RandBothBattleCardOrHero() (iface.ICard, iface.IHero) {

	r := h.GetBattle().GetRand()

	bs := h.GetBattleCards()
	bs = append(bs, h.GetEnemy().GetBattleCards()...)

	rn := r.Intn(len(bs) + 2)

	if rn > len(bs) {
		return nil, h.GetEnemy()
	}

	if rn >= len(bs) {
		return nil, h
	}

	return bs[rn], nil
}

// 随机卡牌，并且排除其中一个卡牌
func (h *Hero) RandExcludeCard(cs []iface.ICard, c iface.ICard) iface.ICard {

	idx := h.GetIdxByCards(c, cs)
	if idx == -1 || len(cs) <= 1 {
		return nil
	}

	ridx := h.GetBattle().GetRand().Intn(len(cs) - 1)
	if ridx >= idx {
		ridx += 1
	}

	return cs[ridx]
}

// 获得本回合出牌次数
func (h *Hero) GetReleaseCardTimes() int {
	return h.releaseCardTimes
}

// 设置本回合出牌次数
func (h *Hero) SetReleaseCardTimes(t int) {
	h.releaseCardTimes = t
}

// 获得事件
func (h *Hero) GetEvent() map[string][]iface.ICard {
	return h.events
}

// 获得事件卡牌
func (h *Hero) GetEventCards(e string) []iface.ICard {
	cs, ok := h.events[e]
	if ok {
		return cs
	}

	return make([]iface.ICard, 0)
}

// 获得双方的事件卡牌
func (h *Hero) GetBothEventCards(e string) []iface.ICard {
	cs := h.GetEventCards(e)
	return append(cs, h.GetEnemy().GetEventCards(e)...)
}

// 添加卡牌到事件
func (h *Hero) AddCardToEvent(c iface.ICard, e string) {
	_, ok := h.events[e]
	if !ok {
		h.events[e] = make([]iface.ICard, 0)
	}

	h.events[e] = append(h.events[e], c)
}

// 删除一个卡牌事件
func (h *Hero) RemoveCardFromEvent(c iface.ICard, e string) {
	es, ok := h.events[e]
	if !ok {
		h.events[e] = make([]iface.ICard, 0)
		return
	}

	for idx, v := range es {
		if v.GetId() == c.GetId() {
			_, h.events[e] = help.DeleteCardFromCardsByIdx(es, idx)
		}
	}
}

// 删除卡牌从双方的事件中
func (h *Hero) RemoveCardFromBothEvent(c iface.ICard) {

	for e := range h.events {
		h.RemoveCardFromEvent(c, e)
	}

	for e := range h.GetEnemy().GetEvent() {
		h.GetEnemy().RemoveCardFromEvent(c, e)
	}
}

// 触发战斗开始
func (h *Hero) TrickBattleBegin() {
	for _, v := range h.GetBothAllCards() {
		v.OnBattleBegin()
	}
}

// 触发战吼
func (h *Hero) TrickRelease(c iface.ICard, choiceId, pidx int, rc iface.ICard, rh iface.IHero) {
	c.OnRelease(choiceId, pidx, rc, rh)
}

// 触发回合开始
func (h *Hero) TrickRoundBegin() {
	for _, v := range h.GetBothEventCards("OnNRRoundBegin") {
		v.OnNRRoundBegin()
	}
}

// 触发回合结束
func (h *Hero) TrickRoundEnd() {
	for _, v := range h.GetBothEventCards("OnNRRoundEnd") {
		v.OnNRRoundEnd()
	}
}

// 触发得到事件
func (h *Hero) TrickGetCardEvent(c iface.ICard) {
	c.OnGet()
}

// 触发销毁事件
func (h *Hero) TrickDevastateCardEvent(c iface.ICard) {
	c.OnDevastate()
}

// 触发攻击后事件
func (h *Hero) TrickAfterAttackEvent(c, ec iface.ICard, eh iface.IHero, trueCostHp int) {

	// 攻击者事件
	if trueCostHp > 0 {
		if ec != nil {
			if ec.GetHaveEffectHp() == 0 && trueCostHp > 0 && !c.IsSilent() {
				c.OnHonorAnnihilate(ec)
			} else if ec.GetHaveEffectHp() < 0 && !c.IsSilent() {
				c.OnOverflowAnnihilate(ec)
			} else if ec.GetHaveEffectHp() > 0 && c.IsHaveTraits(define.CardTraitsHighlyToxic, c) && !c.IsSilent() {
				push.PushAutoLog(h, push.GetCardLogString(c)+" 触发剧毒，"+push.GetCardLogString(ec)+"直接死亡")
				ec.GetOwner().DieCard(ec)
			}
		} else if eh != nil {
			if ec.GetHaveEffectHp() == 0 && !c.IsSilent() {
				c.OnHonorAnnihilate(ec)
			} else if ec.GetHaveEffectHp() < 0 && !c.IsSilent() {
				c.OnOverflowAnnihilate(ec)
			}
		}
	}

}

// 触发死亡事件
func (h *Hero) TrickDieCardEvent(c iface.ICard, bidx int) {

	if !c.IsSilent() {
		c.OnDie(bidx)
	}

	for _, v := range h.GetBothEventCards("OnNROtherDie") {
		v.OnNROtherDie(c)
	}
}

// 触发步入战场事件
func (h *Hero) TrickPutToBattleEvent(c iface.ICard, bidx int) {
	c.OnPutToBattle(bidx)
	for _, v := range h.GetBothEventCards("OnNRPutToBattle") {
		v.OnNRPutToBattle(c)
	}
}

// 触发离开战场事件
func (h *Hero) TrickOutBattleEvent(c iface.ICard) {

	if !c.IsSilent() {
		c.OnOutBattle()
	}

}
