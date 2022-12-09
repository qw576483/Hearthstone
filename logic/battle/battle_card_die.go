package battle

import (
	"fmt"
	"hs/logic/iface"
	"sort"
)

// 收集死亡
func (b *Battle) RecordCardDie(c iface.ICard) {
	b.recordCardsDie[c.GetReleaseId()] = c
}

// 触发死亡
func (b *Battle) TrickCardDie() {

	if len(b.recordCardsDie) <= 0 {
		return
	}

	rcd := b.recordCardsDie
	b.recordCardsDie = make(map[int]iface.ICard, 0)

	var keys []int
	for key := range rcd {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, v := range keys {
		c := rcd[v]
		c.GetOwner().TrickDieCardEvent(c)
	}
}

func (b *Battle) WhileTrickCardDie() {
	for i := 1; i <= 10; i++ {
		if len(b.recordCardsDie) > 0 {
			fmt.Println(b.recordCardsDie)
			b.TrickCardDie()
		} else {
			break
		}
	}
}
