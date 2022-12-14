package bhero

import (
	"errors"
	"hs/logic/config"
	"hs/logic/define"
	"hs/logic/help"
	"hs/logic/iface"
	"hs/logic/push"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/name5566/leaf/gate"
)

type Hero struct {
	gateAgnet        gate.Agent         // 连接
	battle           iface.IBattle      // 战斗句柄
	config           *config.HeroConfig // 配置数据
	skill            iface.ICard        // 英雄技能
	head             iface.ICard        // 我的实体卡牌
	weapon           iface.ICard        // 武器
	enemy            iface.IHero        // 敌人
	preCards         []iface.ICard      // 预存卡牌
	handCards        []iface.ICard      // 手牌
	libCards         []iface.ICard      // 牌库
	graveCards       []iface.ICard      // 坟场
	battleCards      []iface.ICard      // 战场
	allCards         []iface.ICard      // 全部卡牌
	secretCards      []iface.ICard      // 奥秘
	mona             int                // 法力值
	monaMax          int                // 最大法力
	lockMona         int                // 锁定法力值
	lockMonaCache    int                // 下回合锁定的法力值
	maxHandCardsNum  int                // 手牌上限数量
	fatigue          int                // 疲劳伤害
	releaseCardTimes int                // 本回合出牌次数
	timer            *time.Timer        // 定时

	roundDieCards []iface.ICard // 回合死亡卡牌
}

func (h *Hero) NewPoint() iface.IHero {
	return &Hero{}
}

func (h *Hero) GetId() int {
	return h.GetHead().GetId()
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
func (h *Hero) Init(ih iface.IHero, cards []iface.ICard, b iface.IBattle) {

	nc := iface.GetCardFact().GetCard(define.HeroId)
	nc.Init(nc, define.InCardsTypeHead, h, b)
	nc.SetHpMax(h.config.HpMax)
	nc.SetHp(h.config.Hp)

	h.battle = b
	h.head = nc
	h.enemy = nil
	h.preCards = make([]iface.ICard, 0)
	h.handCards = make([]iface.ICard, 0)
	h.libCards = cards
	h.graveCards = make([]iface.ICard, 0)
	h.allCards = make([]iface.ICard, 0)
	h.secretCards = make([]iface.ICard, 0)
	h.mona = h.config.Mona
	h.monaMax = h.config.Mona
	h.weapon = nil
	h.maxHandCardsNum = 10

	skill := iface.GetCardFact().GetCard(h.config.HeroSkillId)
	skill.Init(skill, define.InCardsTypeNone, h, b)
	h.skill = skill

	for _, v := range h.libCards {
		v.Init(v, define.InCardsTypeLib, h, b)
	}
}

// 获得战斗句柄
func (h *Hero) GetBattle() iface.IBattle {
	return h.battle
}

func (h *Hero) GetHead() iface.ICard {
	return h.head
}

// 是否是我的回合
func (h *Hero) IsRoundHero() bool {
	return h.GetBattle().GetRoundHero().GetId() == h.GetId()
}

// 设置英雄技能
func (h *Hero) SetHeroSkill(c iface.ICard) {
	h.skill = c
}

// 获得英雄技能
func (h *Hero) GetHeroSkill() iface.ICard {
	return h.skill
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
		if v.IsHaveTraits(define.CardTraitsTaunt) && !v.IsHaveTraits(define.CardTraitsSneak) {
			tsid = append(tsid, v.GetId())
		}
	}
	return tsid
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

// 回合死亡卡牌
func (h *Hero) GetRoundDieCards() []iface.ICard {
	return h.roundDieCards
}

// 卡牌复制出来新的一份内存地址
func (h *Hero) CardsToNewInstance(cs []iface.ICard) []iface.ICard {
	cs2 := make([]iface.ICard, 0)

	return append(cs2, cs...)
}

// 获得可选择的卡牌，根据id
func (h *Hero) GetCanSelectCardId(id int) iface.ICard {

	if h.head != nil && h.head.GetId() == id {
		return h.head
	}

	return h.GetBattleCardById(id)
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

// 获得法术伤害
func (h *Hero) GetApDamage() int {

	d := 0
	for _, v := range h.GetBattleCards() {
		d += v.GetApDamage()
	}

	for _, v := range h.GetBattle().GetEventCards("OnNROtherGetApDamage") {
		d += v.OnNROtherGetApDamage(h)
	}

	return d
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

	if h.mona < 0 {
		h.mona = 0
	}
}

// 获得法力值
func (h *Hero) GetMona() int {
	mona := h.mona - h.lockMona
	if mona < 0 {
		mona = 0
	}
	return mona
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

// 获得锁定法力值
func (h *Hero) GetLockMona() int {
	return h.lockMona
}

// 设置锁定法力值
func (h *Hero) SetLockMona(lm int) {
	h.lockMona = lm
}

// 获得锁定法力值缓存
func (h *Hero) GetLockMonaCache() int {
	return h.lockMonaCache
}

// 设置锁定法力值缓存
func (h *Hero) SetLockMonaCache(lmc int) {
	h.lockMonaCache = lmc
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
func (h *Hero) MoveToHand(c iface.ICard) bool {

	if len(h.handCards) >= h.GetMaxHandCardsNum() {
		push.PushAutoLog(h, "手牌满了")
		h.DieCard(c, false)
		return false
	}

	// 如果从战场中移回手牌，需要还原数据
	if c.GetCardInCardsPos() == define.InCardsTypeBattle {
		h.MoveOutBattleOnlyBattleCards(c)
		c.Reset()
	}

	// 添加到手牌
	h.handCards = append(h.handCards, c)
	c.SetCardInCardsPos(define.InCardsTypeHand)

	return true
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
func (h *Hero) MoveToBattle(c iface.ICard, bidx int) {

	if c.GetReleaseRound() == 0 {
		c.SetReleaseRound(h.GetBattle().GetIncrRoundId())
	}

	h.MoveOutHandOnlyHandCards(c)

	// 满员了
	if len(h.GetBattleCards()) >= define.MaxBattleNum {

		// 如果从战场上移动到战场上，触发死亡
		if c.GetCardInCardsPos() == define.InCardsTypeBattle {
			push.PushAutoLog(h, "战场满了")
			h.DieCard(c, true)
		}
		return
	}

	// 位置修正
	if bidx < 0 || bidx > len(h.GetBattleCards()) {
		bidx = len(h.GetBattleCards())
	}

	// 添加到战场
	h.battleCards = help.AddCardToCardsByIdx(h.GetBattleCards(), bidx, c)
	c.SetCardInCardsPos(define.InCardsTypeBattle)

	// 触发效果
	h.TrickPutToBattleEvent(c, bidx)
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

// 夺取卡牌
func (h *Hero) CaptureCard(c iface.ICard, bidx int) {

	if c == nil {
		return
	}

	eh := h.GetEnemy()
	c.SetOwner(h)
	if c.GetCardInCardsPos() == define.InCardsTypeBattle {

		eh.MoveOutBattleOnlyBattleCards(c)
		h.MoveToBattle(c, bidx)
	}

	// 未来在做夺取牌库
	// else if c.GetCardInCardsPos() == define.InCardsTypeLib {

	// }
}

// 丢弃手牌
func (h *Hero) DiscardCard(c iface.ICard) {
	h.MoveOutHandOnlyHandCards(c)
	c.OnAfterDisCard()

	push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"被丢弃")
}

// 卡牌死亡
func (h *Hero) DieCard(c iface.ICard, immediatelyDie bool) {

	if c.GetType() == define.CardTypeEntourage {
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"死亡")
	} else if c.GetType() == define.CardTypeWeapon {
		push.PushAutoLog(c.GetOwner(), push.GetCardLogString(c)+"破碎")
	} else if c.GetType() == define.CardTypeHero {
		h.Die()
		return
	}

	// 如果在身上或者在场上触发死亡效果
	if c.GetCardInCardsPos() == define.InCardsTypeBattle || c.GetCardInCardsPos() == define.InCardsTypeBody {

		// 如果在身上就卸下
		if c.GetCardInCardsPos() == define.InCardsTypeBody {
			c.GetOwner().SetWeapon(nil)
		} else if c.GetCardInCardsPos() == define.InCardsTypeBattle {
			// 如果在场上，则移出场
			battleIdx := h.MoveOutBattleOnlyBattleCards(c)
			c.SetAfterDieBidx(battleIdx)
		}

		// 是否立即触发
		if immediatelyDie {
			h.TrickDieCardEvent(c)
		} else {
			h.GetBattle().RecordCardDie(c)
		}

		// 回合结束卡牌
		h.roundDieCards = append(h.roundDieCards, c)

		// 进入坟场
		// c.Reset()
		h.graveCards = append(h.graveCards, c)
	}

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
func (h *Hero) DrawByTimes(t int) []iface.ICard {

	push.PushAutoLog(h, "抽了"+strconv.Itoa(t)+"张牌")
	dcs := make([]iface.ICard, 0)
	for i := 1; i <= t; i++ {
		lcn := len(h.libCards)
		if lcn <= 0 {

			// 增加疲劳伤害
			f := h.GetFatigue()
			h.SetFatigue(f + 1)

			// 扣血
			h.GetHead().CostHp(h.GetHead(), f+1)

			push.PushAutoLog(h, "牌库没有牌了！受到了疲劳伤害"+strconv.Itoa(f+1))
			continue
		}

		// 随机一张卡
		var card iface.ICard
		idx := h.GetBattle().GetRand().Intn(lcn)
		card, h.libCards = help.DeleteCardFromCardsByIdx(h.GetLibCards(), idx)

		h.MoveToHand(card)

		dcs = append(dcs, card)
	}

	return dcs
}

// 抽卡，根据卡牌
func (h *Hero) DrawByCard(dc iface.ICard) {

	for idx, v := range h.libCards {
		if v.GetId() == dc.GetId() {
			_, h.libCards = help.DeleteCardFromCardsByIdx(h.GetLibCards(), idx)
			h.MoveToHand(v)
			return
		}
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
func (h *Hero) Release(c iface.ICard, choiceId, putidx int, rc iface.ICard, trickRelease bool) error {

	cType := c.GetType()

	if putidx == -1 {
		putidx = len(h.GetBattleCards())
	}

	valid := true

	// 其他卡牌释放前
	for _, v := range h.GetBattle().GetEventCards("OnNROtherBeforeRelease") {
		rc, valid = v.OnNROtherBeforeRelease(c, rc)
		if !valid {
			break
		}
	}
	h.GetBattle().WhileTrickCardDie()

	// 如果触发效果
	if trickRelease {

		// 如果法术被拦截
		if cType == define.CardTypeSorcery && !valid {

		} else {
			// 战吼不拦截
			h.TrickRelease(c, choiceId, putidx, rc)
		}
	}

	if valid {
		if cType == define.CardTypeEntourage { // 随从
			h.MoveToBattle(c, putidx)
		} else if cType == define.CardTypeWeapon { // 武器
			h.OnlyReleaseWeapon(c)
		} else if cType == define.CardTypeSorcery || cType == define.CardTypeHeroCanRelease { // 法术 , 英雄卡

			// 如果法术不在身上，强制置为战场上
			if c.GetCardInCardsPos() != define.InCardsTypeBody {
				c.SetCardInCardsPos(define.InCardsTypeBattle)
			}
			h.MoveOutHandOnlyHandCards(c)
		}
	}

	// 如果触发效果
	if trickRelease {

		// 如果法术被拦截
		if cType == define.CardTypeSorcery && !valid {

		} else {
			// 战吼不拦截
			h.TrickRelease2(c, choiceId, putidx, rc)
		}
	}

	// 其他卡牌释放后
	for _, v := range h.GetBattle().GetEventCards("OnNROtherAfterRelease") {
		v.OnNROtherAfterRelease(c)
	}
	h.GetBattle().WhileTrickCardDie()

	c.SetReleaseRound(c.GetOwner().GetBattle().GetIncrRoundId())

	// 如果是法术，没被拦截，则法术进入坟场
	if cType == define.CardTypeSorcery || !valid {
		h.DieCard(c, false)
	}

	// if !valid{
	// 移出手牌
	h.MoveOutHandOnlyHandCards(c)
	// }

	h.GetBattle().WhileTrickCardDie()

	return nil
}

// 仅仅是装备武器
func (h *Hero) OnlyReleaseWeapon(c iface.ICard) {

	push.PushAutoLog(h, "装备了"+push.GetCardLogString(c))

	w := h.GetWeapon()
	if w != nil {
		h.DieCard(w, false)
	}
	h.SetWeapon(c)
	h.MoveOutHandOnlyHandCards(c)
	c.SetCardInCardsPos(define.InCardsTypeBody)
	c.OnWear()
}

// 进攻 ， 这里不减次数， 放在battle那边
func (h *Hero) Attack(c, ec iface.ICard) error {

	if ec == nil {
		return errors.New("未找到敌人")
	}

	// 攻击前
	ec = h.TrickBeforeAttackEvent(c, ec)
	if ec == nil {
		h.TrickAfterAttackEvent(c, ec, 0)
		return nil
	}

	h.GetBattle().WhileTrickCardDie()

	// 伤害可能和血量挂钩，所以先取
	dmg := c.GetHaveEffectDamage()
	dmg2 := ec.GetHaveEffectDamage()

	// 是否有狂战斧
	var cLeft iface.ICard
	var dmgLeft int

	var cRight iface.ICard
	var dmgRight int
	if (c.GetType() == define.CardTypeEntourage && c.IsHaveTraits(define.CardTraitsBattlefury)) ||
		(c.GetType() == define.CardTypeHero && h.GetWeapon() != nil && h.GetWeapon().IsHaveTraits(define.CardTraitsBattlefury)) {

		// 对方不是英雄
		if ec.GetType() != define.CardTypeHero {

			eh := ec.GetOwner()
			bcs := eh.GetBattleCards()

			ecIdx := eh.GetIdxByCards(ec, bcs)

			if (ecIdx - 1) >= 0 {
				cLeft = bcs[ecIdx-1]
			}
			if (ecIdx + 1) < len(bcs) {
				cRight = bcs[ecIdx+1]
			}
		}
	}

	// logs
	dmg = ec.CostHp(c, dmg)

	// 狂战斧效果
	if cLeft != nil {
		push.PushAutoLog(h, "触发狂战斧，分裂目标："+push.GetCardLogString(cLeft))
		dmgLeft = cLeft.CostHp(c, dmg)
	}

	if cRight != nil {
		push.PushAutoLog(h, "触发狂战斧，分裂目标："+push.GetCardLogString(cRight))
		dmgRight = cRight.CostHp(c, dmg)
	}

	// 反击
	if dmg2 > 0 && ec.GetType() != define.CardTypeHero {
		push.PushAutoLog(h.GetEnemy(), push.GetCardLogString(ec)+"反击")
		c.CostHp(ec, dmg2)
	}

	// 攻击后
	h.TrickAfterAttackEvent(c, ec, dmg)

	if cLeft != nil {
		h.TrickAfterAttackEvent(c, ec, dmgLeft)
	}
	if cRight != nil {
		h.TrickAfterAttackEvent(c, ec, dmgRight)
	}

	if c.GetType() == define.CardTypeHero {
		if h.GetWeapon() != nil {
			h.GetWeapon().CostHp(h.GetWeapon(), 1)
		}
	}

	h.GetBattle().WhileTrickCardDie()

	return nil
}

// 死亡
func (h *Hero) Die() {

	// 设置战斗状态
	h.battle.SetBattleStatus(define.BattleStatusEnd)

	// log
	push.PushAutoLog(h, "已死亡")
}

// 推送数据
func (h *Hero) Push(data interface{}) {
	a := h.GetGateAgent()
	if a != nil {
		a.WriteMsg(data)
	}
}

// 随机战场上的卡牌
func (h *Hero) RandBothBattleCard() iface.ICard {

	cs := h.GetBattleCards()
	cs = append(cs, h.GetEnemy().GetBattleCards()...)

	return h.RandCard(cs)
}

// 随机战场上的卡牌或者英雄
func (h *Hero) RandBattleCardOrHero() iface.ICard {

	cs := h.GetBattleCards()
	cs = append(cs, h.GetHead())
	return h.RandCard(cs)
}

// 随机战场上的卡牌或者英雄
func (h *Hero) RandBothBattleCardOrHero() iface.ICard {

	cs := h.GetBattleCards()
	cs = append(cs, h.GetEnemy().GetBattleCards()...)
	cs = append(cs, h.GetHead())
	cs = append(cs, h.GetEnemy().GetHead())

	return h.RandCard(cs)
}

// 随机战场上的受伤的卡牌或者英雄
func (h *Hero) RandBothInjuredBattleCardOrHero() iface.ICard {
	var cs []iface.ICard

	for _, v := range h.GetBattleCards() {
		if v.GetHaveEffectHp() < v.GetHaveEffectHpMax() {
			cs = append(cs, v)
		}
	}

	for _, v := range h.GetEnemy().GetBattleCards() {
		if v.GetHaveEffectHp() < v.GetHaveEffectHpMax() {
			cs = append(cs, v)
		}
	}

	if h.GetHead().GetHaveEffectHp() < h.GetHead().GetHaveEffectHpMax() {
		cs = append(cs, h.GetHead())
	}

	if h.GetEnemy().GetHead().GetHaveEffectHp() < h.GetEnemy().GetHead().GetHaveEffectHpMax() {
		cs = append(cs, h.GetEnemy().GetHead())
	}

	return h.RandCard(cs)
}

// 随机卡牌
func (h *Hero) RandCard(cs []iface.ICard) iface.ICard {

	r := h.GetBattle().GetRand()
	if len(cs) <= 0 {
		return nil
	}

	idx := r.Intn(len(cs))

	return cs[idx]
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

// 获得全部奥秘
func (h *Hero) GetSecrets() []iface.ICard {
	return h.secretCards
}

// 是否能释放奥秘
func (h *Hero) CanReleaseSecret(ic iface.ICard) bool {

	// 奥秘满了
	if len(h.secretCards) >= 5 {
		return false
	}

	// 奥秘有相同的
	for _, v := range h.secretCards {
		if v.GetConfig().Id == ic.GetConfig().Id {
			return false
		}
	}

	return true
}

// 仅仅释放奥秘
func (h *Hero) OnlyReleaseSecret(ic iface.ICard) bool {

	if !h.CanReleaseSecret(ic) {
		return false
	}

	h.secretCards = append(h.secretCards, ic)

	return true
}

// 删除奥秘 , 是否是触发奥秘的删除
func (h *Hero) DeleteSecret(ic iface.ICard, isTigger bool) {

	for idx, v := range h.secretCards {
		if v.GetId() == ic.GetId() {
			_, h.secretCards = help.DeleteCardFromCardsByIdx(h.secretCards, idx)
		}
	}

	if isTigger {

		push.PushLog(h, "触发了奥秘"+ic.GetConfig().Name)
		for _, v := range h.GetBattle().GetEventCards("OnNROtherSecretTigger") {
			v.OnNROtherSecretTigger(ic)
		}
	}

	h.GetBattle().RemoveCardFromAllEvent(ic)
}

// 一个新的倒计时
func (h *Hero) NewCountDown(second int) {

	oldTimer := h.timer
	if oldTimer != nil {
		oldTimer.Stop()
	}

	nTimer := time.NewTimer(time.Duration(second) * time.Second)
	h.timer = nTimer

	go func(h *Hero) {
		<-nTimer.C
		h.FixRoundEnd()
	}(h)
}

// 关闭倒计时
func (h *Hero) CloseCountDown() {
	if h.timer == nil {
		return
	}

	h.timer.Stop()
	h.timer = nil
}

// 变身到卡牌
func (h *Hero) Henshin(c iface.ICard) {

	push.PushAutoLog(h, "变身成了"+c.GetConfig().Name)

	heroId := c.GetConfig().IntParam1

	// 替换config
	h.config = config.GetHeroConfig(heroId)

	// 加护盾
	h.GetHead().SetShield(h.GetHead().GetShield() + c.GetHp())

	// 替换英雄技能
	skill := iface.GetCardFact().GetCard(h.config.HeroSkillId)
	skill.Init(skill, define.InCardsTypeNone, h, h.GetBattle())
	h.skill = skill
}
