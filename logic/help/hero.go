package help

import "hs/logic/iface"

// 从切片中删除一张卡
func DeleteCardFromCardsByIdx(cards []iface.ICard, idx int) (iface.ICard, []iface.ICard) {

	card := cards[idx]

	cards = append(cards[:idx], cards[idx+1:]...)

	return card, cards
}

// 补充一张卡到切片
func AddCardToCardsByIdx(cards []iface.ICard, idx int, card iface.ICard) []iface.ICard {
	return append(cards[:idx], append([]iface.ICard{card}, cards[idx:]...)...)
}
