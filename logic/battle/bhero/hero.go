package bhero

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
	gateAgnet        gate.Agent         // 连接
	battle           iface.IBattle      // 战斗句柄
	id               int                // 唯一id
	realization      iface.IHero        // 实现
	config           *config.HeroConfig // 配置数据
	skill            iface.ICard        // 英雄技能
	enemy            iface.IHero        // 敌人
	preCards         []iface.ICard      // 预存卡牌
	handCards        []iface.ICard      // 手牌
	libCards         []iface.ICard      // 牌库
	graveCards       []iface.ICard      // 坟场
	battleCards      []iface.ICard      // 战场
	allCards         []iface.ICard      // 全部卡牌
	secretCards      []iface.ICard      // 奥秘
	damage           int                // 攻击力
	attackTimes      int                // 攻击次数
	hp               int                // 血量
	hpMax            int                // 最大血量
	mona             int                // 法力值
	monaMax          int                // 最大法力
	lockMona         int                // 锁定法力值
	lockMonaCache    int                // 下回合锁定的法力值
	shield           int                // 护盾
	weapon           iface.ICard        // 武器
	maxHandCardsNum  int                // 手牌上限数量
	fatigue          int                // 疲劳伤害
	releaseCardTimes int                // 本回合出牌次数
	subCards         []iface.ICard      // buff
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
func (h *Hero) Init(ih iface.IHero, cards []iface.ICard, b iface.IBattle) {

	h.battle = b
	h.id = b.GetIncrCardId()
	h.enemy = nil
	h.preCards = make([]iface.ICard, 0)
	h.handCards = make([]iface.ICard, 0)
	h.libCards = cards
	h.graveCards = make([]iface.ICard, 0)
	h.allCards = make([]iface.ICard, 0)
	h.secretCards = make([]iface.ICard, 0)
	h.hp = h.config.Hp
	h.hpMax = h.config.HpMax
	h.mona = h.config.Mona
	h.monaMax = h.config.Mona
	h.shield = h.config.Shield
	h.weapon = nil
	h.maxHandCardsNum = 10
	h.realization = ih

	skill := iface.GetCardFact().GetCard(h.config.HeroSkillId)
	skill.Init(skill, define.InCardsTypeNone, h, b)
	h.skill = skill

	for _, v := range h.libCards {
		v.Init(v, define.InCardsTypeLib, h, b)
	}
}

// 获得实现
func (h *Hero) GetRealization() iface.IHero {
	return h.realization
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

	return h.damage + w.GetHaveEffectDamage()
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

	if num > 0 && h.IsHaveTraits(define.CardTraitsImmune) {
		num = 0
		push.PushAutoLog(h, push.GetHeroLogString(h)+"具有免疫，伤害无效")
	}

	if num > 0 {
		if h.shield >= num {
			h.shield -= num
		} else {
			num = num - h.shield
			h.shield = 0
			h.hp -= num
		}
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
		push.PushAutoLog(h, "手牌满了")
		h.DieCard(c, false)
		return
	}

	// 如果从战场中移回手牌，需要还原数据
	if c.GetCardInCardsPos() == define.InCardsTypeBattle {
		h.MoveOutBattleOnlyBattleCards(c)
		c.Reset()
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
func (h *Hero) MoveToBattle(c iface.ICard, bidx int) {

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

// 卡牌死亡
func (h *Hero) DieCard(c iface.ICard, immediatelyDie bool) {

	// 如果在身上或者在场上触发死亡效果
	if c.GetCardInCardsPos() == define.InCardsTypeBattle || c.GetCardInCardsPos() == define.InCardsTypeBody {

		// 如果在身上就卸下
		if c.GetCardInCardsPos() == define.InCardsTypeBody {
			c.GetOwner().SetWeapon(nil)
		} else {
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
func (h *Hero) DrawByTimes(t int) {

	for i := 1; i <= t; i++ {
		lcn := len(h.libCards)
		if lcn <= 0 {

			// 增加疲劳伤害
			f := h.GetFatigue()
			h.SetFatigue(f + 1)

			// 扣血
			h.CostHp(f + 1)

			push.PushAutoLog(h, "牌库没有牌了！受到了疲劳伤害"+strconv.Itoa(f+1))
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

	cType := c.GetType()

	if putidx == -1 {
		putidx = len(h.GetBattleCards())
	}

	valid := true

	// 其他卡牌释放前
	for _, v := range h.GetBattle().GetEventCards("OnNROtherBeforeRelease") {
		rc, rh, valid = v.OnNROtherBeforeRelease(c, rc, rh)
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
			h.TrickRelease(c, choiceId, putidx, rc, rh)
		}
	}

	if valid {
		if cType == define.CardTypeEntourage { // 随从
			h.MoveToBattle(c, putidx)
		} else if cType == define.CardTypeWeapon { // 武器
			h.OnlyReleaseWeapon(c)
		} else if cType == define.CardTypeSorcery { // 法术

			// 如果法术不在身上，强制置为战场上
			if c.GetCardInCardsPos() != define.InCardsTypeBody {
				c.SetCardInCardsPos(define.InCardsTypeBattle)
			}
			h.MoveOutHandOnlyHandCards(c)
		}
	}

	// 其他卡牌释放后
	for _, v := range h.GetBattle().GetEventCards("OnNROtherAfterRelease") {
		v.OnNROtherAfterRelease(c)
	}
	h.GetBattle().WhileTrickCardDie()

	c.SetReleaseRound(c.GetOwner().GetBattle().GetIncrRoundId())

	if cType == define.CardTypeSorcery || !valid {
		h.DieCard(c, false)
	}

	h.GetBattle().WhileTrickCardDie()

	return nil
}

// 仅仅是装备武器
func (h *Hero) OnlyReleaseWeapon(c iface.ICard) {
	w := h.GetWeapon()
	if w != nil {
		h.DieCard(w, false)
	}
	h.SetWeapon(c)
	h.MoveOutHandOnlyHandCards(c)
	c.SetCardInCardsPos(define.InCardsTypeBody)
}

// 进攻 ， 这里不减次数， 放在battle那边
func (h *Hero) Attack(c, ec iface.ICard, eh iface.IHero) error {

	// 攻击前
	for _, v := range h.GetBattle().GetEventCards("OnNROtherBeforeAttack") {
		ec, eh = v.OnNROtherBeforeAttack(c, ec, eh)
	}
	h.GetBattle().WhileTrickCardDie()

	dmg := c.GetHaveEffectDamage()
	if ec != nil { // 如果对手是卡牌

		// 伤害可能和血量挂钩，所以先取
		dmg2 := ec.GetHaveEffectDamage()

		// logs
		push.PushAutoLog(h, push.GetCardLogString(c)+" 对"+push.GetCardLogString(ec)+"造成了"+strconv.Itoa(dmg)+"点伤害")
		dmg = ec.CostHp(dmg)

		push.PushAutoLog(h.GetEnemy(), push.GetCardLogString(ec)+" 对"+push.GetCardLogString(c)+"反击"+strconv.Itoa(dmg2)+"点伤害")
		c.CostHp(dmg2)

	} else if eh != nil { // 如果对手是英雄

		// logs
		push.PushAutoLog(h, push.GetCardLogString(c)+" 对"+push.GetHeroLogString(eh)+"造成了"+strconv.Itoa(dmg)+"伤害")
		dmg = eh.CostHp(dmg)
	}

	// 攻击后
	if ec != nil && eh != nil {
		h.TrickAfterAttackEvent(c, ec, eh, dmg)
	}

	h.GetBattle().WhileTrickCardDie()

	return nil
}

func (h *Hero) HAttack(ec iface.ICard, eh iface.IHero) error {

	dmg := h.GetDamage()

	if ec != nil { // 如果对手是卡牌

		// 伤害可能和血量挂钩，所以先取
		dmg2 := ec.GetHaveEffectDamage()

		// logs
		push.PushAutoLog(h, push.GetHeroLogString(h)+"对"+push.GetCardLogString(ec)+"造成了"+strconv.Itoa(dmg)+"点伤害")
		dmg = ec.CostHp(dmg)

		push.PushAutoLog(h.GetEnemy(), push.GetCardLogString(ec)+"对"+push.GetHeroLogString(h)+"反击"+strconv.Itoa(dmg2)+"点伤害")
		h.CostHp(dmg2)

	} else if eh != nil { // 如果对手是英雄

		// logs
		push.PushAutoLog(h, push.GetHeroLogString(h)+"对"+push.GetHeroLogString(eh)+"造成了"+strconv.Itoa(dmg)+"伤害")

		dmg = eh.CostHp(dmg)
	}

	if h.GetWeapon() != nil {
		c := h.GetWeapon()
		h.TrickAfterAttackEvent(c, ec, eh, dmg)
	}

	if h.GetWeapon() != nil {
		h.GetWeapon().CostHp(1)
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

// 删除奥秘
func (h *Hero) DeleteSecret(ic iface.ICard, istTigger bool) {
	for idx, v := range h.secretCards {
		if v.GetId() == ic.GetId() {
			_, h.secretCards = help.DeleteCardFromCardsByIdx(h.secretCards, idx)
		}
	}
}

// 获得子卡牌
func (h *Hero) GetSubCards() []iface.ICard {
	return h.subCards
}

// 获得特质
func (h *Hero) GetTraits() []define.CardTraits {

	// 获得子卡牌的特质
	ts := make([]define.CardTraits, 0)
	for _, v := range h.GetSubCards() {
		for _, ct := range v.GetTraits() {

			if !help.InArray(ct, ts) {
				ts = append(ts, ct)
			}
		}
	}

	// 获得光环影响
	for _, v := range h.GetBattle().GetEventCards("OnNROtherHeroGetTraits") {
		for _, v2 := range v.OnNROtherHeroGetTraits(h) {
			if !help.InArray(v2, ts) {
				ts = append(ts, v2)
			}
		}
	}

	return ts
}

// 是否拥有某种特质
func (h *Hero) IsHaveTraits(ct define.CardTraits) bool {
	return help.InArray(ct, h.GetTraits())
}
