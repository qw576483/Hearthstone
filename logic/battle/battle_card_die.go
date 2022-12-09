package battle

import (
	"hs/logic/iface"
	"sort"
)

// 收集死亡
func (b *Battle) RecordCardDie(c iface.ICard) {

	b.recordCardDie[c.GetReleaseId()] = c
}

// 触发死亡
func (b *Battle) TrickCardDie() {

	var i []int
	for key := range b.recordCardDie {
		i = append(i, key)
	}
	sort.Ints(i)

	for _, v := range i {
		c := b.recordCardDie[v]
		c.GetOwner().TrickDieCardEvent(c)
	}

	b.recordCardDie = make(map[int]iface.ICard, 0)
}
