package battle

import (
	"hs/logic/help"
	"hs/logic/iface"
)

// 获得事件
func (b *Battle) GetEvent() map[string][]iface.ICard {
	return b.events
}

// 获得事件卡牌
func (b *Battle) GetEventCards(e string) []iface.ICard {
	cs, ok := b.events[e]
	if ok {
		return cs
	}

	return make([]iface.ICard, 0)
}

// 添加卡牌到事件
func (b *Battle) AddCardToEvent(c iface.ICard, e string) {
	_, ok := b.events[e]
	if !ok {
		b.events[e] = make([]iface.ICard, 0)
	}

	b.events[e] = append(b.events[e], c)
}

// 删除一个卡牌事件
func (b *Battle) RemoveCardFromEvent(c iface.ICard, e string) {
	es, ok := b.events[e]
	if !ok {
		b.events[e] = make([]iface.ICard, 0)
		return
	}

	for idx, v := range es {
		if v.GetId() == c.GetId() {
			_, b.events[e] = help.DeleteCardFromCardsByIdx(es, idx)
		}
	}
}

// 删除卡牌从双方的事件中
func (b *Battle) RemoveCardFromAllEvent(c iface.ICard) {
	for e := range b.events {
		b.RemoveCardFromEvent(c, e)
	}
}
