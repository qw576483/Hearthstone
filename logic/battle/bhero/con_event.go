package bhero

import (
	"hs/logic/help"
	"hs/logic/iface"
)

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
